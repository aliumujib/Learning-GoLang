// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: maxapi.proto

package proto

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

type MaxApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *MaxApiRequest) Reset() {
	*x = MaxApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maxapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaxApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxApiRequest) ProtoMessage() {}

func (x *MaxApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_maxapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxApiRequest.ProtoReflect.Descriptor instead.
func (*MaxApiRequest) Descriptor() ([]byte, []int) {
	return file_maxapi_proto_rawDescGZIP(), []int{0}
}

func (x *MaxApiRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type MaxApiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Max int32 `protobuf:"varint,1,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *MaxApiResponse) Reset() {
	*x = MaxApiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maxapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaxApiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxApiResponse) ProtoMessage() {}

func (x *MaxApiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_maxapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxApiResponse.ProtoReflect.Descriptor instead.
func (*MaxApiResponse) Descriptor() ([]byte, []int) {
	return file_maxapi_proto_rawDescGZIP(), []int{1}
}

func (x *MaxApiResponse) GetMax() int32 {
	if x != nil {
		return x.Max
	}
	return 0
}

var File_maxapi_proto protoreflect.FileDescriptor

var file_maxapi_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x6d, 0x61, 0x78, 0x5f, 0x61, 0x70, 0x69, 0x22, 0x27, 0x0a, 0x0d, 0x4d, 0x61, 0x78, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x22, 0x0a, 0x0e, 0x4d, 0x61, 0x78, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x6d, 0x61, 0x78, 0x32, 0x4b, 0x0a, 0x0d, 0x4d, 0x61, 0x78, 0x41, 0x50, 0x49, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x03, 0x4d, 0x61, 0x78, 0x12, 0x16, 0x2e, 0x6d,
	0x61, 0x78, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x78, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4d,
	0x61, 0x78, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x47, 0x6f, 0x4c, 0x61, 0x6e, 0x67, 0x57, 0x65, 0x62, 0x44, 0x65, 0x76, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_maxapi_proto_rawDescOnce sync.Once
	file_maxapi_proto_rawDescData = file_maxapi_proto_rawDesc
)

func file_maxapi_proto_rawDescGZIP() []byte {
	file_maxapi_proto_rawDescOnce.Do(func() {
		file_maxapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_maxapi_proto_rawDescData)
	})
	return file_maxapi_proto_rawDescData
}

var file_maxapi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_maxapi_proto_goTypes = []interface{}{
	(*MaxApiRequest)(nil),  // 0: max_api.MaxApiRequest
	(*MaxApiResponse)(nil), // 1: max_api.MaxApiResponse
}
var file_maxapi_proto_depIdxs = []int32{
	0, // 0: max_api.MaxAPIService.Max:input_type -> max_api.MaxApiRequest
	1, // 1: max_api.MaxAPIService.Max:output_type -> max_api.MaxApiResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_maxapi_proto_init() }
func file_maxapi_proto_init() {
	if File_maxapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_maxapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaxApiRequest); i {
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
		file_maxapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaxApiResponse); i {
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
			RawDescriptor: file_maxapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_maxapi_proto_goTypes,
		DependencyIndexes: file_maxapi_proto_depIdxs,
		MessageInfos:      file_maxapi_proto_msgTypes,
	}.Build()
	File_maxapi_proto = out.File
	file_maxapi_proto_rawDesc = nil
	file_maxapi_proto_goTypes = nil
	file_maxapi_proto_depIdxs = nil
}