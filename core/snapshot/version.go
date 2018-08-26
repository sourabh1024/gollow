package snapshot

import (
	"fmt"
	"gollow/core/storage"
)

//Version represents the version information for all the keys
type Version interface {
	//GetVersion returns the current version for the given keyName
	GetVersion(keyName string) (string, error)

	//UpdateVersion updates the version for the given keyName
	UpdateVersion(keyName, newVersion string) error

	//ParseVersion parses the VersionImpl number from the snapshotName
	ParseVersionNumber(fileName string) (int64, error)
}

//InitVersionStorage initialises the Version Storage
//Every producer and consumer must initialise it while starting
//Currently, it initialises it with the Storage and passed fileName
func InitVersionStorage(announcedVersion string) {
	if announcedVersion == "" {
		panic("Cannot initialise Storage with nil config")
	}
	versionStorage, err := storage.NewStorage(announcedVersion)
	if err != nil {
		panic(fmt.Errorf("cannot initialise storage err %v", err))
	}
	Init(versionStorage)
}
