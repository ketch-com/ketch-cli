package commands

import (
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/apps"
	"go.ketch.com/cli/ketch-cli/flags"
)

func init() {
	var validate = &cobra.Command{
		RunE:  apps.Validate,
		Use:   "validate",
		Short: "validate an app manifest file",
	}

	validate.Flags().StringP(flags.File, "f", "ketch-manifest.yml", "app config name")

	rootCmd.AddCommand(validate)
}
