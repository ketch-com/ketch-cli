package commands

import (
	"context"
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/config"
	"go.ketch.com/cli/ketch-cli/flags"
	"go.ketch.com/cli/ketch-cli/version"
	"os"
)

var rootCmd = &cobra.Command{
	Use:              version.Name,
	Short:            version.Description,
	Version:          version.String(),
	TraverseChildren: true,
}

// Execute executes the command.
func Execute(ctx context.Context, cfg *config.Config) error {
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

	rootCmd.PersistentFlags().String(flags.Token, os.Getenv("KETCH_TOKEN"), "auth token")
	rootCmd.SilenceUsage = true

	ctx = config.AddToContext(ctx, cfg)
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		return err
	}

	return nil
}
