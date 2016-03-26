package command

import (
	"fmt"
	"time"

	"github.com/codegangsta/cli"
)

// Default displays the default message
func Default(c *cli.Context) {
	fmt.Printf("Good %s friend! Ask for -help if you ever need it. \n", timeOfDay(time.Now()))
}

func timeOfDay(t time.Time) string {
	hour := t.Local().Hour()

	switch {
	case hour >= 0 && hour <= 11:
		return "morning"
	case hour > 11 && hour <= 17:
		return "afternoon"
	case hour > 17 && hour <= 23:
		return "evening"
	}
	return "day"
}
