package data


// Entity is the base interface for all data objects
type Entity interface {

	// returns the id of the data
	GetID() string

	// sets the ID
	SetID(string)

	// NewEntity creates an instance of the entity
	NewEntity() Entity
}

