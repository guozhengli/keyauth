//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/policy/pb/enum.proto

package policy

import (
	proto "github.com/golang/protobuf/proto"
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

type PolicyType int32

const (
	PolicyType_NULL PolicyType = 0
	// CustomPolicy (custom) 用户自己定义的策略
	PolicyType_CUSTOM PolicyType = 1
	// BuildInPolicy (build_in) 系统内部逻辑, 不允许用户看到并修改
	PolicyType_BUILD_IN PolicyType = 2
)

// Enum value maps for PolicyType.
var (
	PolicyType_name = map[int32]string{
		0: "NULL",
		1: "CUSTOM",
		2: "BUILD_IN",
	}
	PolicyType_value = map[string]int32{
		"NULL":     0,
		"CUSTOM":   1,
		"BUILD_IN": 2,
	}
)

func (x PolicyType) Enum() *PolicyType {
	p := new(PolicyType)
	*p = x
	return p
}

func (x PolicyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PolicyType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_policy_pb_enum_proto_enumTypes[0].Descriptor()
}

func (PolicyType) Type() protoreflect.EnumType {
	return &file_pkg_policy_pb_enum_proto_enumTypes[0]
}

func (x PolicyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PolicyType.Descriptor instead.
func (PolicyType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_policy_pb_enum_proto_rawDescGZIP(), []int{0}
}

var File_pkg_policy_pb_enum_proto protoreflect.FileDescriptor

var file_pkg_policy_pb_enum_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x62, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6b, 0x65, 0x79, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2a, 0x30, 0x0a, 0x0a, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x55, 0x4c, 0x4c,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x49, 0x4e, 0x10, 0x02, 0x42, 0x2a, 0x5a, 0x28,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_policy_pb_enum_proto_rawDescOnce sync.Once
	file_pkg_policy_pb_enum_proto_rawDescData = file_pkg_policy_pb_enum_proto_rawDesc
)

func file_pkg_policy_pb_enum_proto_rawDescGZIP() []byte {
	file_pkg_policy_pb_enum_proto_rawDescOnce.Do(func() {
		file_pkg_policy_pb_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_policy_pb_enum_proto_rawDescData)
	})
	return file_pkg_policy_pb_enum_proto_rawDescData
}

var file_pkg_policy_pb_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_policy_pb_enum_proto_goTypes = []interface{}{
	(PolicyType)(0), // 0: keyauth.policy.PolicyType
}
var file_pkg_policy_pb_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_policy_pb_enum_proto_init() }
func file_pkg_policy_pb_enum_proto_init() {
	if File_pkg_policy_pb_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_policy_pb_enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_policy_pb_enum_proto_goTypes,
		DependencyIndexes: file_pkg_policy_pb_enum_proto_depIdxs,
		EnumInfos:         file_pkg_policy_pb_enum_proto_enumTypes,
	}.Build()
	File_pkg_policy_pb_enum_proto = out.File
	file_pkg_policy_pb_enum_proto_rawDesc = nil
	file_pkg_policy_pb_enum_proto_goTypes = nil
	file_pkg_policy_pb_enum_proto_depIdxs = nil
}
