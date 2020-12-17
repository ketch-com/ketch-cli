package apps

import (
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/flags"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path"
)

func Validate(cmd *cobra.Command, args []string) error {
	root, err := cmd.Flags().GetString(flags.Root)
	if err != nil {
		return err
	}

	appConfig, err := cmd.Flags().GetString(flags.File)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadFile(path.Join(root, appConfig))
	if err != nil {
		return err
	}

	var publishAppConfig PublishAppConfig
	if err := yaml.Unmarshal(b, &publishAppConfig); err != nil {
		return err
	}

	if err := validateAppConfig(publishAppConfig); err != nil {
		return err
	}

	return nil
}
