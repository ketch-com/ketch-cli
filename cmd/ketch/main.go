package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.ketch.com/cli/ketch-cli/cli"
	"go.ketch.com/cli/ketch-cli/flags"
	"go.ketch.com/cli/ketch-cli/version"
	"go.ketch.com/lib/orlop/log"
	stdlog "log"
	"os"
	"path"
)

func tlsFlags(f *pflag.FlagSet) {
	f.Bool(flags.TLSInsecure, os.Getenv("KETCH_TLS_INSECURE") == "true", "set true to skip certificate verification")
	f.String(flags.TLSCert, os.Getenv("KETCH_TLS_CERT_FILE"), "TLS client certificate")
	f.String(flags.TLSKey, os.Getenv("KETCH_TLS_KEY_FILE"), "TLS private key")
	f.String(flags.TLSCACert, os.Getenv("KETCH_TLS_CACERT_FILE"), "TLS root CA certificate")
	f.String(flags.TLSServerName, os.Getenv("KETCH_TLS_SERVER_NAME"), "override the TLS server name")
}

// Execute executes the command.
func Execute(ctx context.Context) error {
	var rootCmd = &cobra.Command{
		Use:              version.Name,
		Short:            version.Description,
		Version:          version.String(),
		TraverseChildren: true,
		SilenceUsage:     true,
	}

	rootCmd.SetHelpCommand(&cobra.Command{
		Use:   "help [command]",
		Short: "help about any command",
		Long: `Help provides help for any command in the application.
Simply type ` + rootCmd.Name() + ` help [path to command] for full details.`,

		Run: func(c *cobra.Command, args []string) {
			cmd, _, e := c.Root().Find(args)
			if cmd == nil || e != nil {
				c.Printf("Unknown help topic %#q\n", args)
				c.Root().Usage()
			} else {
				cmd.InitDefaultHelpFlag() // make possible 'help' flag to be shown
				cmd.Help()
			}
		},
	})

	var publish = &cobra.Command{
		RunE:  cli.PublishApp,
		Use:   "publish",
		Short: "publish an App",
	}

	publish.Flags().StringP(flags.File, "f", "ketch-manifest.yml", "app object file")
	publish.Flags().String(flags.Version, os.Getenv("KETCH_VERSION"), "app version")
	publish.Flags().String(flags.Token, os.Getenv("KETCH_TOKEN"), "token for Ketch API")
	publish.Flags().String(flags.URL, os.Getenv("KETCH_URL"), "url to Ketch API")
	publish.Flags().String(flags.Plugin, os.Getenv("KETCH_PLUGIN"), "path to the plugin.js file")
	publish.Flags().String(flags.Objects, os.Getenv("KETCH_OBJECTS"), "path to the objects directory")
	publish.Flags().String(flags.Assets, os.Getenv("KETCH_ASSETS"), "path to the assets directory")

	tlsFlags(publish.Flags())

	rootCmd.AddCommand(publish)

	var validate = &cobra.Command{
		RunE:  cli.ValidateApp,
		Use:   "validate",
		Short: "validate an App",
	}

	validate.Flags().StringP(flags.File, "f", "ketch-manifest.yml", "app object file")

	rootCmd.AddCommand(validate)

	var webhookCmd = &cobra.Command{
		Use:   "webhook",
		Short: "commands related to WebHooks",
	}

	webhookCmd.PersistentFlags().Uint64(flags.QPS, 10, "maximum QPS")
	webhookCmd.PersistentFlags().String(flags.Auth, "", "authorization header")
	webhookCmd.PersistentFlags().BytesBase64(flags.Secret, nil, "shared secret")
	tlsFlags(webhookCmd.PersistentFlags())

	var webhookSendCmd = &cobra.Command{
		RunE:  cli.SendWebhookEvent,
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
		RunE:  cli.ValidateWebhook,
		Use:   "validate",
		Short: "validate the given webhook",
		Args:  cobra.ExactArgs(1),
	}

	webhookCmd.AddCommand(webhookValidateCmd)

	rootCmd.AddCommand(webhookCmd)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		return err
	}

	return nil
}

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
		return
	}

	envFiles := []string{".env", path.Join(homeDir, ".ketchrc"), ".ketchrc"}
	for _, file := range envFiles {
		if _, err = os.Stat(file); err == nil {
			_ = godotenv.Overload(file)
		}
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})

	stdlog.SetOutput(logrus.New().Writer())

	if err = Execute(context.Background()); err != nil {
		os.Exit(1)
	}
}
