package storer

import (
	"context"
	"io"

	pb "github.com/karolhrdina/misc/hw/pb.go"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCStorer struct {
	c pb.PortdomainClient
}

func NewGRPCStorer(client pb.PortdomainClient) *GRPCStorer {
	return &GRPCStorer{c: client}
}

func (g *GRPCStorer) Store(ctx context.Context, port *pb.Port) error {
	_, err := g.c.Snapshot(ctx, port)
	return err
}

func (g *GRPCStorer) List(ctx context.Context) ([]*pb.Port, error) {
	stream, err := g.c.List(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	rv := make([]*pb.Port, 0)
	for {
		port, err := stream.Recv()
		if err == io.EOF {
			break
		}
		// handle error
		if err != nil {
			return nil, err
		}

		rv = append(rv, port)
	}
	return rv, nil
}
