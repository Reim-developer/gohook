package whsendflags

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gohook/core"
	"os"
)

type ToJsonContext struct {
	IsEnableMode bool
}

func (context *ToJsonContext) HandleExportToJson(payload *core.DiscordWebhook) {
	var buffer bytes.Buffer
	var encoder = json.NewEncoder(&buffer)

	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", " ")

	err := encoder.Encode(payload)
	if err != nil {

		core.CriticalShow("Could not decode JSON: %s", err)
		os.Exit(core.JsonDecodeError)
	}

	var timeNow = core.GetTimeNow()
	var filePath = fmt.Sprintf("%s.json", timeNow)
	var contentBytes = buffer.Bytes()

	write_err := core.WriteTo(filePath, contentBytes)
	if write_err != nil {
		core.CriticalShow("Export to JSON FAILED with error: %s", write_err)
		os.Exit(core.WriteJsonFailed)
	}

	core.InfoShow("Successfully export your payload to: %s", filePath)
}
