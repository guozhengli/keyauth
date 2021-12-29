// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: app/endpoint/pb/service.proto

package endpoint

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	DescribeEndpoint(ctx context.Context, in *DescribeEndpointRequest, opts ...grpc.CallOption) (*Endpoint, error)
	QueryEndpoints(ctx context.Context, in *QueryEndpointRequest, opts ...grpc.CallOption) (*Set, error)
	RegistryEndpoint(ctx context.Context, in *RegistryRequest, opts ...grpc.CallOption) (*RegistryResponse, error)
	DeleteEndpoint(ctx context.Context, in *DeleteEndpointRequest, opts ...grpc.CallOption) (*Endpoint, error)
	QueryResources(ctx context.Context, in *QueryResourceRequest, opts ...grpc.CallOption) (*ResourceSet, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) DescribeEndpoint(ctx context.Context, in *DescribeEndpointRequest, opts ...grpc.CallOption) (*Endpoint, error) {
	out := new(Endpoint)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.endpoint.Service/DescribeEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryEndpoints(ctx context.Context, in *QueryEndpointRequest, opts ...grpc.CallOption) (*Set, error) {
	out := new(Set)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.endpoint.Service/QueryEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) RegistryEndpoint(ctx context.Context, in *RegistryRequest, opts ...grpc.CallOption) (*RegistryResponse, error) {
	out := new(RegistryResponse)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.endpoint.Service/RegistryEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteEndpoint(ctx context.Context, in *DeleteEndpointRequest, opts ...grpc.CallOption) (*Endpoint, error) {
	out := new(Endpoint)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.endpoint.Service/DeleteEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryResources(ctx context.Context, in *QueryResourceRequest, opts ...grpc.CallOption) (*ResourceSet, error) {
	out := new(ResourceSet)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.endpoint.Service/QueryResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	DescribeEndpoint(context.Context, *DescribeEndpointRequest) (*Endpoint, error)
	QueryEndpoints(context.Context, *QueryEndpointRequest) (*Set, error)
	RegistryEndpoint(context.Context, *RegistryRequest) (*RegistryResponse, error)
	DeleteEndpoint(context.Context, *DeleteEndpointRequest) (*Endpoint, error)
	QueryResources(context.Context, *QueryResourceRequest) (*ResourceSet, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) DescribeEndpoint(context.Context, *DescribeEndpointRequest) (*Endpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeEndpoint not implemented")
}
func (UnimplementedServiceServer) QueryEndpoints(context.Context, *QueryEndpointRequest) (*Set, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryEndpoints not implemented")
}
func (UnimplementedServiceServer) RegistryEndpoint(context.Context, *RegistryRequest) (*RegistryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistryEndpoint not implemented")
}
func (UnimplementedServiceServer) DeleteEndpoint(context.Context, *DeleteEndpointRequest) (*Endpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEndpoint not implemented")
}
func (UnimplementedServiceServer) QueryResources(context.Context, *QueryResourceRequest) (*ResourceSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryResources not implemented")
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

func _Service_DescribeEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DescribeEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.endpoint.Service/DescribeEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DescribeEndpoint(ctx, req.(*DescribeEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.endpoint.Service/QueryEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryEndpoints(ctx, req.(*QueryEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_RegistryEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).RegistryEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.endpoint.Service/RegistryEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).RegistryEndpoint(ctx, req.(*RegistryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.endpoint.Service/DeleteEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteEndpoint(ctx, req.(*DeleteEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.endpoint.Service/QueryResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryResources(ctx, req.(*QueryResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.keyauth.endpoint.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeEndpoint",
			Handler:    _Service_DescribeEndpoint_Handler,
		},
		{
			MethodName: "QueryEndpoints",
			Handler:    _Service_QueryEndpoints_Handler,
		},
		{
			MethodName: "RegistryEndpoint",
			Handler:    _Service_RegistryEndpoint_Handler,
		},
		{
			MethodName: "DeleteEndpoint",
			Handler:    _Service_DeleteEndpoint_Handler,
		},
		{
			MethodName: "QueryResources",
			Handler:    _Service_QueryResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/endpoint/pb/service.proto",
}
