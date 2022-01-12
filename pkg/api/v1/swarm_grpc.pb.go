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

// SwarmServiceClient is the client API for SwarmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SwarmServiceClient interface {
	// StartLocalEngine starts a Swarm Engine on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the swarm/config.yaml
	//   3. all bytes constituting the Engine YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalEngine(ctx context.Context, opts ...grpc.CallOption) (SwarmService_StartLocalEngineClient, error)
	// StartFromPreviousEngine starts a new Engine based on a previous one.
	// If the previous Engine does not have the can-replay condition set this call will result in an error.
	StartFromPreviousEngine(ctx context.Context, in *StartFromPreviousEngineRequest, opts ...grpc.CallOption) (*StartEngineResponse, error)
	// StartEngineRequest starts a new Engine based on its specification.
	StartEngine(ctx context.Context, in *StartEngineRequest, opts ...grpc.CallOption) (*StartEngineResponse, error)
	// Searches for Engine(s) known to this Engine
	ListEngines(ctx context.Context, in *ListEnginesRequest, opts ...grpc.CallOption) (*ListEnginesResponse, error)
	// Subscribe listens to new Engine(s) updates
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (SwarmService_SubscribeClient, error)
	// GetEngine retrieves details of a single Engine
	GetEngine(ctx context.Context, in *GetEngineRequest, opts ...grpc.CallOption) (*GetEngineResponse, error)
	// Listen listens to Engine updates and log output of a running Engine
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (SwarmService_ListenClient, error)
	// StopEngine stops a currently running Engine
	StopEngine(ctx context.Context, in *StopEngineRequest, opts ...grpc.CallOption) (*StopEngineResponse, error)
}

type swarmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSwarmServiceClient(cc grpc.ClientConnInterface) SwarmServiceClient {
	return &swarmServiceClient{cc}
}

func (c *swarmServiceClient) StartLocalEngine(ctx context.Context, opts ...grpc.CallOption) (SwarmService_StartLocalEngineClient, error) {
	stream, err := c.cc.NewStream(ctx, &SwarmService_ServiceDesc.Streams[0], "/v1.SwarmService/StartLocalEngine", opts...)
	if err != nil {
		return nil, err
	}
	x := &swarmServiceStartLocalEngineClient{stream}
	return x, nil
}

type SwarmService_StartLocalEngineClient interface {
	Send(*StartLocalEngineRequest) error
	CloseAndRecv() (*StartEngineResponse, error)
	grpc.ClientStream
}

type swarmServiceStartLocalEngineClient struct {
	grpc.ClientStream
}

