package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/config"
	"go.ketch.com/lib/orlop"
	"go.ketch.com/lib/orlop/errors"
	"net/http"
	"net/url"
	"time"
)

const ClientID = "TODO"

type GetDeviceCodeRequest struct {
	// The client identifier
	ClientID string `json:"client_id,omitempty"`

	// The scope of the access request
	Scope string `json:"scope,omitempty"`
}

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
	RefreshToken string `json:"refresh_token,omitempty"`
}

func Login(cmd *cobra.Command, args []string) error {
	var err error

	ctx := cmd.Context()
	cfg := config.GetFromContext(ctx)

	codeRequest := make(url.Values)
	codeRequest.Set("client_id", ClientID)

	buf := bytes.NewReader([]byte(codeRequest.Encode()))

	tp := http.DefaultTransport.(*http.Transport).Clone()
	if tp.TLSClientConfig, err = orlop.NewClientTLSConfigContext(ctx, cfg.Rest.TLS, cfg.Vault); err != nil {
		return err
	}

	cli := http.Client{
		Transport: tp,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, cfg.Rest.URL+"/device/code", buf)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		defer resp.Body.Close()

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
				tokenRequest := make(url.Values)
				tokenRequest.Set("grant_type", "urn:ietf:params:oauth:grant-type:device_code")
				tokenRequest.Set("client_id", ClientID)
				tokenRequest.Set("device_code", dc.DeviceCode)

				buf := bytes.NewReader([]byte(codeRequest.Encode()))

				tp := http.DefaultTransport.(*http.Transport).Clone()
				if tp.TLSClientConfig, err = orlop.NewClientTLSConfigContext(ctx, cfg.Rest.TLS, cfg.Vault); err != nil {
					return err
				}

				cli := http.Client{
					Transport: tp,
				}

				req, err := http.NewRequestWithContext(ctx, http.MethodPost, cfg.Rest.URL+"/token", buf)
				if err != nil {
					return err
				}

				req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				resp, err := cli.Do(req)
				if err != nil {
					return err
				}

				fmt.Println(resp.Body)
				continue

			case <-timeout:
				return errors.New("login: timed out")
			}
		}
	} else {
		return errors.Errorf("login: status %s", resp.Status)
	}
}
