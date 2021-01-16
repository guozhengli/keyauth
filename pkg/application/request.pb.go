//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/application/pb/request.proto

package application

import (
	proto "github.com/golang/protobuf/proto"
	page "github.com/infraboard/mcube/pb/page"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ClientType int32

const (
	ClientType_CONFIDENTIAL ClientType = 0
	ClientType_PUBLIC       ClientType = 1
)

// Enum value maps for ClientType.
var (
	ClientType_name = map[int32]string{
		0: "CONFIDENTIAL",
		1: "PUBLIC",
	}
	ClientType_value = map[string]int32{
		"CONFIDENTIAL": 0,
		"PUBLIC":       1,
	}
)

func (x ClientType) Enum() *ClientType {
	p := new(ClientType)
	*p = x
	return p
}

func (x ClientType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_application_pb_request_proto_enumTypes[0].Descriptor()
}

func (ClientType) Type() protoreflect.EnumType {
	return &file_pkg_application_pb_request_proto_enumTypes[0]
}

func (x ClientType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientType.Descriptor instead.
func (ClientType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_application_pb_request_proto_rawDescGZIP(), []int{0}
}

// CreateApplicatonRequest 创建应用请求
type CreateApplicatonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 应用名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 应用的网站地址
	Website string `protobuf:"bytes,2,opt,name=website,proto3" json:"website,omitempty"`
	// 应用的LOGO
	LogoImage string `protobuf:"bytes,3,opt,name=logo_image,json=logoImage,proto3" json:"logo_image,omitempty"`
	// 应用简单的描述
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// 应用重定向URI, Oauht2时需要该参数
	RedirectUri string `protobuf:"bytes,5,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	// 应用申请的token的过期时间
	AccessTokenExpireSecond int64 `protobuf:"varint,6,opt,name=access_token_expire_second,json=accessTokenExpireSecond,proto3" json:"access_token_expire_second,omitempty"`
	// 刷新token过期时间
	RefreshTokenExpireSecond int64 `protobuf:"varint,7,opt,name=refresh_token_expire_second,json=refreshTokenExpireSecond,proto3" json:"refresh_token_expire_second,omitempty"`
	// 客户端类型
	ClientType ClientType `protobuf:"varint,8,opt,name=client_type,json=clientType,proto3,enum=keyauth.application.ClientType" json:"client_type,omitempty"`
}

func (x *CreateApplicatonRequest) Reset() {
	*x = CreateApplicatonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_application_pb_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApplicatonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApplicatonRequest) ProtoMessage() {}

func (x *CreateApplicatonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_application_pb_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApplicatonRequest.ProtoReflect.Descriptor instead.
func (*CreateApplicatonRequest) Descriptor() ([]byte, []int) {
	return file_pkg_application_pb_request_proto_rawDescGZIP(), []int{0}
}

func (x *CreateApplicatonRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateApplicatonRequest) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *CreateApplicatonRequest) GetLogoImage() string {
	if x != nil {
		return x.LogoImage
	}
	return ""
}

func (x *CreateApplicatonRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateApplicatonRequest) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

func (x *CreateApplicatonRequest) GetAccessTokenExpireSecond() int64 {
	if x != nil {
		return x.AccessTokenExpireSecond
	}
	return 0
}

func (x *CreateApplicatonRequest) GetRefreshTokenExpireSecond() int64 {
	if x != nil {
		return x.RefreshTokenExpireSecond
	}
	return 0
}

func (x *CreateApplicatonRequest) GetClientType() ClientType {
	if x != nil {
		return x.ClientType
	}
	return ClientType_CONFIDENTIAL
}

// DescribeApplicationRequest 查询应用详情
type DescribeApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *DescribeApplicationRequest) Reset() {
	*x = DescribeApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_application_pb_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeApplicationRequest) ProtoMessage() {}

func (x *DescribeApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_application_pb_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeApplicationRequest.ProtoReflect.Descriptor instead.
func (*DescribeApplicationRequest) Descriptor() ([]byte, []int) {
	return file_pkg_application_pb_request_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeApplicationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DescribeApplicationRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

// DeleteApplicationRequest 查询应用详情
type DeleteApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteApplicationRequest) Reset() {
	*x = DeleteApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_application_pb_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApplicationRequest) ProtoMessage() {}

func (x *DeleteApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_application_pb_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApplicationRequest.ProtoReflect.Descriptor instead.
func (*DeleteApplicationRequest) Descriptor() ([]byte, []int) {
	return file_pkg_application_pb_request_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteApplicationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// QueryApplicationRequest 查询应用列表
type QueryApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    *page.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	UserId  string            `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Account string            `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *QueryApplicationRequest) Reset() {
	*x = QueryApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_application_pb_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryApplicationRequest) ProtoMessage() {}

func (x *QueryApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_application_pb_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryApplicationRequest.ProtoReflect.Descriptor instead.
func (*QueryApplicationRequest) Descriptor() ([]byte, []int) {
	return file_pkg_application_pb_request_proto_rawDescGZIP(), []int{3}
}

func (x *QueryApplicationRequest) GetPage() *page.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryApplicationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *QueryApplicationRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

// GetBuildInApplicationRequest 获取内建应用
type GetBuildInApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetBuildInApplicationRequest) Reset() {
	*x = GetBuildInApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_application_pb_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBuildInApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBuildInApplicationRequest) ProtoMessage() {}

func (x *GetBuildInApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_application_pb_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBuildInApplicationRequest.ProtoReflect.Descriptor instead.
func (*GetBuildInApplicationRequest) Descriptor() ([]byte, []int) {
	return file_pkg_application_pb_request_proto_rawDescGZIP(), []int{4}
}

func (x *GetBuildInApplicationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_pkg_application_pb_request_proto protoreflect.FileDescriptor

var file_pkg_application_pb_request_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d,
	0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x02, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x6f, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75,
	0x72, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x55, 0x72, 0x69, 0x12, 0x3b, 0x0a, 0x1a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x12, 0x3d, 0x0a, 0x1b, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x18, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x49, 0x0a, 0x1a, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x2a,
	0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x73, 0x0a, 0x17, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x32, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x2a, 0x2a, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x41,
	0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x01, 0x42,
	0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_application_pb_request_proto_rawDescOnce sync.Once
	file_pkg_application_pb_request_proto_rawDescData = file_pkg_application_pb_request_proto_rawDesc
)

func file_pkg_application_pb_request_proto_rawDescGZIP() []byte {
	file_pkg_application_pb_request_proto_rawDescOnce.Do(func() {
		file_pkg_application_pb_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_application_pb_request_proto_rawDescData)
	})
	return file_pkg_application_pb_request_proto_rawDescData
}

var file_pkg_application_pb_request_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_application_pb_request_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_application_pb_request_proto_goTypes = []interface{}{
	(ClientType)(0),                      // 0: keyauth.application.ClientType
	(*CreateApplicatonRequest)(nil),      // 1: keyauth.application.CreateApplicatonRequest
	(*DescribeApplicationRequest)(nil),   // 2: keyauth.application.DescribeApplicationRequest
	(*DeleteApplicationRequest)(nil),     // 3: keyauth.application.DeleteApplicationRequest
	(*QueryApplicationRequest)(nil),      // 4: keyauth.application.QueryApplicationRequest
	(*GetBuildInApplicationRequest)(nil), // 5: keyauth.application.GetBuildInApplicationRequest
	(*page.PageRequest)(nil),             // 6: page.PageRequest
}
var file_pkg_application_pb_request_proto_depIdxs = []int32{
	0, // 0: keyauth.application.CreateApplicatonRequest.client_type:type_name -> keyauth.application.ClientType
	6, // 1: keyauth.application.QueryApplicationRequest.page:type_name -> page.PageRequest
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_application_pb_request_proto_init() }
func file_pkg_application_pb_request_proto_init() {
	if File_pkg_application_pb_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_application_pb_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApplicatonRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_application_pb_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeApplicationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_application_pb_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteApplicationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_application_pb_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryApplicationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_application_pb_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBuildInApplicationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_application_pb_request_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_application_pb_request_proto_goTypes,
		DependencyIndexes: file_pkg_application_pb_request_proto_depIdxs,
		EnumInfos:         file_pkg_application_pb_request_proto_enumTypes,
		MessageInfos:      file_pkg_application_pb_request_proto_msgTypes,
	}.Build()
	File_pkg_application_pb_request_proto = out.File
	file_pkg_application_pb_request_proto_rawDesc = nil
	file_pkg_application_pb_request_proto_goTypes = nil
	file_pkg_application_pb_request_proto_depIdxs = nil
}
