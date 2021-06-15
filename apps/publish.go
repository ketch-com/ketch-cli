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
	"go.ketch.com/cli/ketch-cli/version"
	"go.ketch.com/lib/orlop"
	"go.ketch.com/lib/orlop/errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
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

	if len(versionCliArg) > 0 {
		manifestInputs.Version = versionCliArg
	}

	if err = validateAppConfig(manifestInputs); err != nil {
		return err
	}

	cfg := &config.Config{}
	if err = orlop.Unmarshal(version.Name, cfg); err != nil {
		return err
	}

	cfg.URL = rootUrl
	cfg.TLS = *t

	app, err := NewApp(manifestInputs)
	if err != nil {
		return err
	}

	if len(app.Version) == 0 {
		return errors.New("app version must be specified via cli --version or via manifest")
	}

	app, err = createApp(ctx, cfg, token, app)
	if err != nil {
		return err
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

func createApp(ctx context.Context, cfg *config.Config, token string, app *App) (*App, error) {
	createAppURL := fmt.Sprintf("organizations/%s/apps", app.OrgCode)

	body, err := json.Marshal(&PutAppRequest{
		App: app,
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

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var appResp PutAppResponse
	err = json.Unmarshal(b, &appResp)
	if err != nil {
		return nil, err
	}

	if appResp.App == nil || len(appResp.App.ID) == 0 {
		return nil, errors.Errorf("app not created. statusCode %v, body %v", resp.StatusCode, string(b))
	}

	fmt.Printf("app created successfully:\napp: %v", string(body))

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

	if a.AppMarketplaceEntry == nil || len(a.AppMarketplaceEntry.AppID) == 0 {
		return nil, errors.Errorf("app marketplace entry not created. statusCode %v, body %v", resp.StatusCode, string(respBody))
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

	if publishAppConfig.Type != AppTypeCustom {
		return nil
	}

	if len(publishAppConfig.IdentitySpaces) > 0 {
		codes := make(map[string]interface{}, len(publishAppConfig.IdentitySpaces))
		for _, identitySpace := range publishAppConfig.IdentitySpaces {
			if _, ok := codes[identitySpace.Code]; ok {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"identitySpaces.code "+identitySpace.Code+" is not unique"))
			}

			if !isEntityCodeValid(publishAppConfig.Code, identitySpace.Code) {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"identitySpaces.code must start with \""+publishAppConfig.Code+".\""))
			}

			codes[identitySpace.Code] = struct{}{}
		}
	}

	if len(publishAppConfig.PurposeTemplates) > 0 {
		codes := make(map[string]interface{}, len(publishAppConfig.PurposeTemplates))
		for _, purposeTemplate := range publishAppConfig.PurposeTemplates {
			if _, ok := codes[purposeTemplate.Code]; ok {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"purposeTemplates.code "+purposeTemplate.Code+" is not unique"))
			}

			if !isEntityCodeValid(publishAppConfig.Code, purposeTemplate.Code) {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"purposeTemplates.code must start with \""+publishAppConfig.Code+".\""))
			}

			codes[purposeTemplate.Code] = struct{}{}
		}
	}

	if len(publishAppConfig.Purposes) > 0 {
		codes := make(map[string]interface{}, len(publishAppConfig.Purposes))
		for _, purpose := range publishAppConfig.Purposes {
			if _, ok := codes[purpose.Code]; ok {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"purposes.code "+purpose.Code+" is not unique"))
			}

			if !isEntityCodeValid(publishAppConfig.Code, purpose.Code) {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"purposes.code must start with \""+publishAppConfig.Code+".\""))
			}

			codes[purpose.Code] = struct{}{}
		}
	}

	if len(publishAppConfig.PolicyScopes) > 0 {
		codes := make(map[string]interface{}, len(publishAppConfig.PolicyScopes))
		for _, policyScope := range publishAppConfig.PolicyScopes {
			if _, ok := codes[policyScope.Code]; ok {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"policyScopes.code "+policyScope.Code+" is not unique"))
			}

			if !isEntityCodeValid(publishAppConfig.Code, policyScope.Code) {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"policyScopes.code must start with \""+publishAppConfig.Code+".\""))
			}

			codes[policyScope.Code] = struct{}{}
		}
	}

	if len(publishAppConfig.LegalBases) > 0 {
		codes := make(map[string]interface{}, len(publishAppConfig.LegalBases))
		for _, legalBasis := range publishAppConfig.LegalBases {
			if _, ok := codes[legalBasis.Code]; ok {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"legalBases.code "+legalBasis.Code+" is not unique"))
			}

			if !isEntityCodeValid(publishAppConfig.Code, legalBasis.Code) {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"legalBases.code must start with \""+publishAppConfig.Code+".\""))
			}

			codes[legalBasis.Code] = struct{}{}
		}
	}

	if len(publishAppConfig.Themes) > 0 {
		codes := make(map[string]interface{}, len(publishAppConfig.Themes))
		for _, theme := range publishAppConfig.Themes {
			if _, ok := codes[theme.Code]; ok {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"themes.code "+theme.Code+" is not unique"))
			}

			if !isEntityCodeValid(publishAppConfig.Code, theme.Code) {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"themes.code must start with \""+publishAppConfig.Code+".\""))
			}

			codes[theme.Code] = struct{}{}
		}
	}

	if len(publishAppConfig.Rights) > 0 {
		codes := make(map[string]interface{}, len(publishAppConfig.Rights))
		for _, right := range publishAppConfig.Rights {
			if _, ok := codes[right.Code]; ok {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"rights.code "+right.Code+" is not unique"))
			}

			if !isEntityCodeValid(publishAppConfig.Code, right.Code) {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"right.code must start with \""+publishAppConfig.Code+".\""))
			}

			codes[right.Code] = struct{}{}
		}
	}

	if len(publishAppConfig.Regulations) > 0 {
		codes := make(map[string]interface{}, len(publishAppConfig.Regulations))
		for _, regulation := range publishAppConfig.Regulations {
			if _, ok := codes[regulation.Code]; ok {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"regulations.code "+regulation.Code+" is not unique"))
			}

			if !isEntityCodeValid(publishAppConfig.Code, regulation.Code) {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"regulation.code must start with \""+publishAppConfig.Code+".\""))
			}

			codes[regulation.Code] = struct{}{}
		}
	}

	return nil
}

func isEntityCodeValid(appCode, entityCode string) bool {
	if strings.HasPrefix(entityCode, appCode+".") {
		return true
	}

	return false
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
