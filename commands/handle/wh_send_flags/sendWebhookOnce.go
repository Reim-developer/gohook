package wh_send_flags

import (
	"gohook/core"
	"gohook/core/discord_api"
	"gohook/utils"
	"os"
)

type webhookSendContext struct {
	isDryMode      bool
	isExplicitMode bool
	envURL         string
	loopCount      int
	configToml     *core.DiscordWebhookConfig
}

func NewWebhookSendOnce(
	isDryMode bool, isExplicitMode bool,
	envUrl string, loopCount int, configToml *core.DiscordWebhookConfig) *webhookSendContext {

	webhookSend := webhookSendContext{
		isDryMode: isDryMode, isExplicitMode: isExplicitMode,
		envURL: envUrl, loopCount: loopCount,
		configToml: configToml,
	}

	return &webhookSend
}

func (context *webhookSendContext) HandleWebhookSendOnce(payload *core.DiscordWebhook) {
	var webhookEnv string
	var useEnv = false

	if !context.isDryMode && !context.isExplicitMode && context.loopCount == 1 {
		if val := os.Getenv(context.envURL); val != "" {
			webhookEnv = val
			useEnv = true
		} else {
			webhookEnv = *context.configToml.Webhook.URL
			useEnv = false
		}

		var err = discord_api.SendWebhook(&webhookEnv, payload)

		if err != nil {
			utils.CriticalShow("Critical error: %s\n", err)
			os.Exit(core.WebhookSendFailed)
		}

		utils.InfoShow("Successfully send webhook")
		if useEnv {
			utils.InfoShow("This action use environment: %s", context.configToml)
		}
	}
}
