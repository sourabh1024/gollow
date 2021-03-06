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

// Package data provides method required for loading data from data sources
package data

// Entity is the base interface for all data objects
// All data objects must implement all the methods
// All the methods required to fetch from any data source should be here
// Should this be divided up based on db ?
type Entity interface {

	// GetPrimaryKey returns the primaryKey for the given struct
	// it must be unique and collision would lead to unexpected data
	GetPrimaryKey() string

	// NewEntity returns a new object of entity
	// used for storing the result after mysql
	NewEntity() Entity
}
