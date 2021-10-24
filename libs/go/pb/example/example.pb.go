//*
// Example entity Protocol Buffers
// Author: Zverev Valeriy <zverevvaleriy@gmail.com>

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: protos/example/example.proto

package example

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExampleMessage_Gender int32

const (
	ExampleMessage_NONE   ExampleMessage_Gender = 0
	ExampleMessage_MALE   ExampleMessage_Gender = 1
	ExampleMessage_FEMALE ExampleMessage_Gender = 2
)

// Enum value maps for ExampleMessage_Gender.
var (
	ExampleMessage_Gender_name = map[int32]string{
		0: "NONE",
		1: "MALE",
		2: "FEMALE",
	}
	ExampleMessage_Gender_value = map[string]int32{
		"NONE":   0,
		"MALE":   1,
		"FEMALE": 2,
	}
)

func (x ExampleMessage_Gender) Enum() *ExampleMessage_Gender {
	p := new(ExampleMessage_Gender)
	*p = x
	return p
}

func (x ExampleMessage_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExampleMessage_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_example_example_proto_enumTypes[0].Descriptor()
}

func (ExampleMessage_Gender) Type() protoreflect.EnumType {
	return &file_protos_example_example_proto_enumTypes[0]
}

func (x ExampleMessage_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExampleMessage_Gender.Descriptor instead.
func (ExampleMessage_Gender) EnumDescriptor() ([]byte, []int) {
	return file_protos_example_example_proto_rawDescGZIP(), []int{0, 0}
}

type ExampleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id          int32                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email       string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	Surname     string                 `protobuf:"bytes,5,opt,name=surname,proto3" json:"surname,omitempty"`
	DateOfBirth *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Gender      ExampleMessage_Gender  `protobuf:"varint,7,opt,name=gender,proto3,enum=Example.ExampleMessage_Gender" json:"gender,omitempty"`
}

func (x *ExampleMessage) Reset() {
	*x = ExampleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_example_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleMessage) ProtoMessage() {}

func (x *ExampleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protos_example_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleMessage.ProtoReflect.Descriptor instead.
func (*ExampleMessage) Descriptor() ([]byte, []int) {
	return file_protos_example_example_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExampleMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ExampleMessage) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ExampleMessage) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *ExampleMessage) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *ExampleMessage) GetDateOfBirth() *timestamppb.Timestamp {
	if x != nil {
		return x.DateOfBirth
	}
	return nil
}

func (x *ExampleMessage) GetGender() ExampleMessage_Gender {
	if x != nil {
		return x.Gender
	}
	return ExampleMessage_NONE
}

var File_protos_example_example_proto protoreflect.FileDescriptor

var file_protos_example_example_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x02, 0x0a, 0x0e, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e,
	0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x36,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e,
	0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x28, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41,
	0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02,
	0x42, 0x27, 0x5a, 0x0b, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0xaa,
	0x02, 0x0a, 0x50, 0x62, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0xca, 0x02, 0x0a, 0x50,
	0x62, 0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protos_example_example_proto_rawDescOnce sync.Once
	file_protos_example_example_proto_rawDescData = file_protos_example_example_proto_rawDesc
)

func file_protos_example_example_proto_rawDescGZIP() []byte {
	file_protos_example_example_proto_rawDescOnce.Do(func() {
		file_protos_example_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_example_example_proto_rawDescData)
	})
	return file_protos_example_example_proto_rawDescData
}

var file_protos_example_example_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_example_example_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_example_example_proto_goTypes = []interface{}{
	(ExampleMessage_Gender)(0),    // 0: Example.ExampleMessage.Gender
	(*ExampleMessage)(nil),        // 1: Example.ExampleMessage
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_protos_example_example_proto_depIdxs = []int32{
	2, // 0: Example.ExampleMessage.last_updated:type_name -> google.protobuf.Timestamp
	2, // 1: Example.ExampleMessage.date_of_birth:type_name -> google.protobuf.Timestamp
	0, // 2: Example.ExampleMessage.gender:type_name -> Example.ExampleMessage.Gender
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_example_example_proto_init() }
func file_protos_example_example_proto_init() {
	if File_protos_example_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_example_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleMessage); i {
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
			RawDescriptor: file_protos_example_example_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_example_example_proto_goTypes,
		DependencyIndexes: file_protos_example_example_proto_depIdxs,
		EnumInfos:         file_protos_example_example_proto_enumTypes,
		MessageInfos:      file_protos_example_example_proto_msgTypes,
	}.Build()
	File_protos_example_example_proto = out.File
	file_protos_example_example_proto_rawDesc = nil
	file_protos_example_example_proto_goTypes = nil
	file_protos_example_example_proto_depIdxs = nil
}
