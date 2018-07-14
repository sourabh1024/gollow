package api

import (
	"golang.org/x/net/context"
	"gollow/interactors/announced_versions"
	"gollow/logging"
	"log"
)

// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &PingMessage{Greeting: "bar"}, nil
}

// SayHello generates response to a Ping request
func (s *Server) GetAnnouncedVersion(ctx context.Context, in *AnnouncedVersionRequest) (*AnnouncedVersionResponse, error) {
	version, err := announced_versions.GetAnnouncedVersions(ctx, in.Namespace, in.Entity)
	if err != nil {
		logging.GetLogger().Error("Error in getting announced version : ", err)
		return &AnnouncedVersionResponse{}, err
	}
	return &AnnouncedVersionResponse{Currentversion: version}, nil
}
