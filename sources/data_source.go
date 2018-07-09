package sources

import (
	"encoding/gob"
	"gollow/data"
	"gollow/logging"
)

// DataSource is the interface every datamodel needs to implement
// to be fetched by source data retriever i.e., producer and consumer
// it needs to implement enity , producer interface
type DataModel interface {
	NewDataRef() DataModel
	data.Entity
	DataProducer
}

type DataModelS struct {
	DataModel
}

// interfaceEncode encodes the interface value into the encoder.
func InterfaceEncode(enc *gob.Encoder, p DataModel) {
	// The encode will fail unless the concrete type has been
	// registered. We registered it in the calling function.

	// Pass pointer to interface so Encode sees (and hence sends) a value of
	// interface type. If we passed p directly it would see the concrete type instead.
	// See the blog post, "The Laws of Reflection" for background.
	err := enc.Encode(&p)
	if err != nil {
		logging.GetLogger().Error("encode:", err)
	}
}

// interfaceDecode decodes the next interface value from the stream and returns it.
func InterfaceDecode(dec *gob.Decoder) (DataModel, error) {
	// The decode will fail unless the concrete type on the wire has been
	// registered. We registered it in the calling function.
	var p DataModel
	err := dec.Decode(&p)
	if err != nil {
		logging.GetLogger().Error("decode:", err)
		return nil, err
	}
	return p, nil
}
