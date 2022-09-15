package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/pkg/cli"
	"go.ketch.com/cli/ketch-cli/pkg/flags"
	"go.ketch.com/cli/ketch-cli/pkg/transponder"
	"go.ketch.com/cli/ketch-cli/version"
	"log"
	"os"
	"path"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})

	log.SetOutput(logrus.New().Writer())

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

	rootCmd.PersistentFlags().String(flags.Config, ".ketchrc", "environment file")
	rootCmd.PersistentFlags().String(flags.Token, "", "Ketch API authorization token")
	rootCmd.PersistentFlags().String(flags.URL, "", "url to Ketch API")
	rootCmd.SilenceUsage = true

	//
	// Login
	//

	var loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Connect the CLI to your Ketch account by logging in to persist your secret key locally.",
		Args:  cobra.NoArgs,
		RunE:  cli.Login,
	}

	rootCmd.AddCommand(loginCmd)

	//
	// Transponder
	//

	var transponderCmd = &cobra.Command{
		Use:              "transponder",
		Short:            "transponder commands",
		TraverseChildren: true,
	}

	rootCmd.AddCommand(transponderCmd)

	var transponderRotateCmd = &cobra.Command{
		Use:   "rotate",
		Short: "rotate API key",
		RunE:  transponder.Rotate,
	}

	transponderCmd.AddCommand(transponderRotateCmd)

	var transponderConnListCmd = &cobra.Command{
		Use:     "ls",
		Short:   "list connections",
		RunE:    transponder.List,
		Aliases: []string{"list"},
	}

	transponderCmd.AddCommand(transponderConnListCmd)

	var transponderConnConfigureCmd = &cobra.Command{
		Use:     "configure",
		Short:   "configure a connection",
		RunE:    transponder.Configure,
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"conf", "config"},
	}

	transponderConnConfigureCmd.Flags().StringToStringP(flags.Parameter, "P", nil, "parameter key/value")

	transponderCmd.AddCommand(transponderConnConfigureCmd)

	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		os.Exit(1)
	}
}
