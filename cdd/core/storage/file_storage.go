package storage

import (
	"io/ioutil"
	"os"
)

const (
	BaseSnapshotPath = "/Users/sourabh.suman/gopath/src/gollow/cdd/snapshots/"
)

type FileStorage struct {
	fullPath string
	fileName string
}

func NewFileStorage(fileName string) *FileStorage {
	return &FileStorage{
		fullPath: BaseSnapshotPath + fileName,
		fileName: fileName,
	}
}

func (f *FileStorage) IsExist() bool {
	if _, err := os.Stat(f.fullPath); os.IsExist(err) {
		return true
	}
	return false
}

func (f *FileStorage) Write(data []byte) (int, error) {

	file, err := os.Create(f.fullPath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.Write(data)
}

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
