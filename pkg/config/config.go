package config

import (
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/pkg/flags"
	"go.ketch.com/cli/ketch-cli/version"
	"go.ketch.com/lib/orlop"
)

type Config struct {
	Auth0Domain string
	ClientID    string
	Audience    string
	URL         string
	Token       string
	TLS         orlop.TLSConfig
	Vault       orlop.VaultConfig
}

func NewConfig(cmd *cobra.Command) (*Config, error) {
	cfg := &Config{
		Auth0Domain: "ketch-staging.us.auth0.com",
		ClientID:    "UERD5nrErbgvgKvqJdwzNaoWkYdxW8v9",
		Audience:    "https://dev.ketchapi.com/rest",
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

	if b, err := cmd.Flags().GetBool(flags.TLSInsecure); err != nil {
		return nil, err
	} else {
		cfg.TLS.Insecure = b
	}

	if s, err := cmd.Flags().GetString(flags.TLSCert); err != nil {
		return nil, err
	} else if len(s) > 0 {
		cfg.TLS.Cert.File = s
	}

	if s, err := cmd.Flags().GetString(flags.TLSKey); err != nil {
		return nil, err
	} else if len(s) > 0 {
		cfg.TLS.Key.File = s
	}

	if s, err := cmd.Flags().GetString(flags.TLSCACert); err != nil {
		return nil, err
	} else if len(s) > 0 {
		cfg.TLS.RootCA.File = s
	}

	if s, err := cmd.Flags().GetString(flags.TLSServerName); err != nil {
		return nil, err
	} else if len(s) > 0 {
		cfg.TLS.Override = s
	}

	if len(cfg.TLS.Cert.File) > 0 || cfg.TLS.Insecure {
		cfg.TLS.Enabled = true
	}

	return cfg, nil
}
