// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.1
// source: frequentationPB/frequentation.proto

package frequentationPB

import (
	_ "github.com/golang/protobuf/ptypes/any"
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

type FrequentationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Zipcode string `protobuf:"bytes,1,opt,name=zipcode,proto3" json:"zipcode,omitempty"`
}

func (x *FrequentationRequest) Reset() {
	*x = FrequentationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frequentationPB_frequentation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrequentationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrequentationRequest) ProtoMessage() {}

func (x *FrequentationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_frequentationPB_frequentation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrequentationRequest.ProtoReflect.Descriptor instead.
func (*FrequentationRequest) Descriptor() ([]byte, []int) {
	return file_frequentationPB_frequentation_proto_rawDescGZIP(), []int{0}
}

func (x *FrequentationRequest) GetZipcode() string {
	if x != nil {
		return x.Zipcode
	}
	return ""
}

type FrequentationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UicCode string `protobuf:"bytes,1,opt,name=uic_code,json=uicCode,proto3" json:"uic_code,omitempty"`
	Zipcode string `protobuf:"bytes,2,opt,name=zipcode,proto3" json:"zipcode,omitempty"`
	A2015   int32  `protobuf:"varint,3,opt,name=a2015,proto3" json:"a2015,omitempty"`
	A2016   int32  `protobuf:"varint,4,opt,name=a2016,proto3" json:"a2016,omitempty"`
	A2017   int32  `protobuf:"varint,5,opt,name=a2017,proto3" json:"a2017,omitempty"`
	A2018   int32  `protobuf:"varint,6,opt,name=a2018,proto3" json:"a2018,omitempty"`
	A2019   int32  `protobuf:"varint,7,opt,name=a2019,proto3" json:"a2019,omitempty"`
	A2020   int32  `protobuf:"varint,8,opt,name=a2020,proto3" json:"a2020,omitempty"`
	A2021   int32  `protobuf:"varint,9,opt,name=a2021,proto3" json:"a2021,omitempty"`
}

func (x *FrequentationResponse) Reset() {
	*x = FrequentationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frequentationPB_frequentation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrequentationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrequentationResponse) ProtoMessage() {}

func (x *FrequentationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_frequentationPB_frequentation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrequentationResponse.ProtoReflect.Descriptor instead.
func (*FrequentationResponse) Descriptor() ([]byte, []int) {
	return file_frequentationPB_frequentation_proto_rawDescGZIP(), []int{1}
}

func (x *FrequentationResponse) GetUicCode() string {
	if x != nil {
		return x.UicCode
	}
	return ""
}

func (x *FrequentationResponse) GetZipcode() string {
	if x != nil {
		return x.Zipcode
	}
	return ""
}

func (x *FrequentationResponse) GetA2015() int32 {
	if x != nil {
		return x.A2015
	}
	return 0
}

func (x *FrequentationResponse) GetA2016() int32 {
	if x != nil {
		return x.A2016
	}
	return 0
}

func (x *FrequentationResponse) GetA2017() int32 {
	if x != nil {
		return x.A2017
	}
	return 0
}

func (x *FrequentationResponse) GetA2018() int32 {
	if x != nil {
		return x.A2018
	}
	return 0
}

func (x *FrequentationResponse) GetA2019() int32 {
	if x != nil {
		return x.A2019
	}
	return 0
}

func (x *FrequentationResponse) GetA2020() int32 {
	if x != nil {
		return x.A2020
	}
	return 0
}

func (x *FrequentationResponse) GetA2021() int32 {
	if x != nil {
		return x.A2021
	}
	return 0
}

type FrequentationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []*FrequentationResponse `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty"`
}

func (x *FrequentationList) Reset() {
	*x = FrequentationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frequentationPB_frequentation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrequentationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrequentationList) ProtoMessage() {}

func (x *FrequentationList) ProtoReflect() protoreflect.Message {
	mi := &file_frequentationPB_frequentation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrequentationList.ProtoReflect.Descriptor instead.
func (*FrequentationList) Descriptor() ([]byte, []int) {
	return file_frequentationPB_frequentation_proto_rawDescGZIP(), []int{2}
}

func (x *FrequentationList) GetValue() []*FrequentationResponse {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_frequentationPB_frequentation_proto protoreflect.FileDescriptor

var file_frequentationPB_frequentation_proto_rawDesc = []byte{
	0x0a, 0x23, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x42, 0x2f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x30, 0x0a, 0x14, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0xe6, 0x01, 0x0a, 0x15, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x75,
	0x69, 0x63, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75,
	0x69, 0x63, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x32, 0x30, 0x31, 0x35, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x61, 0x32, 0x30, 0x31, 0x35, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x32, 0x30, 0x31, 0x36, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x32, 0x30, 0x31, 0x36, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x32, 0x30, 0x31, 0x37, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x32, 0x30,
	0x31, 0x37, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x32, 0x30, 0x31, 0x38, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x61, 0x32, 0x30, 0x31, 0x38, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x32, 0x30, 0x31,
	0x39, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x32, 0x30, 0x31, 0x39, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x32, 0x30, 0x32, 0x30, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61,
	0x32, 0x30, 0x32, 0x30, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x32, 0x30, 0x32, 0x31, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x32, 0x30, 0x32, 0x31, 0x22, 0x4f, 0x0a, 0x11, 0x46, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x3a, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x66, 0x0a, 0x0d, 0x46,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x0c,
	0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x2e, 0x66,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x42, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_frequentationPB_frequentation_proto_rawDescOnce sync.Once
	file_frequentationPB_frequentation_proto_rawDescData = file_frequentationPB_frequentation_proto_rawDesc
)

func file_frequentationPB_frequentation_proto_rawDescGZIP() []byte {
	file_frequentationPB_frequentation_proto_rawDescOnce.Do(func() {
		file_frequentationPB_frequentation_proto_rawDescData = protoimpl.X.CompressGZIP(file_frequentationPB_frequentation_proto_rawDescData)
	})
	return file_frequentationPB_frequentation_proto_rawDescData
}

var file_frequentationPB_frequentation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_frequentationPB_frequentation_proto_goTypes = []interface{}{
	(*FrequentationRequest)(nil),  // 0: frequentation.FrequentationRequest
	(*FrequentationResponse)(nil), // 1: frequentation.FrequentationResponse
	(*FrequentationList)(nil),     // 2: frequentation.FrequentationList
}
var file_frequentationPB_frequentation_proto_depIdxs = []int32{
	1, // 0: frequentation.FrequentationList.Value:type_name -> frequentation.FrequentationResponse
	0, // 1: frequentation.Frequentation.ReadStations:input_type -> frequentation.FrequentationRequest
	2, // 2: frequentation.Frequentation.ReadStations:output_type -> frequentation.FrequentationList
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_frequentationPB_frequentation_proto_init() }
func file_frequentationPB_frequentation_proto_init() {
	if File_frequentationPB_frequentation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_frequentationPB_frequentation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrequentationRequest); i {
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
		file_frequentationPB_frequentation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrequentationResponse); i {
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
		file_frequentationPB_frequentation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrequentationList); i {
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
			RawDescriptor: file_frequentationPB_frequentation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_frequentationPB_frequentation_proto_goTypes,
		DependencyIndexes: file_frequentationPB_frequentation_proto_depIdxs,
		MessageInfos:      file_frequentationPB_frequentation_proto_msgTypes,
	}.Build()
	File_frequentationPB_frequentation_proto = out.File
	file_frequentationPB_frequentation_proto_rawDesc = nil
	file_frequentationPB_frequentation_proto_goTypes = nil
	file_frequentationPB_frequentation_proto_depIdxs = nil
}
