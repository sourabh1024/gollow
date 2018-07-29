package core

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gollow/cdd/core/storage"
	"gollow/cdd/sources"
	"testing"
)

func TestDiff_GetDiffBetweenModels(t *testing.T) {

	oldData := []sources.DataModel{
		&datamodel.DummyData{
			ID:        1,
			FirstName: "abc",
		},
		&datamodel.DummyData{
			ID:        2,
			FirstName: "abc",
		},
		&datamodel.DummyData{
			ID:        4,
			FirstName: "abc",
		},
	}

	newData := []sources.DataModel{
		&datamodel.DummyData{
			ID:        1,
			FirstName: "abc",
		},
		&datamodel.DummyData{
			ID:        2,
			FirstName: "def",
		},
		&datamodel.DummyData{
			ID:        3,
			FirstName: "abc",
		},
	}

	delta := getDiffBetweenModels(oldData, newData)

	newObjects := make([]sources.DataModel, 0)
	changedObjects := make([]sources.DataModel, 0)
	missingKeys := make([]string, 0)

	newObjects = append(newObjects, &datamodel.DummyData{
		ID:        3,
		FirstName: "abc",
	})

	changedObjects = append(changedObjects, &datamodel.DummyData{
		ID:        2,
		FirstName: "def",
	})

	missingKeys = append(missingKeys, "4")

	assert.Equal(t, delta.NewObjects, newObjects)
	assert.Equal(t, delta.ChangedObjects, changedObjects)
	assert.Equal(t, delta.MissingKeys, missingKeys)
}

func TestDiffObject_createDiff(t *testing.T) {

	dataModelMock := new(sources.MockDataSource)
	storageMock := new(storage.MockStorage)

	data := []sources.DataModel{
		&datamodel.DummyData{
			ID:        1,
			FirstName: "abc",
		},
		&datamodel.DummyData{
			ID:        2,
			FirstName: "def",
		},
		&datamodel.DummyData{
			ID:        3,
			FirstName: "ghi",
		},
		&datamodel.DummyData{
			ID:        4,
			FirstName: "jkl",
		},
		&datamodel.DummyData{
			ID:        5,
			FirstName: "mno",
		},
		&datamodel.DummyData{
			ID:        1,
			FirstName: "why",
		},
	}

	scenarios := []struct {
		desc          string
		setup         func()
		oldData       []sources.DataModel
		newData       []sources.DataModel
		expectedError error
	}{
		{
			desc: "Happy Path with diff to be generated, 1 missing key",
			setup: func() {
				dataModelMock.On("GetNameSpace").Return("test_namespace").Once()
				dataModelMock.On("GetDataName").Return("test_dataname").Once()
				storageMock.On("Write", mock.Anything).Return(1, nil).Once()
			},
			oldData:       []sources.DataModel{data[0], data[1], data[2]},
			newData:       []sources.DataModel{data[0], data[1]},
			expectedError: nil,
		},
		{
			desc: "Happy Path with diff to be generated, 1 new object , 1 changed , 1 missing",
			setup: func() {
				dataModelMock.On("GetNameSpace").Return("test_namespace").Once()
				dataModelMock.On("GetDataName").Return("test_dataname").Once()
				storageMock.On("Write", mock.Anything).Return(1, nil).Once()
			},
			oldData:       []sources.DataModel{data[0], data[1], data[2]},
			newData:       []sources.DataModel{data[3], data[5]},
			expectedError: nil,
		},
		{
			desc: "Happy Path with no diff to be generated",
			setup: func() {
				dataModelMock.On("GetNameSpace").Return("test_namespace").Once()
				dataModelMock.On("GetDataName").Return("test_dataname").Once()
			},
			oldData:       []sources.DataModel{data[0], data[1]},
			newData:       []sources.DataModel{data[0], data[1]},
			expectedError: nil,
		},
		{
			desc: "Unhappy path with error in writing",
			setup: func() {
				dataModelMock.On("GetNameSpace").Return("test_namespace").Once()
				dataModelMock.On("GetDataName").Return("test_dataname").Once()
				storageMock.On("Write", mock.Anything).Return(1, errors.New("error in writing")).Once()

			},
			oldData:       []sources.DataModel{data[0]},
			newData:       []sources.DataModel{data[0], data[1]},
			expectedError: errors.New("error in writing"),
		},
	}

	for _, scenario := range scenarios {
		scenario.setup()
		_, err := DiffObjectDao.createDiff(dataModelMock, scenario.oldData, scenario.newData, 1, 2, storageMock)
		mock.AssertExpectationsForObjects(t)
		assert.Equal(t, err, scenario.expectedError)
	}

}
