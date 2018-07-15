package cache

import (
	"encoding/json"
	"golang.org/x/net/context"
	"gollow/api"
	"gollow/core"
	"gollow/core/snapshot"
	"gollow/logging"
	"gollow/producer"
	"gollow/sources"
	"gollow/storage"
	"gollow/util"
	"sync"
	"time"
)

var ins *clientSnapshots
var onc sync.Once

type clientSnapshots struct {
	sync.RWMutex
	Snapshots map[string]string
}

func GetClientSnapshots() *clientSnapshots {
	onc.Do(func() {
		ins = newClientSnapshot()
	})
	return ins
}

func newClientSnapshot() *clientSnapshots {
	return &clientSnapshots{
		Snapshots: make(map[string]string),
	}
}

func FetchSnapshot(c api.PingClient, source sources.DataModel) {

	announcedVersion, err := c.GetAnnouncedVersion(context.Background(),
		&api.AnnouncedVersionRequest{Namespace: source.GetNameSpace(), Entity: source.GetDataName()})

	if err != nil {
		logging.GetLogger().Error("Error in fetching current snapshot version for : ", source.GetNameSpace())
		return
	}

	currentSnapshotVersion := GetClientSnapshots().GetCurrentSnapshot(getSnapshotKey(source))

	if currentSnapshotVersion == "" {
		logging.GetLogger().Info("Building cache for dirst time for : ", source.GetNameSpace())
		err := BuildSnapshot(announcedVersion.Currentversion)
		if err != nil {
			logging.GetLogger().Error("Error in building snapshots : ", err)
		}
		GetClientSnapshots().Snapshots[getSnapshotKey(source)] = announcedVersion.Currentversion

		return
	}

	if currentSnapshotVersion != announcedVersion.Currentversion {
		diffs := getDiffBetweenVersions(source, currentSnapshotVersion, announcedVersion.Currentversion)

		for _, diffName := range diffs {
			logging.GetLogger().Info("Reading diff : ", diffName)
			diffManager := storage.NewStorage(diffName)
			data, err := diffManager.Read()
			d := &core.DiffObject{}
			err = json.Unmarshal(data, &d)
			if err != nil {
				logging.GetLogger().Error("Error in Unmarshalling : ", err)
				continue
			}

			applyDiff(source, d)
			GetClientSnapshots().Snapshots[getSnapshotKey(source)] = announcedVersion.Currentversion
		}
	}
}

func applyDiff(source sources.DataModel, d *core.DiffObject) {

	defer util.Duration(time.Now(), "applydiff")
	logging.GetLogger().Info("applying diff : ", d.Namespace)

	newObjectsInterface, ok := d.NewObjects.([]interface{})
	if !ok {
		logging.GetLogger().Error("Error in typecasting the interface ")
		return
	}
	newObjects, err := producer.UnMarshalInterfaceToModel(newObjectsInterface, source)

	if err != nil {
		logging.GetLogger().Error("Error in marshalling to sources datamodel objects : ", err)
		return
	}

	for _, object := range newObjects {
		GetHeatMapDataInstance().SetValue(object.GetPrimaryKey(), object)
	}

	logging.GetLogger().Info("New Objects udated in the map")

	changedObjectsInterface, ok := d.ChangedObjects.([]interface{})
	if !ok {
		logging.GetLogger().Error("Error in typecasting the interface ")
		return
	}
	changedObjects, err := producer.UnMarshalInterfaceToModel(changedObjectsInterface, source)

	if err != nil {
		logging.GetLogger().Error("Error in marshalling to sources datamodel objects : ", err)
		return
	}

	for _, object := range changedObjects {
		GetHeatMapDataInstance().SetValue(object.GetPrimaryKey(), object)
	}

	logging.GetLogger().Info("Changed Objects updated in the map")

	missingKeys := d.MissingKeys

	for _, key := range missingKeys {
		GetHeatMapDataInstance().DeleteValue(key)
	}

	logging.GetLogger().Info("Deleted Objects  in the map")

}

func getSnapshotKey(source sources.DataModel) string {
	return source.GetNameSpace() + "-" + source.GetDataName()
}

func getDiffBetweenVersions(source sources.DataModel, version1, version2 string) []string {

	v1 := snapshot.GetVersionNumber(version1)
	v2 := snapshot.GetVersionNumber(version2)

	diffs := make([]string, 0)
	for i := v1; i < v2; i++ {
		diffs = append(diffs, core.DiffObjectDao.GetDiffName(source, v1, i+1))
		v1 = i
	}

	return diffs
}
func (c *clientSnapshots) GetCurrentSnapshot(key string) string {
	c.RLock()
	defer c.RUnlock()
	val, ok := c.Snapshots[key]
	if !ok {
		return ""
	}
	return val
}

func (c *clientSnapshots) UpdateCurrentSnapshot(key string, version string) {
	c.Lock()
	defer c.Unlock()
	c.Snapshots[key] = version
}
