package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"gollow/cdd/core"
	"gollow/cdd/core/snapshot"
	"gollow/cdd/core/storage"
	"gollow/cdd/logging"
	"gollow/cdd/sources"
	"gollow/cdd/util"
	"time"
)

var (
	ErrTypeCasting = errors.New("error in typecasting the interface")
)

func fetchAnnouncedVersion(ctx context.Context, model sources.DataModel) (string, error) {
	announcedVersion, err := snapshot.SnapshotImpl.GetLatestAnnouncedVersion(model.GetDataName())
	if err != nil {
		logging.GetLogger().Error("Error in fetching current snapshot version for : %s : %+v", model.GetDataName(), err)
		return "", err
	}
	return announcedVersion, nil
}

func FetchFullSnapshot(currentVersion string, source sources.DataModel, cache GollowCache) error {

	logging.GetLogger().Info("Building cache with full snapshot  : %s ", currentVersion)
	data, err := loadSnapshot(currentVersion, source)
	if err != nil {
		logging.GetLogger().Error("Error in fetching full snapshot : %+v ", err)
		return err
	}
	BuildCache(data, cache)
	return nil
}

func FetchSnapshot(ctx context.Context, source sources.DataModel, cache GollowCache) {

	defer util.Duration(time.Now(),
		fmt.Sprintf("fetch snapshot : %s", source.GetDataName()))

	currentVersion, ok := GetSnapshotVersion().getSnapshotVersion(source)

	announcedVersion, err := fetchAnnouncedVersion(ctx, source)
	if err != nil || announcedVersion == "" {
		logging.GetLogger().Error("Error in fetching the announced version , err : %+v", err)
		return
	}

	logging.GetLogger().Info("Current Version is : %s ", currentVersion)
	if currentVersion == "" || !ok {
		logging.GetLogger().Info("Fetching full snapshot : %s", announcedVersion)
		err := FetchFullSnapshot(announcedVersion, source, cache)
		if err != nil {
			logging.GetLogger().Error("Error in fetching the full snapshot , err : %+v", err)
			return
		}
		GetSnapshotVersion().updateSnapshotVersion(source, announcedVersion)
		return
	}

	if currentVersion != announcedVersion {
		if currentVersion < announcedVersion {
			logging.GetLogger().Error("current version is less than announced version , currVersion : %d , snapshotVersion : %d", currentVersion, announcedVersion)
		}
		diffs := getDiffBetweenVersions(source, currentVersion, announcedVersion)
		err := applyAllDiffs(diffs, source, cache)
		if err != nil {
			logging.GetLogger().Error("Error in applying diff , err : %+v ", err)
			return
		}
		GetSnapshotVersion().updateSnapshotVersion(source, announcedVersion)
	}
}

func applyAllDiffs(diffs []string, source sources.DataModel, cache GollowCache) error {

	for _, diffName := range diffs {

		logging.GetLogger().Info("Applying  diff : %s", diffName)

		diff, err := getDiffObject(diffName, source)
		if err != nil {
			logging.GetLogger().Error("Error in Unmarshalling : ", err)
			return err
		}

		err = applyDiff(source, diff, cache)
		if err != nil {
			return err
		}
	}
	return nil
}

func getDiffObject(diffName string, model sources.DataModel) (*core.DiffObject, error) {
	diffManager := storage.NewStorage(diffName)
	data, err := diffManager.Read()
	if err != nil {
		return nil, err
	}
	d := &core.DiffObject{}
	d.ChangedObjects = model.NewBag()
	d.NewObjects = model.NewBag()
	err = json.Unmarshal(data, &d)

	return d, err
}

// i don't like this method but I am making peace with it now.. -_-
func applyDiff(model sources.DataModel, d *core.DiffObject, cache GollowCache) error {

	defer util.Duration(time.Now(), "apply-diff")
	logging.GetLogger().Info("applying diff : ", d.Namespace)

	for _, object := range d.NewObjects.GetEntries() {
		cache.Set(object.GetPrimaryID(), object)
	}

	logging.GetLogger().Info("New Objects updated in the map")

	for _, object := range d.ChangedObjects.GetEntries() {
		cache.Set(object.GetPrimaryID(), object)
	}

	logging.GetLogger().Info("Changed Objects updated in the map")

	missingKeys := d.MissingKeys

	for _, key := range missingKeys {
		cache.Delete(key)
	}

	logging.GetLogger().Info("Deleted Objects  in the map")

	return nil

}

//getDiffBetweenVersions returns all the diff required to reach from version1 to version2
func getDiffBetweenVersions(source sources.DataModel, version1, version2 string) []string {

	v1 := snapshot.GetVersionNumber(version1)
	v2 := snapshot.GetVersionNumber(version2)

	diffs := make([]string, 0)
	for i := v1; i < v2; i++ {
		diffs = append(diffs, core.DiffObjectDao.GetDiffName(source, i, i+1))
	}

	return diffs
}

func loadSnapshot(lastAnnouncedSnapshot string, model sources.DataModel) (sources.Bag, error) {

	dataBytes, err := snapshot.ReadSnapshot(lastAnnouncedSnapshot)
	if err != nil {
		return nil, err
	}

	data := model.NewBag()
	err = proto.Unmarshal(dataBytes, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
