package core

//import (
//	"errors"
//	"github.com/go/Godeps/_workspace/src/github.com/stretchr/testify/assert"
//	"gollow/sources"
//	"gollow/sources/datamodel"
//	"gollow/storage"
//	"testing"
//)
//
//func TestDiff_GetDiffBetweenModels(t *testing.T) {
//
//	oldData := []sources.DataModel{
//		&datamodel.HeatMapData{
//			ID:      1,
//			Geohash: "abc",
//		},
//		&datamodel.HeatMapData{
//			ID:      2,
//			Geohash: "abc",
//		},
//		&datamodel.HeatMapData{
//			ID:      4,
//			Geohash: "abc",
//		},
//	}
//
//	newData := []sources.DataModel{
//		&datamodel.HeatMapData{
//			ID:      1,
//			Geohash: "abc",
//		},
//		&datamodel.HeatMapData{
//			ID:      2,
//			Geohash: "def",
//		},
//		&datamodel.HeatMapData{
//			ID:      3,
//			Geohash: "abc",
//		},
//	}
//
//	delta := getDiffBetweenModels(oldData, newData)
//
//	newObjects := make([]sources.DataModel, 0)
//	changedObjects := make([]sources.DataModel, 0)
//	missingKeys := make([]string, 0)
//
//	newObjects = append(newObjects, &datamodel.HeatMapData{
//		ID:      3,
//		Geohash: "abc",
//	})
//
//	changedObjects = append(changedObjects, &datamodel.HeatMapData{
//		ID:      2,
//		Geohash: "def",
//	})
//
//	missingKeys = append(missingKeys, "4")
//
//	assert.Equal(t, delta.NewObjects, newObjects)
//	assert.Equal(t, delta.ChangedObjects, changedObjects)
//	assert.Equal(t, delta.MissingKeys, missingKeys)
//}
//
//func TestDiffObject_createDiff(t *testing.T) {
//
//	dataModelMock := new(sources.MockDataSource)
//	storageMock := new(storage.MockStorage)
//
//	data := []sources.DataModel{
//		&datamodel.HeatMapData{
//			ID:      1,
//			Geohash: "abc",
//		},
//		&datamodel.HeatMapData{
//			ID:      2,
//			Geohash: "def",
//		},
//		&datamodel.HeatMapData{
//			ID:      3,
//			Geohash: "ghi",
//		},
//		&datamodel.HeatMapData{
//			ID:      4,
//			Geohash: "jkl",
//		},
//		&datamodel.HeatMapData{
//			ID:      5,
//			Geohash: "mno",
//		},
//		&datamodel.HeatMapData{
//			ID:      1,
//			Geohash: "why",
//		},
//	}
//
//	scenarios := []struct {
//		desc          string
//		setup         func()
//		oldData       []sources.DataModel
//		newData       []sources.DataModel
//		expectedError error
//	}{
//		{
//			desc: "Happy Path with diff to be generated, 1 missing key",
//			setup: func() {
//				dataModelMock.On("GetNameSpace").Return("test_namespace").Once()
//				dataModelMock.On("GetDataName").Return("test_dataname").Once()
//				storageMock.On("Write", mock.Anything).Return(1, nil).Once()
//			},
//			oldData:       []sources.DataModel{data[0], data[1], data[2]},
//			newData:       []sources.DataModel{data[0], data[1]},
//			expectedError: nil,
//		},
//		{
//			desc: "Happy Path with diff to be generated, 1 new object , 1 changed , 1 missing",
//			setup: func() {
//				dataModelMock.On("GetNameSpace").Return("test_namespace").Once()
//				dataModelMock.On("GetDataName").Return("test_dataname").Once()
//				storageMock.On("Write", mock.Anything).Return(1, nil).Once()
//			},
//			oldData:       []sources.DataModel{data[0], data[1], data[2]},
//			newData:       []sources.DataModel{data[3], data[5]},
//			expectedError: nil,
//		},
//		{
//			desc: "Happy Path with no diff to be generated",
//			setup: func() {
//				dataModelMock.On("GetNameSpace").Return("test_namespace").Once()
//				dataModelMock.On("GetDataName").Return("test_dataname").Once()
//			},
//			oldData:       []sources.DataModel{data[0], data[1]},
//			newData:       []sources.DataModel{data[0], data[1]},
//			expectedError: errors.New("no new data for diff"),
//		},
//		{
//			desc: "Unhappy path with error in writing",
//			setup: func() {
//				dataModelMock.On("GetNameSpace").Return("test_namespace").Once()
//				dataModelMock.On("GetDataName").Return("test_dataname").Once()
//				storageMock.On("Write", mock.Anything).Return(1, errors.New("error in writing")).Once()
//
//			},
//			oldData:       []sources.DataModel{data[0]},
//			newData:       []sources.DataModel{data[0], data[1]},
//			expectedError: errors.New("error in writing"),
//		},
//	}
//
//	for _, scenario := range scenarios {
//		scenario.setup()
//		err := DiffObjectDao.createDiff(dataModelMock, scenario.oldData, scenario.newData, 1, 2, storageMock)
//		mock.AssertExpectationsForObjects(t)
//		assert.Equal(t, err, scenario.expectedError)
//	}
//
//}
