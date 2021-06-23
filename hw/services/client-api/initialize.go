// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	pb "github.com/karolhrdina/misc/hw/pb.go"
	"github.com/karolhrdina/misc/hw/pkg/env"
	"google.golang.org/grpc"

	"github.com/google/wire"
)

func initializeService() (*Service, error) {
	wire.Build(
		providePortDomainClient,
		NewV1PortsImportAPI,
		NewService,
	)
	return &Service{}, nil
}

func providePortDomainClient() (pb.PortdomainClient, error) {
	addr := env.GetVar("PORTDOMAIN_GRPC_ADDR", ":9082")
	c, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return pb.NewPortdomainClient(c), nil
}
