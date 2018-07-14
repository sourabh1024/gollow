package storage

type SnapshotWriter interface {

	//IsExist checks whether the file exists
	IsExist(path string) bool

	//Write writes the given byte array
	Write(path string, data []byte) (int, error)
}

type SnapshotReader interface {

	//Read reads the given snapshot
	Read(path string) ([]byte, error)
}

type SnapshotReaderWriter interface {
	SnapshotReader
	SnapshotWriter
}

func WriteSnapshot(writer SnapshotWriter, path string, data []byte) (int, error) {

	if ok := writer.IsExist(path); ok {
		return 0, ErrFileAlreadyExists
	}

	return writer.Write(path, data)
}

func ReadSnapshot(reader SnapshotReader, path string) ([]byte, error) {

	return reader.Read(path)
}
