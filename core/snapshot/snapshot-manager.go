package snapshot

import (
	"encoding/json"
	"gollow/logging"
	"gollow/storage"
	regexp "regexp"
	"strconv"
)

const (
	DefaultVersionNumber = 1
)

var (
	r, _ = regexp.Compile(".*-([0-9]+)$")
)

type Snapshot struct {
	AnnouncedSnapshot map[string]string
}

func GetLatestAnnouncedVersion(announcedVersionStorage storage.Storage, keyName string) (string, error) {

	data, err := announcedVersionStorage.Read()

	if err != nil {
		return "", err
	}

	var snapshot Snapshot
	err = json.Unmarshal(data, &snapshot)

	if err != nil {
		logging.GetLogger().Error("Error in unmarshalling announced version :", err)
		return "", err
	}

	announcedVersion, ok := snapshot.AnnouncedSnapshot[keyName]

	if !ok {
		// not produced till now
		return "", nil
	}

	return announcedVersion, nil

}

func UpdateLatestAnnouncedVersion(announcedVersionStorage storage.Storage, keyName, newVersion string) error {

	data, err := announcedVersionStorage.Read()

	if err != nil {
		return err
	}

	var snapshot Snapshot
	err = json.Unmarshal(data, &snapshot)

	if err != nil {
		logging.GetLogger().Error("Error in unmarshalling announced version :", err)
		return err
	}

	logging.GetLogger().Info("Updating the version for key : " + keyName + " with version : " + newVersion)
	snapshot.AnnouncedSnapshot[keyName] = newVersion

	data, err = json.Marshal(snapshot)
	if err != nil {
		logging.GetLogger().Error("Marshalling Error : ", err)
	}
	_, err = announcedVersionStorage.Write(data)
	return err
}

func GetVersionNumber(fileName string) int64 {
	versionNumber, err := strconv.ParseInt(r.FindStringSubmatch(fileName)[1], 10, 64)
	if err != nil {
		logging.GetLogger().Error("Error in extracting version number for : ", fileName)
	}
	return versionNumber
}

func WriteNewSnapshot(fileName string, data []byte) error {
	snapshotStorage := storage.NewStorage(fileName)
	_, err := snapshotStorage.Write(data)
	if err != nil {
		logging.GetLogger().Error("Error in producing snapshot : ", fileName)
	}
	return err
}
