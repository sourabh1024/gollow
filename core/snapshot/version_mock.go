package snapshot

import "github.com/stretchr/testify/mock"

// MockVersion mocks the Version interface implementation
type MockVersion struct {
	mock.Mock
}

// GetVersion implements the Version interface
func (m *MockVersion) GetVersion(keyName string) (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

// UpdateVersion implements the Version interface
func (m *MockVersion) UpdateVersion(keyName, newVersion string) error {
	args := m.Called()
	return args.Error(0)
}

// ParseVersionNumber implements the Version interface
func (m *MockVersion) ParseVersionNumber(fileName string) (int64, error) {
	args := m.Called()
	return int64(args.Int(0)), args.Error(1)
}
