package snapshot

import "github.com/stretchr/testify/mock"

type MockVersion struct {
	mock.Mock
}

func (m *MockVersion) GetVersion(keyName string) (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockVersion) UpdateVersion(keyName, newVersion string) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockVersion) ParseVersionNumber(fileName string) (int64, error) {
	args := m.Called()
	return int64(args.Int(0)), args.Error(1)
}
