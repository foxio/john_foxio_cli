package command

import (
	"github.com/codegangsta/cli"
	"github.com/foxio/john_foxio_cli/services"
)

// SetSlackStatus Sets the slack rooms status
func SetSlackStatus(c *cli.Context) {
	status := c.String("status")
	emoji := c.String("emoji")

	slackService := services.SlackService{}
	slackService.SetStatus(status, emoji)
}

// ClearSlackStatus clears the slack status
func ClearSlackStatus(c *cli.Context) {
	slackService := services.SlackService{}
	slackService.SetStatus("", "")
}
