package ports

import (
	"context"

	pb "github.com/karolhrdina/misc/hw/pb.go"
)

// Storer is the interface that wraps port operations,
// i.e. storing and listing of ports.
//
// Store must either store new port or update an existing one if port
// with given `Id` already exists.
// Store must not modify `port`.
//
// List must return all stored ports.
type Storer interface {
	Store(ctx context.Context, port *pb.Port) error
	List(ctx context.Context) ([]*pb.Port, error)
}
