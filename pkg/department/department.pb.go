//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/department/pb/department.proto

package department

import (
	proto "github.com/golang/protobuf/proto"
	role "github.com/infraboard/keyauth/pkg/role"
	_ "github.com/infraboard/protoc-gen-go-ext/extension/tag"
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

// Department user's department
type Department struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 部门ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 路径
	ParentPath string `protobuf:"bytes,2,opt,name=parent_path,json=parentPath,proto3" json:"parent_path" bson:"parent_path"`
	// 部门编号
	Number uint64 `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty" bson:"number"`
	// 部门创建时间
	CreateAt int64 `protobuf:"varint,4,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty" bson:"create_at"`
	// 更新时间
	UpdateAt int64 `protobuf:"varint,5,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty" bson:"update_at"`
	// 创建人
	Creater string `protobuf:"bytes,6,opt,name=creater,proto3" json:"creater,omitempty" bson:"creater"`
	// 部门所属域
	Domain string `protobuf:"bytes,7,opt,name=domain,proto3" json:"domain,omitempty" bson:"domain"`
	// 第几级部门, 由层数决定
	Grade int32 `protobuf:"varint,8,opt,name=grade,proto3" json:"grade,omitempty" bson:"grade"`
	// 子部门数量
	SubCount int64 `protobuf:"varint,9,opt,name=sub_count,json=subCount,proto3" json:"sub_count,omitempty" bson:"-"`
	// 部门所有用户数量
	UserCount int64 `protobuf:"varint,10,opt,name=user_count,json=userCount,proto3" json:"user_count,omitempty" bson:"-"`
	// 具体数据
	Data *CreateDepartmentRequest `protobuf:"bytes,11,opt,name=data,proto3" json:"data" bson:"data"`
	// 默认角色
	DefaultRole *role.Role `protobuf:"bytes,12,opt,name=default_role,json=defaultRole,proto3" json:"default_role,omitempty" bson:"-"`
}

func (x *Department) Reset() {
	*x = Department{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_department_pb_department_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Department) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Department) ProtoMessage() {}

func (x *Department) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_department_pb_department_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Department.ProtoReflect.Descriptor instead.
func (*Department) Descriptor() ([]byte, []int) {
	return file_pkg_department_pb_department_proto_rawDescGZIP(), []int{0}
}

func (x *Department) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Department) GetParentPath() string {
	if x != nil {
		return x.ParentPath
	}
	return ""
}

func (x *Department) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Department) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Department) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Department) GetCreater() string {
	if x != nil {
		return x.Creater
	}
	return ""
}

func (x *Department) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Department) GetGrade() int32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *Department) GetSubCount() int64 {
	if x != nil {
		return x.SubCount
	}
	return 0
}

func (x *Department) GetUserCount() int64 {
	if x != nil {
		return x.UserCount
	}
	return 0
}

func (x *Department) GetData() *CreateDepartmentRequest {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Department) GetDefaultRole() *role.Role {
	if x != nil {
		return x.DefaultRole
	}
	return nil
}

type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64         `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	Items []*Department `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_department_pb_department_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_department_pb_department_proto_msgTypes[1]
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
	return file_pkg_department_pb_department_proto_rawDescGZIP(), []int{1}
}

func (x *Set) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Set) GetItems() []*Department {
	if x != nil {
		return x.Items
	}
	return nil
}

