// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: main.proto

package adminv1

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

// protoc --proto_path=. --go_out=paths=source_relative:. main.proto
// 性能参数请求定义
type CreatePerfReportReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PopName string `protobuf:"bytes,1,opt,name=pop_name,json=popName,proto3" json:"pop_name,omitempty"`
	// 如果是公共服务，并非产品，那么无需传递租户ID
	TenantId *int64 `protobuf:"varint,2,opt,name=tenant_id,json=tenantId,proto3,oneof" json:"tenant_id,omitempty"`
	// 如果是公共服务，并非产品，那么无需传递租户名
	TenantName *string `protobuf:"bytes,3,opt,name=tenant_name,json=tenantName,proto3,oneof" json:"tenant_name,omitempty"`
	ModuleName string  `protobuf:"bytes,4,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	// 产品私有参数
	ExtendJson *string `protobuf:"bytes,12,opt,name=extend_json,json=extendJson,proto3,oneof" json:"extend_json,omitempty"`
	// 备注
	Remark *string `protobuf:"bytes,13,opt,name=remark,proto3,oneof" json:"remark,omitempty"`
	XAbc   int32   `protobuf:"varint,14,opt,name=_abc,json=Abc,proto3" json:"_abc,omitempty"`
	XAaa   *int32  `protobuf:"varint,15,opt,name=__aaa,json=Aaa,proto3,oneof" json:"__aaa,omitempty"`
	XAbaa_ int32   `protobuf:"varint,16,opt,name=__abaa_,json=Abaa,proto3" json:"__abaa_,omitempty"`
}

func (x *CreatePerfReportReq) Reset() {
	*x = CreatePerfReportReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePerfReportReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePerfReportReq) ProtoMessage() {}

func (x *CreatePerfReportReq) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePerfReportReq.ProtoReflect.Descriptor instead.
func (*CreatePerfReportReq) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePerfReportReq) GetPopName() string {
	if x != nil {
		return x.PopName
	}
	return ""
}

func (x *CreatePerfReportReq) GetTenantId() int64 {
	if x != nil && x.TenantId != nil {
		return *x.TenantId
	}
	return 0
}

func (x *CreatePerfReportReq) GetTenantName() string {
	if x != nil && x.TenantName != nil {
		return *x.TenantName
	}
	return ""
}

func (x *CreatePerfReportReq) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *CreatePerfReportReq) GetExtendJson() string {
	if x != nil && x.ExtendJson != nil {
		return *x.ExtendJson
	}
	return ""
}

func (x *CreatePerfReportReq) GetRemark() string {
	if x != nil && x.Remark != nil {
		return *x.Remark
	}
	return ""
}

func (x *CreatePerfReportReq) GetXAbc() int32 {
	if x != nil {
		return x.XAbc
	}
	return 0
}

func (x *CreatePerfReportReq) GetXAaa() int32 {
	if x != nil && x.XAaa != nil {
		return *x.XAaa
	}
	return 0
}

func (x *CreatePerfReportReq) GetXAbaa_() int32 {
	if x != nil {
		return x.XAbaa_
	}
	return 0
}

var File_main_proto protoreflect.FileDescriptor

var file_main_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x73, 0x6b,
	0x79, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0xe2, 0x02, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x66, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x09, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00,
	0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a,
	0x0b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x6a,
	0x73, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x4a, 0x73, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x11, 0x0a, 0x04, 0x5f, 0x61, 0x62, 0x63, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x41, 0x62, 0x63, 0x12, 0x17, 0x0a, 0x05, 0x5f, 0x5f,
	0x61, 0x61, 0x61, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x48, 0x04, 0x52, 0x03, 0x41, 0x61, 0x61,
	0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x07, 0x5f, 0x5f, 0x61, 0x62, 0x61, 0x61, 0x5f, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x41, 0x62, 0x61, 0x61, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x42, 0x08, 0x0a, 0x06, 0x58, 0x5f, 0x5f, 0x61, 0x61, 0x61, 0x42, 0x6c, 0x5a,
	0x6a, 0x67, 0x69, 0x74, 0x63, 0x64, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x73, 0x6b, 0x79, 0x67, 0x75,
	0x61, 0x72, 0x64, 0x6d, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6b, 0x79, 0x67, 0x75,
	0x61, 0x72, 0x64, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x6b, 0x79, 0x67, 0x75, 0x61, 0x72, 0x64,
	0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x6b, 0x79, 0x67, 0x75, 0x61, 0x72, 0x64,
	0x2f, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x76, 0x31, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_main_proto_rawDescOnce sync.Once
	file_main_proto_rawDescData = file_main_proto_rawDesc
)

func file_main_proto_rawDescGZIP() []byte {
	file_main_proto_rawDescOnce.Do(func() {
		file_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_main_proto_rawDescData)
	})
	return file_main_proto_rawDescData
}

var file_main_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_main_proto_goTypes = []interface{}{
	(*CreatePerfReportReq)(nil), // 0: skyguard.gatorcloud.healthchecker.admin.v1.CreatePerfReportReq
}
var file_main_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_main_proto_init() }
func file_main_proto_init() {
	if File_main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_main_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePerfReportReq); i {
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
	file_main_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_main_proto_goTypes,
		DependencyIndexes: file_main_proto_depIdxs,
		MessageInfos:      file_main_proto_msgTypes,
	}.Build()
	File_main_proto = out.File
	file_main_proto_rawDesc = nil
	file_main_proto_goTypes = nil
	file_main_proto_depIdxs = nil
}
