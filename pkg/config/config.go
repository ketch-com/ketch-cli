package config

import (
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/pkg/flags"
	"go.ketch.com/cli/ketch-cli/version"
	"go.ketch.com/lib/orlop/v2"
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
	}

	err := orlop.Unmarshal(version.Name, cfg)
	if err != nil {
		return nil, err
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

	return cfg, nil
}
