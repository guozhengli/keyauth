// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: app/permission/pb/service.proto

package permission

import (
	context "context"
	role "github.com/infraboard/keyauth/app/role"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	QueryPermission(ctx context.Context, in *QueryPermissionRequest, opts ...grpc.CallOption) (*role.PermissionSet, error)
	QueryRole(ctx context.Context, in *QueryRoleRequest, opts ...grpc.CallOption) (*role.Set, error)
	CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*role.Permission, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) QueryPermission(ctx context.Context, in *QueryPermissionRequest, opts ...grpc.CallOption) (*role.PermissionSet, error) {
	out := new(role.PermissionSet)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.permission.Service/QueryPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryRole(ctx context.Context, in *QueryRoleRequest, opts ...grpc.CallOption) (*role.Set, error) {
	out := new(role.Set)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.permission.Service/QueryRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*role.Permission, error) {
	out := new(role.Permission)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.permission.Service/CheckPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	QueryPermission(context.Context, *QueryPermissionRequest) (*role.PermissionSet, error)
	QueryRole(context.Context, *QueryRoleRequest) (*role.Set, error)
	CheckPermission(context.Context, *CheckPermissionRequest) (*role.Permission, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) QueryPermission(context.Context, *QueryPermissionRequest) (*role.PermissionSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPermission not implemented")
}
func (UnimplementedServiceServer) QueryRole(context.Context, *QueryRoleRequest) (*role.Set, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRole not implemented")
}
func (UnimplementedServiceServer) CheckPermission(context.Context, *CheckPermissionRequest) (*role.Permission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermission not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_QueryPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.permission.Service/QueryPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryPermission(ctx, req.(*QueryPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.permission.Service/QueryRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryRole(ctx, req.(*QueryRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CheckPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.permission.Service/CheckPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CheckPermission(ctx, req.(*CheckPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.keyauth.permission.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryPermission",
			Handler:    _Service_QueryPermission_Handler,
		},
		{
			MethodName: "QueryRole",
			Handler:    _Service_QueryRole_Handler,
		},
		{
			MethodName: "CheckPermission",
			Handler:    _Service_CheckPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/permission/pb/service.proto",
}
