package dummy

import (
	"gollow/cdd/sources"
	"strconv"
	"time"
)

var _ sources.Bag = &DummyDataBag{}
var _ sources.Message = &DummyData{}
var _ sources.DataModel = &DummyData{}

// GetPartitionID implements the Message interface
func (message DummyData) GetPrimaryID() string {
	return strconv.FormatInt(message.ID, 10)
}

func (message DummyData) LoadAll() (sources.Bag, error) {
	return DummyDataRef.LoadAll()
}

func (message DummyData) CacheDuration() int64 {
	return int64(time.Duration(2 * time.Minute))
	//return DummyDataRef.CacheDuration()
}

func (message DummyData) GetDataName() string {
	return DummyDataRef.GetDataName()
}

func (message DummyData) NewBag() sources.Bag {
	return &DummyDataBag{}
}

// AddEntry implements the Bag interface
func (data *DummyDataBag) AddEntry(record sources.Message) {
	data.Entries = append(data.Entries, record.(*DummyData))
}

// GetEntries implements the Bag interface
func (data *DummyDataBag) GetEntries() []sources.Message {
	out := make([]sources.Message, len(data.Entries))

	for index, entry := range data.Entries {
		out[index] = entry
	}

	return out
}

//NewBag implements the Bag interface
func (data *DummyDataBag) NewBag() sources.Bag {
	return &DummyDataBag{}
}
