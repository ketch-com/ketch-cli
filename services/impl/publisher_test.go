package impl

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.ketch.com/cli/ketch-cli/config"
	"go.ketch.com/lib/app"
	"testing"
)

func TestPublish(t *testing.T) {
	ctx := context.Background()

	cfg := &config.PublisherConfig{
		AccessToken: "",
		URL:         "",
	}

	manifest := &app.App{}

	p := NewPublisher(cfg)
	err := p.Publish(ctx, manifest)
	assert.NoError(t, err)
}
