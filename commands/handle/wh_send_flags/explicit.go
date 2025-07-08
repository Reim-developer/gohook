package wh_send_flags

import (
	"gohook/core"
	"gohook/core/discord_api"
	"gohook/utils"
	"os"
)

type explicitContext struct {
	enableExplicit bool
	envUrlName     string
	config         *core.DiscordWebhookConfig
}

func NewExplicit(enable bool, envUrl string, config *core.DiscordWebhookConfig) *explicitContext {
	explicitContext := explicitContext{
		enableExplicit: enable,
		envUrlName:     envUrl,
		config:         config,
	}

	return &explicitContext
}

func (context *explicitContext) HandleExplicitMode(payload *core.DiscordWebhook) {
	if context.enableExplicit {
		var webhookEnv string
		var useEnv = false

		if val := os.Getenv(context.envUrlName); val != "" {
			webhookEnv = val
			useEnv = true
		} else {
			webhookEnv = *context.config.Webhook.URL
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
			utils.InfoShow("This action use environment: %s", context.envUrlName)
		}
	}
}
