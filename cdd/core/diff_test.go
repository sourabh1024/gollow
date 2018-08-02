package core

import (
	"github.com/stretchr/testify/assert"
	"gollow/cdd/sources"
	"gollow/cdd/sources/datamodel/dummy"
	"testing"
)

func TestDiff_GetDiffBetweenModels(t *testing.T) {

	oldData := []sources.Message{
		&dummy.DummyData{
			ID:        1,
			FirstName: "abc",
		},
		&dummy.DummyData{
			ID:        2,
			FirstName: "abc",
		},
		&dummy.DummyData{
			ID:        4,
			FirstName: "abc",
		},
	}

	newData := []sources.Message{
		&dummy.DummyData{
			ID:        1,
			FirstName: "abc",
		},
		&dummy.DummyData{
			ID:        2,
			FirstName: "def",
		},
		&dummy.DummyData{
			ID:        3,
			FirstName: "abc",
		},
	}

	oldBag := dummy.DummyDataBag{}
	for _, message := range oldData {
		oldBag.AddEntry(message)
	}

	newBag := dummy.DummyDataBag{}
	for _, message := range newData {
		newBag.AddEntry(message)
	}
	delta := getDiffBetweenModels(&oldBag, &newBag, dummy.DummyDataRef)

	assert.Equal(t, delta.NewObjects.GetEntries()[0].GetPrimaryID(), "3")
	assert.Equal(t, delta.ChangedObjects.GetEntries()[0].GetPrimaryID(), "2")
	assert.Equal(t, delta.MissingKeys[0], "4")
}
