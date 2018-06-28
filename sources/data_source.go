package sources

import (
	"gollow/data"
)

// DataSource is the interface every datamodel needs to implement
// to be fetched by source data retriever i.e., producer and consumer
// it needs to implement enity , producer interface
type DataModel interface {
	data.Entity

	DataProducer
}
