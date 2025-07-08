package variables

import (
	"gohook/core"
	"gohook/dsl"
	"gohook/dsl/variables/git_vars"
	"gohook/dsl/variables/utils_vars"
)

func ParseVariables(context *dsl.ModeContext) core.VariablesFunc {
	var GoHookVariables = core.VariablesFunc{
		"$TIME_NOW":         utils_vars.TimeNow(),
		"$USER_HOME":        utils_vars.UserHome(),
		"$USER_OS":          utils_vars.UserOS(),
		"$USER_HOSTNAME":    utils_vars.HostName(),
		"$LAST_COMMIT_HASH": git_vars.GetLastCommitHash(context),
		"$GIT_BRANCH":       git_vars.GetBranchName(context),
	}

	return GoHookVariables
}
