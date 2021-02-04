package apps

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xeipuuv/gojsonschema"
	"go.ketch.com/cli/ketch-cli/assets"
	"go.ketch.com/cli/ketch-cli/config"
	"go.ketch.com/cli/ketch-cli/flags"
	"go.ketch.com/lib/orlop"
	"go.ketch.com/lib/orlop/errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

func Publish(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	appConfig, err := cmd.Flags().GetString(flags.File)
	if err != nil {
		return err
	}

	versionCliArg, err := cmd.Flags().GetString(flags.Version)
	if err != nil {
		return err
	}

	create, err := cmd.Flags().GetBool(flags.Create)
	if err != nil {
		return err
	}

	token, err := cmd.Flags().GetString(flags.Token)
	if err != nil {
		return err
	}

	rootUrl, err := cmd.Flags().GetString(flags.URL)
	if err != nil {
		return err
	}

	t, err := config.GetTLSConfig(cmd)
	if err != nil {
		return err
	}

	appConfig, err = filepath.Abs(appConfig)
	if err != nil {
		return err
	}

	basePath := filepath.Dir(appConfig)
	if err = os.Chdir(basePath); err != nil {
		return err
	}

	b, err := ioutil.ReadFile(appConfig)
	if err != nil {
		return err
	}

	b = []byte(os.ExpandEnv(string(b)))

	var manifestInputs ManifestInputs
	if err := yaml.Unmarshal(b, &manifestInputs); err != nil {
		return err
	}

	if err := validateAppConfig(manifestInputs); err != nil {
		return err
	}

	cfg := config.GetFromContext(ctx)

	cfg.URL = rootUrl
	cfg.TLS = *t

	app, err := NewApp(manifestInputs)
	if err != nil {
		return err
	}

	var versionBumpType AppVersionBumpType
	if len(versionCliArg) > 0 {
		val, ok := AppVersionBumpTypeLookup[versionCliArg]
		if !ok {
			return errors.New(fmt.Sprintf("invalid version bump type '%s' (accepts: patch, minor, major)", versionCliArg))
		}

		versionBumpType = val
	} else if len(manifestInputs.VersionBumpType) > 0 {
		versionBumpType = AppVersionBumpTypeLookup[manifestInputs.VersionBumpType]
	}

	if len(app.ID) == 0 {
		if create {
			app, err = createApp(ctx, cfg, token, app, versionBumpType)
			if err != nil {
				return err
			}
		} else {
			return errors.New("app ID must be specified unless creating")
		}
	}

	marketplaceEntry := NewAppMarketplaceEntry(manifestInputs)
	marketplaceEntry.AppID = app.ID

	if len(manifestInputs.Logo.Link) > 0 {
		marketplaceEntry.Logo = Image{
			Title:  manifestInputs.Logo.Title,
			Width:  manifestInputs.Logo.Width,
			Height: manifestInputs.Logo.Height,
		}

		if marketplaceEntry.Logo.Data, err = getFileData(manifestInputs.Logo.Link); err != nil {
			return err
		}
	}

	if len(manifestInputs.Previews) > 0 {
		var previews []*Image

		for _, preview := range manifestInputs.Previews {
			previewImage := &Image{
				Title:  preview.Title,
				Width:  preview.Width,
				Height: preview.Height,
			}

			if previewImage.Data, err = getFileData(preview.Link); err != nil {
				return err
			}

			previews = append(previews, previewImage)
		}

		marketplaceEntry.Previews = previews
	}

	published, err := publishApp(ctx, cfg, token, app, marketplaceEntry, manifestInputs.Webhook)
	if err != nil {
		return err
	}

	fmt.Printf("app published successfully:\nappCode: %s\nappID: %s\nappVersion: %s\n", app.Code, app.ID, published.Version)

	return nil
}

func handleRestResponseError(resp *http.Response) (*Error, error) {
	var respErr *ErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&respErr)
	if err != nil {
		return nil, err
	}

	return respErr.Error, nil
}

func callRest(ctx context.Context, cfg *config.Config, method, urlPath, token string, body []byte) (*http.Response, error) {
	u, err := url.Parse(cfg.URL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "rest", "v1", urlPath)

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	tp := http.DefaultTransport.(*http.Transport).Clone()
	tp.TLSClientConfig, err = orlop.NewClientTLSConfigContext(ctx, cfg.TLS, cfg.Vault)
	if err != nil {
		return nil, err
	}

	cli := &http.Client{
		Transport: tp,
	}

	return cli.Do(req)
}

func createApp(ctx context.Context, cfg *config.Config, token string, app *App, versionBumpType AppVersionBumpType) (*App, error) {
	createAppURL := fmt.Sprintf("organizations/%s/apps", app.OrgCode)

	body, err := json.Marshal(&PutAppRequest{
		App:             app,
		VersionBumpType: versionBumpType,
	})
	if err != nil {
		return nil, err
	}

	resp, err := callRest(ctx, cfg, http.MethodPost, createAppURL, token, body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respErr, err := handleRestResponseError(resp)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(respErr.Message)
	}

	var appResp PutAppResponse
	err = json.NewDecoder(resp.Body).Decode(&appResp)
	if err != nil {
		return nil, err
	}

	return appResp.App, nil
}

func publishApp(ctx context.Context, cfg *config.Config, token string, app *App, marketplaceEntry *AppMarketplaceEntry, webhook *Webhook) (*AppMarketplaceEntry, error) {
	publishURL := fmt.Sprintf("organizations/%s/apps/%s/publish", app.OrgCode, app.ID)

	body, err := json.Marshal(&PublishAppRequest{
		AppMarketplaceEntry: marketplaceEntry,
		Webhook:             webhook,
	})
	if err != nil {
		return nil, err
	}

	resp, err := callRest(ctx, cfg, http.MethodPost, publishURL, token, body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respErr, err := handleRestResponseError(resp)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(respErr.Message)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var a PublishAppResponse
	err = json.Unmarshal(respBody, &a)
	if err != nil {
		return nil, err
	}

	return a.AppMarketplaceEntry, nil
}

func validateAppConfig(publishAppConfig ManifestInputs) error {
	manifestSchema := assets.GetAsset("/schemas/manifest.json")
	schemaLoader := gojsonschema.NewStringLoader(manifestSchema)

	appConfigLoader := gojsonschema.NewGoLoader(publishAppConfig)

	result, err := gojsonschema.Validate(schemaLoader, appConfigLoader)
	if err != nil {
		return err
	}

	if !result.Valid() {
		var errs []string
		for _, resultError := range result.Errors() {
			errs = append(errs, resultError.String()+"\n")
		}

		return errors.New(fmt.Sprintf("app config invalid: %s", errs))
	}

	return nil
}

func isRemoteLink(link string) bool {
	_, err := url.ParseRequestURI(link)
	if err != nil {
		return false
	}

	u, err := url.Parse(link)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func getFileData(link string) ([]byte, error) {
	if isRemoteLink(link) {
		return getRemoteFileData(link)
	}

	return getLocalFileData(link)
}

func getRemoteFileData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func getLocalFileData(link string) ([]byte, error) {
	if filepath.IsAbs(link) {
		return ioutil.ReadFile(link)
	}

	return ioutil.ReadFile(link)
}
