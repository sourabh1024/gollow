package snapshot

import (
	"encoding/json"
	"gollow/logging"
	"gollow/write"
)

type Snapshot struct {
	AnnouncedSnapshot map[string]string
}

func GetLatestAnnouncedVersion(io write.SnapshotReaderWriter, announcedVersionPath string, keyName string) (string, error) {

	data, err := io.Read(announcedVersionPath)

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

//TODO : synchronize this method
func UpdateLatestAnnouncedVersion(io write.SnapshotReaderWriter, announcedVersionPath, keyName, newVersion string) error {

	data, err := io.Read(announcedVersionPath)

	if err != nil {
		return err
	}

	var snapshot Snapshot
	err = json.Unmarshal(data, &snapshot)

	if err != nil {
		logging.GetLogger().Error("Error in unmarshalling announced version :", err)
		return err
	}

	snapshot.AnnouncedSnapshot[keyName] = newVersion

	data, err = json.Marshal(snapshot)
	if err != nil {
		logging.GetLogger().Error("Marshalling Error : ", err)
	}
	_, err = io.Write(announcedVersionPath, data)
	return err
}
