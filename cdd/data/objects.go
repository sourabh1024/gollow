package data

/*
	Entity is the base interface for all data objects
	All data objects must implement all the methods
*/
type Entity interface {

	// returns the primaryKey for the given struct
	// it must be unique and collision would lead to unexpected data
	GetPrimaryKey() string

	NewEntity() Entity
}
