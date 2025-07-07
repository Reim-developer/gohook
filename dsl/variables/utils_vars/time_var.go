package utils_vars

import (
	"gohook/core"
	"time"
)

func TimeNow() core.StrFunc {
	var function = func() string {
		return time.Now().Format("15:04:05")
	}

	return function
}
