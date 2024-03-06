// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: idgen/v1/snowflake.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GenerateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GenerateRequest) Reset() {
	*x = GenerateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idgen_v1_snowflake_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateRequest) ProtoMessage() {}

func (x *GenerateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idgen_v1_snowflake_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateRequest.ProtoReflect.Descriptor instead.
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return file_idgen_v1_snowflake_proto_rawDescGZIP(), []int{0}
}

type GenerateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GenerateReply) Reset() {
	*x = GenerateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idgen_v1_snowflake_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateReply) ProtoMessage() {}

func (x *GenerateReply) ProtoReflect() protoreflect.Message {
	mi := &file_idgen_v1_snowflake_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateReply.ProtoReflect.Descriptor instead.
func (*GenerateReply) Descriptor() ([]byte, []int) {
	return file_idgen_v1_snowflake_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_idgen_v1_snowflake_proto protoreflect.FileDescriptor

var file_idgen_v1_snowflake_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x64, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6e, 0x6f, 0x77, 0x66,
	0x6c, 0x61, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x64, 0x67, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0x7d, 0x0a, 0x09, 0x53, 0x6e,
	0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x12, 0x70, 0x0a, 0x08, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x67, 0x65, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x67, 0x65, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64,
	0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x69, 0x64, 0x67,
	0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idgen_v1_snowflake_proto_rawDescOnce sync.Once
	file_idgen_v1_snowflake_proto_rawDescData = file_idgen_v1_snowflake_proto_rawDesc
)

func file_idgen_v1_snowflake_proto_rawDescGZIP() []byte {
	file_idgen_v1_snowflake_proto_rawDescOnce.Do(func() {
		file_idgen_v1_snowflake_proto_rawDescData = protoimpl.X.CompressGZIP(file_idgen_v1_snowflake_proto_rawDescData)
	})
	return file_idgen_v1_snowflake_proto_rawDescData
}

var file_idgen_v1_snowflake_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_idgen_v1_snowflake_proto_goTypes = []interface{}{
	(*GenerateRequest)(nil), // 0: api.idgen.v1.GenerateRequest
	(*GenerateReply)(nil),   // 1: api.idgen.v1.GenerateReply
}
var file_idgen_v1_snowflake_proto_depIdxs = []int32{
	0, // 0: api.idgen.v1.Snowflake.Generate:input_type -> api.idgen.v1.GenerateRequest
	1, // 1: api.idgen.v1.Snowflake.Generate:output_type -> api.idgen.v1.GenerateReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_idgen_v1_snowflake_proto_init() }
func file_idgen_v1_snowflake_proto_init() {
	if File_idgen_v1_snowflake_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idgen_v1_snowflake_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateRequest); i {
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
		file_idgen_v1_snowflake_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateReply); i {
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
			RawDescriptor: file_idgen_v1_snowflake_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idgen_v1_snowflake_proto_goTypes,
		DependencyIndexes: file_idgen_v1_snowflake_proto_depIdxs,
		MessageInfos:      file_idgen_v1_snowflake_proto_msgTypes,
	}.Build()
	File_idgen_v1_snowflake_proto = out.File
	file_idgen_v1_snowflake_proto_rawDesc = nil
	file_idgen_v1_snowflake_proto_goTypes = nil
	file_idgen_v1_snowflake_proto_depIdxs = nil
}
