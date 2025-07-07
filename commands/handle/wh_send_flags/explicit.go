package wh_send_flags

import (
	"gohook/core"
	"gohook/core/discord_api"
	"gohook/utils"
	"os"
)

type ExplicitContext struct {
	EnableExplicit bool
	EnvUrlName     string
	Config         *core.DiscordWebhookConfig
}

func (context *ExplicitContext) HandleExplicitMode(payload *core.DiscordWebhook) {
	if context.EnableExplicit {
		var webhookEnv string
		var useEnv = false

		if val := os.Getenv(context.EnvUrlName); val != "" {
			webhookEnv = val
			useEnv = true
		} else {
			webhookEnv = *context.Config.Webhook.URL
			useEnv = false
		}

		var result, err = discord_api.ExplicitSendWebhook(&webhookEnv, payload)
		if err != nil {
			utils.CriticalShow("Could not send webhook with error: %s", err)
			os.Exit(core.WebhookSendFailed)
		}

		utils.InfoShow("Use Explicit Mode:")
		utils.InfoShow("Successfully send webhook")
		utils.InfoShow("Message ID: %s", result.MessageID)
		utils.InfoShow("Channel ID: %s", result.ChannelID)

		if useEnv {
			utils.InfoShow("This action use environment: %s", context.EnvUrlName)
		}
	}
}
