package wh_send_flags

import (
	"gohook/commands/handle/wh_send_flags/helper"
	"gohook/core"
	"gohook/core/discord_api"
	"gohook/utils"
)

type webhookSendContext struct {
	isDryMode      bool
	isExplicitMode bool
	envUrlName     string
	loopCount      int
	configToml     *core.DiscordWebhookConfig
}

func NewWebhookSendOnce(
	isDryMode bool, isExplicitMode bool,
	envUrl string, loopCount int, configToml *core.DiscordWebhookConfig) *webhookSendContext {

	webhookSend := webhookSendContext{
		isDryMode: isDryMode, isExplicitMode: isExplicitMode,
		envUrlName: envUrl, loopCount: loopCount,
		configToml: configToml,
	}

	return &webhookSend
}

func (context *webhookSendContext) HandleWebhookSendOnce(payload *core.DiscordWebhook) {
	if !context.isDryMode && !context.isExplicitMode && context.loopCount == 1 {
		var envName = context.envUrlName
		var fallbackDefault = context.configToml.Webhook.URL

		helper.NewWebhookUrl(fallbackDefault).TryHandleNil()
		webhookURL, usedEnv := helper.NewEnvironment(envName, *fallbackDefault).TryGetEnv()

		var err = discord_api.SendWebhook(&webhookURL, payload)
		if err != nil {
			utils.CriticalShow("Could not send webhook with error: %s", err)
			return
		}

		if usedEnv {
			utils.InfoShow("This action use environment: %s", context.envUrlName)
		}

		utils.InfoShow("Successfully send webhook")
	}
}
