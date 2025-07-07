package wh_send_flags

import "gohook/core"

type WebhookSendFlagInterface interface {
	HandleDryRun(payload *core.DiscordWebhook)
	HandleWebhookSendOnce(payload *core.DiscordWebhook)
	HandleVerbose(payload *core.DiscordWebhook)
	HandleExplicit(payload *core.DiscordWebhook)
	HandleLoopSend(payload *core.DiscordWebhook)
	HandleExportToJson(payload *core.DiscordWebhook)
}
