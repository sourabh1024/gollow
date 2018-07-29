package util

import (
	"fmt"
	"gollow/cdd/logging"
	"time"
)

func Duration(invocation time.Time, name string) {
	elapsed := time.Since(invocation)
	logging.GetLogger().Info("%s lasted %s", name, elapsed)
}

func GetCurrentTimeString() string {
	return fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond))
}
