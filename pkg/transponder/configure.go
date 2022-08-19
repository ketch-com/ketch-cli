package transponder

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/pkg/config"
	"go.ketch.com/cli/ketch-cli/pkg/flags"
	"go.ketch.com/lib/orlop/v2/errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
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

	// if value starts with @, then read file, otherwise it's a string value
	in, err := cmd.Flags().GetStringToString(flags.Parameter)
	if err != nil {
		return err
	}

	for key, value := range in {
		if strings.HasPrefix(value, "@") {
			// load from file
			fileName := strings.TrimPrefix(value, "@")
			data, err := os.ReadFile(fileName)
			if err != nil {
				return err
			}
			in[key] = base64.URLEncoding.EncodeToString(data)
		}
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

	fmt.Println("Initiating database connection")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
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

	successMsg := "Successfully configured the connection. Please run the transponder list command to verify that the connection is active"

	if resp.StatusCode == http.StatusNoContent {
		fmt.Println(successMsg)
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
		fmt.Println(string(buf.Bytes()))
		return err
	}

	fmt.Println(successMsg)

	return nil
}
