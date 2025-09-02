package util

import (
	"encoding/json"
	"fmt"
	"os"

	"emv/src/pkg/logger"
	// "lib/logger"
)

/*
Load a config file located at filePath
Pass a pointer to whatever config variable you need.
Make sure that file matches the config type.

It's ok to use logger before loading it's config. It still can print but won't save log line to a file

logLevelOverride is used to print all messages (by default log level is 50 so only first message is printed)
pass -1 if you don't want to change logging level from it's default value
need this since config has no effect yet during LoadConfig
*/
func LoadConfig(filePath string, config any, logLevelOverride logger.LogLevel) (err error, errMsg string) {
	logger.Log(logger.Notice, logger.BlueColor, "%s config '%s'", "Loading", filePath)
	// override log level here first in order for common.LoadConfig to be able to print Info1 and above
	// within LoadConfig function
	if logLevelOverride != logger.DontOverride {
		logger.Log(logger.Cfg.LogLevel, logger.CyanColor, "Log level was switched from '%s' to '%s'", logger.Cfg.LogLevel, logLevelOverride)
		logger.Cfg.LogLevel = logLevelOverride
	}

	byteValue, err := os.ReadFile(filePath)
	if err != nil {
		return err, fmt.Sprintf("Unable to read file: '%s'", filePath)
	}

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return err, errMsg
	}

	logger.Log(logger.Notice1, logger.GreenColor, "%s config '%s'", "Loaded", filePath)
	return err, errMsg
}
