package services

import (
	"context"
	"go.ketch.com/cli/ketch-cli/config"
	"go.ketch.com/lib/app"
)

//go:generate mockery --all

// Reporter provides an interface to report messages
type Reporter interface {
	// Report a message to the user
	Report(ctx context.Context, format string, args ...interface{})
}

// Loader provides an interface to load an App based on the specified configuration
type Loader interface {
	Load(ctx context.Context, cfg *config.LoaderConfig) (*app.App, error)
	LoadExternalFiles(ctx context.Context, cfg *config.LoaderConfig, manifest *app.App) error
}

// Validator provides an interface to validate an App either in bytes form or object form
type Validator interface {
	Validate(ctx context.Context, app *app.App) error
}

// Publisher provides an interface to publish an App
type Publisher interface {
	Publish(ctx context.Context, app *app.App) error
}
