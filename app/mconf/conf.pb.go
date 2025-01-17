// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: app/mconf/pb/conf.proto

package mconf

import (
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

// Micro is service provider
type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 组名称
	// @gotags: bson:"_id" json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"_id"`
	// 组类型
	// @gotags: bson:"type" json:"type"
	Type Type `protobuf:"varint,2,opt,name=type,proto3,enum=infraboard.keyauth.mconf.Type" json:"type" bson:"type"`
	// 创建人
	// @gotags: bson:"creater" json:"creater"
	Creater string `protobuf:"bytes,3,opt,name=creater,proto3" json:"creater" bson:"creater"`
	// 创建的时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,4,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 描述信息
	// @gotags: bson:"description" json:"description"
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description" bson:"description"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_mconf_pb_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_app_mconf_pb_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_app_mconf_pb_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_GLOBAL
}

func (x *Group) GetCreater() string {
	if x != nil {
		return x.Creater
	}
	return ""
}

func (x *Group) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Group) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Item 健值项
type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 项ID
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 建的名称
	// @gotags: bson:"key" json:"key"
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key" bson:"key"`
	// 关联的组
	// @gotags: bson:"group" json:"group"
	Group string `protobuf:"bytes,3,opt,name=group,proto3" json:"group" bson:"group"`
	// 创建人
	// @gotags: bson:"creater" json:"creater"
	Creater string `protobuf:"bytes,4,opt,name=creater,proto3" json:"creater" bson:"creater"`
	// 创建的时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,5,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 创建的时间
	// @gotags: bson:"updater" json:"updater,omitempty"
	Updater int64 `protobuf:"varint,6,opt,name=updater,proto3" json:"updater,omitempty" bson:"updater"`
	// 更新时间
	// @gotags: bson:"update_at" json:"update_at,omitempty"
	UpdateAt int64 `protobuf:"varint,7,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty" bson:"update_at"`
	// 值对应的值
	// @gotags: bson:"value" json:"value"
	Value string `protobuf:"bytes,8,opt,name=value,proto3" json:"value" bson:"value"`
	// 描述信息
	// @gotags: bson:"description" json:"description,omitempty"
	Description string `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty" bson:"description"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_mconf_pb_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_app_mconf_pb_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_app_mconf_pb_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Item) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Item) GetCreater() string {
	if x != nil {
		return x.Creater
	}
	return ""
}

func (x *Item) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Item) GetUpdater() int64 {
	if x != nil {
		return x.Updater
	}
	return 0
}

func (x *Item) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Item) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Item) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ItemSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *ItemSet) Reset() {
	*x = ItemSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_mconf_pb_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemSet) ProtoMessage() {}

func (x *ItemSet) ProtoReflect() protoreflect.Message {
	mi := &file_app_mconf_pb_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemSet.ProtoReflect.Descriptor instead.
func (*ItemSet) Descriptor() ([]byte, []int) {
	return file_app_mconf_pb_conf_proto_rawDescGZIP(), []int{2}
}

func (x *ItemSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ItemSet) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type GroupSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*Group `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *GroupSet) Reset() {
	*x = GroupSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_mconf_pb_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupSet) ProtoMessage() {}

func (x *GroupSet) ProtoReflect() protoreflect.Message {
	mi := &file_app_mconf_pb_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupSet.ProtoReflect.Descriptor instead.
func (*GroupSet) Descriptor() ([]byte, []int) {
	return file_app_mconf_pb_conf_proto_rawDescGZIP(), []int{3}
}

func (x *GroupSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GroupSet) GetItems() []*Group {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_app_mconf_pb_conf_proto protoreflect.FileDescriptor

var file_app_mconf_pb_conf_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x2f, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x70, 0x62, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x63,
	0x6f, 0x6e, 0x66, 0x1a, 0x17, 0x61, 0x70, 0x70, 0x2f, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x70,
	0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a,
	0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x63,
	0x6f, 0x6e, 0x66, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe4, 0x01, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x55,
	0x0a, 0x07, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x34, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x57, 0x0a, 0x08, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x35, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x63, 0x6f, 0x6e,
	0x66, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x29,
	0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x61, 0x70, 0x70, 0x2f, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_app_mconf_pb_conf_proto_rawDescOnce sync.Once
	file_app_mconf_pb_conf_proto_rawDescData = file_app_mconf_pb_conf_proto_rawDesc
)

func file_app_mconf_pb_conf_proto_rawDescGZIP() []byte {
	file_app_mconf_pb_conf_proto_rawDescOnce.Do(func() {
		file_app_mconf_pb_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_mconf_pb_conf_proto_rawDescData)
	})
	return file_app_mconf_pb_conf_proto_rawDescData
}

var file_app_mconf_pb_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_app_mconf_pb_conf_proto_goTypes = []interface{}{
	(*Group)(nil),    // 0: infraboard.keyauth.mconf.Group
	(*Item)(nil),     // 1: infraboard.keyauth.mconf.Item
	(*ItemSet)(nil),  // 2: infraboard.keyauth.mconf.ItemSet
	(*GroupSet)(nil), // 3: infraboard.keyauth.mconf.GroupSet
	(Type)(0),        // 4: infraboard.keyauth.mconf.Type
}
var file_app_mconf_pb_conf_proto_depIdxs = []int32{
	4, // 0: infraboard.keyauth.mconf.Group.type:type_name -> infraboard.keyauth.mconf.Type
	1, // 1: infraboard.keyauth.mconf.ItemSet.items:type_name -> infraboard.keyauth.mconf.Item
	0, // 2: infraboard.keyauth.mconf.GroupSet.items:type_name -> infraboard.keyauth.mconf.Group
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_app_mconf_pb_conf_proto_init() }
func file_app_mconf_pb_conf_proto_init() {
	if File_app_mconf_pb_conf_proto != nil {
		return
	}
	file_app_mconf_pb_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_app_mconf_pb_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_app_mconf_pb_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_app_mconf_pb_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemSet); i {
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
		file_app_mconf_pb_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupSet); i {
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
			RawDescriptor: file_app_mconf_pb_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_mconf_pb_conf_proto_goTypes,
		DependencyIndexes: file_app_mconf_pb_conf_proto_depIdxs,
		MessageInfos:      file_app_mconf_pb_conf_proto_msgTypes,
	}.Build()
	File_app_mconf_pb_conf_proto = out.File
	file_app_mconf_pb_conf_proto_rawDesc = nil
	file_app_mconf_pb_conf_proto_goTypes = nil
	file_app_mconf_pb_conf_proto_depIdxs = nil
}
