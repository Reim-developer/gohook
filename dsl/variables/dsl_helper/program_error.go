package dsl_helper

import (
	"gohook/core"
	"gohook/dsl"
	"gohook/utils"
)

type programErrorContext struct {
	withError   error
	contextMode *dsl.ModeContext
	programName string
}

func NewProgramError(err error, contextMode *dsl.ModeContext, programName string) *programErrorContext {
	context := programErrorContext{
		withError:   err,
		contextMode: contextMode,
		programName: programName,
	}

	return &context
}

func (programContext *programErrorContext) HandleProgramError() {
	var err = programContext.withError
	var isStrictMode = programContext.contextMode.StrictMode
	var programName = programContext.programName

	if err != nil {
		if isStrictMode {
			var statusCode = core.RunProgramFailed

			// [!] Exit now. Don't return anything if strict mode is enabled.
			utils.FatalNow(statusCode, "Could not run program '%s' with error: %s.", programName, err)
		}

		utils.CriticalShow("Could not run program: '%s'.", programName)
		utils.CriticalShow("Details error: %s.", err)
		utils.InfoShow("Strict mode is disabled. GoHook will return empty string to your variable.")
	}
}
