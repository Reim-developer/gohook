package wh_send_flags

import (
	"gohook/commands/handle/wh_send_flags/helper"
	"gohook/core"
	"gohook/core/discord_api"
	"gohook/utils"
	"time"
)

type LoopSendContext struct {
	DryMode      bool
	ExplicitMode bool
	LoopCount    int
	DelayTime    int
	EnvUrlName   string
	Config       *core.DiscordWebhookConfig
}

func NewLoopSend(dryMode bool, explicitMode bool,
	loopCount int, envUrlName string, delayTime int,
	config *core.DiscordWebhookConfig) *LoopSendContext {

	loopSendContext := LoopSendContext{
		DryMode:      dryMode,
		ExplicitMode: explicitMode,
		LoopCount:    loopCount,
		DelayTime:    delayTime,
		EnvUrlName:   envUrlName,
		Config:       config,
	}

	return &loopSendContext
}

func (context *LoopSendContext) HandleLoopSend(payload *core.DiscordWebhook) {
	/*
	* [!] Only run if
	* [x] Loop count > 1.
	* [x] Explicit mode is not enabled.
	* [x] Dry mode is not enabled.
	 */
	if context.LoopCount > 1 && !context.ExplicitMode && !context.DryMode {
		var successCount = 0
		var failedCount = 0
		var envName = context.EnvUrlName
		var fallbackDefault = context.Config.Webhook.URL

		helper.NewWebhookUrl(fallbackDefault).TryHandleNil()
		webhookUrl, usedEnv := helper.NewEnvironment(envName, *fallbackDefault).TryGetEnv()

		for index := range context.LoopCount {

			err := discord_api.SendWebhook(&webhookUrl, payload)

			if err != nil {
				utils.CriticalShow("Send webhook failed (%d) time(s) %s\n", index, err)

				failedCount += 1
				time.Sleep(time.Duration(context.DelayTime) * time.Second)
				continue
			}

			index = index + 1
			utils.InfoShow("Send webhook success (%d) time(s), delay time: %d", index, context.DelayTime)

			successCount += 1
			time.Sleep(time.Duration(context.DelayTime) * time.Second)
		}

		utils.InfoShow("Success count: %d time(s)", successCount)
		utils.InfoShow("Failed count: %d time(s)", failedCount)

		if usedEnv {
			utils.InfoShow("This action use environment: %s", context.EnvUrlName)
		}

	}
}
