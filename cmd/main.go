package main

import (
	"os"

	"github.com/jeewangue/oss-licenses/internal/completion"
	"github.com/jeewangue/oss-licenses/internal/conda"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {

	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)

	app := &cli.App{
		EnableBashCompletion: true,
		Name:                 "oss-licenses",
		Usage:                "export licenses of open source packages in the environment.",
		HideHelpCommand:      true,
		Commands: []*cli.Command{
			conda.Command,
			completion.Command,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}
