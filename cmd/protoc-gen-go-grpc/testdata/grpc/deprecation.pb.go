// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/deprecation.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	sync "sync"
)

const (
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 0)
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(0 - protoimpl.MinVersion)
)

var File_grpc_deprecation_proto protoreflect.FileDescriptor

var file_grpc_deprecation_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x0f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x6c,
	0x0a, 0x11, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x03, 0x88, 0x02, 0x01, 0x1a, 0x03, 0x88, 0x02, 0x01, 0x42, 0x3c, 0x5a, 0x3a,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_grpc_deprecation_proto_rawDescOnce sync.Once
	file_grpc_deprecation_proto_rawDescData = file_grpc_deprecation_proto_rawDesc
)

func file_grpc_deprecation_proto_rawDescGZIP() []byte {
	file_grpc_deprecation_proto_rawDescOnce.Do(func() {
		file_grpc_deprecation_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_deprecation_proto_rawDescData)
	})
	return file_grpc_deprecation_proto_rawDescData
}

var file_grpc_deprecation_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: goproto.protoc.grpc.Request
	(*Response)(nil), // 1: goproto.protoc.grpc.Response
}
var file_grpc_deprecation_proto_depIdxs = []int32{
	0, // goproto.protoc.grpc.DeprecatedService.DeprecatedCall:input_type -> goproto.protoc.grpc.Request
	1, // goproto.protoc.grpc.DeprecatedService.DeprecatedCall:output_type -> goproto.protoc.grpc.Response
	1, // starting offset of method output_type sub-list
	0, // starting offset of method input_type sub-list
	0, // starting offset of extension type_name sub-list
	0, // starting offset of extension extendee sub-list
	0, // starting offset of field type_name sub-list
}

func init() { file_grpc_deprecation_proto_init() }
func file_grpc_deprecation_proto_init() {
	if File_grpc_deprecation_proto != nil {
		return
	}
	file_grpc_grpc_proto_init()
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			RawDescriptor: file_grpc_deprecation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_deprecation_proto_goTypes,
		DependencyIndexes: file_grpc_deprecation_proto_depIdxs,
	}.Build()
	File_grpc_deprecation_proto = out.File
	file_grpc_deprecation_proto_rawDesc = nil
	file_grpc_deprecation_proto_goTypes = nil
	file_grpc_deprecation_proto_depIdxs = nil
}
