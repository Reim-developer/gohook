package core

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type DiscordWebhook struct {
	Content  string `json:"content,omitempty"`
	Username string `json:"username,omitempty"`
	Avatar   string `json:"avatar_url,omitempty"`
}

func SendWebhook(URL string, discord_webhook *DiscordWebhook) error {
	payload, _ := json.Marshal(discord_webhook)

	response, err := http.Post(URL, "application/json", bytes.NewBuffer(payload))

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 204 {
		return err
	}

	return nil
}
