// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: app/session/pb/session.proto

package session

import (
	token "github.com/infraboard/keyauth/app/token"
	types "github.com/infraboard/keyauth/app/user/types"
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

// UserAgent todo
type UserAgent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"os" json:"os"
	Os string `protobuf:"bytes,1,opt,name=os,proto3" json:"os" bson:"os"`
	// @gotags: bson:"platform" json:"platform"
	Platform string `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform" bson:"platform"`
	// @gotags: bson:"engine_name" json:"engine_name"
	EngineName string `protobuf:"bytes,3,opt,name=engine_name,json=engineName,proto3" json:"engine_name" bson:"engine_name"`
	// @gotags: bson:"engine_version" json:"engine_version"
	EngineVersion string `protobuf:"bytes,4,opt,name=engine_version,json=engineVersion,proto3" json:"engine_version" bson:"engine_version"`
	// @gotags: bson:"browser_name" json:"browser_name"
	BrowserName string `protobuf:"bytes,5,opt,name=browser_name,json=browserName,proto3" json:"browser_name" bson:"browser_name"`
	// @gotags: bson:"browser_version" json:"browser_version"
	BrowserVersion string `protobuf:"bytes,6,opt,name=browser_version,json=browserVersion,proto3" json:"browser_version" bson:"browser_version"`
}

func (x *UserAgent) Reset() {
	*x = UserAgent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_session_pb_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAgent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAgent) ProtoMessage() {}

func (x *UserAgent) ProtoReflect() protoreflect.Message {
	mi := &file_app_session_pb_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAgent.ProtoReflect.Descriptor instead.
func (*UserAgent) Descriptor() ([]byte, []int) {
	return file_app_session_pb_session_proto_rawDescGZIP(), []int{0}
}

func (x *UserAgent) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *UserAgent) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *UserAgent) GetEngineName() string {
	if x != nil {
		return x.EngineName
	}
	return ""
}

func (x *UserAgent) GetEngineVersion() string {
	if x != nil {
		return x.EngineVersion
	}
	return ""
}

func (x *UserAgent) GetBrowserName() string {
	if x != nil {
		return x.BrowserName
	}
	return ""
}

func (x *UserAgent) GetBrowserVersion() string {
	if x != nil {
		return x.BrowserVersion
	}
	return ""
}

// IPInfo todo
type IPInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"city_id" json:"city_id"
	CityId int64 `protobuf:"varint,1,opt,name=city_id,json=cityId,proto3" json:"city_id" bson:"city_id"`
	// @gotags: bson:"country" json:"country"
	Country string `protobuf:"bytes,2,opt,name=country,proto3" json:"country" bson:"country"`
	// @gotags: bson:"region" json:"region"
	Region string `protobuf:"bytes,3,opt,name=region,proto3" json:"region" bson:"region"`
	// @gotags: bson:"province" json:"province"
	Province string `protobuf:"bytes,4,opt,name=province,proto3" json:"province" bson:"province"`
	// @gotags: bson:"city" json:"city"
	City string `protobuf:"bytes,5,opt,name=city,proto3" json:"city" bson:"city"`
	// @gotags: bson:"isp" json:"isp"
	Isp string `protobuf:"bytes,6,opt,name=isp,proto3" json:"isp" bson:"isp"`
}

func (x *IPInfo) Reset() {
	*x = IPInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_session_pb_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPInfo) ProtoMessage() {}

func (x *IPInfo) ProtoReflect() protoreflect.Message {
	mi := &file_app_session_pb_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPInfo.ProtoReflect.Descriptor instead.
func (*IPInfo) Descriptor() ([]byte, []int) {
	return file_app_session_pb_session_proto_rawDescGZIP(), []int{1}
}

func (x *IPInfo) GetCityId() int64 {
	if x != nil {
		return x.CityId
	}
	return 0
}

func (x *IPInfo) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *IPInfo) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *IPInfo) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *IPInfo) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *IPInfo) GetIsp() string {
	if x != nil {
		return x.Isp
	}
	return ""
}

