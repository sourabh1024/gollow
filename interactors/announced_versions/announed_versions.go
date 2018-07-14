package announced_versions

import (
	"golang.org/x/net/context"
	"gollow/snapshot"
	"gollow/storage"
	"gollow/util"
	"time"
)

func GetAnnouncedVersions(ctx context.Context, namespace, entityName string) (string, error) {

	defer util.Duration(time.Now(), "GetAnnouncedVersionApi")

	announcedFileName := "announced.version"
	announcedVersionStorage := storage.NewFileStorage(announcedFileName)

	return snapshot.GetLatestAnnouncedVersion(announcedVersionStorage, namespace+"-"+entityName)

}
