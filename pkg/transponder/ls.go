package transponder

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/pkg/config"
	"go.ketch.com/lib/orlop/v2/errors"
	"net/http"
	"net/url"
)

func List(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	cfg, err := config.NewConfig(cmd)
	if err != nil {
		return err
	}

	u, err := url.Parse(cfg.URL)
	if err != nil {
		return err
	}

	u.Path = "/captain/connections"

	client := http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	if len(cfg.Token) > 0 {
		req.Header.Add("Authorization", "Bearer "+cfg.Token)
	} else if len(cfg.ApiKey) > 0 {
		req.Header.Add("Authorization", "Ketch-Api-Key "+cfg.ApiKey)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		out := &ErrorResponseBody{}

		err = json.NewDecoder(resp.Body).Decode(&out)
		if err != nil {
			return err
		}

		if len(out.Errors) > 0 {
			return errors.New(out.Errors[0].Detail)
		}

		return errors.New("failed to list connections")
	}

	out := &FindConnectionsResponseBody{}

	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil {
		return err
	}

	fmt.Println(out)
	return nil
}
