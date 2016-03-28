package command

import (
	"fmt"
	"time"

	"github.com/codegangsta/cli"
	"github.com/deckarep/gosx-notifier"
)

var doneChan chan bool
var duration int

// PomodoroStart starts a pom
func PomodoroStart(c *cli.Context) {
	doneChan = make(chan bool)

	fmt.Println("Pom starting")
	duration = c.Int("duration")
	displayNotification(fmt.Sprintf("Pom started for %d mintues", duration))

	go runTimer(duration, pomCompleted)

	fmt.Println("Waiting")
	<-doneChan
	fmt.Println("done")
	pomStartBreak()
}

// PomodoroStop stops a pom
func PomodoroStop(c *cli.Context) {
	fmt.Println("Ending pom ...")
	displayNotification("Pom stopped")
}

func pomCompleted() {
	displayNotification("Pom complete.")
	doneChan <- true
}

func pomStartBreak() {
	fmt.Println("Break starting")

	displayNotification("Break Time!")

	breakDuration := int(float64(duration) * 0.2)
	go runTimer(breakDuration, pomBreakOver)
	<-doneChan
}

func pomBreakOver() {
	displayNotification("Break Over.")
	doneChan <- true
}

func runTimer(maxMinutes int, callback func()) {
	startTime := time.Now()

	tick := time.NewTicker(1 * time.Minute)
	for now := range tick.C {

		minute := int(now.Sub(startTime).Minutes())
		fmt.Printf("%d minute\n", minute)
		if minute >= maxMinutes {
			tick.Stop()
			callback()
		}
	}
}

func displayNotification(message string) {
	note := gosxnotifier.NewNotification(message)
	note.Title = "John Foxio"
	note.AppIcon = "notification_icon.png"
	note.Group = "com.foxio.john_foio.pomodoro"
	note.Push()
}
