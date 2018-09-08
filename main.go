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

package main

import (
	"context"
	"github.com/gollow/server"
	"github.com/gollow/core/snapshot"
	"github.com/gollow/config"
	"github.com/gollow/producer"
	"github.com/gollow/sources/datamodel/dummy"
)

func main() {
	//Register Data Model to be produced
	RegisterDataModel()
	//initialise all you want
	Init(context.Background())
	//initialise server
	server.Init()
}

// Init initialises everything here
// should initialise which storage to use
// schedule the producers
func Init(ctx context.Context) {
	//initialise everything here

	snapshot.InitVersionStorage(config.GlobalConfig.Storage.AnnouncedVersion)

	go producer.ScheduleProducers()
}

// RegisterDataModel registers the data model for production
func RegisterDataModel() {
	producer.Register(&dummy.DummyData{}, struct{}{})
}
