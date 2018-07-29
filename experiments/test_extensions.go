package experiments

// GetPartitionID implements the Message interface
func (message Test) GetPrimaryID() int64 {
	return message.Id
}

// NewBag implements the Message interface
func (message Test) NewBag() Bag {
	return &TestBag{}
}

// AddEntry implements the Message interface
func (data *TestBag) AddEntry(record Message) {
	data.Entries = append(data.Entries, record.(*Test))
}

// GetEntries implements the Message interface
func (data *TestBag) GetEntries() []Message {
	out := make([]Message, len(data.Entries))

	for index, entry := range data.Entries {
		out[index] = entry
	}

	return out
}
