// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	hwDB "github.com/karolhrdina/misc/hw/pkg/db"
)

func initializeService() (*Service, error) {
	wire.Build(
		hwDB.ProvidePG,
		NewService,
	)
	return &Service{}, nil
}