func (x *swarmServiceStartLocalEngineClient) Send(m *StartLocalEngineRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *swarmServiceStartLocalEngineClient) CloseAndRecv() (*StartEngineResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StartEngineResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *swarmServiceClient) StartFromPreviousEngine(ctx context.Context, in *StartFromPreviousEngineRequest, opts ...grpc.CallOption) (*StartEngineResponse, error) {
	out := new(StartEngineResponse)
	err := c.cc.Invoke(ctx, "/v1.SwarmService/StartFromPreviousEngine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swarmServiceClient) StartEngine(ctx context.Context, in *StartEngineRequest, opts ...grpc.CallOption) (*StartEngineResponse, error) {
	out := new(StartEngineResponse)
	err := c.cc.Invoke(ctx, "/v1.SwarmService/StartEngine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swarmServiceClient) ListEngines(ctx context.Context, in *ListEnginesRequest, opts ...grpc.CallOption) (*ListEnginesResponse, error) {
	out := new(ListEnginesResponse)
	err := c.cc.Invoke(ctx, "/v1.SwarmService/ListEngines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swarmServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (SwarmService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &SwarmService_ServiceDesc.Streams[1], "/v1.SwarmService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &swarmServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SwarmService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type swarmServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *swarmServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *swarmServiceClient) GetEngine(ctx context.Context, in *GetEngineRequest, opts ...grpc.CallOption) (*GetEngineResponse, error) {
	out := new(GetEngineResponse)
	err := c.cc.Invoke(ctx, "/v1.SwarmService/GetEngine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swarmServiceClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (SwarmService_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &SwarmService_ServiceDesc.Streams[2], "/v1.SwarmService/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &swarmServiceListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SwarmService_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type swarmServiceListenClient struct {
	grpc.ClientStream
}

func (x *swarmServiceListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *swarmServiceClient) StopEngine(ctx context.Context, in *StopEngineRequest, opts ...grpc.CallOption) (*StopEngineResponse, error) {
	out := new(StopEngineResponse)
	err := c.cc.Invoke(ctx, "/v1.SwarmService/StopEngine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SwarmServiceServer is the server API for SwarmService service.
// All implementations must embed UnimplementedSwarmServiceServer
// for forward compatibility
type SwarmServiceServer interface {
	// StartLocalEngine starts a Swarm Engine on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the swarm/config.yaml
	//   3. all bytes constituting the Engine YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalEngine(SwarmService_StartLocalEngineServer) error
	// StartFromPreviousEngine starts a new Engine based on a previous one.
	// If the previous Engine does not have the can-replay condition set this call will result in an error.
	StartFromPreviousEngine(context.Context, *StartFromPreviousEngineRequest) (*StartEngineResponse, error)
	// StartEngineRequest starts a new Engine based on its specification.
	StartEngine(context.Context, *StartEngineRequest) (*StartEngineResponse, error)
	// Searches for Engine(s) known to this Engine
	ListEngines(context.Context, *ListEnginesRequest) (*ListEnginesResponse, error)
	// Subscribe listens to new Engine(s) updates
	Subscribe(*SubscribeRequest, SwarmService_SubscribeServer) error
	// GetEngine retrieves details of a single Engine
	GetEngine(context.Context, *GetEngineRequest) (*GetEngineResponse, error)
	// Listen listens to Engine updates and log output of a running Engine
	Listen(*ListenRequest, SwarmService_ListenServer) error
	// StopEngine stops a currently running Engine
	StopEngine(context.Context, *StopEngineRequest) (*StopEngineResponse, error)
	mustEmbedUnimplementedSwarmServiceServer()
}

// UnimplementedSwarmServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSwarmServiceServer struct {
}

func (UnimplementedSwarmServiceServer) StartLocalEngine(SwarmService_StartLocalEngineServer) error {
	return status.Errorf(codes.Unimplemented, "method StartLocalEngine not implemented")
}
func (UnimplementedSwarmServiceServer) StartFromPreviousEngine(context.Context, *StartFromPreviousEngineRequest) (*StartEngineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFromPreviousEngine not implemented")
}
func (UnimplementedSwarmServiceServer) StartEngine(context.Context, *StartEngineRequest) (*StartEngineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartEngine not implemented")
}
func (UnimplementedSwarmServiceServer) ListEngines(context.Context, *ListEnginesRequest) (*ListEnginesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEngines not implemented")
}
func (UnimplementedSwarmServiceServer) Subscribe(*SubscribeRequest, SwarmService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedSwarmServiceServer) GetEngine(context.Context, *GetEngineRequest) (*GetEngineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEngine not implemented")
}
func (UnimplementedSwarmServiceServer) Listen(*ListenRequest, SwarmService_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedSwarmServiceServer) StopEngine(context.Context, *StopEngineRequest) (*StopEngineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopEngine not implemented")
}
func (UnimplementedSwarmServiceServer) mustEmbedUnimplementedSwarmServiceServer() {}

// UnsafeSwarmServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SwarmServiceServer will
// result in compilation errors.
type UnsafeSwarmServiceServer interface {
	mustEmbedUnimplementedSwarmServiceServer()
}

func RegisterSwarmServiceServer(s grpc.ServiceRegistrar, srv SwarmServiceServer) {
	s.RegisterService(&SwarmService_ServiceDesc, srv)
}

func _SwarmService_StartLocalEngine_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SwarmServiceServer).StartLocalEngine(&swarmServiceStartLocalEngineServer{stream})
}

type SwarmService_StartLocalEngineServer interface {
	SendAndClose(*StartEngineResponse) error
	Recv() (*StartLocalEngineRequest, error)
	grpc.ServerStream
}

type swarmServiceStartLocalEngineServer struct {
	grpc.ServerStream
}

func (x *swarmServiceStartLocalEngineServer) SendAndClose(m *StartEngineResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *swarmServiceStartLocalEngineServer) Recv() (*StartLocalEngineRequest, error) {
	m := new(StartLocalEngineRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SwarmService_StartFromPreviousEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFromPreviousEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwarmServiceServer).StartFromPreviousEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.SwarmService/StartFromPreviousEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwarmServiceServer).StartFromPreviousEngine(ctx, req.(*StartFromPreviousEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwarmService_StartEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwarmServiceServer).StartEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.SwarmService/StartEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwarmServiceServer).StartEngine(ctx, req.(*StartEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwarmService_ListEngines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnginesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwarmServiceServer).ListEngines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.SwarmService/ListEngines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwarmServiceServer).ListEngines(ctx, req.(*ListEnginesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwarmService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SwarmServiceServer).Subscribe(m, &swarmServiceSubscribeServer{stream})
}

type SwarmService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type swarmServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *swarmServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SwarmService_GetEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwarmServiceServer).GetEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.SwarmService/GetEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwarmServiceServer).GetEngine(ctx, req.(*GetEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwarmService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SwarmServiceServer).Listen(m, &swarmServiceListenServer{stream})
}

type SwarmService_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type swarmServiceListenServer struct {
	grpc.ServerStream
}

func (x *swarmServiceListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SwarmService_StopEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwarmServiceServer).StopEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.SwarmService/StopEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwarmServiceServer).StopEngine(ctx, req.(*StopEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SwarmService_ServiceDesc is the grpc.ServiceDesc for SwarmService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SwarmService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.SwarmService",
	HandlerType: (*SwarmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartFromPreviousEngine",
			Handler:    _SwarmService_StartFromPreviousEngine_Handler,
		},
		{
			MethodName: "StartEngine",
			Handler:    _SwarmService_StartEngine_Handler,
		},
		{
			MethodName: "ListEngines",
			Handler:    _SwarmService_ListEngines_Handler,
		},
		{
			MethodName: "GetEngine",
			Handler:    _SwarmService_GetEngine_Handler,
		},
		{
			MethodName: "StopEngine",
			Handler:    _SwarmService_StopEngine_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartLocalEngine",
			Handler:       _SwarmService_StartLocalEngine_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _SwarmService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Listen",
			Handler:       _SwarmService_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "swarm.proto",
}
