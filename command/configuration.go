package command

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/codegangsta/cli"
)

var (
	configurationFile            = ".john_foxio"
	configurationFilePermissions = os.FileMode(0644)
)

// Configuration represents the config file
type Configuration struct {
	Version  string
	Pomodoro PomodoroConfiguration
}

// FirstTimeSetup checks if a config file exists
func FirstTimeSetup() bool {
	if _, err := os.Stat(configurationFilePath()); os.IsNotExist(err) {
		return true
	}
	return false
}

// ReadConfiguration reads the config file
func ReadConfiguration() *Configuration {
	dat, err := ioutil.ReadFile(configurationFilePath())
	check(err)

	var c Configuration
	err = json.Unmarshal(dat, &c)
	check(err)

	return &c
}

// WriteConfigurationFile writes default config file
func WriteConfigurationFile(app *cli.App) {
	pom := PomodoroConfiguration{25, 5}
	configuration := Configuration{app.Version, pom}

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
