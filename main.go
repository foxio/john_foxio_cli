package main

import (
	"log"
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

	if command.FirstTimeSetup() {
		command.WriteConfigurationFile(app)
	}
	config := command.ReadConfiguration()

	app.Commands = []cli.Command{
		{
			Name:   "update",
			Usage:  "Update Johnny",
			Action: update,
		},
		{
			Name:    "init",
			Aliases: []string{"in"},
			Usage:   "Setup user defaults",
			Action: func(c *cli.Context) {
				command.InitUser(c)
			},
		},
		{
			Name:    "slack",
			Aliases: []string{"slack"},
			Usage:   "Interact with Slack",
			Subcommands: []cli.Command{
				{
					Name:  "status",
					Usage: "Set the user's slack status",
					Action: func(c *cli.Context) {
						command.SetSlackStatus(c)
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "status, s",
							Usage: "the status to send to slack",
						},
						cli.StringFlag{
							Name:  "emoji, e",
							Usage: "the emoji to use",
						},
					},
				},
				{
					Name:   "clear",
					Usage:  "Clear the user's slack status",
					Action: command.ClearSlackStatus,
				},
			},
		},
		{
			Name:    "pomodoro",
			Aliases: []string{"pom"},
			Usage:   "Starts and stops pomodoros",
			Subcommands: []cli.Command{
				{
					Name:  "start",
					Usage: "stars a new pom",
					Action: func(c *cli.Context) {
						command.PomodoroStart(c, config)
					},
					Flags: []cli.Flag{
						cli.IntFlag{
							Name:  "duration, d",
							Usage: "pom duration in minutes"}},
				},
				{
					Name:   "stop",
					Usage:  "stops a pom",
					Action: command.PomodoroStop,
				},
			},
		},
	}

	app.Run(os.Args)
}

func update(c *cli.Context) {
	log.Println("Updating ... ")

	command.WriteConfigurationFile(c.App)

	log.Println("Updating complete")
}
