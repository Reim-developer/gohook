package wh_send_flags

import (
	"bytes"
	"encoding/json"
	"gohook/core"
	"gohook/utils"
	"os"
)

type DryRunContext struct {
	EnableMode bool
}

func (context *DryRunContext) HandleDryRun(payload *core.DiscordWebhook) {
	if context.EnableMode {
		var buffer bytes.Buffer
		var encoder = json.NewEncoder(&buffer)

		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", " ")

		err := encoder.Encode(payload)
		if err != nil {

			utils.CriticalShow("Could not decode JSON: %s", err)
			os.Exit(core.JsonDecodeError)
		}

		utils.InfoShow("Running in Dry Mode:")
		utils.InfoShow("Your webhook payload:\n%s", buffer.String())
	}
}
