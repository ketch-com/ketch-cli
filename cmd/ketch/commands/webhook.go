package commands

import (
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/flags"
	"go.ketch.com/cli/ketch-cli/webhooks"
)

func init() {
	var webhookCmd = &cobra.Command{
		Use:   "webhook",
		Short: "commands related to WebHooks",
	}

	webhookCmd.PersistentFlags().Uint64(flags.QPS, 10, "maximum QPS")
	webhookCmd.PersistentFlags().String(flags.Auth, "", "authorization header")
	webhookCmd.PersistentFlags().BytesBase64(flags.Secret, nil, "shared secret")
	webhookCmd.PersistentFlags().Bool(flags.TLSInsecure, false, "set true to skip certificate verification")
	webhookCmd.PersistentFlags().String(flags.TLSCert, "", "TLS client certificate")
	webhookCmd.PersistentFlags().String(flags.TLSKey, "", "TLS private key")
	webhookCmd.PersistentFlags().String(flags.TLSCACert, "", "TLS root CA certificate")
	webhookCmd.PersistentFlags().String(flags.TLSServerName, "", "override the TLS server name")

	var webhookSendCmd = &cobra.Command{
		RunE:  webhooks.SendWebhookEvent,
		Use:   "send",
		Short: "send event to the webhook specified",
		Args:  cobra.ExactArgs(1),
	}

	webhookSendCmd.Flags().String(flags.EventType, "", "event type")
	webhookSendCmd.Flags().String(flags.EventSource, "", "event source")
	webhookSendCmd.Flags().StringP(flags.File, "f", "", "filename")
	webhookSendCmd.Flags().String(flags.Org, "", "organization code")
	webhookSendCmd.Flags().String(flags.AppID, "", "appID")

	webhookCmd.AddCommand(webhookSendCmd)

	var webhookValidateCmd = &cobra.Command{
		RunE:  webhooks.Validate,
		Use:   "validate",
		Short: "validate the given webhook",
		Args:  cobra.ExactArgs(1),
	}

	webhookCmd.AddCommand(webhookValidateCmd)

	rootCmd.AddCommand(webhookCmd)
}
