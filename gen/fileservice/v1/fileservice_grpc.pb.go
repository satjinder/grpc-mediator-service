// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: fileservice/v1/fileservice.proto

package fileservicev1

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

// FileAPIClient is the client API for FileAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileAPIClient interface {
	GetJson(ctx context.Context, in *GetJsonRequest, opts ...grpc.CallOption) (*GetJsonResponse, error)
	UnsupportedHandler(ctx context.Context, in *GetJsonRequest, opts ...grpc.CallOption) (*GetJsonResponse, error)
}

type fileAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewFileAPIClient(cc grpc.ClientConnInterface) FileAPIClient {
	return &fileAPIClient{cc}
}

func (c *fileAPIClient) GetJson(ctx context.Context, in *GetJsonRequest, opts ...grpc.CallOption) (*GetJsonResponse, error) {
	out := new(GetJsonResponse)
	err := c.cc.Invoke(ctx, "/med8r.schemas.samples.fileservice.v1.FileAPI/GetJson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileAPIClient) UnsupportedHandler(ctx context.Context, in *GetJsonRequest, opts ...grpc.CallOption) (*GetJsonResponse, error) {
	out := new(GetJsonResponse)
	err := c.cc.Invoke(ctx, "/med8r.schemas.samples.fileservice.v1.FileAPI/UnsupportedHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileAPIServer is the server API for FileAPI service.
// All implementations must embed UnimplementedFileAPIServer
// for forward compatibility
type FileAPIServer interface {
	GetJson(context.Context, *GetJsonRequest) (*GetJsonResponse, error)
	UnsupportedHandler(context.Context, *GetJsonRequest) (*GetJsonResponse, error)
	mustEmbedUnimplementedFileAPIServer()
}

// UnimplementedFileAPIServer must be embedded to have forward compatible implementations.
type UnimplementedFileAPIServer struct {
}

func (UnimplementedFileAPIServer) GetJson(context.Context, *GetJsonRequest) (*GetJsonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJson not implemented")
}
func (UnimplementedFileAPIServer) UnsupportedHandler(context.Context, *GetJsonRequest) (*GetJsonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsupportedHandler not implemented")
}
func (UnimplementedFileAPIServer) mustEmbedUnimplementedFileAPIServer() {}

// UnsafeFileAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileAPIServer will
// result in compilation errors.
type UnsafeFileAPIServer interface {
	mustEmbedUnimplementedFileAPIServer()
}

func RegisterFileAPIServer(s grpc.ServiceRegistrar, srv FileAPIServer) {
	s.RegisterService(&FileAPI_ServiceDesc, srv)
}

func _FileAPI_GetJson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJsonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileAPIServer).GetJson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/med8r.schemas.samples.fileservice.v1.FileAPI/GetJson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileAPIServer).GetJson(ctx, req.(*GetJsonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileAPI_UnsupportedHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJsonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileAPIServer).UnsupportedHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/med8r.schemas.samples.fileservice.v1.FileAPI/UnsupportedHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileAPIServer).UnsupportedHandler(ctx, req.(*GetJsonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileAPI_ServiceDesc is the grpc.ServiceDesc for FileAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "med8r.schemas.samples.fileservice.v1.FileAPI",
	HandlerType: (*FileAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJson",
			Handler:    _FileAPI_GetJson_Handler,
		},
		{
			MethodName: "UnsupportedHandler",
			Handler:    _FileAPI_UnsupportedHandler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fileservice/v1/fileservice.proto",
}
