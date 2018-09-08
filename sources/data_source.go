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

package sources

import (
	"github.com/golang/protobuf/proto"
	"github.com/gollow/data"
)

//Message represents a single entity of the data
type Message interface {
	proto.Message

	//GetUniqueKey returns the primaryID of the string
	GetUniqueKey() string
}

//Bag interface represents a collection of message
type Bag interface {
	proto.Message

	//AddEntry provides method to add MESSAGE TO Bag
	AddEntry(Message)

	//GetEntries returns list of all messages in the Bag
	GetEntries() []Message

	//NewBag returns a newBag of Message
	NewBag() Bag
}

//DTO represents the interface every interface needs to implement
type DTO interface {
	data.Entity
	//ToPB provides methods to convert DTO to proto-buf Message
	ToPB() Message
}

//DataModel represents the datamodel being produced
type DataModel interface {

	//NewBag returns a newBag of Message
	NewBag() Bag

	//CacheDuration provides cache duration in time.NanoSeconds
	CacheDuration() int64

	//LoadAll provides the method to load all the data
	LoadAll() (Bag, error)

	//GetDataName provides the data-name and should be unique
	GetDataName() string
}
