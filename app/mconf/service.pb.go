// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: app/mconf/pb/service.proto

package mconf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_app_mconf_pb_service_proto protoreflect.FileDescriptor

var file_app_mconf_pb_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x70, 0x2f, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x70, 0x62, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x1a, 0x1a, 0x61, 0x70, 0x70, 0x2f, 0x6d, 0x63, 0x6f, 0x6e,
	0x66, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x61, 0x70, 0x70, 0x2f, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x70, 0x62,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd6, 0x04, 0x0a, 0x07,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x63, 0x6f, 0x6e,
	0x66, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x5d, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x53, 0x65, 0x74, 0x12, 0x5c, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x2c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b,
	0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x5a, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x74, 0x12, 0x64,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x2f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x41, 0x64, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b,
	0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x53, 0x65, 0x74, 0x12, 0x6e, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x34, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b,
	0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x53, 0x65, 0x74, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_app_mconf_pb_service_proto_goTypes = []interface{}{
	(*CreateGroupRequest)(nil),         // 0: infraboard.keyauth.mconf.CreateGroupRequest
	(*QueryGroupRequest)(nil),          // 1: infraboard.keyauth.mconf.QueryGroupRequest
	(*DeleteGroupRequest)(nil),         // 2: infraboard.keyauth.mconf.DeleteGroupRequest
	(*QueryItemRequest)(nil),           // 3: infraboard.keyauth.mconf.QueryItemRequest
	(*AddItemToGroupRequest)(nil),      // 4: infraboard.keyauth.mconf.AddItemToGroupRequest
	(*RemoveItemFromGroupRequest)(nil), // 5: infraboard.keyauth.mconf.RemoveItemFromGroupRequest
	(*Group)(nil),                      // 6: infraboard.keyauth.mconf.Group
	(*GroupSet)(nil),                   // 7: infraboard.keyauth.mconf.GroupSet
	(*ItemSet)(nil),                    // 8: infraboard.keyauth.mconf.ItemSet
}
var file_app_mconf_pb_service_proto_depIdxs = []int32{
	0, // 0: infraboard.keyauth.mconf.Service.CreateGroup:input_type -> infraboard.keyauth.mconf.CreateGroupRequest
	1, // 1: infraboard.keyauth.mconf.Service.QueryGroup:input_type -> infraboard.keyauth.mconf.QueryGroupRequest
	2, // 2: infraboard.keyauth.mconf.Service.DeleteGroup:input_type -> infraboard.keyauth.mconf.DeleteGroupRequest
	3, // 3: infraboard.keyauth.mconf.Service.QueryItem:input_type -> infraboard.keyauth.mconf.QueryItemRequest
	4, // 4: infraboard.keyauth.mconf.Service.AddItemToGroup:input_type -> infraboard.keyauth.mconf.AddItemToGroupRequest
	5, // 5: infraboard.keyauth.mconf.Service.RemoveItemFromGroup:input_type -> infraboard.keyauth.mconf.RemoveItemFromGroupRequest
	6, // 6: infraboard.keyauth.mconf.Service.CreateGroup:output_type -> infraboard.keyauth.mconf.Group
	7, // 7: infraboard.keyauth.mconf.Service.QueryGroup:output_type -> infraboard.keyauth.mconf.GroupSet
	6, // 8: infraboard.keyauth.mconf.Service.DeleteGroup:output_type -> infraboard.keyauth.mconf.Group
	8, // 9: infraboard.keyauth.mconf.Service.QueryItem:output_type -> infraboard.keyauth.mconf.ItemSet
	8, // 10: infraboard.keyauth.mconf.Service.AddItemToGroup:output_type -> infraboard.keyauth.mconf.ItemSet
	8, // 11: infraboard.keyauth.mconf.Service.RemoveItemFromGroup:output_type -> infraboard.keyauth.mconf.ItemSet
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_mconf_pb_service_proto_init() }
func file_app_mconf_pb_service_proto_init() {
	if File_app_mconf_pb_service_proto != nil {
		return
	}
	file_app_mconf_pb_request_proto_init()
	file_app_mconf_pb_conf_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_mconf_pb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_mconf_pb_service_proto_goTypes,
		DependencyIndexes: file_app_mconf_pb_service_proto_depIdxs,
	}.Build()
	File_app_mconf_pb_service_proto = out.File
	file_app_mconf_pb_service_proto_rawDesc = nil
	file_app_mconf_pb_service_proto_goTypes = nil
	file_app_mconf_pb_service_proto_depIdxs = nil
}
