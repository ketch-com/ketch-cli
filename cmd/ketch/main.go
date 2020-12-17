package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.ketch.com/cli/ketch-cli/cmd/ketch/commands"
	"go.ketch.com/cli/ketch-cli/config"
	"go.ketch.com/cli/ketch-cli/version"
	"go.ketch.com/lib/orlop"
	"go.ketch.com/lib/orlop/log"
	stdlog "log"
	"os"
	"path"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
		return
	}

	envFiles := []string{".env", path.Join(homeDir, ".ketchrc"), ".ketchrc"}
	for _, file := range envFiles {
		if _, err = os.Stat(file); err == nil {
			_ = godotenv.Overload(file)
		}
	}

	cfg := &config.Config{}
	if err = orlop.Unmarshal(version.Name, cfg); err != nil {
		log.Fatal(err)
		return
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})

	stdlog.SetOutput(logrus.New().Writer())

	if err = commands.Execute(context.Background(), cfg); err != nil {
		os.Exit(1)
	}
}
