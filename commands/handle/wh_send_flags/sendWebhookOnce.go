package wh_send_flags

import (
	"gohook/core"
	"gohook/core/discord_api"
	"gohook/utils"
	"os"
)

type WebhookSendContext struct {
	IsDryMode      bool
	IsExplicitMode bool
	EnvURL         string
	LoopCount      int
	ConfigToml     *core.DiscordWebhookConfig
}

func (context *WebhookSendContext) HandleWebhookSendOnce(payload *core.DiscordWebhook) {
	var webhookEnv string
	var useEnv = false

	if !context.IsDryMode && !context.IsExplicitMode && context.LoopCount == 1 {
		if val := os.Getenv(context.EnvURL); val != "" {
			webhookEnv = val
			useEnv = true
		} else {
			webhookEnv = *context.ConfigToml.Webhook.URL
			useEnv = false
		}

		var err = discord_api.SendWebhook(&webhookEnv, payload)

		if err != nil {
			utils.CriticalShow("Critical error: %s\n", err)
			os.Exit(core.WebhookSendFailed)
		}

		utils.InfoShow("Successfully send webhook")
		if useEnv {
			utils.InfoShow("This action use environment: %s", context.ConfigToml)
		}
	}
}
