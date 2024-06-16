package utils

import (
	"events/internal/config"
	logger "log"
)

var isDevelopment = config.Envs.Env == "development"

/* Logs information in the console, works only in development environment */
func Log(message ...any) {
	if isDevelopment {
		logger.Println(message...)
	}
}

/* Logs information in the console with an specified format, works only in development environment */
func Logf(format string, message ...any) {
	if isDevelopment {
		logger.Printf(format, message...)
	}
}

