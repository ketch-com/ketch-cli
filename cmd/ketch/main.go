package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.ketch.com/cli/ketch-cli/cmd/ketch/commands"
	stdlog "log"
	"os"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})

	stdlog.SetOutput(logrus.New().Writer())

	if err := commands.Execute(context.Background()); err != nil {
		os.Exit(1)
	}
}
