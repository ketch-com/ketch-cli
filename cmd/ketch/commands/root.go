package commands

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/flags"
	"go.ketch.com/cli/ketch-cli/version"
	"os"
	"path"
)

var rootCmd = &cobra.Command{
	Use:              version.Name,
	Short:            version.Description,
	Version:          version.String(),
	TraverseChildren: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		configFile, err := cmd.Flags().GetString(flags.Config)
		if err != nil {
			return err
		}

		envFiles := []string{".env", path.Join(homeDir, ".ketchrc"), ".ketchrc", configFile}
		for _, file := range envFiles {
			if f, err := os.Stat(file); err == nil && !f.IsDir() {
				_ = godotenv.Overload(file)
			}
		}

		return nil
	},
}

// Execute executes the command.
func Execute(ctx context.Context) error {
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
	rootCmd.PersistentFlags().String(flags.Config, ".ketchrc", "environment file")
	rootCmd.SilenceUsage = true

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		return err
	}

	return nil
}
