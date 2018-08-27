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
