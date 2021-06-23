package main

import (
	"context"
	"strings"

	pb "github.com/karolhrdina/misc/hw/pb.go"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// PortStorer is the interface that wraps persistence of ports,
// i.e. storing and listing of stored ports.
//
// Store must either persist new port or update an existing one if port
// with given `Id` already exists.
// Store must not modify `port`.
//
// List must return all stored ports.
type PortStorer interface {
	Store(ctx context.Context, port *pb.Port) error
	List(ctx context.Context) ([]*pb.Port, error)
}

// portGrpcServer implements gRPC `Portdomain` service
type portGrpcServer struct {
	pb.UnimplementedPortdomainServer
	store PortStorer
}

func (s *portGrpcServer) Snapshot(ctx context.Context, in *pb.Port) (*emptypb.Empty, error) {
	// input checks
	switch {
	case strings.TrimSpace(in.Id) == "":
		// Note: could be improved with errdetails
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "Message field 'id' must not be empty.")
	}

	err := s.store.Store(ctx, in)
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "Storing message failed: %s", err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *portGrpcServer) List(in *emptypb.Empty, stream pb.Portdomain_ListServer) error {

	ctx := stream.Context()

	ports, err := s.store.List(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Storing message failed: %s", err.Error())
	}

	for _, cpy := range ports {
		err = stream.Send(cpy)
		if err != nil {
			return status.Errorf(codes.Internal, "Error sending response: %s", err.Error())
		}
	}
	return nil
}

// Note:
// Sometimes it's desirable with gRPC rpc calls, to distinguish between type of error,
// especially context cancelation & deadline exceeded. This is how it can be done:
//      switch {
//      case errors.Is(ctx.Err(), context.Canceled):
//          ... handle context Canceled ...
//
//      case errors.Is(ctx.Err(), context.DeadlineExceeded):
//          ...
//
//      case err != nil:
//          ...
//      }
