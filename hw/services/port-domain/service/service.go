package service

import (
	"database/sql"

	"github.com/karolhrdina/misc/hw/model/ports"
	pb "github.com/karolhrdina/misc/hw/pb.go"
	"github.com/karolhrdina/misc/hw/pkg/storer"
	pdGrpc "github.com/karolhrdina/misc/hw/services/port-domain/grpc"

	"google.golang.org/grpc"
)

// Service is a container to hold all dependencies of port-domain app.
type Service struct {
	DB     *sql.DB
	storer ports.Storer
}

// InitializeServices provides new Service injected with dependencies
func InitializeService() (*Service, error) {
	return initializeService()
}

// NewService returns new Service struct
func NewService(db *sql.DB, storer ports.Storer) *Service {
	return &Service{
		DB:     db,
		storer: storer,
	}
}

// GRPC returns initialized grpc server.
func (s *Service) GRPC() (*grpc.Server, error) {
	// Note:
	// For real prod environments, I'd add:
	//  * grpc middleware for logger, stats support, tracing support
	//  * unary and stream server interceptor
	//  * grpc health server

	server := grpc.NewServer()
	portGRPCServer := pdGrpc.New(storer.NewPostgresStorer(s.DB))
	pb.RegisterPortdomainServer(server, portGRPCServer)
	return server, nil
}

// Close closes all dependencies of the service.
func (c *Service) Close() error {
	return nil
}
