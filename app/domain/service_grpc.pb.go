// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: app/domain/pb/service.proto

package domain

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
	CreateDomain(ctx context.Context, in *CreateDomainRequest, opts ...grpc.CallOption) (*Domain, error)
	UpdateDomain(ctx context.Context, in *UpdateDomainInfoRequest, opts ...grpc.CallOption) (*Domain, error)
	DescribeDomain(ctx context.Context, in *DescribeDomainRequest, opts ...grpc.CallOption) (*Domain, error)
	QueryDomain(ctx context.Context, in *QueryDomainRequest, opts ...grpc.CallOption) (*Set, error)
	DeleteDomain(ctx context.Context, in *DeleteDomainRequest, opts ...grpc.CallOption) (*Domain, error)
	UpdateDomainSecurity(ctx context.Context, in *UpdateDomainSecurityRequest, opts ...grpc.CallOption) (*SecuritySetting, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CreateDomain(ctx context.Context, in *CreateDomainRequest, opts ...grpc.CallOption) (*Domain, error) {
	out := new(Domain)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.domain.Service/CreateDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateDomain(ctx context.Context, in *UpdateDomainInfoRequest, opts ...grpc.CallOption) (*Domain, error) {
	out := new(Domain)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.domain.Service/UpdateDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DescribeDomain(ctx context.Context, in *DescribeDomainRequest, opts ...grpc.CallOption) (*Domain, error) {
	out := new(Domain)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.domain.Service/DescribeDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryDomain(ctx context.Context, in *QueryDomainRequest, opts ...grpc.CallOption) (*Set, error) {
	out := new(Set)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.domain.Service/QueryDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteDomain(ctx context.Context, in *DeleteDomainRequest, opts ...grpc.CallOption) (*Domain, error) {
	out := new(Domain)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.domain.Service/DeleteDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateDomainSecurity(ctx context.Context, in *UpdateDomainSecurityRequest, opts ...grpc.CallOption) (*SecuritySetting, error) {
	out := new(SecuritySetting)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.domain.Service/UpdateDomainSecurity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	CreateDomain(context.Context, *CreateDomainRequest) (*Domain, error)
	UpdateDomain(context.Context, *UpdateDomainInfoRequest) (*Domain, error)
	DescribeDomain(context.Context, *DescribeDomainRequest) (*Domain, error)
	QueryDomain(context.Context, *QueryDomainRequest) (*Set, error)
	DeleteDomain(context.Context, *DeleteDomainRequest) (*Domain, error)
	UpdateDomainSecurity(context.Context, *UpdateDomainSecurityRequest) (*SecuritySetting, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) CreateDomain(context.Context, *CreateDomainRequest) (*Domain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDomain not implemented")
}
func (UnimplementedServiceServer) UpdateDomain(context.Context, *UpdateDomainInfoRequest) (*Domain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDomain not implemented")
}
func (UnimplementedServiceServer) DescribeDomain(context.Context, *DescribeDomainRequest) (*Domain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeDomain not implemented")
}
func (UnimplementedServiceServer) QueryDomain(context.Context, *QueryDomainRequest) (*Set, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDomain not implemented")
}
func (UnimplementedServiceServer) DeleteDomain(context.Context, *DeleteDomainRequest) (*Domain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDomain not implemented")
}
func (UnimplementedServiceServer) UpdateDomainSecurity(context.Context, *UpdateDomainSecurityRequest) (*SecuritySetting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDomainSecurity not implemented")
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

func _Service_CreateDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.domain.Service/CreateDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateDomain(ctx, req.(*CreateDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDomainInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.domain.Service/UpdateDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateDomain(ctx, req.(*UpdateDomainInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DescribeDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DescribeDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.domain.Service/DescribeDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DescribeDomain(ctx, req.(*DescribeDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.domain.Service/QueryDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryDomain(ctx, req.(*QueryDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.domain.Service/DeleteDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteDomain(ctx, req.(*DeleteDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateDomainSecurity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDomainSecurityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateDomainSecurity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.domain.Service/UpdateDomainSecurity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateDomainSecurity(ctx, req.(*UpdateDomainSecurityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.keyauth.domain.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDomain",
			Handler:    _Service_CreateDomain_Handler,
		},
		{
			MethodName: "UpdateDomain",
			Handler:    _Service_UpdateDomain_Handler,
		},
		{
			MethodName: "DescribeDomain",
			Handler:    _Service_DescribeDomain_Handler,
		},
		{
			MethodName: "QueryDomain",
			Handler:    _Service_QueryDomain_Handler,
		},
		{
			MethodName: "DeleteDomain",
			Handler:    _Service_DeleteDomain_Handler,
		},
		{
			MethodName: "UpdateDomainSecurity",
			Handler:    _Service_UpdateDomainSecurity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/domain/pb/service.proto",
}
