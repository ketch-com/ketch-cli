package transponder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/pkg/config"
	"go.ketch.com/lib/orlop/v2/errors"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
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

	u.Path = path.Join(u.Path, "/captain", "connections")

	client := http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	if len(cfg.Token) > 0 {
		req.Header.Add("Authorization", "Bearer "+cfg.Token)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		out := &ErrorResponseBody{}

		b := bytes.NewBuffer(nil)
		_, err = io.Copy(b, resp.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(b.Bytes(), &out)
		if err != nil {
			fmt.Println(string(b.Bytes()))
			return err
		}

		if len(out.Errors) > 0 {
			return errors.New(out.Errors[0].Detail)
		}

		return errors.New("failed to list connections")
	}

	out := &FindConnectionsResponseBody{}

	b := bytes.NewBuffer(nil)
	_, err = io.Copy(b, resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b.Bytes(), &out)
	if err != nil {
		fmt.Println(string(b.Bytes()))
		return err
	}

	columns := []int{20, 40, 20, 20, 10}
	fmt.Println(padLeft("code", columns[0]), padLeft("name", columns[1]), padLeft("provider", columns[2]), padLeft("technology", columns[3]), padLeft("status", columns[4]))
	fmt.Println(strings.Repeat("=", columns[0]), strings.Repeat("=", columns[1]), strings.Repeat("=", columns[2]), strings.Repeat("=", columns[3]), strings.Repeat("=", columns[4]))

	for _, conn := range out.Data {
		fmt.Println(padLeft(conn.Code, columns[0]), padLeft(conn.Name, columns[1]), padLeft(conn.Provider, columns[2]), padLeft(conn.Technology, columns[3]), padLeft(string(conn.Status), columns[4]))
	}

	return nil
}

func padLeft(s string, n int) string {
	if len(s) >= n {
		return s[0:n]
	}

	return s + strings.Repeat(" ", n-len(s))
}

func padRight(s string, n int) string {
	if len(s) >= n {
		return s[0:n]
	}

	return strings.Repeat(" ", n-len(s)) + s
}
