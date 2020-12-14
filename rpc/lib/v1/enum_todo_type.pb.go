// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: enum_todo_type.proto

package lib

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

type TodoType int32

const (
	TodoType_TODO_TYPE_UNSPECIFIED TodoType = 0
	TodoType_TODO_TYPE_OK          TodoType = 1
)

// Enum value maps for TodoType.
var (
	TodoType_name = map[int32]string{
		0: "TODO_TYPE_UNSPECIFIED",
		1: "TODO_TYPE_OK",
	}
	TodoType_value = map[string]int32{
		"TODO_TYPE_UNSPECIFIED": 0,
		"TODO_TYPE_OK":          1,
	}
)

func (x TodoType) Enum() *TodoType {
	p := new(TodoType)
	*p = x
	return p
}

func (x TodoType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TodoType) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_todo_type_proto_enumTypes[0].Descriptor()
}

func (TodoType) Type() protoreflect.EnumType {
	return &file_enum_todo_type_proto_enumTypes[0]
}

func (x TodoType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TodoType.Descriptor instead.
func (TodoType) EnumDescriptor() ([]byte, []int) {
	return file_enum_todo_type_proto_rawDescGZIP(), []int{0}
}

var File_enum_todo_type_proto protoreflect.FileDescriptor

var file_enum_todo_type_proto_rawDesc = []byte{
	0x0a, 0x14, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x69, 0x62, 0x2e, 0x76, 0x31, 0x2a, 0x37,
	0x0a, 0x08, 0x54, 0x6f, 0x64, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x4f,
	0x44, 0x4f, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x4f, 0x44, 0x4f, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4f, 0x4b, 0x10, 0x01, 0x42, 0x1c, 0x5a, 0x1a, 0x74, 0x6f, 0x64, 0x6f, 0x61,
	0x70, 0x70, 0x2d, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x76,
	0x31, 0x3b, 0x6c, 0x69, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enum_todo_type_proto_rawDescOnce sync.Once
	file_enum_todo_type_proto_rawDescData = file_enum_todo_type_proto_rawDesc
)

func file_enum_todo_type_proto_rawDescGZIP() []byte {
	file_enum_todo_type_proto_rawDescOnce.Do(func() {
		file_enum_todo_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_enum_todo_type_proto_rawDescData)
	})
	return file_enum_todo_type_proto_rawDescData
}

var file_enum_todo_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enum_todo_type_proto_goTypes = []interface{}{
	(TodoType)(0), // 0: lib.v1.TodoType
}
var file_enum_todo_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enum_todo_type_proto_init() }
func file_enum_todo_type_proto_init() {
	if File_enum_todo_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enum_todo_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enum_todo_type_proto_goTypes,
		DependencyIndexes: file_enum_todo_type_proto_depIdxs,
		EnumInfos:         file_enum_todo_type_proto_enumTypes,
	}.Build()
	File_enum_todo_type_proto = out.File
	file_enum_todo_type_proto_rawDesc = nil
	file_enum_todo_type_proto_goTypes = nil
	file_enum_todo_type_proto_depIdxs = nil
}
