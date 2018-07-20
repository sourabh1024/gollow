package snapshot

//import (
//	"encoding/json"
//	"errors"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"gollow/storage"
//	"testing"
//)
//
//func TestSnapshotImpl_GetLatestAnnouncedVersion(t *testing.T) {
//
//	mockStorage := &storage.MockStorage{}
//	originalSnapshotStorage := snapshotStorage
//	originalAnnouncedVersion := SnapshotImpl.announcedVersion
//	defer func() {
//		snapshotStorage = originalSnapshotStorage
//		SnapshotImpl.announcedVersion = originalAnnouncedVersion
//	}()
//
//	dummyAnnouncedVersion := map[string]string{
//		"file1": "file1-1",
//		"file2": "file2-10",
//	}
//	dummyAnnouncedVersionBytes, _ := json.Marshal(dummyAnnouncedVersion)
//	Init(mockStorage)
//
//	SnapshotImpl.announcedVersion = dummyAnnouncedVersion
//
//	version, err := SnapshotImpl.GetLatestAnnouncedVersion("file1")
//	assert.Nil(t, err, "no error should be there for valid keyname")
//	assert.Equal(t, dummyAnnouncedVersion["file1"], version, "version name should match to map data")
//
//	version, err = SnapshotImpl.GetLatestAnnouncedVersion("filebla")
//	assert.Nil(t, err, "error should not be there for invalid keyname")
//	assert.Equal(t, "", version, "version name should be empty for invalid keyname")
//
//	SnapshotImpl.announcedVersion = nil
//	mockStorage.On("Read").Return(dummyAnnouncedVersionBytes, nil).Once()
//	version, err = SnapshotImpl.GetLatestAnnouncedVersion("file1")
//	assert.Nil(t, err, "no error should be there for valid keyname")
//	assert.Equal(t, dummyAnnouncedVersion["file1"], version, "version name should match to map data")
//
//	SnapshotImpl.announcedVersion = nil
//	mockStorage.On("Read").Return([]byte("blabla"), nil).Once()
//	version, err = SnapshotImpl.GetLatestAnnouncedVersion("file1")
//	assert.NotNil(t, err, "error should be there for invalid data from snapshot storage")
//	assert.Equal(t, "", version, "version name should be empty")
//
//	SnapshotImpl.announcedVersion = nil
//	mockStorage.On("Read").Return([]byte("blabla"), errors.New("oops")).Once()
//	version, err = SnapshotImpl.GetLatestAnnouncedVersion("file1")
//	assert.NotNil(t, err, "error should be there for error from snapshot storage")
//	assert.Equal(t, "", version, "version name should be empty")
//}
//
//func TestSnapshotImpl_UpdateLatestAnnouncedVersion(t *testing.T) {
//
//	mockStorage := &storage.MockStorage{}
//	originalSnapshotStorage := snapshotStorage
//	originalAnnouncedVersion := SnapshotImpl.announcedVersion
//	defer func() {
//		snapshotStorage = originalSnapshotStorage
//		SnapshotImpl.announcedVersion = originalAnnouncedVersion
//	}()
//
//	dummyAnnouncedVersion := map[string]string{
//		"file1": "file1-1",
//		"file2": "file2-10",
//	}
//	//dummyAnnouncedVersionBytes, _ := json.Marshal(dummyAnnouncedVersion)
//	Init(mockStorage)
//
//	SnapshotImpl.announcedVersion = dummyAnnouncedVersion
//	mockStorage.On("Write", mock.Anything).Return(1, nil).Once()
//	err := SnapshotImpl.UpdateLatestAnnouncedVersion("file1", "file1-22")
//	assert.Nil(t, err, "error should be nil for happy path")
//
//	SnapshotImpl.announcedVersion = dummyAnnouncedVersion
//	mockStorage.On("Write", mock.Anything).Return(1, errors.New("error in writing snapshot")).Once()
//	oldversion, _ := SnapshotImpl.GetLatestAnnouncedVersion("file1")
//	err = SnapshotImpl.UpdateLatestAnnouncedVersion("file1", "file1-2")
//	newversion, _ := SnapshotImpl.GetLatestAnnouncedVersion("file1")
//	assert.NotNil(t, err, "error should not be nil for error in writing to snapshot")
//	assert.Equal(t, oldversion, newversion, "snapshots version should not be updated")
//
//	SnapshotImpl.announcedVersion = nil
//	mockStorage.On("Read").Return([]byte("blabla"), errors.New("oops")).Once()
//	err = SnapshotImpl.UpdateLatestAnnouncedVersion("file1", "file1-2")
//	assert.NotNil(t, err, "error should not be nil for error in lading the snapshot")
//}
//
//func TestAnnouncedVersionKeyName(t *testing.T) {
//	versionname := AnnouncedVersionKeyName("namespace", "entityname")
//	assert.Equal(t, "namespace-entityname", versionname, "TestAnnouncedVersionKeyName")
//}
//
//func TestGetVersionNumber(t *testing.T) {
//	versionnumber := GetVersionNumber("file1-1")
//	assert.Equal(t, int64(1), versionnumber, "TestGetVersionNumber")
//}
