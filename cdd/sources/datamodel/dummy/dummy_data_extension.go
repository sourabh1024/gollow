package dummy

import (
	"gollow/cdd/sources"
	"strconv"
	"time"
)

var _ sources.Bag = &DummyDataBag{}

// GetPartitionID implements the Message interface
func (message DummyData) GetPrimaryID() string {
	return strconv.FormatInt(message.ID, 10)
}

// NewBag implements the Message interface
func (message DummyData) NewBag() sources.Bag {
	return &DummyDataBag{}
}

func (message DummyData) LoadAll() (sources.Bag, error) {
	x := &DummyDataDTO{}
	return x.LoadAll()
}

func (message DummyData) CacheDuration() int64 {
	return int64(time.Duration(2 * time.Minute))
}

func (message DummyData) GetDataName() string {
	return "dummy_data"
}

// AddEntry implements the Message interface
func (data *DummyDataBag) AddEntry(record sources.Message) {
	data.Entries = append(data.Entries, record.(*DummyData))
}

// GetEntries implements the Message interface
func (data *DummyDataBag) GetEntries() []sources.Message {
	out := make([]sources.Message, len(data.Entries))

	for index, entry := range data.Entries {
		out[index] = entry
	}

	return out
}
