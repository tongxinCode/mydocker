package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "mydocker",
		Usage: "mydocker is a simple container runtime implementation.",
	}

	app.Commands = []*cli.Command{
		&runCommand,
		&initCommand,
		&commitCommand,
		&listCommand,
		&logCommand,
		&execCommand,
		&stopCommand,
		&removeCommand,
		&networkCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
