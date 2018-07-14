package diff

import (
	"github.com/go/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"gollow/sources"
	"gollow/sources/datamodel"
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

	delta := GetNewDiffObj("test", "test")
	delta.GetDiffBetweenModels(oldData, newData)

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
