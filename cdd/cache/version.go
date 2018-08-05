package cache

import (
	"gollow/cdd/sources"
	"sync"
)

var version *snapshotVersion
var versionOnce sync.Once

type snapshotVersion struct {
	sync.RWMutex
	snapshotsVersions map[string]string
}

// GetSnapshotVersion returns a singleton object of SnapshotVersion of Cache
// all the go routines should access the same map , hence singleton
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
