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
	configurationFile            = ".john_foxio"
	configurationFilePermissions = os.FileMode(0644)
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
	log.Println("Updating ... ")

	writeConfigurationFile(c.App)

	log.Println("Updating complete")
}

func firstTimeSetup() bool {
	if _, err := os.Stat(configurationFilePath()); os.IsNotExist(err) {
		return true
	}
	return false
}

func readConfiguration() *Configuration {
	dat, err := ioutil.ReadFile(configurationFilePath())
	check(err)

	var c Configuration
	err = json.Unmarshal(dat, &c)
	check(err)

	return &c
}

func writeConfigurationFile(app *cli.App) {
	configuration := Configuration{app.Version}
	configurationJSON, err := json.Marshal(configuration)
	check(err)

	err = ioutil.WriteFile(configurationFilePath(), configurationJSON, configurationFilePermissions)
	check(err)
}

func configurationFilePath() string {
	usr, err := user.Current()
	check(err)

	return strings.Join([]string{usr.HomeDir, configurationFile}, string(os.PathSeparator))
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
