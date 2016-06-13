package services

import (
	"fmt"
	"net/http"

	"os"

	"github.com/tbruyelle/hipchat-go/hipchat"
)

// HipchatService services object
type HipchatService struct{}

// Available check to see if this service is available to use
func (h HipchatService) Available() bool {
	apiKey := os.Getenv("HIPCHAT_API_KEY")
	userID := os.Getenv("HIPCHAT_USER_ID")

	return apiKey != "" && userID != ""
}

// SetStatus sets the status for the given user
func (h *HipchatService) SetStatus(userPresence hipchat.UpdateUserPresenceRequest) (*http.Response, error) {
	apiKey := os.Getenv("HIPCHAT_API_KEY")
	userID := os.Getenv("HIPCHAT_USER_ID")
	client := hipchat.NewClient(apiKey)

	user, _, err := client.User.View(userID)
	if err != nil {
		fmt.Println("\n", err)
		return nil, err
	}

	userRequest := hipchat.UpdateUserRequest{
		Name:        user.Name,
		MentionName: user.MentionName,
		Email:       user.Email,
		Presence:    userPresence,
	}
	resp, err := client.User.Update(userID, &userRequest)
	if err != nil {
		fmt.Println("\n", err)
		return nil, err
	}

	if resp.StatusCode == 204 {
		fmt.Println("\nHipChat Status: ", userPresence.Status)
	}

	return resp, nil
}
