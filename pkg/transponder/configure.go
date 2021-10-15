package transponder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/pkg/config"
	"go.ketch.com/cli/ketch-cli/pkg/flags"
	"go.ketch.com/lib/orlop/errors"
	"net/http"
	"net/url"
)

func Configure(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	cfg, err := config.NewConfig(cmd)
	if err != nil {
		return err
	}

	u, err := url.Parse(cfg.URL)
	if err != nil {
		return err
	}

	u.Path = fmt.Sprintf("/captain/connections/%s", args[0])

	client := http.Client{}

	in, err := cmd.Flags().GetStringToString(flags.Parameter)
	if err != nil {
		return err
	}

	b, err := json.Marshal(in)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(b)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, u.String(), buf)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

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
