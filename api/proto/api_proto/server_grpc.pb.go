// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.2
// source: server.proto

package api_proto

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

// LocationServiceClient is the client API for LocationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocationServiceClient interface {
	// Unary RPC method
	GetLocation(ctx context.Context, in *LocationRequest, opts ...grpc.CallOption) (*LocationResponse, error)
	// Server-side streaming RPC method
	StreamLocations(ctx context.Context, in *LocationRequest, opts ...grpc.CallOption) (LocationService_StreamLocationsClient, error)
	// Bidirectional streaming RPC method
	TrackLocations(ctx context.Context, opts ...grpc.CallOption) (LocationService_TrackLocationsClient, error)
}

type locationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationServiceClient(cc grpc.ClientConnInterface) LocationServiceClient {
	return &locationServiceClient{cc}
}

func (c *locationServiceClient) GetLocation(ctx context.Context, in *LocationRequest, opts ...grpc.CallOption) (*LocationResponse, error) {
	out := new(LocationResponse)
	err := c.cc.Invoke(ctx, "/api_proto.LocationService/GetLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) StreamLocations(ctx context.Context, in *LocationRequest, opts ...grpc.CallOption) (LocationService_StreamLocationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &LocationService_ServiceDesc.Streams[0], "/api_proto.LocationService/StreamLocations", opts...)
	if err != nil {
		return nil, err
	}
	x := &locationServiceStreamLocationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LocationService_StreamLocationsClient interface {
	Recv() (*LocationResponse, error)
	grpc.ClientStream
}

type locationServiceStreamLocationsClient struct {
	grpc.ClientStream
}

func (x *locationServiceStreamLocationsClient) Recv() (*LocationResponse, error) {
	m := new(LocationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *locationServiceClient) TrackLocations(ctx context.Context, opts ...grpc.CallOption) (LocationService_TrackLocationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &LocationService_ServiceDesc.Streams[1], "/api_proto.LocationService/TrackLocations", opts...)
	if err != nil {
		return nil, err
	}
	x := &locationServiceTrackLocationsClient{stream}
	return x, nil
}

type LocationService_TrackLocationsClient interface {
	Send(*LocationRequest) error
	Recv() (*LocationResponse, error)
	grpc.ClientStream
}

type locationServiceTrackLocationsClient struct {
	grpc.ClientStream
}

func (x *locationServiceTrackLocationsClient) Send(m *LocationRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *locationServiceTrackLocationsClient) Recv() (*LocationResponse, error) {
	m := new(LocationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LocationServiceServer is the server API for LocationService service.
// All implementations must embed UnimplementedLocationServiceServer
// for forward compatibility
type LocationServiceServer interface {
	// Unary RPC method
	GetLocation(context.Context, *LocationRequest) (*LocationResponse, error)
	// Server-side streaming RPC method
	StreamLocations(*LocationRequest, LocationService_StreamLocationsServer) error
	// Bidirectional streaming RPC method
	TrackLocations(LocationService_TrackLocationsServer) error
	mustEmbedUnimplementedLocationServiceServer()
}

// UnimplementedLocationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLocationServiceServer struct {
}

func (UnimplementedLocationServiceServer) GetLocation(context.Context, *LocationRequest) (*LocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocation not implemented")
}
func (UnimplementedLocationServiceServer) StreamLocations(*LocationRequest, LocationService_StreamLocationsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamLocations not implemented")
}
func (UnimplementedLocationServiceServer) TrackLocations(LocationService_TrackLocationsServer) error {
	return status.Errorf(codes.Unimplemented, "method TrackLocations not implemented")
}
func (UnimplementedLocationServiceServer) mustEmbedUnimplementedLocationServiceServer() {}

// UnsafeLocationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocationServiceServer will
// result in compilation errors.
type UnsafeLocationServiceServer interface {
	mustEmbedUnimplementedLocationServiceServer()
}

func RegisterLocationServiceServer(s grpc.ServiceRegistrar, srv LocationServiceServer) {
	s.RegisterService(&LocationService_ServiceDesc, srv)
}

func _LocationService_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_proto.LocationService/GetLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).GetLocation(ctx, req.(*LocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_StreamLocations_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LocationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LocationServiceServer).StreamLocations(m, &locationServiceStreamLocationsServer{stream})
}

type LocationService_StreamLocationsServer interface {
	Send(*LocationResponse) error
	grpc.ServerStream
}

type locationServiceStreamLocationsServer struct {
	grpc.ServerStream
}

func (x *locationServiceStreamLocationsServer) Send(m *LocationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LocationService_TrackLocations_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LocationServiceServer).TrackLocations(&locationServiceTrackLocationsServer{stream})
}

type LocationService_TrackLocationsServer interface {
	Send(*LocationResponse) error
	Recv() (*LocationRequest, error)
	grpc.ServerStream
}

type locationServiceTrackLocationsServer struct {
	grpc.ServerStream
}

func (x *locationServiceTrackLocationsServer) Send(m *LocationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *locationServiceTrackLocationsServer) Recv() (*LocationRequest, error) {
	m := new(LocationRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LocationService_ServiceDesc is the grpc.ServiceDesc for LocationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api_proto.LocationService",
	HandlerType: (*LocationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLocation",
			Handler:    _LocationService_GetLocation_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLocations",
			Handler:       _LocationService_StreamLocations_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TrackLocations",
			Handler:       _LocationService_TrackLocations_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "server.proto",
}