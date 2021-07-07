// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package service

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/karolhrdina/misc/hw/model/ports"
	hwDB "github.com/karolhrdina/misc/hw/pkg/db"
	"github.com/karolhrdina/misc/hw/pkg/storer"
)

func initializeService() (*Service, error) {
	wire.Build(
		hwDB.ProvidePG,
		NewService,
		provideStorer,
	)
	return &Service{}, nil
}

func provideStorer(db *sql.DB) ports.Storer {
	return storer.NewPostgresStorer(db)
}
