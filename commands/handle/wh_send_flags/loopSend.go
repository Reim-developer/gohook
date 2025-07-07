package wh_send_flags

import (
	"gohook/core"
	"gohook/core/discord_api"
	"gohook/utils"
	"os"
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
		var webhookEnv string
		var useEnv = false

		if val := os.Getenv(context.EnvUrlName); val != "" {
			webhookEnv = val
			useEnv = true
		} else {
			webhookEnv = *context.Config.Webhook.URL
			useEnv = false
		}

		for index := range context.LoopCount {

			err := discord_api.SendWebhook(&webhookEnv, payload)

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

		if useEnv {
			utils.InfoShow("This action use environment: %s", context.EnvUrlName)
		}

	}
}
