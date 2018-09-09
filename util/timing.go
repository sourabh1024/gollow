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

package util

import (
	"fmt"
	"github.com/sourabh1024/gollow/logging"
	"time"
)

// Duration prints the time elasped from the invocation time
func Duration(invocation time.Time, name string) {
	elapsed := time.Since(invocation)
	logging.GetLogger().Info("%s lasted %s", name, elapsed)
}

// GetCurrentTimeString returns current time stamp in milli seconds in string format
func GetCurrentTimeString() string {
	return fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond))
}
