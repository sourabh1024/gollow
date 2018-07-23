package producer

import (
	"gollow/logging"
	"gollow/sources"
	"time"
)

/**
ProducerWorker is Worker method , it consumes the jobs from the channel
@param workerID : workerID of the worker
@param jobs : jobs channel in which all DataModel to be produced are pushed
@param results : result channel in which the results are pushed back
*/
func ProducerWorker(workerID int, jobs <-chan sources.DataModel, results chan<- interface{}) {

	//for j := range jobs {
	//
	//}
}

func ScheduleProducers() {
	models := GetRegisteredModels()
	for model, _ := range models {
		logging.GetLogger().Info("Starting go routine for model : %s", model.GetDataName())
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
