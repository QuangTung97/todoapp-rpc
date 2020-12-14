// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: todoapp.proto

package todoapp

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	v1 "todoapp-rpc/rpc/lib/v1"
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

// TodoListRequest
type TodoListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// Optional
	Page int64 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	// Optional
	PageSize int64 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *TodoListRequest) Reset() {
	*x = TodoListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todoapp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoListRequest) ProtoMessage() {}

func (x *TodoListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todoapp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoListRequest.ProtoReflect.Descriptor instead.
func (*TodoListRequest) Descriptor() ([]byte, []int) {
	return file_todoapp_proto_rawDescGZIP(), []int{0}
}

func (x *TodoListRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *TodoListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *TodoListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// TodoData
type TodoData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	//
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *TodoData) Reset() {
	*x = TodoData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todoapp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoData) ProtoMessage() {}

func (x *TodoData) ProtoReflect() protoreflect.Message {
	mi := &file_todoapp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoData.ProtoReflect.Descriptor instead.
func (*TodoData) Descriptor() ([]byte, []int) {
	return file_todoapp_proto_rawDescGZIP(), []int{1}
}

func (x *TodoData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TodoData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TodoData) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

// TodoListResponse
type TodoListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	Todos []*TodoData `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
	//
	Total int64 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *TodoListResponse) Reset() {
	*x = TodoListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todoapp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoListResponse) ProtoMessage() {}

func (x *TodoListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todoapp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoListResponse.ProtoReflect.Descriptor instead.
func (*TodoListResponse) Descriptor() ([]byte, []int) {
	return file_todoapp_proto_rawDescGZIP(), []int{2}
}

func (x *TodoListResponse) GetTodos() []*TodoData {
	if x != nil {
		return x.Todos
	}
	return nil
}

func (x *TodoListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

// TodoGetRequest
type TodoGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type v1.TodoType `protobuf:"varint,1,opt,name=type,proto3,enum=lib.v1.TodoType" json:"type,omitempty"`
}

func (x *TodoGetRequest) Reset() {
	*x = TodoGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todoapp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoGetRequest) ProtoMessage() {}

func (x *TodoGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todoapp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoGetRequest.ProtoReflect.Descriptor instead.
func (*TodoGetRequest) Descriptor() ([]byte, []int) {
	return file_todoapp_proto_rawDescGZIP(), []int{3}
}

func (x *TodoGetRequest) GetType() v1.TodoType {
	if x != nil {
		return x.Type
	}
	return v1.TodoType_TODO_TYPE_UNSPECIFIED
}

var File_todoapp_proto protoreflect.FileDescriptor

var file_todoapp_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x6f, 0x64, 0x6f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x74, 0x6f, 0x64, 0x6f, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6c, 0x69, 0x62, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0f, 0x54, 0x6f, 0x64, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0x69, 0x0a, 0x08, 0x54, 0x6f, 0x64, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x54, 0x0a, 0x10,
	0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2a, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x22, 0x36, 0x0a, 0x0e, 0x54, 0x6f, 0x64, 0x6f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x32, 0x60, 0x0a, 0x0b, 0x54, 0x6f,
	0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1b, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x42, 0x24, 0x5a, 0x22,
	0x74, 0x6f, 0x64, 0x6f, 0x61, 0x70, 0x70, 0x2d, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x74, 0x6f, 0x64, 0x6f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x6f, 0x64, 0x6f, 0x61,
	0x70, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todoapp_proto_rawDescOnce sync.Once
	file_todoapp_proto_rawDescData = file_todoapp_proto_rawDesc
)

func file_todoapp_proto_rawDescGZIP() []byte {
	file_todoapp_proto_rawDescOnce.Do(func() {
		file_todoapp_proto_rawDescData = protoimpl.X.CompressGZIP(file_todoapp_proto_rawDescData)
	})
	return file_todoapp_proto_rawDescData
}

var file_todoapp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_todoapp_proto_goTypes = []interface{}{
	(*TodoListRequest)(nil),     // 0: todoapp.v1.TodoListRequest
	(*TodoData)(nil),            // 1: todoapp.v1.TodoData
	(*TodoListResponse)(nil),    // 2: todoapp.v1.TodoListResponse
	(*TodoGetRequest)(nil),      // 3: todoapp.v1.TodoGetRequest
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(v1.TodoType)(0),            // 5: lib.v1.TodoType
}
var file_todoapp_proto_depIdxs = []int32{
	4, // 0: todoapp.v1.TodoData.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: todoapp.v1.TodoListResponse.todos:type_name -> todoapp.v1.TodoData
	5, // 2: todoapp.v1.TodoGetRequest.type:type_name -> lib.v1.TodoType
	0, // 3: todoapp.v1.TodoService.List:input_type -> todoapp.v1.TodoListRequest
	2, // 4: todoapp.v1.TodoService.List:output_type -> todoapp.v1.TodoListResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_todoapp_proto_init() }
func file_todoapp_proto_init() {
	if File_todoapp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todoapp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoListRequest); i {
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
		file_todoapp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoData); i {
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
		file_todoapp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoListResponse); i {
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
		file_todoapp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoGetRequest); i {
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
			RawDescriptor: file_todoapp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todoapp_proto_goTypes,
		DependencyIndexes: file_todoapp_proto_depIdxs,
		MessageInfos:      file_todoapp_proto_msgTypes,
	}.Build()
	File_todoapp_proto = out.File
	file_todoapp_proto_rawDesc = nil
	file_todoapp_proto_goTypes = nil
	file_todoapp_proto_depIdxs = nil
}
