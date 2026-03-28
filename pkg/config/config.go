package config

import (
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/pkg/flags"
	"go.ketch.com/cli/ketch-cli/version"
	"go.ketch.com/lib/orlop/v2/config"
	"go.ketch.com/lib/orlop/v2/errors"
	"net/url"
)

type Config struct {
	Auth0Domain string
	ClientID    string
	Audience    string
	URL         string
	Token       string
}

func NewConfig(cmd *cobra.Command) (*Config, error) {
	cfg := &Config{
		Auth0Domain: "ketch.us.auth0.com",
		ClientID:    "j9gemizsXis5IcUg931sBjGoyGSxbT1a",
		Audience:    "https://global.ketchapi.com/rest",
		URL:         config.GetEnv(version.Name, "URL"),
		Token:       config.GetEnv(version.Name, "TOKEN"),
	}

	if domain := config.GetEnv(version.Name, "AUTH0_DOMAIN"); len(domain) > 0 {
		cfg.Auth0Domain = domain
	}

	if clientID := config.GetEnv(version.Name, "CLIENT_ID"); len(clientID) > 0 {
		cfg.ClientID = clientID
	}

	if audience := config.GetEnv(version.Name, "AUDIENCE"); len(audience) > 0 {
		cfg.Audience = audience
	}

	if s, err := cmd.Flags().GetString(flags.URL); err != nil {
		return nil, err
	} else if len(s) > 0 {
		cfg.URL = s
	}

	if s, err := cmd.Flags().GetString(flags.Token); err != nil {
		return nil, err
	} else if len(s) > 0 {
		cfg.Token = s
	}

	if len(cfg.URL) == 0 {
		return nil, errors.Invalidf("url is required. either specify --url or set KETCH_URL environment variable")
	}

	u, err := url.Parse(cfg.URL)
	if err != nil || u.Scheme != "https" {
		return nil, errors.Invalidf("url is invalid. check the value of --url or KETCH_URL (url provided is '%s')", cfg.URL)
	}

	return cfg, nil
}