// Session 登录回话
type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 唯一ID
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 所处域
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 用户名称
	// @gotags: bson:"account" json:"account"
	Account string `protobuf:"bytes,3,opt,name=account,proto3" json:"account" bson:"account"`
	// 用户类型
	// @gotags: bson:"user_type" json:"user_type"
	UserType types.UserType `protobuf:"varint,4,opt,name=user_type,json=userType,proto3,enum=infraboard.keyauth.user.UserType" json:"user_type" bson:"user_type"`
	// 用户通过哪个端登录的
	// @gotags: bson:"application_id" json:"application_id"
	ApplicationId string `protobuf:"bytes,5,opt,name=application_id,json=applicationId,proto3" json:"application_id" bson:"application_id"`
	// 用户通过哪个端登录的
	// @gotags: bson:"application_name" json:"application_name"
	ApplicationName string `protobuf:"bytes,6,opt,name=application_name,json=applicationName,proto3" json:"application_name" bson:"application_name"`
	// 登录方式
	// @gotags: bson:"grant_type" json:"grant_type"
	GrantType token.GrantType `protobuf:"varint,7,opt,name=grant_type,json=grantType,proto3,enum=infraboard.keyauth.token.GrantType" json:"grant_type" bson:"grant_type"`
	// 登录时间
	// @gotags: bson:"login_at" json:"login_at"
	LoginAt int64 `protobuf:"varint,8,opt,name=login_at,json=loginAt,proto3" json:"login_at" bson:"login_at"`
	// 登录IP
	// @gotags: bson:"login_ip" json:"login_ip"
	LoginIp string `protobuf:"bytes,9,opt,name=login_ip,json=loginIp,proto3" json:"login_ip" bson:"login_ip"`
	// 登出时间
	// @gotags: bson:"logout_at" json:"logout_at"
	LogoutAt int64 `protobuf:"varint,10,opt,name=logout_at,json=logoutAt,proto3" json:"logout_at" bson:"logout_at"`
	// 应用的网站地址
	// @gotags: bson:"access_token" json:"access_token"
	AccessToken string `protobuf:"bytes,11,opt,name=access_token,json=accessToken,proto3" json:"access_token" bson:"access_token"`
	// @gotags: bson:"user_agent" json:"user_agent"
	UserAgent *UserAgent `protobuf:"bytes,12,opt,name=user_agent,json=userAgent,proto3" json:"user_agent" bson:"user_agent"`
	// @gotags: bson:"ip_info" json:"ip_info"
	IpInfo *IPInfo `protobuf:"bytes,13,opt,name=ip_info,json=ipInfo,proto3" json:"ip_info" bson:"ip_info"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_session_pb_session_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_app_session_pb_session_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_app_session_pb_session_proto_rawDescGZIP(), []int{2}
}

func (x *Session) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Session) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Session) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Session) GetUserType() types.UserType {
	if x != nil {
		return x.UserType
	}
	return types.UserType(0)
}

func (x *Session) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *Session) GetApplicationName() string {
	if x != nil {
		return x.ApplicationName
	}
	return ""
}

func (x *Session) GetGrantType() token.GrantType {
	if x != nil {
		return x.GrantType
	}
	return token.GrantType(0)
}

func (x *Session) GetLoginAt() int64 {
	if x != nil {
		return x.LoginAt
	}
	return 0
}

func (x *Session) GetLoginIp() string {
	if x != nil {
		return x.LoginIp
	}
	return ""
}

func (x *Session) GetLogoutAt() int64 {
	if x != nil {
		return x.LogoutAt
	}
	return 0
}

func (x *Session) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Session) GetUserAgent() *UserAgent {
	if x != nil {
		return x.UserAgent
	}
	return nil
}

func (x *Session) GetIpInfo() *IPInfo {
	if x != nil {
		return x.IpInfo
	}
	return nil
}

type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64      `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Session `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_session_pb_session_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_app_session_pb_session_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set.ProtoReflect.Descriptor instead.
func (*Set) Descriptor() ([]byte, []int) {
	return file_app_session_pb_session_proto_rawDescGZIP(), []int{3}
}

func (x *Set) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Set) GetItems() []*Session {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_app_session_pb_session_proto protoreflect.FileDescriptor

var file_app_session_pb_session_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x62,
	0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x17, 0x61, 0x70, 0x70, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61, 0x70, 0x70, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70,
	0x62, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x01,
	0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x21, 0x0a, 0x0c, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x72, 0x6f,
	0x77, 0x73, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x95, 0x01, 0x0a, 0x06,
	0x49, 0x50, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x73, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x69, 0x73, 0x70, 0x22, 0x9a, 0x04, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x3e, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x67, 0x72,
	0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x41, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x70, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x70, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x44, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b,
	0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x07, 0x69, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x49, 0x50, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x69, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x56, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x39, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_session_pb_session_proto_rawDescOnce sync.Once
	file_app_session_pb_session_proto_rawDescData = file_app_session_pb_session_proto_rawDesc
)

func file_app_session_pb_session_proto_rawDescGZIP() []byte {
	file_app_session_pb_session_proto_rawDescOnce.Do(func() {
		file_app_session_pb_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_session_pb_session_proto_rawDescData)
	})
	return file_app_session_pb_session_proto_rawDescData
}

var file_app_session_pb_session_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_app_session_pb_session_proto_goTypes = []interface{}{
	(*UserAgent)(nil),    // 0: infraboard.keyauth.session.UserAgent
	(*IPInfo)(nil),       // 1: infraboard.keyauth.session.IPInfo
	(*Session)(nil),      // 2: infraboard.keyauth.session.Session
	(*Set)(nil),          // 3: infraboard.keyauth.session.Set
	(types.UserType)(0),  // 4: infraboard.keyauth.user.UserType
	(token.GrantType)(0), // 5: infraboard.keyauth.token.GrantType
}
var file_app_session_pb_session_proto_depIdxs = []int32{
	4, // 0: infraboard.keyauth.session.Session.user_type:type_name -> infraboard.keyauth.user.UserType
	5, // 1: infraboard.keyauth.session.Session.grant_type:type_name -> infraboard.keyauth.token.GrantType
	0, // 2: infraboard.keyauth.session.Session.user_agent:type_name -> infraboard.keyauth.session.UserAgent
	1, // 3: infraboard.keyauth.session.Session.ip_info:type_name -> infraboard.keyauth.session.IPInfo
	2, // 4: infraboard.keyauth.session.Set.items:type_name -> infraboard.keyauth.session.Session
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_app_session_pb_session_proto_init() }
func file_app_session_pb_session_proto_init() {
	if File_app_session_pb_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_session_pb_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAgent); i {
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
		file_app_session_pb_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPInfo); i {
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
		file_app_session_pb_session_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_app_session_pb_session_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Set); i {
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
			RawDescriptor: file_app_session_pb_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_session_pb_session_proto_goTypes,
		DependencyIndexes: file_app_session_pb_session_proto_depIdxs,
		MessageInfos:      file_app_session_pb_session_proto_msgTypes,
	}.Build()
	File_app_session_pb_session_proto = out.File
	file_app_session_pb_session_proto_rawDesc = nil
	file_app_session_pb_session_proto_goTypes = nil
	file_app_session_pb_session_proto_depIdxs = nil
}
