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

package storage

import (
	"github.com/stretchr/testify/mock"
)

//MockStorage provides mock Storage implementation
type MockStorage struct {
	mock.Mock
}

// IsExist implements the Storage interface
func (m *MockStorage) IsExist() bool {
	args := m.Called()
	return args.Bool(0)
}

// Write implements the Storage interface
func (m *MockStorage) Write(data []byte) (int, error) {
	args := m.Called()

	return args.Int(0), args.Error(1)
}

// Read implements the storage interface
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
