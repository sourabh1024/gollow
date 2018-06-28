package data

/*
	Entity is the base interface for all data objects
	All data objects must implement all the methods
*/
type Entity interface {

	// returns the namespace to which this entity belongs
	GetNameSpace() string

	// returns the name of the data source
	// should be unique in a namespace
	GetDataName() string

	// returns the primaryKey for the given struct
	// it must be unique and collision would lead to unexpected data
	GetPrimaryKey() string

	// NewEntity creates an instance of the entity
	NewEntity() Entity
}
