package logger

import "github.com/pterm/pterm"

var Verbose bool

func NewLogger() *pterm.Logger {
	if Verbose {
		return pterm.DefaultLogger.WithLevel(pterm.LogLevelDebug).WithCaller()
	}

	return pterm.DefaultLogger.WithLevel(pterm.LogLevelInfo)
}
