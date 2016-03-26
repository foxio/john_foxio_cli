package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/foxio/john_foxio_cli/command"
)

func main() {
	app := cli.NewApp()
	app.Name = "John"
	app.Usage = "John Foxio, he's an extension of your team."
	app.Version = "0.0.1"
	app.Action = command.Default

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"in"},
			Usage:   "Setup user defaults",
			Action:  command.InitUser,
		},
	}

	app.Run(os.Args)
}
