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
