package impl

import (
	"context"
	"fmt"
	"go.ketch.com/cli/ketch-cli/services"
)

type reporter struct {}

func NewReporter() services.Reporter {
	return &reporter{}
}

func (reporter) Report(ctx context.Context, format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args...))
}

type nilReporter struct {}

func NewNilReporter() services.Reporter {
	return &nilReporter{}
}

func (nilReporter) Report(ctx context.Context, format string, args ...interface{}) {}
