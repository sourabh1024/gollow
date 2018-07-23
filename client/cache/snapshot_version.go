package cache

import (
	"gollow/core/snapshot"
	"gollow/sources"
	"sync"
)

var version *snapshotVersion
var versionOnce sync.Once

type snapshotVersion struct {
	sync.RWMutex
	snapshotsVersions map[string]string
}

func GetSnapshotVersion() *snapshotVersion {
	versionOnce.Do(func() {
		version = &snapshotVersion{}
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
	return snapshot.AnnouncedVersionKeyName(model.GetNameSpace(), model.GetDataName())
}
