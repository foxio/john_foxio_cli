package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/foxio/john_foxio_cli/command"
)

var (
	configurationFile = ".john_foxio"
)

// Configuration represents the config file
type Configuration struct {
	Version string
}

func main() {
	app := cli.NewApp()
	app.Name = "John"
	app.Usage = "John Foxio, he's an extension of your team."
	app.Version = "0.0.1"
	app.Action = command.Default

	if firstTimeSetup() {
		writeConfigurationFile(app)
	} else {
		c := readConfiguration()
		log.Println(c)
	}

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
			Action:  command.InitUser,
		},
	}

	app.Run(os.Args)
}

func update(c *cli.Context) {
	println("Updating ... ")
	writeConfigurationFile(c.App)
	println("Updating complete")
}

func firstTimeSetup() bool {
	if _, err := os.Stat(configurationFilePath()); os.IsNotExist(err) {
		return true
	}
	return false
}

func readConfiguration() *Configuration {
	dat, err := ioutil.ReadFile(configurationFilePath())
	if err != nil {
		log.Fatal(err)
	}

	var c Configuration
	err = json.Unmarshal(dat, &c)
	if err != nil {
		log.Fatal(err)
	}

	return &c
}

func writeConfigurationFile(app *cli.App) {
	configuration := Configuration{app.Version}
	configurationJSON, err := json.Marshal(configuration)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(configurationFilePath(), configurationJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func configurationFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return strings.Join([]string{usr.HomeDir, configurationFile}, string(os.PathSeparator))
}
