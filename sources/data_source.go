package sources

import (
	"github.com/golang/protobuf/proto"
	"gollow/data"
)

//Message represents a single entity of the data
type Message interface {
	proto.Message

	//GetUniqueKey returns the primaryID of the string
	GetUniqueKey() string
}

//Bag interface represents a collection of message
type Bag interface {
	proto.Message

	//AddEntry provides method to add MESSAGE TO Bag
	AddEntry(Message)

	//GetEntries returns list of all messages in the Bag
	GetEntries() []Message

	//NewBag returns a newBag of Message
	NewBag() Bag
}

//DTO represents the interface every interface needs to implement
type DTO interface {
	data.Entity
	//ToPB provides methods to convert DTO to proto-buf Message
	ToPB() Message
}

//DataModel represents the datamodel being produced
type DataModel interface {

	//NewBag returns a newBag of Message
	NewBag() Bag

	//CacheDuration provides cache duration in time.NanoSeconds
	CacheDuration() int64

	//LoadAll provides the method to load all the data
	LoadAll() (Bag, error)

	//GetDataName provides the data-name and should be unique
	GetDataName() string
}
