//Copyright 2018 Sourabh Suman ( https://github.com/sourabh1024 )
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package producer

import (
	"gollow/logging"
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
