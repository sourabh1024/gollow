package write

import "errors"

var (
	ErrFileAlreadyExists = errors.New(" File Path already exists. Can't override")
)
