package grpc

import (
	"context"
	"strings"

	"github.com/karolhrdina/misc/hw/model/ports"
	pb "github.com/karolhrdina/misc/hw/pb.go"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// PortGrpcServer implements gRPC `Portdomain` service
type Server struct {
	pb.UnimplementedPortdomainServer
	storer ports.Storer
}

func New(storer ports.Storer) *Server {
	return &Server{
		storer: storer,
	}
}

func (s *Server) Snapshot(ctx context.Context, in *pb.Port) (*emptypb.Empty, error) {
	// input checks
	switch {
	case strings.TrimSpace(in.Id) == "":
		// Note: could be improved with errdetails
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "Message field 'id' must not be empty.")
	}

	err := s.storer.Store(ctx, in)
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "Storing message failed: %s", err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) List(in *emptypb.Empty, stream pb.Portdomain_ListServer) error {

	ctx := stream.Context()

	ports, err := s.storer.List(ctx)
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
