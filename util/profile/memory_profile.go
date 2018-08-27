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

package profile

import (
	"gollow/logging"
	"runtime"
)

func GetMemoryProfile() {

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	logging.GetLogger().Info("------------------------------------------")
	logging.GetLogger().Info("mem alloc : ", mem.Alloc)
	logging.GetLogger().Info("mem total alloc : ", mem.TotalAlloc)
	logging.GetLogger().Info("mem heap alloc : ", mem.HeapAlloc)
	logging.GetLogger().Info("mem heap size", mem.HeapSys)
	logging.GetLogger().Info("------------------------------------------")

}
