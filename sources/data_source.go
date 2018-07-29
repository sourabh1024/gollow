package sources

import (
	"github.com/golang/protobuf/proto"
)

//DataProducer interfaces all the methods a data should implement
//for being produced

//type DataProducer interface {
//
//	//CacheDuration should return time.Duration value
//	// Its used to decide frequency at which producer or consumer will run
//	CacheDuration() int64
//
//	//Load All loads all the dataModel
//	LoadAll() ([]DataModel, error)
//}
//
//// DataSource is the interface every datamodel needs to implement
//// to be fetched by source data retriever i.e., producer and consumer
//// it needs to implement enity , producer interface
//type DataModel interface {
//
//	// NewDataRef() gives the new DataModel Object
//	NewDataRef() DataModel
//
//	// Implement the data Entity interface to satisfy methods for DB
//	data.Entity
//
//	// Implement the DataProducer interface to satisfy producer
//	DataProducer
//}

//type UniversalDTO struct {
//	Data interface{} `json:"data"`
//}

type Message interface {
	proto.Message

	GetPrimaryID() string

	NewBag() Bag
}

type Bag interface {
	proto.Message

	AddEntry(Message)

	GetEntries() []Message
}

type ProtoDataModel interface {
	Message
	CacheDuration() int64
	LoadAll() (Bag, error)
	GetDataName() string
}

type DTO interface {
	ToPB() Message
}
