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

	version, err := cmd.Flags().GetString(flags.Version)
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

	var publishAppConfig PublishAppConfig
	if err := yaml.Unmarshal(b, &publishAppConfig); err != nil {
		return err
	}

	if err := validateAppConfig(publishAppConfig); err != nil {
		return err
	}

	cfg := config.GetFromContext(ctx)

	cfg.URL = rootUrl
	cfg.TLS = *t

	app, err := NewApp(publishAppConfig)
	if err != nil {
		return err
	}

	if len(version) > 0 {
		app.Version = version
	}

	if len(app.ID) == 0 {
		if create {
			app.ID, err = createApp(ctx, cfg, token, app)
			if err != nil {
				return err
			}
		} else {
			return errors.New("app ID must be specified unless creating")
		}
	} else {
		err = updateApp(ctx, cfg, token, app)
		if err != nil {
			return err
		}
	}

	marketplaceEntry := NewAppMarketplaceEntry(publishAppConfig)
	marketplaceEntry.AppID = app.ID

	if len(publishAppConfig.Logo.Link) > 0 {
		marketplaceEntry.Logo = Image{
			Title:  publishAppConfig.Logo.Title,
			Width:  publishAppConfig.Logo.Width,
			Height: publishAppConfig.Logo.Height,
		}

		if marketplaceEntry.Logo.Data, err = getFileData(publishAppConfig.Logo.Link); err != nil {
			return err
		}
	}

	if len(publishAppConfig.Previews) > 0 {
		var previews []*Image

		for _, preview := range publishAppConfig.Previews {
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

	err = publishApp(ctx, cfg, token, app, marketplaceEntry, publishAppConfig.Webhook)
	if err != nil {
		return err
	}

	fmt.Printf("app published successfully: %s\n", app.ID)

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

func createApp(ctx context.Context, cfg *config.Config, token string, app *App) (string, error) {
	createAppURL := fmt.Sprintf("organizations/%s/apps", app.OrgCode)

	body, err := json.Marshal(app)
	if err != nil {
		return "", err
	}

	resp, err := callRest(ctx, cfg, http.MethodPost, createAppURL, token, body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		respErr, err := handleRestResponseError(resp)
		if err != nil {
			return "", err
		}

		return "", errors.New(respErr.Message)
	}

	var appResp PutAppResponse
	err = json.NewDecoder(resp.Body).Decode(&appResp)
	if err != nil {
		return "", err
	}

	return appResp.App.ID, nil
}

func updateApp(ctx context.Context, cfg *config.Config, token string, app *App) error {
	updateAppURL := fmt.Sprintf("organizations/%s/apps/%s", app.OrgCode, app.ID)

	body, err := json.Marshal(app)
	if err != nil {
		return err
	}

	resp, err := callRest(ctx, cfg, http.MethodPut, updateAppURL, token, body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		respErr, err := handleRestResponseError(resp)
		if err != nil {
			return err
		}

		return errors.New(respErr.Message)
	}

	return nil
}

func publishApp(ctx context.Context, cfg *config.Config, token string, app *App, marketplaceEntry *AppMarketplaceEntry, webhook *Webhook) error {
	publishURL := fmt.Sprintf("organizations/%s/apps/%s/publish", app.OrgCode, app.ID)

	body, err := json.Marshal(&PublishAppRequest{
		AppMarketplaceEntry: marketplaceEntry,
		Webhook:             webhook,
	})
	if err != nil {
		return err
	}

	resp, err := callRest(ctx, cfg, http.MethodPost, publishURL, token, body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		respErr, err := handleRestResponseError(resp)
		if err != nil {
			return err
		}

		return errors.New(respErr.Message)
	}

	return nil
}

func validateAppConfig(publishAppConfig PublishAppConfig) error {
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
