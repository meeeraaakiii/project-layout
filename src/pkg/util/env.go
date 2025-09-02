package util

import (
	"os"

	"emv/src/pkg/logger"
)

// Check if all required environment variables are present. If not all present - print warinig and quit.
func CheckIfEnvVarsPresent(listOfEnvVars []string) {
	for _, envVarName := range listOfEnvVars {
		if os.Getenv(envVarName) == "" {
			logger.Log(logger.Warning, logger.YellowColor, "Env var. '%s' is not set. %s", envVarName, "Check your environment file")
			os.Exit(0)
		}
	}
}
