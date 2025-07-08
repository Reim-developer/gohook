package wh_send_flags

import (
	"gohook/commands/handle/wh_send_flags/helper"
	"gohook/core"
	"gohook/core/discord_api"
	"gohook/core/status_code"
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
		var envName = context.envUrlName
		var fallbackDefault = context.config.Webhook.URL

		helper.NewWebhookUrl(fallbackDefault).TryHandleNil()
		webhookURL, usedEnv := helper.NewEnvironment(envName, *fallbackDefault).TryGetEnv()

		var discordWebhook = discord_api.NewDiscordWebhook(payload.Content, payload.Username, payload.Avatar, payload.Embeds)
		var result, err = discordWebhook.ExplicitWebhookSend(&webhookURL)

		if err != nil {
			utils.CriticalShow("Could not send webhook with error: %s", err)
			os.Exit(status_code.WebhookSendFailed)
		}

		utils.InfoShow("Use Explicit Mode:")
		utils.InfoShow("Successfully send webhook")
		utils.InfoShow("Message ID: %s", result.MessageID)
		utils.InfoShow("Channel ID: %s", result.ChannelID)

		if usedEnv {
			utils.InfoShow("This action use environment: %s", context.envUrlName)
		}
	}
}
