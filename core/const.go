// Package core provides methods which are core to producer and consumer
package core

const (

	//DefaultVersionNumber for the snapshot
	//When the snapshot is produced for the first time it starts with DefaultVersionNumber
	DefaultVersionNumber = 1

	//Separator used for Snapshot and diff name generated
	Separator = "-"

	//DiffPrefix used for generating the diff
	DiffPrefix = "diff-"
)
