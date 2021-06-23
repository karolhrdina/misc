// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb_go

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PortdomainClient is the client API for Portdomain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortdomainClient interface {
	Snapshot(ctx context.Context, in *Port, opts ...grpc.CallOption) (*emptypb.Empty, error)
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Portdomain_ListClient, error)
}

type portdomainClient struct {
	cc grpc.ClientConnInterface
}

func NewPortdomainClient(cc grpc.ClientConnInterface) PortdomainClient {
	return &portdomainClient{cc}
}

func (c *portdomainClient) Snapshot(ctx context.Context, in *Port, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/portdomain.Portdomain/Snapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portdomainClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Portdomain_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Portdomain_ServiceDesc.Streams[0], "/portdomain.Portdomain/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &portdomainListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Portdomain_ListClient interface {
	Recv() (*Port, error)
	grpc.ClientStream
}

type portdomainListClient struct {
	grpc.ClientStream
}

func (x *portdomainListClient) Recv() (*Port, error) {
	m := new(Port)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PortdomainServer is the server API for Portdomain service.
// All implementations must embed UnimplementedPortdomainServer
// for forward compatibility
type PortdomainServer interface {
	Snapshot(context.Context, *Port) (*emptypb.Empty, error)
	List(*emptypb.Empty, Portdomain_ListServer) error
	mustEmbedUnimplementedPortdomainServer()
}

// UnimplementedPortdomainServer must be embedded to have forward compatible implementations.
type UnimplementedPortdomainServer struct {
}

func (UnimplementedPortdomainServer) Snapshot(context.Context, *Port) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Snapshot not implemented")
}
func (UnimplementedPortdomainServer) List(*emptypb.Empty, Portdomain_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPortdomainServer) mustEmbedUnimplementedPortdomainServer() {}

// UnsafePortdomainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortdomainServer will
// result in compilation errors.
type UnsafePortdomainServer interface {
	mustEmbedUnimplementedPortdomainServer()
}

func RegisterPortdomainServer(s grpc.ServiceRegistrar, srv PortdomainServer) {
	s.RegisterService(&Portdomain_ServiceDesc, srv)
}

func _Portdomain_Snapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Port)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortdomainServer).Snapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portdomain.Portdomain/Snapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortdomainServer).Snapshot(ctx, req.(*Port))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portdomain_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PortdomainServer).List(m, &portdomainListServer{stream})
}

type Portdomain_ListServer interface {
	Send(*Port) error
	grpc.ServerStream
}

type portdomainListServer struct {
	grpc.ServerStream
}

func (x *portdomainListServer) Send(m *Port) error {
	return x.ServerStream.SendMsg(m)
}

// Portdomain_ServiceDesc is the grpc.ServiceDesc for Portdomain service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Portdomain_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "portdomain.Portdomain",
	HandlerType: (*PortdomainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Snapshot",
			Handler:    _Portdomain_Snapshot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Portdomain_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "port-domain.proto",
}