package cache

import (
	"gollow/sources"
)

var (
	models modelsImpl
)

type modelsImpl struct {
	modelsList map[sources.ProtoDataModel]GollowCache
}

func init() {
	models.modelsList = make(map[sources.ProtoDataModel]GollowCache)
}

func Register(model sources.ProtoDataModel, cache GollowCache) {
	models.modelsList[model] = cache
}

func GetRegisteredModels() map[sources.ProtoDataModel]GollowCache {
	return models.modelsList
}
