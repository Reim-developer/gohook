package helper

import (
	"gohook/core/status_code"
	"gohook/utils"
	"os"
)

type webhookContext struct {
	url *string
}

func NewWebhookUrl(url *string) *webhookContext {
	var webhook = webhookContext{
		url: url,
	}

	return &webhook
}

func (contextWebhook *webhookContext) TryHandleNil() {
	if contextWebhook.url == nil {
		utils.CriticalShow(`Could not find the field value 'url' in your settings. 
[HINT]
If you use enviroment variable, please leave blank, like: 
[webhook]
url = ""`)

		os.Exit(status_code.FieldNotProvided)
	}
}
