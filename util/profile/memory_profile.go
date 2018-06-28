package profile

import (
	"gollow/logging"
	"runtime"
)

func GetMemoryProfile() {

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	logging.GetLogger().Info("------------------------------------------")
	logging.GetLogger().Info("mem alloc : ", mem.Alloc)
	logging.GetLogger().Info("mem total alloc : ", mem.TotalAlloc)
	logging.GetLogger().Info("mem heap alloc : ", mem.HeapAlloc)
	logging.GetLogger().Info("mem heap size", mem.HeapSys)
	logging.GetLogger().Info("------------------------------------------")

}
