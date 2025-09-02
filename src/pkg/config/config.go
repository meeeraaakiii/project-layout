// this package exists to hold Cfg global var
// can also change default values here for this project
package config

import (
	"os"

	"emv/src/pkg/util"
	"emv/src/pkg/logger"
)

type Config struct {
	// parts present in configuration file (some of the parameters are generated during initilization process)
	Logger   *logger.Config  `json:"logger"`

	// those parametrs are initialized during InitializeConfig()
	CallerProgramName string `json:"caller_program_name,omitempty"`
}

var Cfg Config


func GetDefaultConfig() Config {
	return Config{

	}
}

// this function will quit when encountering error
// make sure to run common.LoadConfig first here since it's the earliest point we overriding log level from
func InitializeConfig(configPath string, logLevelOverride logger.LogLevel, logDirOverride string) {
	if util.FileExists(configPath) {
		logger.QuitIfError(util.LoadConfig(configPath, &Cfg, logLevelOverride))
		// if certain part of the config is present in this project config file - override default config
		if Cfg.Logger != nil {
			logger.QuitIfError(logger.InitializeLogger(*Cfg.Logger, logLevelOverride, logDirOverride))
		}
	} else {
		if configPath == "" {
			logger.Log(logger.Notice, logger.BoldCyanColor, "%s path is '', using %s", "Config", "default config")
			Cfg = GetDefaultConfig()
		} else {
			logger.Log(logger.Notice, logger.BoldYellowColor, "%s path is '%s' but file %s, %s", "Config", configPath, "does not exist", "exiting...")
			os.Exit(0)
		}
	}
	Cfg.CallerProgramName = util.GetCallerProgramNamePanicWrapper(4)
}
