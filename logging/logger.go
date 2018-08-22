// Package logging provides the logging methods
package logging

import "github.com/op/go-logging"

var log *logging.Logger

func init() {
	log = logging.MustGetLogger("gollow")
}

// GetLogger returns the logger
func GetLogger() *logging.Logger {
	return log
}
