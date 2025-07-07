package dsl_tests

import (
	"fmt"
	"gohook/dsl/variables/git_vars"
	"os/exec"
	"testing"
)

func WrongCommand(programName string, args ...string) (string, error) {
	command := exec.Command(programName, args...)

	result, err := command.Output()

	if err != nil {
		/*
		* For better error handling,
		* we will return "" (empty string), and warning user.
		* I'll create a flag '--strict',
		* for stop program now, if CriticalShow,  (stderr).
		 */
		return "", fmt.Errorf("%s", err)
	}

	return string(result), nil
}

func TestWrongCommand(t *testing.T) {
	result, err := WrongCommand("wrong_command", "--wrong-flag")
	if err != nil {
		t.Log(err)
	}

	t.Log(result)
}

func TestGitDsl(t *testing.T) {
	var lastGitCommit = git_vars.GetLastCommitHash()

	t.Log(lastGitCommit())
}
