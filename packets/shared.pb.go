// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.3
// source: messages/shared.proto

package packets

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EmptyMessage) Reset() {
	*x = EmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessage) ProtoMessage() {}

func (x *EmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessage.ProtoReflect.Descriptor instead.
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return file_messages_shared_proto_rawDescGZIP(), []int{0}
}

func (x *EmptyMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var file_messages_shared_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         51234,
		Name:          "is_generic",
		Tag:           "varint,51234,opt,name=is_generic",
		Filename:      "messages/shared.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional bool is_generic = 51234;
	E_IsGeneric = &file_messages_shared_proto_extTypes[0]
)

var File_messages_shared_proto protoreflect.FileDescriptor

var file_messages_shared_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x0c, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x40, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa2, 0x90, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x42, 0x0b, 0x5a, 0x09, 0x2e,
	0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_shared_proto_rawDescOnce sync.Once
	file_messages_shared_proto_rawDescData = file_messages_shared_proto_rawDesc
)

func file_messages_shared_proto_rawDescGZIP() []byte {
	file_messages_shared_proto_rawDescOnce.Do(func() {
		file_messages_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_shared_proto_rawDescData)
	})
	return file_messages_shared_proto_rawDescData
}

var file_messages_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_messages_shared_proto_goTypes = []interface{}{
	(*EmptyMessage)(nil),                // 0: EmptyMessage
	(*descriptorpb.MessageOptions)(nil), // 1: google.protobuf.MessageOptions
}
var file_messages_shared_proto_depIdxs = []int32{
	1, // 0: is_generic:extendee -> google.protobuf.MessageOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_messages_shared_proto_init() }
func file_messages_shared_proto_init() {
	if File_messages_shared_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessage); i {
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
			RawDescriptor: file_messages_shared_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_messages_shared_proto_goTypes,
		DependencyIndexes: file_messages_shared_proto_depIdxs,
		MessageInfos:      file_messages_shared_proto_msgTypes,
		ExtensionInfos:    file_messages_shared_proto_extTypes,
	}.Build()
	File_messages_shared_proto = out.File
	file_messages_shared_proto_rawDesc = nil
	file_messages_shared_proto_goTypes = nil
	file_messages_shared_proto_depIdxs = nil
}
