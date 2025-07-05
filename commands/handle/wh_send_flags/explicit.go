package whsendflags

import (
	"gohook/core"
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

		var result, err = core.ExplicitSendWebhook(&webhookEnv, payload)
		if err != nil {
			core.CriticalShow("Could not send webhook with error: %s", err)
			os.Exit(core.WebhookSendFailed)
		}

		core.InfoShow("Use Explicit Mode:")
		core.InfoShow("Successfully send webhook")
		core.InfoShow("Message ID: %s", result.MessageID)
		core.InfoShow("Channel ID: %s", result.ChannelID)

		if useEnv {
			core.InfoShow("This action use environment: %s", context.EnvUrlName)
		}
	}
}
