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
	"encoding/json"
	"github.com/sourabh1024/gollow/core/storage"
	"github.com/sourabh1024/gollow/data"
	"regexp"
	"strconv"
	"sync"
)

var (
	r, _ = regexp.Compile(".*-([0-9]+)$")
	//VersionImpl implements the Version interface for announced version
	VersionImpl *versionImpl
)

//Init initialises the version
func Init(storage storage.Storage) {
	VersionImpl = &versionImpl{
		storage: storage,
	}
}

type versionImpl struct {
	sync.RWMutex
	storage storage.Storage
}

//GetVersion returns the currentVersion for the given keyName
func (s *versionImpl) GetVersion(keyName string) (string, error) {
	s.RLock()
	defer s.RUnlock()

	var err error
	versionMap, err := loadAnnouncedVersion()
	if err != nil {
		return "", err
	}

	announcedVersion, ok := versionMap[keyName]
	if !ok {
		return "", data.ErrNoData
	}

	return announcedVersion, nil
}

//UpdateVersion updates the version for the given keyName with newVersion
func (s *versionImpl) UpdateVersion(keyName, newVersion string) error {

	s.Lock()
	defer s.Unlock()

	var err error
	versionMap, err := loadAnnouncedVersion()
	if err != nil {
		return err
	}

	versionMap[keyName] = newVersion

	//persist the change
	err = writeAnnouncedVersion(versionMap)
	if err != nil {
		return err
	}
	return nil
}

//ParseVersionNumber parses the version number from the file
//any integer after last - is considered as version number
func (s *versionImpl) ParseVersionNumber(fileName string) (int64, error) {
	versionNumber, err := strconv.ParseInt(r.FindStringSubmatch(fileName)[1], 10, 64)
	if err != nil {
		return -1, err
	}
	return versionNumber, nil
}

//loadAnnouncedVersion loads the announced version
func loadAnnouncedVersion() (map[string]string, error) {

	bytes, err := VersionImpl.storage.Read()

	if err != nil {
		return nil, err
	}
	announcedVersion := make(map[string]string)
	err = json.Unmarshal(bytes, &announcedVersion)
	if err != nil {
		return nil, err
	}
	return announcedVersion, nil
}

//writeAnnouncedVersion writes the announced version map to storage
func writeAnnouncedVersion(snapshot map[string]string) error {
	bytes, err := json.Marshal(snapshot)
	if err != nil {
		return err
	}
	_, err = VersionImpl.storage.Write(bytes)
	return err
}
