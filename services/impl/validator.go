package impl

import (
	"context"
	"github.com/xeipuuv/gojsonschema"
	"go.ketch.com/cli/ketch-cli/config"
	"go.ketch.com/cli/ketch-cli/services"
	"go.ketch.com/lib/app"
	"go.ketch.com/lib/orlop/errors"
)

type validator struct {
	cfg      *config.ValidatorConfig
	reporter services.Reporter
}

func NewValidator(cfg *config.ValidatorConfig, reporter services.Reporter) services.Validator {
	return &validator{
		cfg:      cfg,
		reporter: reporter,
	}
}

func (v *validator) Validate(ctx context.Context, app *app.App) error {
	return v.validate(ctx, gojsonschema.NewGoLoader(app))
}

func (v *validator) validate(ctx context.Context, appConfigLoader gojsonschema.JSONLoader) error {
	manifestSchema, err := app.JSONSchema()
	if err != nil {
		return err
	}

	result, err := gojsonschema.Validate(gojsonschema.NewBytesLoader(manifestSchema), appConfigLoader)
	if err != nil {
		return err
	}

	if !result.Valid() {
		for _, resultError := range result.Errors() {
			v.reporter.Report(ctx, resultError.String())
		}

		return errors.New("app config invalid")
	}

	return nil
}
