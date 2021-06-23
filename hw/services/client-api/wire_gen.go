// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/karolhrdina/misc/hw/pb.go"
	"github.com/karolhrdina/misc/hw/pkg/env"
	"google.golang.org/grpc"
)

// Injectors from initialize.go:

func initializeService() (*Service, error) {
	portdomainClient, err := providePortDomainClient()
	if err != nil {
		return nil, err
	}
	mainV1PortsImport := NewV1PortsImportAPI(portdomainClient)
	service := NewService(mainV1PortsImport)
	return service, nil
}

// initialize.go:

func providePortDomainClient() (pb_go.PortdomainClient, error) {
	addr := env.GetVar("PORTDOMAIN_GRPC_ADDR", ":9082")
	c, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return pb_go.NewPortdomainClient(c), nil
}
