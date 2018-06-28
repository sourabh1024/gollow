package profile

import (
	"gollow/logging"
	"time"
)

func Duration(invocation time.Time, name string) {
	elapsed := time.Since(invocation)
	logging.GetLogger().Info("%s lasted %s", name, elapsed)
}
