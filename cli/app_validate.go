package cli

import (
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/config"
	"go.ketch.com/cli/ketch-cli/flags"
	"go.ketch.com/cli/ketch-cli/services/impl"
	"net/http"
)

func ValidateApp(cmd *cobra.Command, args []string) error {
	var err error

	ctx := cmd.Context()

	validatorConfig := &config.ValidatorConfig{}

	loaderConfig := &config.LoaderConfig{
		Version: "0.0.0",
	}

	if loaderConfig.AppConfigFile, err = cmd.Flags().GetString(flags.File); err != nil {
		return err
	}

	if loaderConfig.Env, err = cmd.Flags().GetStringToString(flags.Env); err != nil {
		return err
	}

	reporter := impl.NewReporter()

	loader := impl.NewLoader(reporter, &http.Client{})

	manifest, err := loader.Load(ctx, loaderConfig)
	if err != nil {
		return err
	}

	if err = impl.NewValidator(validatorConfig, reporter).Validate(ctx, manifest); err != nil {
		return err
	}

	return nil
}
