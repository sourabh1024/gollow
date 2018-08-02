package producer

import "gollow/cdd/sources"

var (
	models modelsImpl
)

type modelsImpl struct {
	modelsList map[sources.DataModel]struct{}
}

func init() {
	models.modelsList = make(map[sources.DataModel]struct{})
}

func Register(model sources.DataModel, val struct{}) {
	models.modelsList[model] = val
}

func GetRegisteredModels() map[sources.DataModel]struct{} {
	return models.modelsList
}
