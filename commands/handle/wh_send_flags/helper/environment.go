package helper

import "os"

const (
	usedEnv   = true
	notUseEnv = false
)

type envroment struct {
	envName         string
	defaultFallback string
}

func NewEnvironment(variableEnv string, defaultFallback string) *envroment {
	var env = envroment{
		envName:         variableEnv,
		defaultFallback: defaultFallback,
	}

	return &env
}

func (env *envroment) TryGetEnv() (string, bool) {
	var envVar = os.Getenv(env.envName)
	var webhookURL string

	if envVar != "" {
		webhookURL = envVar

		return webhookURL, usedEnv
	}

	webhookURL = env.defaultFallback
	return webhookURL, notUseEnv

}
