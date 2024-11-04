// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: identity.proto

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	IdentityService_InsertUser_FullMethodName = "/pb.IdentityService/InsertUser"
	IdentityService_GetUser_FullMethodName    = "/pb.IdentityService/GetUser"
)

// IdentityServiceClient is the client API for IdentityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityServiceClient interface {
	InsertUser(ctx context.Context, in *InsertReq, opts ...grpc.CallOption) (*InsertRes, error)
	GetUser(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRes, error)
}

type identityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityServiceClient(cc grpc.ClientConnInterface) IdentityServiceClient {
	return &identityServiceClient{cc}
}

func (c *identityServiceClient) InsertUser(ctx context.Context, in *InsertReq, opts ...grpc.CallOption) (*InsertRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InsertRes)
	err := c.cc.Invoke(ctx, IdentityService_InsertUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) GetUser(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRes)
	err := c.cc.Invoke(ctx, IdentityService_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServiceServer is the server API for IdentityService service.
// All implementations must embed UnimplementedIdentityServiceServer
// for forward compatibility.
type IdentityServiceServer interface {
	InsertUser(context.Context, *InsertReq) (*InsertRes, error)
	GetUser(context.Context, *GetReq) (*GetRes, error)
	mustEmbedUnimplementedIdentityServiceServer()
}

// UnimplementedIdentityServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedIdentityServiceServer struct{}

func (UnimplementedIdentityServiceServer) InsertUser(context.Context, *InsertReq) (*InsertRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertUser not implemented")
}
func (UnimplementedIdentityServiceServer) GetUser(context.Context, *GetReq) (*GetRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedIdentityServiceServer) mustEmbedUnimplementedIdentityServiceServer() {}
func (UnimplementedIdentityServiceServer) testEmbeddedByValue()                         {}

// UnsafeIdentityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityServiceServer will
// result in compilation errors.
type UnsafeIdentityServiceServer interface {
	mustEmbedUnimplementedIdentityServiceServer()
}

func RegisterIdentityServiceServer(s grpc.ServiceRegistrar, srv IdentityServiceServer) {
	// If the following call pancis, it indicates UnimplementedIdentityServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&IdentityService_ServiceDesc, srv)
}

func _IdentityService_InsertUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).InsertUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityService_InsertUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).InsertUser(ctx, req.(*InsertReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).GetUser(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

// IdentityService_ServiceDesc is the grpc.ServiceDesc for IdentityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdentityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.IdentityService",
	HandlerType: (*IdentityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertUser",
			Handler:    _IdentityService_InsertUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _IdentityService_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity.proto",
}