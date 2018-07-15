package core

import (
	"errors"
	"github.com/go/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gollow/sources"
	"gollow/sources/datamodel"
	"gollow/storage"
	"testing"
)

func TestDiff_GetDiffBetweenModels(t *testing.T) {

	oldData := []sources.DataModel{
		&datamodel.HeatMapData{
			ID:      1,
			Geohash: "abc",
		},
		&datamodel.HeatMapData{
			ID:      2,
			Geohash: "abc",
		},
		&datamodel.HeatMapData{
			ID:      4,
			Geohash: "abc",
		},
	}

	newData := []sources.DataModel{
		&datamodel.HeatMapData{
			ID:      1,
			Geohash: "abc",
		},
		&datamodel.HeatMapData{
			ID:      2,
			Geohash: "def",
		},
		&datamodel.HeatMapData{
			ID:      3,
			Geohash: "abc",
		},
	}

	delta := getDiffBetweenModels(oldData, newData)

	newObjects := make([]sources.DataModel, 0)
	changedObjects := make([]sources.DataModel, 0)
	missingKeys := make([]string, 0)

	newObjects = append(newObjects, &datamodel.HeatMapData{
		ID:      3,
		Geohash: "abc",
	})

	changedObjects = append(changedObjects, &datamodel.HeatMapData{
		ID:      2,
		Geohash: "def",
	})

	missingKeys = append(missingKeys, "4")

	assert.Equal(t, delta.NewObjects, newObjects)
	assert.Equal(t, delta.ChangedObjects, changedObjects)
	assert.Equal(t, delta.MissingKeys, missingKeys)
}

func TestDiffObject_CreateNewDiff(t *testing.T) {

	oldData := []sources.DataModel{
		&datamodel.HeatMapData{
			ID:      1,
			Geohash: "abc",
		},
		&datamodel.HeatMapData{
			ID:      2,
			Geohash: "abc",
		},
		&datamodel.HeatMapData{
			ID:      4,
			Geohash: "ghi",
		},
	}

	newData := []sources.DataModel{
		&datamodel.HeatMapData{
			ID:      1,
			Geohash: "abc",
		},
		&datamodel.HeatMapData{
			ID:      2,
			Geohash: "def",
		},
		&datamodel.HeatMapData{
			ID:      3,
			Geohash: "ghi",
		},
	}

	dataModelMock := new(sources.MockDataSource)

	dataModelMock.On("GetNameSpace").Return("test_namespace")
	dataModelMock.On("GetDataName").Return("test_dataname")
	err := DiffObjectDao.CreateNewDiff(dataModelMock, oldData, newData, 1, 2)

	storageMock := new(storage.MockStorage)
	storageMock.On("Write", mock.Anything).Return(1, errors.New("as"))
	assert.Nil(t, err)

	mock.AssertExpectationsForObjects(t)
}
