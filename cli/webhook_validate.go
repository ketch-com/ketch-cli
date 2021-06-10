package cli

import (
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/flags"
	"go.ketch.com/lib/orlop"
	"go.ketch.com/lib/webhook-client/webhook"
)

func ValidateWebhook(cmd *cobra.Command, args []string) error {
	var err error

	vault := &orlop.VaultConfig{}

	maxQPS, err := cmd.Flags().GetUint64(flags.QPS)
	if err != nil {
		return err
	}

	authToken, err := cmd.Flags().GetString(flags.Auth)
	if err != nil {
		return err
	}

	auth := &orlop.KeyConfig{
		Secret: []byte(authToken),
	}

	secretToken, err := cmd.Flags().GetBytesBase64(flags.Secret)
	if err != nil {
		return err
	}

	secret := &orlop.KeyConfig{
		Secret: secretToken,
	}

	tls, err := getTLSConfig(cmd)
	if err != nil {
		return err
	}

	cli, err := webhook.NewClient(cmd.Context(), webhook.Binary, args[0], maxQPS, tls, auth, secret, vault)
	if err != nil {
		return err
	}

	return cli.Validate(cmd.Context())
}
