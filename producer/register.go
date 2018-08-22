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
