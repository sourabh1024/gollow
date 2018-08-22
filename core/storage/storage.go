// Package storage provides methods related to storage handling
package storage

import (
	"gollow/config"
)

//Storage interface provides all methods required for read/write/checkExistence
//of any snapshot or diff
type Storage interface {
	IsExist() bool

	Write(data []byte) (int, error)

	Read() ([]byte, error)
}

//NewStorage gives the current implementation of Storage
//In future when S3/ other buckets support is added it should be returned from here
//Default it returns file storage
func NewStorage(fileName string) Storage {
	switch storage := config.GlobalConfig.Storage.StorageType; storage {
	case "file":
		return NewFileStorage(fileName)

	default:
		return NewFileStorage(fileName)
	}
}
