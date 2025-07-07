package utils_vars

import (
	"gohook/core"
	"os"
)

func HostName() core.StrFunc {
	var function = func() string {
		hostname, err := os.Hostname()

		if err != nil {
			return "Unknown Host Name"
		}

		return hostname
	}

	return function
}
