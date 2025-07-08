package wh_send_flags

import (
	"bytes"
	"encoding/json"
	"gohook/core"
	"gohook/core/status_code"
	"gohook/utils"
	"os"
)

type dryRunContext struct {
	enableMode bool
}

func NewDryRun(enableMode bool) *dryRunContext {
	dryRun := dryRunContext{
		enableMode: enableMode,
	}

	return &dryRun
}

func (context *dryRunContext) HandleDryRun(payload *core.DiscordWebhook) {
	var enableMode = context.enableMode

	if enableMode {
		var buffer bytes.Buffer
		var encoder = json.NewEncoder(&buffer)

		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", " ")

		err := encoder.Encode(payload)
		if err != nil {

			utils.CriticalShow("Could not decode JSON: %s", err)
			os.Exit(status_code.JsonDecodeError)
		}

		utils.InfoShow("Running in Dry Mode:")
		utils.InfoShow("Your webhook payload:\n%s", buffer.String())
	}
}
