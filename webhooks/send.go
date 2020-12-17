package webhooks

import (
	"encoding/json"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/spf13/cobra"
	"go.ketch.com/cli/ketch-cli/config"
	"go.ketch.com/cli/ketch-cli/flags"
	"go.ketch.com/lib/oid"
	"go.ketch.com/lib/orlop"
	"go.ketch.com/lib/orlop/errors"
	"go.ketch.com/lib/webhook-client/webhook"
	"io/ioutil"
	"time"
)

func SendWebhookEvent(cmd *cobra.Command, args []string) error {
	var err error

	vault := &orlop.VaultConfig{}

	orgCode, err := cmd.Flags().GetString(flags.Org)
	if err != nil {
		return err
	}

	appID, err := cmd.Flags().GetString(flags.AppID)
	if err != nil {
		return err
	}

	eventType, err := cmd.Flags().GetString(flags.EventType)
	if err != nil {
		return err
	}

	if len(eventType) == 0 {
		return errors.New("event type is required")
	}

	eventSource, err := cmd.Flags().GetString(flags.EventSource)
	if err != nil {
		return err
	}

	if len(eventType) == 0 {
		return errors.New("event source is required")
	}

	filename, err := cmd.Flags().GetString(flags.File)
	if err != nil {
		return err
	}

	if len(filename) == 0 {
		return errors.New("filename is required")
	}

	maxQPS, err := cmd.Flags().GetUint64(flags.QPS)
	if err != nil {
		return err
	}

	authToken, err := cmd.Flags().GetString(flags.Auth)
	if err != nil {
		return err
	}

	auth := &orlop.KeyConfig{
		Secret: []byte(authToken),
	}

	secretToken, err := cmd.Flags().GetBytesBase64(flags.Secret)
	if err != nil {
		return err
	}

	secret := &orlop.KeyConfig{
		Secret: secretToken,
	}

	tls, err := config.GetTLSConfig(cmd)
	if err != nil {
		return err
	}

	cli, err := webhook.NewClient(cmd.Context(), webhook.Binary, args[0], maxQPS, tls, auth, secret, vault)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	id, err := oid.NewOID()
	if err != nil {
		return err
	}

	e := cloudevents.NewEvent()
	e.SetID(id)
	e.SetType(eventType)
	e.SetTime(time.Now().UTC())
	e.SetSource(eventSource)

	err = e.SetData(cloudevents.ApplicationJSON, data)
	if err != nil {
		return err
	}

	if len(orgCode) != 0 {
		e.SetExtension("orgcode", orgCode)
	}

	if len(appID) != 0 {
		e.SetExtension("appid", appID)
	}

	if err = cli.Send(cmd.Context(), &e); err != nil && err != webhook.Accepted {
		return err
	}

	return nil
}
