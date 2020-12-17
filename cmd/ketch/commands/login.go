package commands

import (
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/auth"
)

func init() {
	var login = &cobra.Command{
		RunE:  auth.Login,
		Use:   "login",
		Short: "Connect the CLI to your Ketch account by logging in to persist your secret key locally.",
		Args:  cobra.NoArgs,
	}

	rootCmd.AddCommand(login)
}
