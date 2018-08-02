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

func Register(model sources.DataModel, cache GollowCache) {
	models.modelsList[model] = cache
}

func GetRegisteredModels() map[sources.DataModel]GollowCache {
	return models.modelsList
}
