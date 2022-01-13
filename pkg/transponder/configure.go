package transponder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/pkg/config"
	"go.ketch.com/cli/ketch-cli/pkg/flags"
	"go.ketch.com/lib/orlop/errors"
	"io"
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
	if len(cfg.Token) > 0 {
		req.Header.Add("Authorization", "Bearer "+cfg.Token)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK &&  resp.StatusCode != http.StatusNoContent  {
		out := &ErrorResponseBody{}

		buf = bytes.NewBuffer(nil)
		_, err = io.Copy(buf, resp.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(buf.Bytes(), &out)
		if err != nil {
			fmt.Println(string(buf.Bytes()))
			return err
		}

		if len(out.Errors) > 0 {
			return errors.New(out.Errors[0].Detail)
		}

		return errors.New("failed to list connections")
	}

	fmt.Println("Initiating data discovery")
	fmt.Println("Delivering blueberries and pancakes")

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	out := &PutConnectionResponseBody{}

	buf = bytes.NewBuffer(nil)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf.Bytes(), &out)
	if err != nil {
		return err
	}

	return nil
}
