// Package snapshot provides all the methods related to Snapshot and Version handling
package snapshot

import "gollow/sources"

// Snapshot represents the snapshot being produced
// Snapshot interface provides methods to load and
// save the snapshot with the initialised storage
// Storage is needed to create the snapshot object
type Snapshot interface {

	//Load loads the snapshot of given model type into Model Bag from the given storage and file
	Load(model sources.DataModel) (sources.Bag, error)

	//Save saves the Model Bag into the given storage and file name
	Save(sources.Bag) (int, error)
}
