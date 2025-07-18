package wh_send_flags

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gohook/core"
	"gohook/core/status_code"
	"gohook/utils"
	"os"
)

type ToJsonContext struct {
	IsEnableMode bool
}

func NewToJson(enableMode bool) *ToJsonContext {
	toJsonContext := ToJsonContext{
		IsEnableMode: enableMode,
	}

	return &toJsonContext
}

func (context *ToJsonContext) HandleExportToJson(payload *core.DiscordWebhook) {
	if context.IsEnableMode {
		var buffer bytes.Buffer
		var encoder = json.NewEncoder(&buffer)

		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", " ")

		err := encoder.Encode(payload)
		if err != nil {

			utils.CriticalShow("Could not decode JSON: %s", err)
			os.Exit(status_code.JsonDecodeError)
		}

		var timeNow = utils.GetTimeNow()
		var filePath = fmt.Sprintf("%s.json", timeNow)
		var contentBytes = buffer.Bytes()

		write_err := utils.WriteTo(filePath, contentBytes)
		if write_err != nil {
			utils.CriticalShow("Export to JSON FAILED with error: %s", write_err)
			os.Exit(status_code.WriteJsonFailed)
		}

		utils.InfoShow("Successfully export your payload to: %s", filePath)
	}
}
