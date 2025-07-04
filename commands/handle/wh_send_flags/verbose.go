package whsendflags

import (
	"bytes"
	"encoding/json"
	"gohook/core"
	"os"
)

type VerboseContext struct {
	EnableVerbose bool
	EnableDryRun  bool
}

func (context *VerboseContext) HandleVerbose(payload *core.DiscordWebhook) {
	if context.EnableVerbose && !context.EnableDryRun {
		var buffer bytes.Buffer
		var encoder = json.NewEncoder(&buffer)

		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", " ")

		err := encoder.Encode(payload)
		if err != nil {

			core.CriticalShow("Could not decode JSON: %s", err)
			os.Exit(core.JsonDecodeError)
		}

		core.InfoShow("Your webhook payload:\n%s", buffer.String())
	}
}
