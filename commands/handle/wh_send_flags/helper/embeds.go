package helper

import (
	"gohook/core"
	"os"
	"unicode/utf8"
)

type MaxLenCheck struct {
	FieldName string
	Value     string
	MaxLen    int
	ExitCode  int
}

func AssertMaxLen(max *MaxLenCheck) {
	var length = utf8.RuneCountInString(max.Value)

	if length > max.MaxLen {
		core.CriticalShow("Embed [%s] too long (%d characters).", max.FieldName, length)
		core.InfoShow("Discord only allows up to %d characters.", max.MaxLen)
		os.Exit(max.ExitCode)
	}
}
