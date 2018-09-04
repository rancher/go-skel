package main

import (
	"os"

	"github.com/rancher/log"
	"github.com/urfave/cli"
)

var VERSION = "v0.0.0-dev"

func main() {
	app := cli.NewApp()
	app.Name = "%APP%"
	app.Version = VERSION
	app.Usage = "You need help!"
	app.Action = func(c *cli.Context) error {
		logrus.Info("I'm a turkey")
		return nil
	}

	app.Run(os.Args)
}
