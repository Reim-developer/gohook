package discord_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gohook/core/discord_api_error"
	"net/http"
)

// Embed struct to Discord Webhook.
// Example JSON:
//
//	{
//	  "content": "Hello World",
//	  "username": "Kaxtr",
//	  "avatar_url": "https://example.com/image.png",
//	  "embeds": [
//	    {
//	      "title": "Test",
//	      "description": "Test Description",
//	      "color": 16777215,
//	      "footer": {
//	        "text": "test"
//	      },
//	      "image": {
//	        "url": "https://example.com/image.png"
//	      },
//	      "thumbnail": {
//	        "url": "https://example.com/image.png"
//	      }
//	    }
//	  ]
//	}
type Embed struct {
	Title       string          `json:"title,omitempty"`
	Description string          `json:"description,omitempty"`
	Color       int             `json:"color,omitempty"`
	Footer      *EmbedFooter    `json:"footer,omitempty"`
	Image       *EmbedImage     `json:"image,omitempty"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"`
}

type EmbedFooter struct {
	Text string `json:"text"`
}

type EmbedImage struct {
	URL string `json:"url"`
}

type EmbedThumbnail struct {
	URL string `json:"url"`
}

type DiscordWebhook struct {
	Content  string  `json:"content,omitempty"`
	Username string  `json:"username,omitempty"`
	Avatar   string  `json:"avatar_url,omitempty"`
	Embeds   []Embed `json:"embeds,omitempty"`
}

type WebhookResponse struct {
	MessageID string `json:"id"`
	ChannelID string `json:"channel_id"`
}

func NewDiscordWebhook(content string, username string,
	avatarUrl string, embeds []Embed) *DiscordWebhook {

	var discordWebhook = DiscordWebhook{
		Content:  content,
		Username: username,
		Avatar:   avatarUrl,
		Embeds:   embeds,
	}

	return &discordWebhook
}

func (webhook *DiscordWebhook) ExplicitWebhookSend(URL *string) (*WebhookResponse, error) {
	if URL == nil || *URL == "" {
		return nil, errors.New("invalid Webhook URL. Please make sure your webhook URL in TOML setting is valid")
	}

	payload, marshalErr := json.Marshal(webhook)
	if marshalErr != nil {
		return nil, marshalErr
	}

	response, response_err := http.Post(*URL+"?wait=true", "application/json", bytes.NewBuffer(payload))
	if response_err != nil {
		return nil, response_err
	}
	defer response.Body.Close()

	var webhookResponse WebhookResponse
	decodeError := json.NewDecoder(response.Body).Decode(&webhookResponse)
	if decodeError != nil {
		return nil, decodeError
	}

	return &webhookResponse, nil
}

type WhenError = discord_api_error.ApiDiscordError

func SendWebhook(URL *string, discord_webhook *DiscordWebhook) *WhenError {
	if URL == nil || *URL == "" {
		var err = errors.New("invalid Webhook URL. Please make sure your webhook URL in TOML setting is valid")
		var errConext = discord_api_error.MapError(err)

		return errConext
	}

	payload, _ := json.Marshal(discord_webhook)
	response, err := http.Post(*URL, "application/json", bytes.NewBuffer(payload))

	if err != nil {
		var errConext = discord_api_error.MapError(err)

		return errConext
	}

	defer response.Body.Close()

	if response.StatusCode != 204 {
		var err = fmt.Errorf("unexpected response status: %d", response.StatusCode)

		var errConext = discord_api_error.MapError(err)
		return errConext
	}

	// [!] If error is nothing, do nothing
	var errConext = discord_api_error.MapError(nil)

	return errConext
}
