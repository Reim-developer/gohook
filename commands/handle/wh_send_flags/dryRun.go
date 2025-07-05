package whsendflags

import (
	"bytes"
	"encoding/json"
	"gohook/core"
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

			core.CriticalShow("Could not decode JSON: %s", err)
			os.Exit(core.JsonDecodeError)
		}

		core.InfoShow("Running in Dry Mode:")
		core.InfoShow("Your webhook payload:\n%s", buffer.String())

		//HandleExportToJson(params, payload)
	}
}
