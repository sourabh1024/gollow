//Copyright 2018 Sourabh Suman ( https://github.com/sourabh1024 )
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

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
