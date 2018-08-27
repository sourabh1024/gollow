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

package data

import (
	"errors"
)

// all the errors data source can return
var (
	// ErrNoData : no data is found for the query
	ErrNoData = errors.New("no data in result set")

	// ErrTimedOut : response timed out from data source
	ErrTimedOut = errors.New("connection to the data source has timed out")

	// ErrConvert : error when converting interface generally during typecasting
	ErrConvert = errors.New("Error converting interface")
)
