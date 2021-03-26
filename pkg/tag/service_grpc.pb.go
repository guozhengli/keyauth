// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tag

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TagServiceClient is the client API for TagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagServiceClient interface {
	CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*TagKey, error)
	QueryTagKey(ctx context.Context, in *QueryTagKeyRequest, opts ...grpc.CallOption) (*TagKeySet, error)
	QueryTagValue(ctx context.Context, in *QueryTagValueRequest, opts ...grpc.CallOption) (*TagValueSet, error)
}

type tagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagServiceClient(cc grpc.ClientConnInterface) TagServiceClient {
	return &tagServiceClient{cc}
}

func (c *tagServiceClient) CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*TagKey, error) {
	out := new(TagKey)
	err := c.cc.Invoke(ctx, "/keyauth.tag.TagService/CreateTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) QueryTagKey(ctx context.Context, in *QueryTagKeyRequest, opts ...grpc.CallOption) (*TagKeySet, error) {
	out := new(TagKeySet)
	err := c.cc.Invoke(ctx, "/keyauth.tag.TagService/QueryTagKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) QueryTagValue(ctx context.Context, in *QueryTagValueRequest, opts ...grpc.CallOption) (*TagValueSet, error) {
	out := new(TagValueSet)
	err := c.cc.Invoke(ctx, "/keyauth.tag.TagService/QueryTagValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServiceServer is the server API for TagService service.
// All implementations must embed UnimplementedTagServiceServer
// for forward compatibility
type TagServiceServer interface {
	CreateTag(context.Context, *CreateTagRequest) (*TagKey, error)
	QueryTagKey(context.Context, *QueryTagKeyRequest) (*TagKeySet, error)
	QueryTagValue(context.Context, *QueryTagValueRequest) (*TagValueSet, error)
	mustEmbedUnimplementedTagServiceServer()
}

// UnimplementedTagServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTagServiceServer struct {
}

func (UnimplementedTagServiceServer) CreateTag(context.Context, *CreateTagRequest) (*TagKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTag not implemented")
}
func (UnimplementedTagServiceServer) QueryTagKey(context.Context, *QueryTagKeyRequest) (*TagKeySet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTagKey not implemented")
}
func (UnimplementedTagServiceServer) QueryTagValue(context.Context, *QueryTagValueRequest) (*TagValueSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTagValue not implemented")
}
func (UnimplementedTagServiceServer) mustEmbedUnimplementedTagServiceServer() {}

// UnsafeTagServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagServiceServer will
// result in compilation errors.
type UnsafeTagServiceServer interface {
	mustEmbedUnimplementedTagServiceServer()
}

func RegisterTagServiceServer(s *grpc.Server, srv TagServiceServer) {
	s.RegisterService(&_TagService_serviceDesc, srv)
}

func _TagService_CreateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).CreateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.tag.TagService/CreateTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).CreateTag(ctx, req.(*CreateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_QueryTagKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTagKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).QueryTagKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.tag.TagService/QueryTagKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).QueryTagKey(ctx, req.(*QueryTagKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_QueryTagValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTagValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).QueryTagValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.tag.TagService/QueryTagValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).QueryTagValue(ctx, req.(*QueryTagValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TagService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "keyauth.tag.TagService",
	HandlerType: (*TagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTag",
			Handler:    _TagService_CreateTag_Handler,
		},
		{
			MethodName: "QueryTagKey",
			Handler:    _TagService_QueryTagKey_Handler,
		},
		{
			MethodName: "QueryTagValue",
			Handler:    _TagService_QueryTagValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/tag/pb/service.proto",
}
