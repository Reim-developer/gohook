package helper

import (
	"gohook/utils"
	"os"
	"unicode/utf8"
)

type AssertDiscordEmbed struct {
	ExitCode int
}

func NewAssertEmbed(exitCode int) *AssertDiscordEmbed {

	var assert = AssertDiscordEmbed{
		ExitCode: exitCode,
	}

	return &assert
}

func (assertContext *AssertDiscordEmbed) TryAssertEmbedLen(fieldName string, value string, maxChar int) *AssertDiscordEmbed {
	var length = utf8.RuneCountInString(value)

	if length > maxChar {
		utils.CriticalShow("Embed [%s] too long (%d characters).", fieldName, length)
		utils.InfoShow("Discord only allows up to %d characters.", maxChar)
		os.Exit(assertContext.ExitCode)
	}

	return assertContext
}
