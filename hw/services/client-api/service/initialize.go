// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package service

import (
	"github.com/karolhrdina/misc/hw/model/ports"
	pb "github.com/karolhrdina/misc/hw/pb.go"
	"github.com/karolhrdina/misc/hw/pkg/env"
	"github.com/karolhrdina/misc/hw/pkg/storer"
	"github.com/karolhrdina/misc/hw/services/client-api/handlers"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

func initializeService() (*Service, error) {
	wire.Build(
		providePortDomainClient,
		handlers.NewV1PortsImportAPI,
		NewService,
		provideStorer,
	)
	return &Service{}, nil
}

func providePortDomainClient() pb.PortdomainClient {
	addr := env.GetVar("PORTDOMAIN_GRPC_ADDR", ":9082")
	c, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return pb.NewPortdomainClient(c)
}

func provideStorer(client pb.PortdomainClient) ports.Storer {
	return storer.NewGRPCStorer(client)
}
