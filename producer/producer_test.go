package producer

//import (
//	"github.com/magiconair/properties/assert"
//	"gollow/core/snapshot"
//	"gollow/sources"
//	"strconv"
//	"testing"
//)
//
//func Test_getAnnouncedVersionName(t *testing.T) {
//	dataModelMock := new(sources.MockDataSource)
//
//	dataModelMock.On("GetNameSpace").Return("test_namespace")
//	dataModelMock.On("GetDataName").Return("test_dataname")
//
//	assert.Equal(t, getAnnouncedVersionName(dataModelMock, -1),
//		"test_namespace-test_dataname-"+strconv.FormatInt(snapshot.DefaultVersionNumber, 10))
//	assert.Equal(t, getAnnouncedVersionName(dataModelMock, 1),
//		"test_namespace-test_dataname-"+strconv.FormatInt(2, 10))
//
//	dataModelMock.AssertExpectations(t)
//}
//
//func Test_announcedVersionKeyName(t *testing.T) {
//
//	dataModelMock := new(sources.MockDataSource)
//
//	dataModelMock.On("GetNameSpace").Return("test_namespace")
//	dataModelMock.On("GetDataName").Return("test_dataname")
//
//	assert.Equal(t, snapshot.AnnouncedVersionKeyName(dataModelMock.GetNameSpace(), dataModelMock.GetDataName()), "test_namespace-test_dataname")
//	dataModelMock.AssertExpectations(t)
//}
