package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xeipuuv/gojsonschema"
	"go.ketch.com/cli/ketch-cli/assets"
	"go.ketch.com/cli/ketch-cli/pkg/apps"
	"go.ketch.com/cli/ketch-cli/pkg/config"
	"go.ketch.com/cli/ketch-cli/pkg/flags"
	"go.ketch.com/lib/orlop"
	"go.ketch.com/lib/orlop/errors"
	"go.ketch.com/lib/orlop/log"
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
	cfg, err := config.NewConfig(cmd)
	if err != nil {
		return err
	}

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

	appConfig, err = filepath.Abs(appConfig)
	if err != nil {
		return err
	}

	basePath := filepath.Dir(appConfig)
	if err = os.Chdir(basePath); err != nil {
		return err
	}

	var manifests []string
	info, err := os.Stat(appConfig)
	if err != nil {
		return err
	}

	if info.IsDir(){
		files, _ := ioutil.ReadDir(appConfig)
		for _, file := range files {
			if !file.IsDir() {
				manifests = append(manifests, fmt.Sprintf("%s/%s", appConfig, file.Name()))
			}
		}
	} else {
		manifests = []string{appConfig}
	}

	for _, manifest := range manifests {
		b, err := ioutil.ReadFile(manifest)
		if err != nil {
			return err
		}

		b = []byte(os.ExpandEnv(string(b)))

		var manifestInputs apps.ManifestInputs
		if err := yaml.Unmarshal(b, &manifestInputs); err != nil {
			return err
		}

		if len(versionCliArg) > 0 {
			manifestInputs.Version = versionCliArg
		}

		if err = validateAppConfig(manifestInputs); err != nil {
			return err
		}

		app, err := apps.NewApp(manifestInputs)
		if err != nil {
			return err
		}

		if len(app.Version) == 0 {
			return errors.New("app version must be specified via cli --version or via manifest")
		}

		app, err = createApp(ctx, *cfg, token, app)
		if err != nil {
			return err
		}

		marketplaceEntry := apps.NewAppMarketplaceEntry(manifestInputs)
		marketplaceEntry.AppID = app.ID

		if manifestInputs.Logo != nil && len(manifestInputs.Logo.Link) > 0 {
			marketplaceEntry.Logo = apps.Image{
				Title:  manifestInputs.Logo.Title,
				Width:  manifestInputs.Logo.Width,
				Height: manifestInputs.Logo.Height,
			}

			if marketplaceEntry.Logo.Data, err = getFileData(manifestInputs.Logo.Link); err != nil {
				return err
			}
		}

		if len(manifestInputs.Previews) > 0 {
			var previews []*apps.Image

			for _, preview := range manifestInputs.Previews {
				previewImage := &apps.Image{
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

		published, err := publishApp(ctx, *cfg, token, app, marketplaceEntry, manifestInputs.Webhook)
		if err != nil {
			return err
		}

		fmt.Printf("app published successfully:\nappCode: %s\nappID: %s\nappVersion: %s\n", app.Code, app.ID, published.Version)
	}

	return nil
}

func handleRestResponseError(resp *http.Response) (*apps.Error, error) {
	var respErr *apps.ErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&respErr)
	if err != nil {
		return nil, err
	}

	return respErr.Error, nil
}

func callRest(ctx context.Context, cfg config.Config, method, urlPath, token string, body []byte) (*http.Response, error) {
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
	tp.TLSClientConfig, err = orlop.NewClientTLSConfig(ctx, cfg.TLS, cfg.Vault)
	if err != nil {
		return nil, err
	}

	cli := &http.Client{
		Transport: tp,
	}

	return cli.Do(req)
}

func createApp(ctx context.Context, cfg config.Config, token string, app *apps.App) (*apps.App, error) {
	createAppURL := fmt.Sprintf("organizations/%s/apps", app.OrgCode)

	body, err := json.Marshal(&apps.PutAppRequest{
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

	var appResp apps.PutAppResponse
	err = json.Unmarshal(b, &appResp)
	if err != nil {
		log.WithField("resp", string(b)).Error(err)
		return nil, err
	}

	if appResp.App == nil || len(appResp.App.ID) == 0 {
		return nil, errors.Errorf("app not created. statusCode %v, body %v", resp.StatusCode, string(b))
	}

	fmt.Printf("app created successfully:\napp: %v", string(body))

	return appResp.App, nil
}

func publishApp(ctx context.Context, cfg config.Config, token string, app *apps.App, marketplaceEntry *apps.AppMarketplaceEntry, webhook *apps.Webhook) (*apps.AppMarketplaceEntry, error) {
	publishURL := fmt.Sprintf("organizations/%s/apps/%s/publish", app.OrgCode, app.ID)

	body, err := json.Marshal(&apps.PublishAppRequest{
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

	var a apps.PublishAppResponse
	err = json.Unmarshal(respBody, &a)
	if err != nil {
		log.WithField("resp", string(respBody)).Error(err)
		return nil, err
	}

	if a.AppMarketplaceEntry == nil || len(a.AppMarketplaceEntry.AppID) == 0 {
		return nil, errors.Errorf("app marketplace entry not created. statusCode %v, body %v", resp.StatusCode, string(respBody))
	}

	return a.AppMarketplaceEntry, nil
}

func validateAppConfig(publishAppConfig apps.ManifestInputs) error {
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

	if publishAppConfig.Logo != nil && len(publishAppConfig.Logo.Link) > 0 {
		if _, err := filePathExists(publishAppConfig.Logo.Link); err != nil {
			return err
		}
	}

	if len(publishAppConfig.IdentitySpaces) > 0 {
		codes := make(map[string]interface{}, len(publishAppConfig.IdentitySpaces))
		for _, identitySpace := range publishAppConfig.IdentitySpaces {
			if _, ok := codes[identitySpace.Code]; ok {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"identitySpaces.code "+identitySpace.Code+" is not unique"))
			}

			codes[identitySpace.Code] = struct{}{}
		}
	}

	if len(publishAppConfig.Cookies) > 0 {
		codes := make(map[string]interface{}, len(publishAppConfig.Cookies))
		for _, cookie := range publishAppConfig.Cookies {
			if _, ok := codes[cookie.Code]; ok {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"cookies.code "+cookie.Code+" is not unique"))
			}

			if !isEntityCodeValid(publishAppConfig.Code, cookie.Code) {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"cookies.code must start with \""+publishAppConfig.Code+".\""))
			}

			codes[cookie.Code] = struct{}{}
		}
	}

	if len(publishAppConfig.PurposeTemplateCollections) > 0 {
		codes := make(map[string]interface{}, len(publishAppConfig.PurposeTemplateCollections))
		for _, purposeTemplateCollection := range publishAppConfig.PurposeTemplateCollections {
			if _, ok := codes[purposeTemplateCollection.Code]; ok {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"purposeTemplateCollections.code "+purposeTemplateCollection.Code+" is not unique"))
			}

			if !isEntityCodeValid(publishAppConfig.Code, purposeTemplateCollection.Code) {
				return errors.New(fmt.Sprintf("app config invalid: %s",
					"purposeTemplateCollections.code must start with \""+publishAppConfig.Code+".\""))
			}

			codes[purposeTemplateCollection.Code] = struct{}{}
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

func remoteFileDataExists(link string) (bool, error) {
	_, err := getRemoteFileData(link)
	if err != nil {
		return false, err
	}
	return true, nil
}

func localFilePathExists(link string) (bool, error) {

	logoFileInfo, err :=  os.Stat(link)
	if err != nil {
		return false, err
	}
	if logoFileInfo.IsDir() {
		return false, errors.New(fmt.Sprintf("app config invalid: logo.link %s is a directory", link))
	}
	return true, nil
}

func filePathExists(link string) (bool, error) {
	if isRemoteLink(link) {
		return remoteFileDataExists(link)
	}
	return localFilePathExists(link)
}
