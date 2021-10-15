package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/pkg/apps"
	"go.ketch.com/cli/ketch-cli/pkg/flags"
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

	appConfig, err = filepath.Abs(appConfig)
	if err != nil {
		return err
	}

	basePath := filepath.Dir(appConfig)
	if err = os.Chdir(basePath); err != nil {
		return err
	}

	var manifests []string
	info, err := os.Stat(appConfig)
	if err != nil {
		return err
	}

	if info.IsDir(){
		files, _ := ioutil.ReadDir(appConfig)
		for _, file := range files {
			if !file.IsDir() {
				manifests = append(manifests, fmt.Sprintf("%s/%s", appConfig, file.Name()))
			}
		}
	} else {
		manifests = []string{appConfig}
	}

	for _, manifest := range manifests {
		b, err := ioutil.ReadFile(manifest)
		if err != nil {
			return err
		}

		b = []byte(os.ExpandEnv(string(b)))

		var manifestInputs apps.ManifestInputs
		if err := yaml.Unmarshal(b, &manifestInputs); err != nil {
			return err
		}

		if err := validateAppConfig(manifestInputs); err != nil {
			return err
		}
	}

	return nil
}
