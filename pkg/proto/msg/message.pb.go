// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: pkg/proto/msg/message.proto

package msg

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Query string `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Count int64  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Model string `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty"`
	Type  string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_msg_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_msg_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_pkg_proto_msg_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Message) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Message) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Message) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Message) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_pkg_proto_msg_message_proto protoreflect.FileDescriptor

var file_pkg_proto_msg_message_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x73, 0x67, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x73, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x73, 0x61, 0x2d, 0x66, 0x69, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x2f, 0x6d, 0x61, 0x73, 0x61, 0x2d, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_msg_message_proto_rawDescOnce sync.Once
	file_pkg_proto_msg_message_proto_rawDescData = file_pkg_proto_msg_message_proto_rawDesc
)

func file_pkg_proto_msg_message_proto_rawDescGZIP() []byte {
	file_pkg_proto_msg_message_proto_rawDescOnce.Do(func() {
		file_pkg_proto_msg_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_msg_message_proto_rawDescData)
	})
	return file_pkg_proto_msg_message_proto_rawDescData
}

var file_pkg_proto_msg_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_proto_msg_message_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: msg.Message
}
var file_pkg_proto_msg_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_msg_message_proto_init() }
func file_pkg_proto_msg_message_proto_init() {
	if File_pkg_proto_msg_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_msg_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_pkg_proto_msg_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_msg_message_proto_goTypes,
		DependencyIndexes: file_pkg_proto_msg_message_proto_depIdxs,
		MessageInfos:      file_pkg_proto_msg_message_proto_msgTypes,
	}.Build()
	File_pkg_proto_msg_message_proto = out.File
	file_pkg_proto_msg_message_proto_rawDesc = nil
	file_pkg_proto_msg_message_proto_goTypes = nil
	file_pkg_proto_msg_message_proto_depIdxs = nil
}
