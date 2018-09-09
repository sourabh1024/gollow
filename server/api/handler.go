package api

import (
	"github.com/sourabh1024/gollow/core/snapshot"
	"github.com/sourabh1024/gollow/logging"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	logging.GetLogger().Info("Receive message %s", in.Greeting)
	return &PingMessage{Greeting: "bar"}, nil
}

// GetAnnouncedVersion generates response to a Ping request
func (s *Server) GetAnnouncedVersion(ctx context.Context, in *AnnouncedVersionRequest) (*AnnouncedVersionResponse, error) {
	logging.GetLogger().Info("Receive message %s", in.Namespace)
	version, err := snapshot.VersionImpl.GetVersion(in.Namespace)
	if err != nil {
		logging.GetLogger().Error("Error in getting announced version : ", err)
		return &AnnouncedVersionResponse{}, err
	}
	return &AnnouncedVersionResponse{Currentversion: version}, nil
}
