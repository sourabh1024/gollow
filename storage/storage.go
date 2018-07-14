package storage

type Storage interface {
	IsExist() bool

	Write(data []byte) (int, error)

	Read() ([]byte, error)
}