// ApplicationForm todo
type ApplicationForm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 申请单ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 域
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 申请人
	Creater string `protobuf:"bytes,3,opt,name=creater,proto3" json:"creater" bson:"creater"`
	// 创建时间
	CreateAt int64 `protobuf:"varint,4,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	UpdateAt int64 `protobuf:"varint,5,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 状态
	Status ApplicationFormStatus `protobuf:"varint,6,opt,name=status,proto3,enum=keyauth.department.ApplicationFormStatus" json:"status" bson:"status"`
	// 数据
	Data *JoinDepartmentRequest `protobuf:"bytes,7,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *ApplicationForm) Reset() {
	*x = ApplicationForm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_department_pb_department_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationForm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationForm) ProtoMessage() {}

func (x *ApplicationForm) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_department_pb_department_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationForm.ProtoReflect.Descriptor instead.
func (*ApplicationForm) Descriptor() ([]byte, []int) {
	return file_pkg_department_pb_department_proto_rawDescGZIP(), []int{2}
}

func (x *ApplicationForm) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ApplicationForm) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ApplicationForm) GetCreater() string {
	if x != nil {
		return x.Creater
	}
	return ""
}

func (x *ApplicationForm) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *ApplicationForm) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *ApplicationForm) GetStatus() ApplicationFormStatus {
	if x != nil {
		return x.Status
	}
	return ApplicationFormStatus_NULL
}

func (x *ApplicationForm) GetData() *JoinDepartmentRequest {
	if x != nil {
		return x.Data
	}
	return nil
}

type ApplicationFormSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64              `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	Items []*ApplicationForm `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *ApplicationFormSet) Reset() {
	*x = ApplicationFormSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_department_pb_department_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationFormSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationFormSet) ProtoMessage() {}

func (x *ApplicationFormSet) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_department_pb_department_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationFormSet.ProtoReflect.Descriptor instead.
func (*ApplicationFormSet) Descriptor() ([]byte, []int) {
	return file_pkg_department_pb_department_proto_rawDescGZIP(), []int{3}
}

func (x *ApplicationFormSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ApplicationFormSet) GetItems() []*ApplicationForm {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_pkg_department_pb_department_proto protoreflect.FileDescriptor

var file_pkg_department_pb_department_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f, 0x6c,
	0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x92, 0x07, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x2a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xc2, 0xde, 0x1f,
	0x16, 0x0a, 0x14, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4c, 0x0a, 0x0b, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2b, 0xc2, 0xde, 0x1f, 0x27, 0x0a, 0x25, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x22, 0x52, 0x0a, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x43, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x2b, 0xc2, 0xde, 0x1f, 0x27, 0x0a,
	0x25, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x20, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2c, 0x6f, 0x6d, 0x69, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x4e,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x31, 0xc2, 0xde, 0x1f, 0x2d, 0x0a, 0x2b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x4e,
	0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x31, 0xc2, 0xde, 0x1f, 0x2d, 0x0a, 0x2b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x47,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2d, 0xc2, 0xde, 0x1f, 0x29, 0x0a, 0x27, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x72, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x72, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xc2, 0xde, 0x1f, 0x27, 0x0a, 0x25, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x3f, 0x0a, 0x05,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x42, 0x29, 0xc2, 0xde, 0x1f,
	0x25, 0x0a, 0x23, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x67, 0x72, 0x61, 0x64, 0x65, 0x22, 0x20,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2c, 0x6f, 0x6d, 0x69, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x46, 0x0a,
	0x09, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x29, 0xc2, 0xde, 0x1f, 0x25, 0x0a, 0x23, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2d, 0x22,
	0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x08, 0x73, 0x75, 0x62,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x49, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x42, 0x2a, 0xc2, 0xde, 0x1f, 0x26, 0x0a,
	0x24, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2d, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x5e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x1d, 0xc2, 0xde, 0x1f,
	0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x61, 0x74, 0x61, 0x22, 0x20, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x61, 0x74, 0x61, 0x22, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x63, 0x0a, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x2c, 0xc2, 0xde, 0x1f, 0x28,
	0x0a, 0x26, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2d, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x2c, 0x6f, 0x6d,
	0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x35, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x1f, 0xc2, 0xde,
	0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x55, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x87, 0x04, 0x0a, 0x0f,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x12,
	0x2a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xc2, 0xde, 0x1f,
	0x16, 0x0a, 0x14, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xc2, 0xde, 0x1f,
	0x1d, 0x0a, 0x1b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22,
	0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xc2, 0xde, 0x1f, 0x1f, 0x0a, 0x1d, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x22, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x22, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21,
	0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22,
	0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74,
	0x22, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x44, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x27,
	0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x12, 0x64, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x29, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x21, 0xc2, 0xde,
	0x1f, 0x1d, 0x0a, 0x1b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x44,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x1d, 0xc2, 0xde, 0x1f, 0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x61, 0x74, 0x61, 0x22, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa7, 0x01, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x53, 0x65, 0x74, 0x12, 0x35, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x1f, 0xc2, 0xde, 0x1f,
	0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x20,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x5a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_department_pb_department_proto_rawDescOnce sync.Once
	file_pkg_department_pb_department_proto_rawDescData = file_pkg_department_pb_department_proto_rawDesc
)

func file_pkg_department_pb_department_proto_rawDescGZIP() []byte {
	file_pkg_department_pb_department_proto_rawDescOnce.Do(func() {
		file_pkg_department_pb_department_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_department_pb_department_proto_rawDescData)
	})
	return file_pkg_department_pb_department_proto_rawDescData
}

var file_pkg_department_pb_department_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_department_pb_department_proto_goTypes = []interface{}{
	(*Department)(nil),              // 0: keyauth.department.Department
	(*Set)(nil),                     // 1: keyauth.department.Set
	(*ApplicationForm)(nil),         // 2: keyauth.department.ApplicationForm
	(*ApplicationFormSet)(nil),      // 3: keyauth.department.ApplicationFormSet
	(*CreateDepartmentRequest)(nil), // 4: keyauth.department.CreateDepartmentRequest
	(*role.Role)(nil),               // 5: keyauth.role.Role
	(ApplicationFormStatus)(0),      // 6: keyauth.department.ApplicationFormStatus
	(*JoinDepartmentRequest)(nil),   // 7: keyauth.department.JoinDepartmentRequest
}
var file_pkg_department_pb_department_proto_depIdxs = []int32{
	4, // 0: keyauth.department.Department.data:type_name -> keyauth.department.CreateDepartmentRequest
	5, // 1: keyauth.department.Department.default_role:type_name -> keyauth.role.Role
	0, // 2: keyauth.department.Set.items:type_name -> keyauth.department.Department
	6, // 3: keyauth.department.ApplicationForm.status:type_name -> keyauth.department.ApplicationFormStatus
	7, // 4: keyauth.department.ApplicationForm.data:type_name -> keyauth.department.JoinDepartmentRequest
	2, // 5: keyauth.department.ApplicationFormSet.items:type_name -> keyauth.department.ApplicationForm
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_department_pb_department_proto_init() }
func file_pkg_department_pb_department_proto_init() {
	if File_pkg_department_pb_department_proto != nil {
		return
	}
	file_pkg_department_pb_enum_proto_init()
	file_pkg_department_pb_request_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_department_pb_department_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Department); i {
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
		file_pkg_department_pb_department_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_department_pb_department_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationForm); i {
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
		file_pkg_department_pb_department_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationFormSet); i {
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
			RawDescriptor: file_pkg_department_pb_department_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_department_pb_department_proto_goTypes,
		DependencyIndexes: file_pkg_department_pb_department_proto_depIdxs,
		MessageInfos:      file_pkg_department_pb_department_proto_msgTypes,
	}.Build()
	File_pkg_department_pb_department_proto = out.File
	file_pkg_department_pb_department_proto_rawDesc = nil
	file_pkg_department_pb_department_proto_goTypes = nil
	file_pkg_department_pb_department_proto_depIdxs = nil
}
