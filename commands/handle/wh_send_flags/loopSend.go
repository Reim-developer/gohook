package whsendflags

import (
	"gohook/core"
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

			err := core.SendWebhook(&webhookEnv, payload)

			if err != nil {
				core.CriticalShow("Send webhook failed (%d) time(s) %s\n", index, err)

				failedCount += 1
				time.Sleep(time.Duration(context.DelayTime) * time.Second)
				continue
			}
			core.InfoShow("Send webhook success (%d) time(s), delay time: %d", index+1, context.DelayTime)

			successCount += 1
			time.Sleep(time.Duration(context.DelayTime) * time.Second)
		}

		core.InfoShow("Success count: %d time(s)", successCount)
		core.InfoShow("Failed count: %d time(s)", failedCount)
		if useEnv {
			core.InfoShow("This action use environment: %s", context.EnvUrlName)
		}

	}
}
