package utils_vars

import (
	"gohook/core"
	"runtime"
	"strings"
)

func UserOS() core.StrFunc {
	var function = func() string {
		os := runtime.GOOS

		return strings.ToUpper(os[:1]) + strings.ToLower(os[1:])
	}

	return function
}
