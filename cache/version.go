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

package cache

import (
	"gollow/sources"
	"sync"
)

var version *snapshotVersion
var versionOnce sync.Once

type snapshotVersion struct {
	sync.RWMutex
	snapshotsVersions map[string]string
}

// GetSnapshotVersion returns a singleton object of SnapshotVersion of Cache
// all the go routines should access the same map, hence singleton
func GetSnapshotVersion() *snapshotVersion {
	versionOnce.Do(func() {
		version = &snapshotVersion{
			snapshotsVersions: make(map[string]string, 0),
		}
	})
	return version
}

func (s *snapshotVersion) getSnapshotVersion(model sources.DataModel) (string, bool) {
	val, ok := s.snapshotsVersions[getSnapshotKey(model)]
	return val, ok
}

func (s *snapshotVersion) updateSnapshotVersion(model sources.DataModel, newVersion string) {
	s.snapshotsVersions[getSnapshotKey(model)] = newVersion
}

func getSnapshotKey(model sources.DataModel) string {
	return model.GetDataName()
}
