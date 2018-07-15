package core

import (
	"golang.org/x/net/context"
	"gollow/core/snapshot"
	"gollow/storage"
	"gollow/util"
	"time"
)

func GetAnnouncedVersions(_ context.Context, namespace, entityName string) (string, error) {
	defer util.Duration(time.Now(), "GetAnnouncedVersionApi")

	announcedFileName := "announced.version"
	announcedVersionStorage := storage.NewStorage(announcedFileName)

	return snapshot.GetLatestAnnouncedVersion(announcedVersionStorage, namespace+"-"+entityName)
}
