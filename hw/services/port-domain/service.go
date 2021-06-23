package main

import (
	"database/sql"

	pb "github.com/karolhrdina/misc/hw/pb.go"
	"github.com/karolhrdina/misc/hw/services/port-domain/storage"

	"google.golang.org/grpc"
)

// Service is a container to hold all dependencies of port-domain app.
type Service struct {
	DB *sql.DB
}

// NewService returns new Service struct
func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
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
	pb.RegisterPortdomainServer(server, &portGrpcServer{
		store: storage.New(s.DB),
	})
	return server, nil
}

// Close closes all dependencies of the service.
func (c *Service) Close() error {
	return nil
}
