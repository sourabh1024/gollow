package cache

import (
	"gollow/cdd/sources"
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

//GetRegisteredModels() returns you the map of RegisteredModels to cache
//which are populated with this model data
func GetRegisteredModels() map[sources.DataModel]GollowCache {
	return models.modelsList
}
