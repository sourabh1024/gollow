package logging

import "github.com/op/go-logging"

var log *logging.Logger

func init() {
	log = logging.MustGetLogger("gollow")
}

func GetLogger() *logging.Logger {
	return log
}
