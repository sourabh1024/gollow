package snapshot

import "gollow/cdd/sources"

//Snapshot represents the snapshot being produced
//Snapshot is the whole data image being produced
type Snapshot interface {

	//Load loads the snapshot of given model type into Model Bag from the given storage and file
	Load(model sources.DataModel) (sources.Bag, error)

	//Save saves the Model Bag into the given storage and file name
	Save(sources.Bag) (int, error)
}
