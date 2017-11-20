package command

import (
	"log"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/foxio/john_foxio_cli/lib"
)

var (
	rootLogFolder = ".john_logs"
)

// InitUser handles setting up the user defaults
func InitUser(c *cli.Context) {
	println("TODO: init task")
	createLogDirs()
}

func createLogDirs() {
	homePath, err := lib.HomeDir()
	if err != nil {
		log.Println("Could not find home dir: ", err)
		return
	}

	logDirs := []string{lib.PomLogDir}

	for _, logDir := range logDirs {
		dir := filepath.Join(homePath, lib.RootLogFolder, logDir)
		if shouldCreateDir(dir) {
			createLogDir(dir)
		}
	}
}

func shouldCreateDir(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return true
	}
	return false
}

func createLogDir(logDir string) {
	os.MkdirAll(logDir, os.ModePerm)
}
