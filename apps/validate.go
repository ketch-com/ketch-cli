package apps

import (
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/flags"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Validate(cmd *cobra.Command, args []string) error {
	appConfig, err := cmd.Flags().GetString(flags.File)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadFile(appConfig)
	if err != nil {
		return err
	}

	appConfig, err = filepath.Abs(appConfig)
	if err != nil {
		return err
	}

	basePath := filepath.Dir(appConfig)
	if err = os.Chdir(basePath); err != nil {
		return err
	}

	var publishAppConfig ManifestInputs
	if err := yaml.Unmarshal(b, &publishAppConfig); err != nil {
		return err
	}

	if err := validateAppConfig(publishAppConfig); err != nil {
		return err
	}

	return nil
}
