package core_test

import (
	"gohook/core"
	"log"
	"os"
	"testing"

	"github.com/BurntSushi/toml"
)

func TestUpdateWebhook(t *testing.T) {
	var config core.Config
	_, err := toml.DecodeFile("settings_test.toml", &config)
	if err != nil {
		log.Println("[Core_Test Package] Error:", err)
		os.Exit(1)
	}

	payload := core.DiscordWebhook{
		Content:  config.Message.Content,
		Username: config.Base.Username,
		Avatar:   config.Base.Avatar,
	}
	core.SendWebhook(config.Webhook.URL, &payload)
}
