package lib

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	pomStart  = "POM Start"
	pomDone   = "POM Done"
	PomLogDir = "poms"
)

func currentLogFile() (*os.File, error) {
	year, month, day := time.Now().Date()
	homePath, err := HomeDir()
	if err != nil {
		log.Println("Could not find home dir: ", err)
		return nil, err
	}

	fileName := fmt.Sprintf("logs_%d_%d_%d", year, month, day)
	logFile := filepath.Join(homePath, RootLogFolder, PomLogDir, fileName)
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("error opening file: %v", err)
		return nil, err
	}
	return f, nil
}

// CountPomsLogged returns cound of todays logged poms
func CountPomsLogged() int {
	return 0
}

// LogPomStart logs pom start to the pom log file
func LogPomStart() {
	f, err := currentLogFile()
	defer f.Close()

	if err != nil {
		return
	}

	log.SetOutput(f)
	log.Println(pomStart)
}

// LogPomStart logs pom done to the pom log file
func LogPomComplete() {
	f, err := currentLogFile()
	defer f.Close()

	if err != nil {
		return
	}

	log.SetOutput(f)
	log.Println(pomDone)
}
