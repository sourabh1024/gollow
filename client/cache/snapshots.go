package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"gollow/core"
	"gollow/core/snapshot"
	"gollow/core/storage"
	"gollow/logging"
	"gollow/server/api"
	"gollow/sources"
	"gollow/util"
	"time"
)

var (
	ErrTypeCasting = errors.New("error in typecasting the interface")
)

func fetchAnnouncedVersion(client api.PingClient, ctx context.Context, source sources.DataModel) (string, error) {
	announcedVersion, err := client.GetAnnouncedVersion(ctx,
		&api.AnnouncedVersionRequest{Namespace: source.GetNameSpace(), Entity: source.GetDataName()})

	if err != nil {
		logging.GetLogger().Error("Error in fetching current snapshot version for : %s , %s: %+v", source.GetNameSpace(), source.GetDataName(), err)
		return "", err
	}
	return announcedVersion.Currentversion, nil
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

func FetchSnapshot(client api.PingClient, ctx context.Context, source sources.DataModel, cache GollowCache) {

	defer util.Duration(time.Now(),
		fmt.Sprintf("fetch snapshot : %s-%s", source.GetNameSpace(), source.GetDataName()))

	currentVersion, ok := GetSnapshotVersion().getSnapshotVersion(source)

	announcedVersion, err := fetchAnnouncedVersion(client, ctx, source)
	if err != nil || announcedVersion == "" {
		logging.GetLogger().Error("Error in fetching the announced version , err : %+v", err)
		return
	}

	if currentVersion == "" || !ok {
		logging.GetLogger().Info("Fetching full snapshot : %s", announcedVersion)
		err := FetchFullSnapshot(announcedVersion, source, cache)
		if err != nil {
			logging.GetLogger().Error("Error in fetching the full snapshot , err : %+v", err)
			return
		}
		GetSnapshotVersion().updateSnapshotVersion(source, currentVersion)
		return
	}

	if currentVersion != announcedVersion {
		diffs := getDiffBetweenVersions(source, currentVersion, announcedVersion)
		err := applyAllDiffs(diffs, source, cache)
		if err != nil {
			logging.GetLogger().Error("Error in applying diff , err : %+v ", err)
			return
		}
		GetSnapshotVersion().updateSnapshotVersion(source, currentVersion)
	}
}

func applyAllDiffs(diffs []string, source sources.DataModel, cache GollowCache) error {

	for _, diffName := range diffs {

		logging.GetLogger().Info("Applying  diff : %s", diffName)

		diff, err := getDiffObject(diffName)
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

func getDiffObject(diffName string) (*core.DiffObject, error) {
	diffManager := storage.NewStorage(diffName)
	data, err := diffManager.Read()
	if err != nil {
		return nil, err
	}
	d := &core.DiffObject{}
	err = json.Unmarshal(data, &d)

	return d, err
}

// i don't like this method but I am making peace with it now.. -_-
func applyDiff(source sources.DataModel, d *core.DiffObject, cache GollowCache) error {

	defer util.Duration(time.Now(), "apply-diff")
	logging.GetLogger().Info("applying diff : ", d.Namespace)

	newObjectsInterface, ok := d.NewObjects.([]interface{})
	if !ok {
		return ErrTypeCasting
	}

	changedObjectsInterface, ok := d.ChangedObjects.([]interface{})
	if !ok {
		return ErrTypeCasting
	}

	newObjects, err := util.UnMarshalInterfaceToModel(newObjectsInterface, source)
	if err != nil {
		return err
	}
	changedObjects, err := util.UnMarshalInterfaceToModel(changedObjectsInterface, source)
	if err != nil {
		return err
	}

	for _, object := range newObjects {
		cache.Set(object.GetPrimaryKey(), object)
	}

	logging.GetLogger().Info("New Objects updated in the map")

	for _, object := range changedObjects {
		cache.Set(object.GetPrimaryKey(), object)
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
		diffs = append(diffs, core.DiffObjectDao.GetDiffName(source, v1, i+1))
		v1 = i
	}

	return diffs
}

func loadSnapshot(lastAnnouncedSnapshot string, model sources.DataModel) ([]sources.DataModel, error) {

	dataBytes, err := snapshot.ReadSnapshot(lastAnnouncedSnapshot)
	if err != nil {
		return nil, err
	}

	data, err := util.UnMarshalDataModelsBytes(dataBytes, model.NewDataRef())
	if err != nil {
		return nil, err
	}

	return data, nil
}
