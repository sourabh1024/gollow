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

package client_api

import (
	"context"
	"gollow/gollow-demo/cache/client_datamodel"
	"gollow/logging"
	"gollow/sources/datamodel/dummy"
)

// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	logging.GetLogger().Info("Receive message %s", in.Greeting)
	return &PingMessage{Greeting: "bar"}, nil
}

func (s *Server) GetDummyData(ctx context.Context, in *DummyDataRequest) (*DummyDataResponse, error) {
	logging.GetLogger().Info("Received request GetDummyData for %s", in.Keyname)
	response := &DummyDataResponse{}
	val, err := client_datamodel.DummyDataCache.Get(in.Keyname)
	if err != nil {
		logging.GetLogger().Error("error in getting data : %+v", err)
		return response, nil
	}

	dummyData, ok := val.(*dummy.DummyData)
	if !ok {
		logging.GetLogger().Error("error in typecasting dummydata data : %+v", err)
		return response, nil
	}

	response.Id = dummyData.ID
	response.Firstname = dummyData.FirstName
	response.Balance = dummyData.Balance

	return response, nil
}
