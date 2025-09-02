// functions that logger relies on, so we can't put them in common
package logger

import (
	"fmt"
	"os"
	"runtime/debug"
)

// Cfg.LoggerFilePath == "" at the point of logger initialization, so
// it will just print without saving to log file
func CreateDirIfDoesntExist(path string) (err error, errMsg string) {
	Log(Info, BlueColor, "%s dir: '%s'", "Creating", path)
	if path == "" {
		Log(Verbose2, DimBlueColor, "%s", "Dir is an empty string, not creating")
		return nil, ""
	}
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		Log(Verbose, DimCyanColor, "Dir '%s' doesn't exist", path)
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err, fmt.Sprintf("Unable to create dir: %s", path)
		}
	} else {
		Log(Info1, PurpleColor, "Dir '%s' already exists, not creating", path)
		return nil, ""
	}

	Log(Info1, GreenColor, "%s dir: '%s'", "Creating", path)

	return nil, ""
}

// Exit if error is encountered, otherwise do nothing
// this version relies on logger
func QuitIfError(err error, errMsg string) {
	if err != nil {
		Log(Critical, BoldRedColor, "%s", errMsg)
		Log(Critical1, BoldRedColor, "%s", err.Error())
		Log(Critical2, DimRedColor, "%s", string(debug.Stack()))
		os.Exit(1)
	}
}

// Exit if error is encountered, otherwise do nothing
// this version does not rely on logger itself
func QuitIfErrorLoggerIndependent(err error, errMsg string) {
	if err != nil {
		fmt.Printf("\033[0;31m") // bold red
		fmt.Printf("%s\n", errMsg)
		fmt.Printf("%s\n", err.Error())
		fmt.Printf("\033[0m")
		fmt.Printf("\033[2;31m") // dim red
		debug.PrintStack()
		fmt.Printf("\033[0m")
		os.Exit(1)
	}
}

// if error is encountered - print err.Error() and errMsg with Error and Error1 log levels.
// Don't exit though.
// Those errors are to be fixed but most of the tme should not stop the program.
func PrintError(err error, errMsg string) {
	if err != nil {
		Log(Error, RedColor, "%s", errMsg)
		Log(Error1, BrightRedColor, "%s", err.Error())
	}
}

// if error is encountered - print err.Error() and errMsg with Warning and Warning1 log levels.
// Don't exit though.
// Those errors are often don't need to be fixed and most of the tme should not stop the program.
func PrintWarning(err error, errMsg string) {
	if err != nil {
		Log(Warning, YellowColor, "%s", errMsg)
		Log(Warning1, BrightYellowColor, "%s", err.Error())
	}
}

func QuitIfWarning(err error, errMsg string) {
	if err != nil {
		Log(Warning, BoldYellowColor, "%s", errMsg)
		Log(Warning1, BoldYellowColor, "%s", err.Error())
		os.Exit(0)
	}
}

func QuitIfErrorWithContext(err error, errMsg string, context interface{}) {
	if err != nil {
		Log(Critical, BoldRedColor, "ErrMsg: '%s'", errMsg)
		Log(Critical1, BoldRedColor, "Err: '%s'", err.Error())
		Log(Critical2, DimRedColor, "Debug stack:\n```\n%s\n```", string(debug.Stack()))
		Log(Critical3, DimRedColor, "Context:\n```\n%s\n```", context)
		os.Exit(1)
	}
}
