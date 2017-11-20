package lib

import (
	"log"
	"os/user"
)

const (
	RootLogFolder = ".john_logs"
)

// HomeDir gets the users home directory
func HomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return usr.HomeDir, nil
}
