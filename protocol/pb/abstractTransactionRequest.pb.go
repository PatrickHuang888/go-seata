// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: abstractTransactionRequest.proto

package pb

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

// PublishRequest is a publish request.
type AbstractTransactionRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbstractMessage *AbstractMessageProto `protobuf:"bytes,1,opt,name=abstractMessage,proto3" json:"abstractMessage,omitempty"`
}

func (x *AbstractTransactionRequestProto) Reset() {
	*x = AbstractTransactionRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_abstractTransactionRequest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbstractTransactionRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbstractTransactionRequestProto) ProtoMessage() {}

func (x *AbstractTransactionRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_abstractTransactionRequest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbstractTransactionRequestProto.ProtoReflect.Descriptor instead.
func (*AbstractTransactionRequestProto) Descriptor() ([]byte, []int) {
	return file_abstractTransactionRequest_proto_rawDescGZIP(), []int{0}
}

func (x *AbstractTransactionRequestProto) GetAbstractMessage() *AbstractMessageProto {
	if x != nil {
		return x.AbstractMessage
	}
	return nil
}

var File_abstractTransactionRequest_proto protoreflect.FileDescriptor

var file_abstractTransactionRequest_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x15,
	0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x1f, 0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x5a, 0x0a, 0x0f, 0x61, 0x62, 0x73, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x0f, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x4b, 0x0a, 0x26, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x61,
	0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x42, 0x1a,
	0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x01, 0x5a, 0x03, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_abstractTransactionRequest_proto_rawDescOnce sync.Once
	file_abstractTransactionRequest_proto_rawDescData = file_abstractTransactionRequest_proto_rawDesc
)

func file_abstractTransactionRequest_proto_rawDescGZIP() []byte {
	file_abstractTransactionRequest_proto_rawDescOnce.Do(func() {
		file_abstractTransactionRequest_proto_rawDescData = protoimpl.X.CompressGZIP(file_abstractTransactionRequest_proto_rawDescData)
	})
	return file_abstractTransactionRequest_proto_rawDescData
}

var file_abstractTransactionRequest_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_abstractTransactionRequest_proto_goTypes = []interface{}{
	(*AbstractTransactionRequestProto)(nil), // 0: io.seata.protocol.protobuf.AbstractTransactionRequestProto
	(*AbstractMessageProto)(nil),            // 1: io.seata.protocol.protobuf.AbstractMessageProto
}
var file_abstractTransactionRequest_proto_depIdxs = []int32{
	1, // 0: io.seata.protocol.protobuf.AbstractTransactionRequestProto.abstractMessage:type_name -> io.seata.protocol.protobuf.AbstractMessageProto
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_abstractTransactionRequest_proto_init() }
func file_abstractTransactionRequest_proto_init() {
	if File_abstractTransactionRequest_proto != nil {
		return
	}
	file_abstractMessage_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_abstractTransactionRequest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbstractTransactionRequestProto); i {
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
			RawDescriptor: file_abstractTransactionRequest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_abstractTransactionRequest_proto_goTypes,
		DependencyIndexes: file_abstractTransactionRequest_proto_depIdxs,
		MessageInfos:      file_abstractTransactionRequest_proto_msgTypes,
	}.Build()
	File_abstractTransactionRequest_proto = out.File
	file_abstractTransactionRequest_proto_rawDesc = nil
	file_abstractTransactionRequest_proto_goTypes = nil
	file_abstractTransactionRequest_proto_depIdxs = nil
}