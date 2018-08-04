package cache

import (
	"errors"
	"fmt"
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

func FetchSnapshot(ctx context.Context, model sources.DataModel, cache GollowCache) {

	defer util.Duration(time.Now(),
		fmt.Sprintf("fetch snapshot : %s", model.GetDataName()))

	announcedVersion, err := snapshot.VersionImpl.GetVersion(model.GetDataName())
	if err != nil || announcedVersion == "" {
		logging.GetLogger().Error("Error in fetching the announced version , err : %+v", err)
		return
	}

	currentCacheVersion, ok := GetSnapshotVersion().getSnapshotVersion(model)

	logging.GetLogger().Info("Current VersionImpl is : %s ", currentCacheVersion)

	if currentCacheVersion == "" || !ok {
		logging.GetLogger().Info("Fetching full snapshot : %s", announcedVersion)
		loadCompleteSnapshot(announcedVersion, model, cache)
		return
	}

	if currentCacheVersion != announcedVersion {
		if currentCacheVersion > announcedVersion {
			logging.GetLogger().Error("current cache version is greater than announced version , currCacheVersion : %d , snapshotVersion : %d",
				currentCacheVersion, announcedVersion)
			return
		}
		diffs := getDiffVersions(currentCacheVersion, announcedVersion, model)
		err := applyAllDiffs(diffs, model, cache)
		if err != nil {
			logging.GetLogger().Error("Error in applying diff , err : %+v ", err)
			return
		}
		GetSnapshotVersion().updateSnapshotVersion(model, announcedVersion)
	}
}

func applyAllDiffs(diffs []string, source sources.DataModel, cache GollowCache) error {

	for _, diffName := range diffs {

		logging.GetLogger().Info("applying  diff : %s", diffName)
		diff, err := getDiffObject(diffName, source)
		if err != nil {
			logging.GetLogger().Error("error in unMarshalling : ", err)
			return err
		}

		err = applyDiff(diff, cache)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadCompleteSnapshot(announcedVersion string, model sources.DataModel, cache GollowCache) {
	snapshotImpl := snapshot.NewSnapshot(storage.NewStorage(announcedVersion))
	bag, err := snapshotImpl.Load(model)

	if err != nil {
		logging.GetLogger().Error("error in unMarshalling  announced snapshot , err :+v", err)
		return
	}

	buildCache(bag, cache)
	GetSnapshotVersion().updateSnapshotVersion(model, announcedVersion)
}

func getDiffObject(diffName string, model sources.DataModel) (*core.DiffObject, error) {
	params := &core.DiffParams{Model: model}
	return params.LoadDiff(diffName)
}

func applyDiff(d *core.DiffObject, cache GollowCache) error {

	defer util.Duration(time.Now(), "apply-diff")
	logging.GetLogger().Info("applying diff : ", d.DataName)

	for _, object := range d.NewObjects.GetEntries() {
		cache.Set(object.GetPrimaryID(), object)
	}

	for _, object := range d.ChangedObjects.GetEntries() {
		cache.Set(object.GetPrimaryID(), object)
	}

	missingKeys := d.MissingKeys

	for _, key := range missingKeys {
		cache.Delete(key)
	}

	return nil
}

//getDiffVersions returns all the diff required to reach from version1 to version2
func getDiffVersions(version1, version2 string, model sources.DataModel) []string {

	v1, _ := snapshot.VersionImpl.ParseVersionNumber(version1)
	v2, _ := snapshot.VersionImpl.ParseVersionNumber(version2)

	diffs := make([]string, 0)
	for i := v1; i < v2; i++ {
		params := &core.DiffParams{
			Model:       model,
			PrevVersion: i,
			CurrVersion: i + 1,
		}
		diffs = append(diffs, params.GetDiffName())
	}

	return diffs
}
