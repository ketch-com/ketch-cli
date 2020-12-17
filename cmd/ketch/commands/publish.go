package commands

import (
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/apps"
	"go.ketch.com/cli/ketch-cli/flags"
	"os"
)

func init() {
	var publish = &cobra.Command{
		RunE:  apps.Publish,
		Use:   "publish",
		Short: "publish an app",
	}

	publish.Flags().StringP(flags.File, "f", "ketch-manifest.yml", "app config name")
	publish.Flags().String(flags.Version, os.Getenv("KETCH_VERSION"), "app version")
	publish.Flags().Bool(flags.Create, false, "flag that specifies to create an app before publishing")
	publish.Flags().String(flags.Token, os.Getenv("KETCH_TOKEN"), "token for Ketch API")
	publish.Flags().String(flags.URL, os.Getenv("KETCH_URL"), "url to Ketch API")
	publish.Flags().Bool(flags.TLSInsecure, false, "set true to skip certificate verification")
	publish.Flags().String(flags.TLSCert, "", "TLS client certificate")
	publish.Flags().String(flags.TLSKey, "", "TLS private key")
	publish.Flags().String(flags.TLSCACert, "", "TLS root CA certificate")
	publish.Flags().String(flags.TLSServerName, "", "override the TLS server name")

	rootCmd.AddCommand(publish)
}
