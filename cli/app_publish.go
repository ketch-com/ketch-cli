package cli

import (
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/config"
	"go.ketch.com/cli/ketch-cli/flags"
	"go.ketch.com/cli/ketch-cli/services/impl"
	"net/http"
)

func PublishApp(cmd *cobra.Command, args []string) error {
	var err error
	ctx := cmd.Context()

	validatorConfig := &config.ValidatorConfig{}

	loaderConfig := &config.LoaderConfig{}

	if loaderConfig.AppConfigFile, err = cmd.Flags().GetString(flags.File); err != nil {
		return err
	}

	if loaderConfig.Version, err = cmd.Flags().GetString(flags.Version); err != nil {
		return err
	}

	if loaderConfig.PluginFilename, err = cmd.Flags().GetString(flags.Plugin); err != nil {
		return err
	}

	if loaderConfig.ObjectsDir, err = cmd.Flags().GetString(flags.Objects); err != nil {
		return err
	}

	if loaderConfig.AssetsDir, err = cmd.Flags().GetString(flags.Assets); err != nil {
		return err
	}

	publisherConfig := &config.PublisherConfig{}

	if publisherConfig.AccessToken, err = cmd.Flags().GetString(flags.Token); err != nil {
		return err
	}

	if publisherConfig.URL, err = cmd.Flags().GetString(flags.URL); err != nil {
		return err
	}

	validator := impl.NewValidator(validatorConfig, impl.NewReporter())
	loader := impl.NewLoader(impl.NewReporter(), &http.Client{})
	publisher := impl.NewPublisher(publisherConfig)

	manifest, err := loader.Load(ctx, loaderConfig)
	if err != nil {
		return err
	}

	err = validator.Validate(ctx, manifest)
	if err != nil {
		return err
	}

	err = publisher.Publish(ctx, manifest)
	if err != nil {
		return err
	}

	return nil
}

