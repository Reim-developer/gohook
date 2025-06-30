package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type DiscordWebhook struct {
	Content  string `json:"content,omitempty"`
	Username string `json:"username,omitempty"`
	Avatar   string `json:"avatar_url,omitempty"`
}

func SendWebhook(URL *string, discord_webhook *DiscordWebhook) error {
	payload, _ := json.Marshal(discord_webhook)

	if URL == nil || *URL == "" {
		return errors.New("invalid Webhook URL. Please make sure your webhook URL in TOML setting is valid")
	}

	response, err := http.Post(*URL, "application/json", bytes.NewBuffer(payload))

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 204 {
		return fmt.Errorf("unexpected response status: %d", response.StatusCode)
	}

	return nil
}
