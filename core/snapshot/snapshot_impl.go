package snapshot

import (
	"encoding/json"
	"gollow/core/storage"
	"gollow/logging"
	"regexp"
	"strconv"
	"sync"
)

/*
this file could be improved by only acquiring locks while updating the map or loading
*/
var (
	SnapshotImpl    = &snapshotImpl{}
	r, _            = regexp.Compile(".*-([0-9]+)$")
	snapshotStorage *snapshotStorageImpl
)

func Init(storage storage.Storage) {
	snapshotStorage = &snapshotStorageImpl{
		storage: storage,
	}
}

type snapshotStorageImpl struct {
	sync.RWMutex
	storage storage.Storage
}

type snapshotImpl struct {
	sync.RWMutex
	announcedVersion map[string]string
}

func (s *snapshotImpl) GetLatestAnnouncedVersion(keyName string) (string, error) {
	s.RLock()
	defer s.RUnlock()

	var err error
	//if s.announcedVersion == nil {
	s.announcedVersion, err = loadAnnouncedVersion()
	if err != nil {
		logging.GetLogger().Error("error in loading snapshot from storage , err : ", err)
		return "", err
	}
	//}

	announcedVersion, ok := s.announcedVersion[keyName]
	if !ok {
		// not produced till now
		return "", nil
	}

	return announcedVersion, nil

}

func (s *snapshotImpl) UpdateLatestAnnouncedVersion(keyName, newVersion string) error {

	s.Lock()
	defer s.Unlock()

	var err error
	//if s.announcedVersion == nil {
	s.announcedVersion, err = loadAnnouncedVersion()
	if err != nil {
		logging.GetLogger().Error("error in loading snapshot from storage , err : ", err)
		return err
	}
	//}

	logging.GetLogger().Info("Updating the version for key : " + keyName + " with version : " + newVersion)
	oldVersion := s.announcedVersion[keyName]
	s.announcedVersion[keyName] = newVersion

	//persist the change
	err = writeAnnouncedVersion(s.announcedVersion)
	if err != nil {
		//restore the old announced version , since new version could not be updated, hence would cause inconsistencies
		s.announcedVersion[keyName] = oldVersion
		logging.GetLogger().Error("error updating the announced version for keyname : ", keyName, err)
		return err
	}
	return nil
}

func loadAnnouncedVersion() (map[string]string, error) {
	snapshotStorage.RLock()
	defer snapshotStorage.RUnlock()
	data, err := snapshotStorage.storage.Read()

	if err != nil {
		return nil, err
	}
	announcedVersion := make(map[string]string)
	err = json.Unmarshal(data, &announcedVersion)
	if err != nil {
		return nil, err
	}
	return announcedVersion, nil
}

func writeAnnouncedVersion(snapshot map[string]string) error {
	data, err := json.Marshal(snapshot)
	if err != nil {
		logging.GetLogger().Error("Marshalling Error : ", err)
	}
	_, err = snapshotStorage.storage.Write(data)
	return err
}

func GetVersionNumber(fileName string) int64 {
	versionNumber, err := strconv.ParseInt(r.FindStringSubmatch(fileName)[1], 10, 64)
	if err != nil {
		logging.GetLogger().Error("Error in extracting version number for : ", fileName)
		return -1
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

func ReadSnapshot(fileName string) ([]byte, error) {
	snapshotStorage := storage.NewStorage(fileName)
	return snapshotStorage.Read()
}

func AnnouncedVersionKeyName(namespace, entityName string) string {
	return namespace + Separator + entityName
}
