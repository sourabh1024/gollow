package write

import (
	"io/ioutil"
	"os"
)

type FileWriter struct {
	FilePath string
	FileName string
}

func (f *FileWriter) GetFullPath() string {
	return f.FilePath + f.FileName
}

func (f *FileWriter) IsExist(path string) bool {
	if _, err := os.Stat(path); os.IsExist(err) {
		return true
	}
	return false
}

func (f *FileWriter) Write(path string, data []byte) (int, error) {

	file, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.Write(data)
}

func (f *FileWriter) Read(path string) ([]byte, error) {

	file, err := os.Open(path)
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
