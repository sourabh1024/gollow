package storage

import (
	"gollow/cdd/config"
	"io/ioutil"
	"os"
)

// FileStorage implements the Storage interface
type FileStorage struct {
	fullPath string
	fileName string
}

// NewFileStorage returns a new file storage with fullPath
func NewFileStorage(fileName string) *FileStorage {
	return &FileStorage{
		fullPath: config.GlobalConfig.Storage.BaseSnapshotPath + fileName,
		fileName: fileName,
	}
}

// IsExist implements storage interface
// returns whether the file already exists or not
func (f *FileStorage) IsExist() bool {
	if _, err := os.Stat(f.fullPath); os.IsExist(err) {
		return true
	}
	return false
}

// Write implements the storage interface
// Writes the given data bytes at the fileLocation
func (f *FileStorage) Write(data []byte) (int, error) {

	file, err := os.Create(f.fullPath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.Write(data)
}

// Read implements the storage interface
// Reads the bytes from the file storage
func (f *FileStorage) Read() ([]byte, error) {
	file, err := os.Open(f.fullPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var data []byte
	data, err = ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return data, nil
}
