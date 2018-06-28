package write

type SnapshotWriter interface {

	//IsExist checks whether the file exists
	IsExist() bool

	//Write writes the given byte array
	Write(data []byte) (int, error)
}

type SnapshotReader interface {

	//Read reads the given snapshot
	Read() ([]byte, error)
}

func WriteSnapshot(writer SnapshotWriter, data []byte) (int, error) {

	if ok := writer.IsExist(); ok {
		return 0, ErrFileAlreadyExists
	}

	return writer.Write(data)
}

func ReadSnapshot(reader SnapshotReader) ([]byte, error) {

	return reader.Read()
}

