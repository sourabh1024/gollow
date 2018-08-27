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

package cache

import (
	"gollow/sources"
)

var (
	models modelsImpl
)

type modelsImpl struct {
	modelsList map[sources.DataModel]GollowCache
}

func init() {
	models.modelsList = make(map[sources.DataModel]GollowCache)
}

//Register registers the model to pull in cache
//model is the DataModel which is registered for client cache
//model must be registered on Data origination server to be produced else cache cant be build
func Register(model sources.DataModel, cache GollowCache) {
	models.modelsList[model] = cache
}

//GetRegisteredModels returns you the map of RegisteredModels to cache
//which are populated with this model data
func GetRegisteredModels() map[sources.DataModel]GollowCache {
	return models.modelsList
}
