package storage

//Storage interface provides all methods required for read/write/checkExistence
//of any snapshot or diff
type Storage interface {
	IsExist() bool

	Write(data []byte) (int, error)

	Read() ([]byte, error)
}

//GetStorage gives the current implementation of Storage
//In future when S3 support is added it should be returned from here
func NewStorage(fileName string) Storage {
	return NewFileStorage(fileName)
}
