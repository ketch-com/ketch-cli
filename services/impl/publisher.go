package impl

import (
	"context"
	"encoding/json"
	"go.ketch.com/cli/ketch-cli/config"
	"go.ketch.com/cli/ketch-cli/services"
	"go.ketch.com/lib/app"
	"go.ketch.com/lib/orlop/errors"
	"os"
)

type publisher struct {
	cfg *config.PublisherConfig
}

func NewPublisher(cfg *config.PublisherConfig) services.Publisher {
	return &publisher{
		cfg: cfg,
	}
}

func (p *publisher) Publish(ctx context.Context, app *app.App) error {
	en := json.NewEncoder(os.Stdout)
	en.SetIndent("", "  ")
	if err := en.Encode(app); err != nil {
		return err
	}

	return errors.New("implement me")
}
