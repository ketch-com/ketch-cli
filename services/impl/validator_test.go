package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"go.ketch.com/cli/ketch-cli/config"
	"go.ketch.com/cli/ketch-cli/utils"
	"go.ketch.com/lib/app"
	"go.ketch.com/lib/orlop/errors"
	"os"
	"testing"
)

func TestValidate(t *testing.T) {
	ctx := context.Background()

	cfg := &config.ValidatorConfig{}

	p := NewValidator(cfg, NewReporter())

	for _, testCase := range []struct {
		Name     string
		Error    error
	}{
		{
			Name:  "fixture2",
			Error: nil,
		},
		{
			Name:  "fixture3",
			Error: nil,
		},
		{
			Name:  "fixture4",
			Error: nil,
		},
		{
			Name:  "fixture5",
			Error: errors.New("app config invalid"),
		},
		{
			Name:  "fixture6",
			Error: errors.New("app config invalid"),
		},
		{
			Name:  "fixture7",
			Error: nil,
		},
		{
			Name:  "fixture8",
			Error: nil,
		},
		{
			Name:  "fixture9",
			Error: nil,
		},
		{
			Name:  "fixture10",
			Error: errors.New("app config invalid"),
		},
		{
			Name:  "fixture11",
			Error: nil,
		},
		{
			Name:  "fixture12",
			Error: nil,
		},
		{
			Name:  "fixture13",
			Error: nil,
		},
	} {
		t.Run(testCase.Name, func(t *testing.T) {
			b, err := os.ReadFile(fmt.Sprintf("../../test/fixtures/%s/ketch-manifest.yaml", testCase.Name))
			require.NoError(t, err)

			manifest := &app.App{}
			b, err = utils.YAMLtoJSON(b)
			require.NoError(t, err)

			err = json.Unmarshal(b, &manifest)
			require.NoError(t, err)

			err = p.Validate(ctx, manifest)
			if testCase.Error == nil {
				require.NoError(t, err)
			} else if err == nil {
				require.Error(t, err)
			} else {
				require.EqualError(t, err, testCase.Error.Error())
			}
		})
	}
}
