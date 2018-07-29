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

func GetSnapshotVersion() *snapshotVersion {
	versionOnce.Do(func() {
		version = &snapshotVersion{
			snapshotsVersions: make(map[string]string, 0),
		}
	})
	return version
}

func (s *snapshotVersion) getSnapshotVersion(model sources.ProtoDataModel) (string, bool) {
	val, ok := s.snapshotsVersions[getSnapshotKey(model)]
	return val, ok
}

func (s *snapshotVersion) updateSnapshotVersion(model sources.ProtoDataModel, newVersion string) {
	s.snapshotsVersions[getSnapshotKey(model)] = newVersion
}

func getSnapshotKey(model sources.ProtoDataModel) string {
	return model.GetDataName()
}
