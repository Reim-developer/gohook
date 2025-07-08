package discord_api_error

import (
	"gohook/core/status_code"
	"gohook/core/success"
	"gohook/utils"
	"os"
)

type ApiDiscordError struct {
	Err error
}

func MapError(err error) *ApiDiscordError {
	var apiError = ApiDiscordError{
		Err: err,
	}

	return &apiError
}

func (errContext *ApiDiscordError) Try() *success.ThenSuccess {
	if errContext.Err != nil {
		utils.CriticalShow("Could not send webhook with error: %s\n", errContext.Err)
		os.Exit(status_code.WebhookSendFailed)
	}

	var successContext = success.ThenSuccess{
		IsError: nil,
	}

	return &successContext
}
