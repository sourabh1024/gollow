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

// NewStorage gives the current implementation of Storage
// returns storage object if initialised properly
// throws error if storage cannot be initialised
// Default it returns file storage
func NewStorage(key string) (Storage, error) {
	switch storage := config.GlobalConfig.Storage.StorageType; storage {
	case "file":
		return NewFileStorage(key)
	case "s3":
		return NewS3Storage(&Config{
			Region: config.GlobalConfig.Storage.AWSRegion,
			Bucket: config.GlobalConfig.Storage.S3Bucket,
			Key:    key,
		})
	default:
		return NewFileStorage(key)
	}
}
