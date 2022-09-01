// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: usstats/v2/usstats.proto

package usstatsv2

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

// StatsAPIClient is the client API for StatsAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatsAPIClient interface {
	GetStatsData(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error)
}

type statsAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewStatsAPIClient(cc grpc.ClientConnInterface) StatsAPIClient {
	return &statsAPIClient{cc}
}

func (c *statsAPIClient) GetStatsData(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error) {
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, "/med8r.schemas.samples.usstats.v2.StatsAPI/GetStatsData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsAPIServer is the server API for StatsAPI service.
// All implementations must embed UnimplementedStatsAPIServer
// for forward compatibility
type StatsAPIServer interface {
	GetStatsData(context.Context, *GetStatsRequest) (*GetStatsResponse, error)
	mustEmbedUnimplementedStatsAPIServer()
}

// UnimplementedStatsAPIServer must be embedded to have forward compatible implementations.
type UnimplementedStatsAPIServer struct {
}

func (UnimplementedStatsAPIServer) GetStatsData(context.Context, *GetStatsRequest) (*GetStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatsData not implemented")
}
func (UnimplementedStatsAPIServer) mustEmbedUnimplementedStatsAPIServer() {}

// UnsafeStatsAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatsAPIServer will
// result in compilation errors.
type UnsafeStatsAPIServer interface {
	mustEmbedUnimplementedStatsAPIServer()
}

func RegisterStatsAPIServer(s grpc.ServiceRegistrar, srv StatsAPIServer) {
	s.RegisterService(&StatsAPI_ServiceDesc, srv)
}

func _StatsAPI_GetStatsData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsAPIServer).GetStatsData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/med8r.schemas.samples.usstats.v2.StatsAPI/GetStatsData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsAPIServer).GetStatsData(ctx, req.(*GetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StatsAPI_ServiceDesc is the grpc.ServiceDesc for StatsAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatsAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "med8r.schemas.samples.usstats.v2.StatsAPI",
	HandlerType: (*StatsAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatsData",
			Handler:    _StatsAPI_GetStatsData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usstats/v2/usstats.proto",
}