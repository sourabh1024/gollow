package storage

import (
	"github.com/stretchr/testify/mock"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) IsExist() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockStorage) Write(data []byte) (int, error) {
	args := m.Called()

	return args.Int(0), args.Error(1)
}

func (m *MockStorage) Read() ([]byte, error) {
	args := m.Called()

	var r0 []byte
	if rf, ok := args.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).([]byte)
		}
	}

	var r1 error

	if rf, ok := args.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
