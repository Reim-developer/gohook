package utils_vars

import (
	"gohook/core"
	"os/user"
)

func UserHome() core.StrFunc {
	var function = func() string {
		user, err := user.Current()

		if err != nil {
			return "Unknown User Home"
		}

		return user.HomeDir
	}

	return function
}
