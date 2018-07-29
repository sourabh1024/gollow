package producer

import (
	"gollow/logging"
	"time"
)

func ScheduleProducers() {
	models := GetRegisteredModels()
	for model, _ := range models {
		logging.GetLogger().Info("Starting producer go routine for model : %s", model.GetDataName())
		ProduceModel(model)
		ticker := time.NewTicker(time.Duration(model.CacheDuration()))
		quit := make(chan struct{})
		go func() {
			for {
				select {
				case <-ticker.C:
					ProduceModel(model)
				case <-quit:
					ticker.Stop()
					return
				}
			}
		}()
	}
}
