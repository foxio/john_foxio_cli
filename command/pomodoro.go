package command

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/codegangsta/cli"
	notifier "github.com/deckarep/gosx-notifier"
	"github.com/tbruyelle/hipchat-go/hipchat"

	"github.com/foxio/john_foxio_cli/lib"
	"github.com/foxio/john_foxio_cli/services"
)

var doneChan chan bool
var duration int
var breakDuration int

// PomodoroConfiguration represents the pom config file
type PomodoroConfiguration struct {
	RunTime int
	Break   int
}

// PomodoroStart starts a pom
func PomodoroStart(c *cli.Context, config *Configuration) {
	doneChan = make(chan bool)

	duration = c.Int("duration")
	if duration <= 0 {
		duration = config.Pomodoro.RunTime
	}
	breakDuration = config.Pomodoro.Break

	fmt.Printf("Pom started for %d mintues\n", duration)
	displayNotification(fmt.Sprintf("Pom started for %d mintues", duration))

	go runTimer(duration, pomCompleted, pomTick)

	userPresence := hipchat.UpdateUserPresenceRequest{
		Show:   hipchat.UserPresenceShowDnd,
		Status: fmt.Sprintf("In %dm Pom", duration),
	}
	updateHipChatStatus(userPresence)
	updateSlackStatus(fmt.Sprintf("In %dm Pom", duration), ":timer_clock:")
	lib.LogPomStart()

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)
	go func() {
		<-interruptChan
		fmt.Println("")
		PomodoroStop(c)
		updateSlackStatus("", "")
		lib.LogPomInterrupt()
		os.Exit(0)
	}()

	<-doneChan
	fmt.Println("done")
	lib.LogPomComplete()
	pomStartBreak()
}

// PomodoroStop stops a pom
func PomodoroStop(c *cli.Context) {
	fmt.Println("Ending pom ...")
	displayNotification("Pom stopped")
}

// PomodoroStop stops a pom
func PomodoroCount(c *cli.Context) {
	fmt.Println("Poms completed today: ", lib.CountPomsLogged())
}

// PomodoroShow prints today's log file
func PomodoroShow(c *cli.Context) {
	fmt.Println(lib.TodaysPomsLogged())
}

func pomCompleted() {
	displayNotification("Pom complete.")
	doneChan <- true
}

func pomTick(maxMinutes int, minute int) {
	if minute%2 == 0 {
		userPresence := hipchat.UpdateUserPresenceRequest{
			Show:   hipchat.UserPresenceShowDnd,
			Status: fmt.Sprintf("%dm left in Pom", maxMinutes-minute),
		}
		updateHipChatStatus(userPresence)
		updateSlackStatus(fmt.Sprintf("%dm left in Pom", maxMinutes-minute), ":timer_clock:")
	}
}

func breakTick(maxMinutes int, minute int) {

}

func pomStartBreak() {
	fmt.Println("Break starting")

	userPresence := hipchat.UpdateUserPresenceRequest{
		Show:   hipchat.UserPresenceShowChat,
		Status: "",
	}
	updateHipChatStatus(userPresence)
	updateSlackStatus("", "")

	displayNotification("Break Time!")

	go runTimer(breakDuration, pomBreakOver, breakTick)
	<-doneChan
}

func pomBreakOver() {
	displayNotification("Break Over.")
	doneChan <- true
}

func runTimer(maxMinutes int, callback func(), tickCallback func(maxMinutes int, minute int)) {
	startTime := time.Now()

	fmt.Printf("\r0 minute")
	tick := time.NewTicker(1 * time.Minute)
	for now := range tick.C {

		minute := int(now.Sub(startTime).Minutes())
		fmt.Printf("\r%d minute", minute)

		tickCallback(maxMinutes, minute)

		if minute >= maxMinutes {
			fmt.Printf("\n")
			tick.Stop()
			callback()
		}
	}
}

func displayNotification(message string) {
	note := notifier.Notification{
		Title:   "John Foxio",
		Message: message,
		AppIcon: "./command/notification_icon.png",
		Sound:   notifier.Hero,
	}

	note.Push()
}

func updateSlackStatus(status string, emoji string) {
	slackService := services.SlackService{}
	if slackService.Available() {
		slackService.SetStatus(status, emoji)
	}
}

func updateHipChatStatus(userPresence hipchat.UpdateUserPresenceRequest) {
	var service services.Servicer
	hipchatService := services.HipchatService{}
	service = hipchatService
	if service.Available() {
		hipchatService.SetStatus(userPresence)
	}
}
