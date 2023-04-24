// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/matchfunction.proto

package pb

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

// MatchFunctionClient is the client API for MatchFunction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatchFunctionClient interface {
	// DO NOT CALL THIS FUNCTION MANUALLY. USE backend.FetchMatches INSTEAD.
	// Run pulls Tickets that satisfy Profile constraints from QueryService,
	// runs matchmaking logic against them, then constructs and streams back
	// match candidates to the Backend service.
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (MatchFunction_RunClient, error)
}

type matchFunctionClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchFunctionClient(cc grpc.ClientConnInterface) MatchFunctionClient {
	return &matchFunctionClient{cc}
}

func (c *matchFunctionClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (MatchFunction_RunClient, error) {
	stream, err := c.cc.NewStream(ctx, &MatchFunction_ServiceDesc.Streams[0], "/openmatch.MatchFunction/Run", opts...)
	if err != nil {
		return nil, err
	}
	x := &matchFunctionRunClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MatchFunction_RunClient interface {
	Recv() (*RunResponse, error)
	grpc.ClientStream
}

type matchFunctionRunClient struct {
	grpc.ClientStream
}

func (x *matchFunctionRunClient) Recv() (*RunResponse, error) {
	m := new(RunResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MatchFunctionServer is the server API for MatchFunction service.
// All implementations should embed UnimplementedMatchFunctionServer
// for forward compatibility
type MatchFunctionServer interface {
	// DO NOT CALL THIS FUNCTION MANUALLY. USE backend.FetchMatches INSTEAD.
	// Run pulls Tickets that satisfy Profile constraints from QueryService,
	// runs matchmaking logic against them, then constructs and streams back
	// match candidates to the Backend service.
	Run(*RunRequest, MatchFunction_RunServer) error
}

// UnimplementedMatchFunctionServer should be embedded to have forward compatible implementations.
type UnimplementedMatchFunctionServer struct {
}

func (UnimplementedMatchFunctionServer) Run(*RunRequest, MatchFunction_RunServer) error {
	return status.Errorf(codes.Unimplemented, "method Run not implemented")
}

// UnsafeMatchFunctionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatchFunctionServer will
// result in compilation errors.
type UnsafeMatchFunctionServer interface {
	mustEmbedUnimplementedMatchFunctionServer()
}

func RegisterMatchFunctionServer(s grpc.ServiceRegistrar, srv MatchFunctionServer) {
	s.RegisterService(&MatchFunction_ServiceDesc, srv)
}

func _MatchFunction_Run_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RunRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MatchFunctionServer).Run(m, &matchFunctionRunServer{stream})
}

type MatchFunction_RunServer interface {
	Send(*RunResponse) error
	grpc.ServerStream
}

type matchFunctionRunServer struct {
	grpc.ServerStream
}

func (x *matchFunctionRunServer) Send(m *RunResponse) error {
	return x.ServerStream.SendMsg(m)
}

// MatchFunction_ServiceDesc is the grpc.ServiceDesc for MatchFunction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MatchFunction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openmatch.MatchFunction",
	HandlerType: (*MatchFunctionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Run",
			Handler:       _MatchFunction_Run_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/matchfunction.proto",
}
