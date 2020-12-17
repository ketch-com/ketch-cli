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

	publish.Flags().String(flags.Root, "./", "app root directory")
	publish.Flags().StringP(flags.File, "f", "ketch-manifest.yml", "app config name")
	publish.Flags().String(flags.ID, "", "app ID")
	publish.Flags().String(flags.Org, os.Getenv("KETCH_ORG"), "app organization code")

	rootCmd.AddCommand(publish)
}
