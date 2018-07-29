package client_api

import (
	"context"
	"gollow/cdd/logging"
	"gollow/cdd/sources/datamodel/dummy"
	"gollow/gollow-client/cache/client_datamodel"
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
