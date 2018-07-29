package producer

import "gollow/cdd/sources"

var (
	models modelsImpl
)

type modelsImpl struct {
	modelsList map[sources.ProtoDataModel]struct{}
}

func init() {
	models.modelsList = make(map[sources.ProtoDataModel]struct{})
}

func Register(model sources.ProtoDataModel, val struct{}) {
	models.modelsList[model] = val
}

func GetRegisteredModels() map[sources.ProtoDataModel]struct{} {
	return models.modelsList
}
