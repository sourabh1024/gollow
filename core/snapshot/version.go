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
	"fmt"
	"gollow/core/storage"
)

//Version represents the version information for all the keys
type Version interface {
	//GetVersion returns the current version for the given keyName
	GetVersion(keyName string) (string, error)

	//UpdateVersion updates the version for the given keyName
	UpdateVersion(keyName, newVersion string) error

	//ParseVersion parses the VersionImpl number from the snapshotName
	ParseVersionNumber(fileName string) (int64, error)
}

//InitVersionStorage initialises the Version Storage
//Every producer and consumer must initialise it while starting
//Currently, it initialises it with the Storage and passed fileName
func InitVersionStorage(announcedVersion string) {
	if announcedVersion == "" {
		panic("Cannot initialise Storage with nil config")
	}
	versionStorage, err := storage.NewStorage(announcedVersion)
	if err != nil {
		panic(fmt.Errorf("cannot initialise storage err %v", err))
	}

	Init(versionStorage)

	// if versionStorage doesn't exist create one
	if !versionStorage.IsExist() {
		versionMap := make(map[string]string, 0)
		versionMap["version"] = "1.0.0"
		err := writeAnnouncedVersion(versionMap)
		if err != nil {
			panic("version map cannot be initialised")
		}
	}

}
