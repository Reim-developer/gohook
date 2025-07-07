package git_vars

import (
	"gohook/core"
	"gohook/dsl"
	"gohook/utils"
)

func GetLastCommitHash(context *dsl.ModeContext) core.StrFunc {
	var function = func() string {
		var programName = "git"

		result, err := utils.RunProgram(programName, "rev-parse", "HEAD")

		if err != nil {
			if context.StrictMode {
				var statusCode = core.RunProgramFailed

				// [!] Exit now. Don't return anything if strict mode is enabled.
				utils.FatalNow(statusCode, "Could not run program '%s' with error: %s.", programName, err)
			}

			utils.CriticalShow("Could not run program: '%s'.", programName)
			utils.CriticalShow("Details error: %s.", err)
			utils.InfoShow("Strict mode is disabled. GoHook will return empty string to your variable.")
		}

		return string(result)
	}

	return function
}
