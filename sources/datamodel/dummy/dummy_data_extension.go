package dummy

import (
	"gollow/sources"
	"strconv"
	"time"
)

var _ sources.Bag = &DummyDataBag{}
var _ sources.Message = &DummyData{}
var _ sources.DataModel = &DummyData{}

// GetUniqueKey implements the Message interface
// should return the unique key for the message
// can be an combination of 2 or more keys also
func (message DummyData) GetUniqueKey() string {
	return strconv.FormatInt(message.ID, 10)
}

// LoadAll implements the DataModel interface
// LoadAll loads all the data for given message type
// returns the Bag of object and error if any
func (message DummyData) LoadAll() (sources.Bag, error) {
	return DummyDataRef.LoadAll()
}

// CacheDuration implements the DataModel interface
// CacheDuration returns the cache duration for the given data model
func (message DummyData) CacheDuration() int64 {
	return int64(time.Duration(1 * time.Minute))
}

// GetDataName implements the DataModel interface
// GetDataName returns the data name which is used for storing announced version
// should be unique at producer and consumer , collisions should be avoided
// by having a proper descriptive name like www.xyz.org.team1.db1.table1
func (message DummyData) GetDataName() string {
	return DummyDataRef.GetDataName()
}

// NewBag implements the DataModel interface
// returns an empty Bag of Message
func (message DummyData) NewBag() sources.Bag {
	return &DummyDataBag{}
}

// AddEntry implements the Bag interface
// adds the given message to the bag
func (data *DummyDataBag) AddEntry(record sources.Message) {
	data.Bag = append(data.Bag, record.(*DummyData))
}

// GetEntries implements the Bag interface
// returns a list of message
func (data *DummyDataBag) GetEntries() []sources.Message {
	out := make([]sources.Message, len(data.Bag))

	for index, entry := range data.Bag {
		out[index] = entry
	}

	return out
}

//NewBag implements the Bag interface
func (data *DummyDataBag) NewBag() sources.Bag {
	return &DummyDataBag{}
}
