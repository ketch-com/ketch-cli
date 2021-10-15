package config

import (
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/pkg/flags"
	"go.ketch.com/cli/ketch-cli/version"
	"go.ketch.com/lib/orlop"
)

type Config struct {
	URL   string
	TLS   orlop.TLSConfig
	Vault orlop.VaultConfig
}

func NewConfig(cmd *cobra.Command) (*Config, error) {
	cfg := &Config{}

	err := orlop.Unmarshal(version.Name, cfg)
	if err != nil {
		return nil, err
	}

	if s, err := cmd.Flags().GetString(flags.URL); err != nil {
		return nil, err
	} else {
		cfg.URL = s
	}

	if b, err := cmd.Flags().GetBool(flags.TLSInsecure); err != nil {
		return nil, err
	} else {
		cfg.TLS.Insecure = b
	}

	if s, err := cmd.Flags().GetString(flags.TLSCert); err != nil {
		return nil, err
	} else {
		cfg.TLS.Cert.File = s
	}

	if s, err := cmd.Flags().GetString(flags.TLSKey); err != nil {
		return nil, err
	} else {
		cfg.TLS.Key.File = s
	}

	if s, err := cmd.Flags().GetString(flags.TLSCACert); err != nil {
		return nil, err
	} else {
		cfg.TLS.RootCA.File = s
	}

	if s, err := cmd.Flags().GetString(flags.TLSServerName); err != nil {
		return nil, err
	} else {
		cfg.TLS.Override = s
	}

	if len(cfg.TLS.Cert.File) > 0 || cfg.TLS.Insecure {
		cfg.TLS.Enabled = true
	}

	return cfg, nil
}
