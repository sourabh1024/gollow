package storage

import "errors"

var (
	ErrFileAlreadyExists = errors.New(" File Path already exists. Can't override")

	ErrFileNotFound = errors.New("File not found. ")
)