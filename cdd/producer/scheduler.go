package producer

import (
	"gollow/cdd/logging"
	"time"
)

// ScheduleProducers schedules the Producer for being produced
// Fetches the list of Registered Models in Producer
// launches a separate go routine for every dataModel being produced
func ScheduleProducers() {
	models := GetRegisteredModels()
	for model := range models {
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
