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

import "gollow/sources"

var (
	models modelsImpl
)

type modelsImpl struct {
	modelsList map[sources.DataModel]struct{}
}

func init() {
	models.modelsList = make(map[sources.DataModel]struct{})
}

// Register registers the model for being produced
// Any DataModel which needs to be produced should be registered
func Register(model sources.DataModel, val struct{}) {
	models.modelsList[model] = val
}

// GetRegisteredModels returns the map of DataModel being Registered for production
func GetRegisteredModels() map[sources.DataModel]struct{} {
	return models.modelsList
}
