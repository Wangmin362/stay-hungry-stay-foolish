// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: protobuf/proto-syntax/01_basic_demo.proto

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

type StringMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *StringMessage) Reset() {
	*x = StringMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_proto_syntax_01_basic_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringMessage) ProtoMessage() {}

func (x *StringMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_proto_syntax_01_basic_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringMessage.ProtoReflect.Descriptor instead.
func (*StringMessage) Descriptor() ([]byte, []int) {
	return file_protobuf_proto_syntax_01_basic_demo_proto_rawDescGZIP(), []int{0}
}

func (x *StringMessage) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_protobuf_proto_syntax_01_basic_demo_proto protoreflect.FileDescriptor

var file_protobuf_proto_syntax_01_basic_demo_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2f, 0x30, 0x31, 0x5f, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x79, 0x6f, 0x75,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0d, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x32, 0x72, 0x0a, 0x0b, 0x59, 0x6f, 0x75, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x63, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x1e, 0x2e, 0x79, 0x6f, 0x75, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1e, 0x2e, 0x79, 0x6f, 0x75, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x65, 0x63, 0x68, 0x6f, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x6f, 0x75, 0x72, 0x6f, 0x72, 0x67, 0x2f, 0x79, 0x6f, 0x75, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x79, 0x6f,
	0x75, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_proto_syntax_01_basic_demo_proto_rawDescOnce sync.Once
	file_protobuf_proto_syntax_01_basic_demo_proto_rawDescData = file_protobuf_proto_syntax_01_basic_demo_proto_rawDesc
)

func file_protobuf_proto_syntax_01_basic_demo_proto_rawDescGZIP() []byte {
	file_protobuf_proto_syntax_01_basic_demo_proto_rawDescOnce.Do(func() {
		file_protobuf_proto_syntax_01_basic_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_proto_syntax_01_basic_demo_proto_rawDescData)
	})
	return file_protobuf_proto_syntax_01_basic_demo_proto_rawDescData
}

var file_protobuf_proto_syntax_01_basic_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobuf_proto_syntax_01_basic_demo_proto_goTypes = []interface{}{
	(*StringMessage)(nil), // 0: your.service.v1.StringMessage
}
var file_protobuf_proto_syntax_01_basic_demo_proto_depIdxs = []int32{
	0, // 0: your.service.v1.YourService.Echo:input_type -> your.service.v1.StringMessage
	0, // 1: your.service.v1.YourService.Echo:output_type -> your.service.v1.StringMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_proto_syntax_01_basic_demo_proto_init() }
func file_protobuf_proto_syntax_01_basic_demo_proto_init() {
	if File_protobuf_proto_syntax_01_basic_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_proto_syntax_01_basic_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringMessage); i {
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
			RawDescriptor: file_protobuf_proto_syntax_01_basic_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_proto_syntax_01_basic_demo_proto_goTypes,
		DependencyIndexes: file_protobuf_proto_syntax_01_basic_demo_proto_depIdxs,
		MessageInfos:      file_protobuf_proto_syntax_01_basic_demo_proto_msgTypes,
	}.Build()
	File_protobuf_proto_syntax_01_basic_demo_proto = out.File
	file_protobuf_proto_syntax_01_basic_demo_proto_rawDesc = nil
	file_protobuf_proto_syntax_01_basic_demo_proto_goTypes = nil
	file_protobuf_proto_syntax_01_basic_demo_proto_depIdxs = nil
}
