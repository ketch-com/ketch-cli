package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/pkg/config"
	"go.ketch.com/lib/orlop/v2/errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type DeviceCode struct {
	// The device verification code.
	DeviceCode string `json:"device_code,omitempty"`

	// The end-user verification code.
	UserCode string `json:"user_code"`

	// The end-user verification URI on the authorization
	// server.  The URI should be short and easy to remember as end users
	// will be asked to manually type it into their user agent.
	VerificationUri string `json:"verification_uri"`

	// A verification URI that includes the "user_code" (or
	// other information with the same function as the "user_code"),
	// which is designed for non-textual transmission.
	VerificationUriComplete string `json:"verification_uri_complete"`

	// The lifetime in seconds of the "device_code" and "user_code".
	ExpiresInSec int64 `json:"expires_in"`

	// The minimum amount of time in seconds that the client
	// SHOULD wait between polling requests to the token endpoint.  If no
	// value is provided, clients MUST use 5 as the default.
	IntervalInSec int64 `json:"interval"`
}

type Token struct {
	AccessToken  string `json:"access_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	ExpiresInSec int64  `json:"expires_in,omitempty"`
}

func Login(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	cfg, err := config.NewConfig(cmd)
	if err != nil {
		return err
	}

	codeRequest := make(url.Values)
	codeRequest.Set("client_id", cfg.ClientID)
	codeRequest.Set("audience", cfg.Audience)

	buf := bytes.NewReader([]byte(codeRequest.Encode()))

	cli := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("https://%s/oauth/device/code", cfg.Auth0Domain), buf)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var dc DeviceCode

		err := json.NewDecoder(resp.Body).Decode(&dc)
		if err != nil {
			return err
		}

		fmt.Printf("Now, go to %s and enter the following code:\n", dc.VerificationUri)
		fmt.Printf("%s\n", dc.UserCode)

		timeout := time.After(time.Duration(dc.ExpiresInSec) * time.Second)
		ticker := time.NewTicker(time.Duration(dc.IntervalInSec) * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return nil

			case <-ticker.C:
				if tok, ok, err := check(ctx, cfg, dc); err != nil {
					return err
				} else if ok {
					fmt.Println(tok)
					return nil
				}

			case <-timeout:
				return errors.New("login: timed out")
			}
		}
	} else {
		io.Copy(os.Stderr, resp.Body)
		return errors.Errorf("login: status %s", resp.Status)
	}
}

func check(ctx context.Context, cfg *config.Config, dc DeviceCode) (string, bool, error) {
	var err error

	tokenRequest := make(url.Values)
	tokenRequest.Set("grant_type", "urn:ietf:params:oauth:grant-type:device_code")
	tokenRequest.Set("client_id", cfg.ClientID)
	tokenRequest.Set("device_code", dc.DeviceCode)

	buf := bytes.NewReader([]byte(tokenRequest.Encode()))

	cli := http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("https://%s/oauth/token", cfg.Auth0Domain), buf)
	if err != nil {
		return "", false, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := cli.Do(req)
	if err != nil {
		return "", false, err
	}

	defer resp.Body.Close()

	var token Token

	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return "", false, err
	}

	if token.TokenType == "Bearer" {
		return token.AccessToken, true, nil
	}

	return "", false, nil
}
