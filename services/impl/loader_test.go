package impl

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/types"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.ketch.com/cli/ketch-cli/config"
	"go.ketch.com/lib/app"
	"go.ketch.com/lib/orlop/errors"
	"testing"
)

func copyApp(in app.App, modify func(*app.App)) app.App {
	var out app.App
	copier.Copy(&out, &in)
	modify(&out)
	return out
}

func TestLoad(t *testing.T) {
	ctx := context.Background()

	for _, testCase := range []struct {
		Name     string
		Config   *config.LoaderConfig
		Error    error
		Expected app.App
	}{
		{
			Name:   "fixture1",
			Config: &config.LoaderConfig{},
			Error:  errors.New("apiVersion '' is invalid (expected 'v1')"),
		},
		{
			Name:   "fixture2",
			Config: &config.LoaderConfig{},
			Error:  errors.New("kind '' is invalid (expected 'App')"),
		},
		{
			Name:   "fixture3",
			Config: &config.LoaderConfig{},
			Error:  errors.New("metadata is required"),
		},
		{
			Name:   "fixture4",
			Config: &config.LoaderConfig{},
			Error:  errors.New("metadata is required"),
		},
		{
			Name:   "fixture5",
			Config: &config.LoaderConfig{},
			Error:  errors.New("data is required"),
		},
		{
			Name:   "fixture6",
			Config: &config.LoaderConfig{},
			Error:  nil,
		},
		{
			Name:   "fixture7",
			Config: &config.LoaderConfig{},
			Expected: app.App{
				ApiVersion: "v1",
				Kind:       "App",
				Metadata: &app.AppMetadata{
					Code: "fixture",
					Name: "Fixture",
				},
				Data: &app.AppData{
					Version:         "0.0.0",
					PrimaryCategory: "privacy",
				},
			},
		},
		{
			Name:   "fixture8",
			Config: &config.LoaderConfig{},
			Expected: app.App{
				ApiVersion: "v1",
				Kind:       "App",
				Metadata: &app.AppMetadata{
					Code:        "fixture",
					Name:        "Fixture",
					Description: "Fixture description",
				},
				Data: &app.AppData{
					Version:         "0.0.0",
					PrimaryCategory: "privacy",
				},
			},
		},
		{
			Name:   "fixture9",
			Config: &config.LoaderConfig{},
			Expected: app.App{
				ApiVersion: "v1",
				Kind:       "App",
				Metadata: &app.AppMetadata{
					Code:        "fixture",
					Name:        "Fixture",
					Description: "Fixture description",
				},
				Data: &app.AppData{
					Version:           "0.0.0",
					PrimaryCategory:   "privacy",
					SecondaryCategory: "privacy",
				},
			},
		},
	} {
		t.Run(testCase.Name, func(t *testing.T) {
			p := NewLoader(NewNilReporter(), nil)

			testCase.Config.Version = "0.0.0"
			testCase.Config.AppConfigFile = fmt.Sprintf("../../test/fixtures/%s/ketch-manifest.yaml", testCase.Name)

			actual, err := p.Load(ctx, testCase.Config)
			if testCase.Error == nil {
				require.NoError(t, err)
			} else if err == nil {
				require.Error(t, err)
			} else {
				require.EqualError(t, err, testCase.Error.Error())
			}

			if testCase.Expected.ApiVersion != "" {
				assert.EqualValues(t, testCase.Expected, *actual)
			}
		})
	}
}

