package snapshot

import (
	"github.com/golang/protobuf/proto"
	"gollow/cdd/core/storage"
	"gollow/cdd/sources"
	"sync"
)

//snapshotImpl implements the Snapshot
type snapshotImpl struct {
	sync.RWMutex
	storage storage.Storage
}

//NewSnapshot returns the SnapshotImpl initialised with passed storage
func NewSnapshot(storage storage.Storage) *snapshotImpl {
	return &snapshotImpl{
		storage: storage,
	}
}

//Load loads the snapshot of given model type into Model Bag from the given storage and file
func (s *snapshotImpl) Load(model sources.DataModel) (sources.Bag, error) {

	data, err := s.storage.Read()
	if err != nil {
		return nil, err
	}

	bag := model.NewBag()

	err = proto.Unmarshal(data, bag)

	if err != nil {
		return nil, err
	}

	return bag, nil
}

//Save saves the Model Bag into the given storage and file name
func (s *snapshotImpl) Save(data sources.Bag) (int, error) {

	bytes, err := proto.Marshal(data)
	if err != nil {
		return 0, err
	}

	return s.storage.Write(bytes)
}
