package services

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

// SlackService services object
type SlackService struct{}

// Available check to see if this service is available to use
func (s SlackService) Available() bool {
	token := os.Getenv("SLACK_TOKEN")

	return token != ""
}

func (s *SlackService) SetStatus(status string, emoji string) {
	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)

	err := api.SetUserCustomStatus(status, emoji)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}
