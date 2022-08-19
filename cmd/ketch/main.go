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
	"go.ketch.com/lib/orlop/v2/cmd"
	stdlog "log"
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

	stdlog.SetOutput(logrus.New().Writer())

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

	var runner = cmd.NewRunner(version.Name)

	rootCmd.PersistentFlags().String(flags.Token, runner.Getenv("TOKEN"), "auth token")
	rootCmd.PersistentFlags().String(flags.Config, ".ketchrc", "environment file")
	rootCmd.PersistentFlags().String(flags.URL, runner.Getenv("URL"), "url to Ketch API")
	rootCmd.PersistentFlags().Bool(flags.TLSInsecure, runner.Getenv("INSECURE") == "true", "set true to skip certificate verification")
	rootCmd.PersistentFlags().String(flags.TLSCert, runner.Getenv("TLS_CERT_FILE"), "TLS client certificate")
	rootCmd.PersistentFlags().String(flags.TLSKey, runner.Getenv("TLS_KEY_FILE"), "TLS private key")
	rootCmd.PersistentFlags().String(flags.TLSCACert, runner.Getenv("TLS_ROOTCA_FILE"), "TLS root CA certificate")
	rootCmd.PersistentFlags().String(flags.TLSServerName, runner.Getenv("TLS_OVERRIDE"), "override the TLS server name")
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

	var transponderListCmd = &cobra.Command{
		Use:   "ls",
		Short: "list connections",
		RunE:  transponder.List,
	}

	transponderCmd.AddCommand(transponderListCmd)

	var transponderConfigureCmd = &cobra.Command{
		Use:   "configure",
		Short: "configure a connection",
		RunE:  transponder.Configure,
		Args:  cobra.ExactArgs(1),
	}

	transponderConfigureCmd.Flags().StringToStringP(flags.Parameter, "P", nil, "parameter key/value")

	transponderCmd.AddCommand(transponderConfigureCmd)

	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		os.Exit(1)
	}
}
