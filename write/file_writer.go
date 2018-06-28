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

func (f *FileWriter) IsExist() bool {
	if _, err := os.Stat(f.GetFullPath()); os.IsExist(err) {
		return true
	}
	return false
}

func (f *FileWriter) Write(data []byte) (int, error) {

	file, err := os.Create(f.GetFullPath())
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.Write(data)
}

func (f *FileWriter) Read() ([]byte, error) {

	file, err := os.Open(f.GetFullPath())
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
