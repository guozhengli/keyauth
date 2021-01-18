// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package token

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TokenServiceClient is the client API for TokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenServiceClient interface {
	IssueToken(ctx context.Context, in *IssueTokenRequest, opts ...grpc.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*Token, error)
	DescribeToken(ctx context.Context, in *DescribeTokenRequest, opts ...grpc.CallOption) (*Token, error)
	RevolkToken(ctx context.Context, in *RevolkTokenRequest, opts ...grpc.CallOption) (*Token, error)
	BlockToken(ctx context.Context, in *BlockTokenRequest, opts ...grpc.CallOption) (*Token, error)
	QueryToken(ctx context.Context, in *QueryTokenRequest, opts ...grpc.CallOption) (*Set, error)
}

type tokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenServiceClient(cc grpc.ClientConnInterface) TokenServiceClient {
	return &tokenServiceClient{cc}
}

func (c *tokenServiceClient) IssueToken(ctx context.Context, in *IssueTokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/keyauth.token.TokenService/IssueToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/keyauth.token.TokenService/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) DescribeToken(ctx context.Context, in *DescribeTokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/keyauth.token.TokenService/DescribeToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) RevolkToken(ctx context.Context, in *RevolkTokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/keyauth.token.TokenService/RevolkToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) BlockToken(ctx context.Context, in *BlockTokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/keyauth.token.TokenService/BlockToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) QueryToken(ctx context.Context, in *QueryTokenRequest, opts ...grpc.CallOption) (*Set, error) {
	out := new(Set)
	err := c.cc.Invoke(ctx, "/keyauth.token.TokenService/QueryToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenServiceServer is the server API for TokenService service.
// All implementations must embed UnimplementedTokenServiceServer
// for forward compatibility
type TokenServiceServer interface {
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error)
	DescribeToken(context.Context, *DescribeTokenRequest) (*Token, error)
	RevolkToken(context.Context, *RevolkTokenRequest) (*Token, error)
	BlockToken(context.Context, *BlockTokenRequest) (*Token, error)
	QueryToken(context.Context, *QueryTokenRequest) (*Set, error)
	mustEmbedUnimplementedTokenServiceServer()
}

// UnimplementedTokenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTokenServiceServer struct {
}

func (UnimplementedTokenServiceServer) IssueToken(context.Context, *IssueTokenRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueToken not implemented")
}
func (UnimplementedTokenServiceServer) ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedTokenServiceServer) DescribeToken(context.Context, *DescribeTokenRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeToken not implemented")
}
func (UnimplementedTokenServiceServer) RevolkToken(context.Context, *RevolkTokenRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevolkToken not implemented")
}
func (UnimplementedTokenServiceServer) BlockToken(context.Context, *BlockTokenRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockToken not implemented")
}
func (UnimplementedTokenServiceServer) QueryToken(context.Context, *QueryTokenRequest) (*Set, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryToken not implemented")
}
func (UnimplementedTokenServiceServer) mustEmbedUnimplementedTokenServiceServer() {}

// UnsafeTokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenServiceServer will
// result in compilation errors.
type UnsafeTokenServiceServer interface {
	mustEmbedUnimplementedTokenServiceServer()
}

func RegisterTokenServiceServer(s grpc.ServiceRegistrar, srv TokenServiceServer) {
	s.RegisterService(&_TokenService_serviceDesc, srv)
}

func _TokenService_IssueToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).IssueToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.token.TokenService/IssueToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).IssueToken(ctx, req.(*IssueTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.token.TokenService/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).ValidateToken(ctx, req.(*ValidateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_DescribeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).DescribeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.token.TokenService/DescribeToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).DescribeToken(ctx, req.(*DescribeTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_RevolkToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevolkTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).RevolkToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.token.TokenService/RevolkToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).RevolkToken(ctx, req.(*RevolkTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_BlockToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).BlockToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.token.TokenService/BlockToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).BlockToken(ctx, req.(*BlockTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_QueryToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).QueryToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.token.TokenService/QueryToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).QueryToken(ctx, req.(*QueryTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TokenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "keyauth.token.TokenService",
	HandlerType: (*TokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueToken",
			Handler:    _TokenService_IssueToken_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _TokenService_ValidateToken_Handler,
		},
		{
			MethodName: "DescribeToken",
			Handler:    _TokenService_DescribeToken_Handler,
		},
		{
			MethodName: "RevolkToken",
			Handler:    _TokenService_RevolkToken_Handler,
		},
		{
			MethodName: "BlockToken",
			Handler:    _TokenService_BlockToken_Handler,
		},
		{
			MethodName: "QueryToken",
			Handler:    _TokenService_QueryToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/token/pb/service.proto",
}