func TestLoadExternalFiles(t *testing.T) {
	ctx := context.Background()

	baseApp := app.App{
		ApiVersion: "v1",
		Kind:       "App",
		Metadata: &app.AppMetadata{
			Code:        "fixture",
			Name:        "Fixture",
			Description: "Fixture description",
		},
		Data: &app.AppData{
			AutoUpgrade:  true,
			Capabilities: []string{"permitPropagation"},
			Contacts: []*app.AppContact{
				{
					Email: "technical@example.com",
					Name:  "Technical",
					Type:  "technical",
				},
			},
			CustomerSupportURL:  "https://support",
			Depends:             "a, b, c",
			DetailedDescription: "detailed desc",
			DocURL:              "https://docs",
			ExpireUserTokens:    true,
			Form: []*app.FormComponent{
				{
					Data:            nil,
					Default:         "",
					Description:     "",
					Editable:        false,
					MaxLength:       0,
					MinLength:       0,
					Multiple:        false,
					Name:            "input",
					Pattern:         "",
					Placeholder:     "",
					Required:        false,
					ShowOnView:      false,
					ShowValueOnEdit: false,
					Spellcheck:      false,
					Type:            "string",
				},
			},
			HomepageURL: "https://homepage",
			IdentitySpaces: []*app.IdentitySpace{
				{
					Linked: &app.ReferencedIdentitySpace{
						Code:        "linked",
						Description: "linked identity",
						Filters: []*app.IdentitySpaceFilter{
							{
								Format: "raw",
								Type:   "ios_advertising_id",
							},
						},
						Name: "linked idfa",
					},
				},
				{
					Managed: &app.ManagedIdentitySpace{
						Code:        "managed",
						Description: "managed identity",
						Format:      "raw",
						Name:        "managed aaid",
						Type:        "android_advertising_id",
						Variable: &app.IdentitySpaceVariable{
							Jwt: &app.IdentitySpaceVariableJwt{
								Key:      "key",
								Location: "claims",
							},
							Location: "cookie",
							Name:     "myjwt",
						},
					},
				},
			},
			InfoURL:   "https://info",
			Instances: "single",
			Logo: &app.AppImage{
				Code:        "logo",
				ContentType: "image/png",
				Height:      234,
				Title:       "Logo",
				Width:       345,
				Contents:    []byte{},
			},
			Permissions: []string{"scope1"},
			Previews: []*app.AppImage{
				{
					Code:        "preview",
					ContentType: "image/png",
					Height:      123,
					Title:       "Preview 1",
					Width:       456,
					Contents:    []byte{},
				},
			},
			PrimaryCategory:  "privacy",
			PrivacyPolicyURL: "https://privacy",
			Provides: []string{
				"sku1",
				"sku2",
			},
			RedirectOnUpdate: true,
			RefreshInterval:  "2h",
			RequestUserAuth:  true,
			Rules: &app.AppDataRules{
				Install: "to install",
				View:    "to view",
			},
			SecondaryCategory:   "privacy",
			SetupURL:            "https://setup",
			ShortDescription:    "short description",
			StatusURL:           "https://status",
			SupportedLanguages:  []string{"en"},
			SupportedPurposes:   []string{"analytics"},
			SupportedRights:     []string{"delete"},
			TosURL:              "https://tos",
			Type:                "custom",
			UserAuthCallbackURL: "https://auth",
			Version:             "0.0.0",
			Webhook: &app.Webhook{
				Events: []string{"ping"},
				MaxQPS: 0,
				Tls: &app.WebhookTls{
					Insecure: true,
				},
				Url:    "https://webhook",
				Secret: []byte{},
			},
			Assets:        []*app.AppAsset{},
			CustomObjects: []*app.AppObject{},
		},
	}

	for _, testCase := range []struct {
		Name     string
		Config   *config.LoaderConfig
		Error    error
		Expected app.App
	}{
		{
			Name:   "fixture10",
			Config: &config.LoaderConfig{},
			Error: errors.New("open images/logo.png: no such file or directory"),
		},
		{
			Name:   "fixture11",
			Config: &config.LoaderConfig{},
			Expected: copyApp(baseApp, func(a *app.App) {
				a.Data.Logo.Contents = []byte("logo\n")
				a.Data.Previews[0].Contents = []byte("preview1\n")
			}),
		},
		{
			Name:   "fixture12",
			Config: &config.LoaderConfig{},
			Expected: copyApp(baseApp, func(a *app.App) {
				a.Data.Logo.Contents = []byte("logo\n")
				a.Data.Previews[0].Contents = []byte("preview1\n")
				a.Data.Assets = append(a.Data.Assets, &app.AppAsset{
					ContentType: "application/javascript",
					Contents:    []byte("alert(1);\n"),
					Name:        "plugin.js",
				})
			}),
		},
		{
			Name:   "fixture13",
			Config: &config.LoaderConfig{},
			Expected: copyApp(baseApp, func(a *app.App) {
				a.Data.Logo.Contents = []byte("logo\n")
				a.Data.Previews[0].Contents = []byte("preview1\n")
				a.Data.Assets = append(a.Data.Assets, &app.AppAsset{
					ContentType: "application/javascript",
					Contents:    []byte("alert(1);\n"),
					Name:        "plugin.js",
				})
				a.Data.Assets = append(a.Data.Assets, &app.AppAsset{
					ContentType: "image/png",
					Name:        "image.png",
					Contents:    []byte("image\n"),
				})
			}),
		},
		{
			Name:   "fixture14",
			Config: &config.LoaderConfig{},
			Error: errors.New("custom object 'legalBasis1.yaml' is empty"),
		},
		{
			Name:   "fixture15",
			Config: &config.LoaderConfig{},
			Error: errors.New("custom object 'legalBasis1.yaml' is invalid - apiVersion must be 'v1'"),
		},
		{
			Name:   "fixture16",
			Config: &config.LoaderConfig{},
			Error: errors.New("custom object 'legalBasis1.yaml' is invalid - kind is empty"),
		},
		{
			Name:   "fixture17",
			Config: &config.LoaderConfig{},
			Error: errors.New("custom object 'legalBasis1.yaml' is invalid - metadata is empty"),
		},
		{
			Name:   "fixture18",
			Config: &config.LoaderConfig{},
			Error: errors.New("custom object 'legalBasis1.yaml' is invalid - data is empty"),
		},
		{
			Name:   "fixture19",
			Config: &config.LoaderConfig{},
			Expected: copyApp(baseApp, func(a *app.App) {
				a.Data.Logo.Contents = []byte("logo\n")
				a.Data.Previews[0].Contents = []byte("preview1\n")
				a.Data.Assets = append(a.Data.Assets, &app.AppAsset{
					ContentType: "application/javascript",
					Contents:    []byte("alert(1);\n"),
					Name:        "plugin.js",
				})
				a.Data.Assets = append(a.Data.Assets, &app.AppAsset{
					ContentType: "image/png",
					Name:        "image.png",
					Contents:    []byte("image\n"),
				})
				a.Data.CustomObjects = append(a.Data.CustomObjects, &app.AppObject{
					ApiVersion: "v1",
					Kind:       "LegalBasis",
					Metadata: &app.AppMetadata{
						Code: "disclosure",
						Name: "Disclosure",
					},
					Data: &types.Struct{
						Fields: map[string]*types.Value{
							"requiresOptIn": {
								Kind: &types.Value_BoolValue{
									BoolValue: true,
								},
							},
						},
					},
				})
			}),
		},
	} {
		t.Run(testCase.Name, func(t *testing.T) {
			p := NewLoader(NewNilReporter(), nil)

			testCase.Config.Version = "0.0.0"
			testCase.Config.AppConfigFile = fmt.Sprintf("../../test/fixtures/%s/ketch-manifest.yaml", testCase.Name)

			actual, err := p.Load(ctx, testCase.Config)
			require.NoError(t, err)

			err = p.LoadExternalFiles(ctx, testCase.Config, actual)
			if testCase.Error == nil {
				require.NoError(t, err)
			} else if err == nil {
				require.Error(t, err)
			} else {
				require.EqualError(t, err, testCase.Error.Error())
			}

			if testCase.Expected.ApiVersion != "" {
				assert.EqualValues(t, testCase.Expected, *actual)
			}
		})
	}
}
