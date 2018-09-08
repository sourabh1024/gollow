//Copyright 2018 Sourabh Suman ( https://github.com/sourabh1024 )
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package snapshot

import (
	"github.com/golang/protobuf/proto"
	"github.com/sourabh1024/gollow/core/storage"
	"github.com/sourabh1024/gollow/sources"
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
