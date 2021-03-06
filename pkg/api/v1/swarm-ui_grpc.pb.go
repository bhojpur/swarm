// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SwarmUIClient is the client API for SwarmUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SwarmUIClient interface {
	// ListEngineSpecs returns a list of Swarm Engine(s) that can be started through the UI.
	ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (SwarmUI_ListEngineSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type swarmUIClient struct {
	cc grpc.ClientConnInterface
}

func NewSwarmUIClient(cc grpc.ClientConnInterface) SwarmUIClient {
	return &swarmUIClient{cc}
}

func (c *swarmUIClient) ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (SwarmUI_ListEngineSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SwarmUI_ServiceDesc.Streams[0], "/v1.SwarmUI/ListEngineSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &swarmUIListEngineSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SwarmUI_ListEngineSpecsClient interface {
	Recv() (*ListEngineSpecsResponse, error)
	grpc.ClientStream
}

type swarmUIListEngineSpecsClient struct {
	grpc.ClientStream
}

func (x *swarmUIListEngineSpecsClient) Recv() (*ListEngineSpecsResponse, error) {
	m := new(ListEngineSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *swarmUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.SwarmUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SwarmUIServer is the server API for SwarmUI service.
// All implementations must embed UnimplementedSwarmUIServer
// for forward compatibility
type SwarmUIServer interface {
	// ListEngineSpecs returns a list of Swarm Engine(s) that can be started through the UI.
	ListEngineSpecs(*ListEngineSpecsRequest, SwarmUI_ListEngineSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedSwarmUIServer()
}

// UnimplementedSwarmUIServer must be embedded to have forward compatible implementations.
type UnimplementedSwarmUIServer struct {
}

func (UnimplementedSwarmUIServer) ListEngineSpecs(*ListEngineSpecsRequest, SwarmUI_ListEngineSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEngineSpecs not implemented")
}
func (UnimplementedSwarmUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedSwarmUIServer) mustEmbedUnimplementedSwarmUIServer() {}

// UnsafeSwarmUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SwarmUIServer will
// result in compilation errors.
type UnsafeSwarmUIServer interface {
	mustEmbedUnimplementedSwarmUIServer()
}

func RegisterSwarmUIServer(s grpc.ServiceRegistrar, srv SwarmUIServer) {
	s.RegisterService(&SwarmUI_ServiceDesc, srv)
}

func _SwarmUI_ListEngineSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEngineSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SwarmUIServer).ListEngineSpecs(m, &swarmUIListEngineSpecsServer{stream})
}

type SwarmUI_ListEngineSpecsServer interface {
	Send(*ListEngineSpecsResponse) error
	grpc.ServerStream
}

type swarmUIListEngineSpecsServer struct {
	grpc.ServerStream
}

func (x *swarmUIListEngineSpecsServer) Send(m *ListEngineSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SwarmUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwarmUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.SwarmUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwarmUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SwarmUI_ServiceDesc is the grpc.ServiceDesc for SwarmUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SwarmUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.SwarmUI",
	HandlerType: (*SwarmUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _SwarmUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListEngineSpecs",
			Handler:       _SwarmUI_ListEngineSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "swarm-ui.proto",
}
