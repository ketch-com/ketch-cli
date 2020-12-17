package commands

import (
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/apps"
	"go.ketch.com/cli/ketch-cli/flags"
	"os"
)

func init() {
	var validate = &cobra.Command{
		RunE:  apps.Validate,
		Use:   "validate",
		Short: "validate an app manifest file",
	}

	validate.Flags().String(flags.Root, "./", "app root directory")
	validate.Flags().StringP(flags.File, "f", "ketch-manifest.yml", "app config name")
	validate.Flags().String(flags.ID, "", "app ID")
	validate.Flags().String(flags.Org, os.Getenv("KETCH_ORG"), "app organization code")

	rootCmd.AddCommand(validate)
}
