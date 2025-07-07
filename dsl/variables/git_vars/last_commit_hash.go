package git_vars

import (
	"gohook/core"
	"os/exec"
)

func GetLastCommitHash() core.StrFunc {
	var function = func() string {
		command := exec.Command("git", "rev-parse", "HEAD")

		result, err := command.Output()
		if err != nil {
			return "Unknown Commit Hash"
		}

		return string(result)
	}

	return function
}
