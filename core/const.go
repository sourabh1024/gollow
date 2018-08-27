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

// Package core provides methods which are core to producer and consumer
package core

const (

	//DefaultVersionNumber for the snapshot
	//When the snapshot is produced for the first time it starts with DefaultVersionNumber
	DefaultVersionNumber = 1

	//Separator used for Snapshot and diff name generated
	Separator = "-"

	//DiffPrefix used for generating the diff
	DiffPrefix = "diff"
)
