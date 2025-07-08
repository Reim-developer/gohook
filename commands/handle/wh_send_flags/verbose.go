package wh_send_flags

import (
	"bytes"
	"encoding/json"
	"gohook/core"
	"gohook/utils"
	"os"
)

type VerboseContext struct {
	EnableVerbose bool
	EnableDryRun  bool
}

func NewVerbose(enableVerbose bool, enableDryMode bool) *VerboseContext {
	verboseContext := VerboseContext{
		EnableVerbose: enableVerbose,
		EnableDryRun:  enableDryMode,
	}

	return &verboseContext
}

func (context *VerboseContext) HandleVerbose(payload *core.DiscordWebhook) {
	if context.EnableVerbose && !context.EnableDryRun {
		var buffer bytes.Buffer
		var encoder = json.NewEncoder(&buffer)

		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", " ")

		err := encoder.Encode(payload)
		if err != nil {

			utils.CriticalShow("Could not decode JSON: %s", err)
			os.Exit(core.JsonDecodeError)
		}

		utils.InfoShow("Your webhook payload:\n%s", buffer.String())
	}
}
