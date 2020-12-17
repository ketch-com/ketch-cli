package config

import (
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/flags"
	"go.ketch.com/lib/orlop"
)

func GetTLSConfig(cmd *cobra.Command) (*orlop.TLSConfig, error) {
	var err error
	tls := &orlop.TLSConfig{}
	if tls.Insecure, err = cmd.Flags().GetBool(flags.TLSInsecure); err != nil {
		return nil, err
	}

	if tls.Cert.File, err = cmd.Flags().GetString(flags.TLSCert); err != nil {
		return nil, err
	}

	if tls.Key.File, err = cmd.Flags().GetString(flags.TLSKey); err != nil {
		return nil, err
	}

	if tls.RootCA.File, err = cmd.Flags().GetString(flags.TLSCACert); err != nil {
		return nil, err
	}

	if tls.Override, err = cmd.Flags().GetString(flags.TLSServerName); err != nil {
		return nil, err
	}

	if len(tls.Cert.File) > 0 {
		tls.Enabled = true
	}

	return tls, nil
}
