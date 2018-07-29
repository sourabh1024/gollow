package sources

import (
	"github.com/golang/protobuf/proto"
)

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
