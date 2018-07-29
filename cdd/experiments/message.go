package experiments

import (
	"github.com/golang/protobuf/proto"
)

type Message interface {
	proto.Message

	GetPrimaryID() int64

	NewBag() Bag
}

type Bag interface {
	proto.Message

	AddEntry(Message)

	GetEntries() []Message
}
