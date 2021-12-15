// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: abstractTransactionResponse.proto

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
type AbstractTransactionResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbstractResultMessage    *AbstractResultMessageProto   `protobuf:"bytes,1,opt,name=abstractResultMessage,proto3" json:"abstractResultMessage,omitempty"`
	TransactionExceptionCode TransactionExceptionCodeProto `protobuf:"varint,2,opt,name=transactionExceptionCode,proto3,enum=io.seata.protocol.protobuf.TransactionExceptionCodeProto" json:"transactionExceptionCode,omitempty"`
}

func (x *AbstractTransactionResponseProto) Reset() {
	*x = AbstractTransactionResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_abstractTransactionResponse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbstractTransactionResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbstractTransactionResponseProto) ProtoMessage() {}

func (x *AbstractTransactionResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_abstractTransactionResponse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbstractTransactionResponseProto.ProtoReflect.Descriptor instead.
func (*AbstractTransactionResponseProto) Descriptor() ([]byte, []int) {
	return file_abstractTransactionResponse_proto_rawDescGZIP(), []int{0}
}

func (x *AbstractTransactionResponseProto) GetAbstractResultMessage() *AbstractResultMessageProto {
	if x != nil {
		return x.AbstractResultMessage
	}
	return nil
}

func (x *AbstractTransactionResponseProto) GetTransactionExceptionCode() TransactionExceptionCodeProto {
	if x != nil {
		return x.TransactionExceptionCode
	}
	return TransactionExceptionCodeProto_Unknown
}

var File_abstractTransactionResponse_proto protoreflect.FileDescriptor

var file_abstractTransactionResponse_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a,
	0x1b, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x02, 0x0a,
	0x20, 0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x6c, 0x0a, 0x15, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x36, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x62,
	0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x15, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x75, 0x0a, 0x18, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78,
	0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x39, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x18, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x4c, 0x0a, 0x26, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61,
	0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x1b, 0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x01, 0x5a,
	0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_abstractTransactionResponse_proto_rawDescOnce sync.Once
	file_abstractTransactionResponse_proto_rawDescData = file_abstractTransactionResponse_proto_rawDesc
)

func file_abstractTransactionResponse_proto_rawDescGZIP() []byte {
	file_abstractTransactionResponse_proto_rawDescOnce.Do(func() {
		file_abstractTransactionResponse_proto_rawDescData = protoimpl.X.CompressGZIP(file_abstractTransactionResponse_proto_rawDescData)
	})
	return file_abstractTransactionResponse_proto_rawDescData
}

var file_abstractTransactionResponse_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_abstractTransactionResponse_proto_goTypes = []interface{}{
	(*AbstractTransactionResponseProto)(nil), // 0: io.seata.protocol.protobuf.AbstractTransactionResponseProto
	(*AbstractResultMessageProto)(nil),       // 1: io.seata.protocol.protobuf.AbstractResultMessageProto
	(TransactionExceptionCodeProto)(0),       // 2: io.seata.protocol.protobuf.TransactionExceptionCodeProto
}
var file_abstractTransactionResponse_proto_depIdxs = []int32{
	1, // 0: io.seata.protocol.protobuf.AbstractTransactionResponseProto.abstractResultMessage:type_name -> io.seata.protocol.protobuf.AbstractResultMessageProto
	2, // 1: io.seata.protocol.protobuf.AbstractTransactionResponseProto.transactionExceptionCode:type_name -> io.seata.protocol.protobuf.TransactionExceptionCodeProto
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_abstractTransactionResponse_proto_init() }
func file_abstractTransactionResponse_proto_init() {
	if File_abstractTransactionResponse_proto != nil {
		return
	}
	file_abstractResultMessage_proto_init()
	file_transactionExceptionCode_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_abstractTransactionResponse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbstractTransactionResponseProto); i {
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
			RawDescriptor: file_abstractTransactionResponse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_abstractTransactionResponse_proto_goTypes,
		DependencyIndexes: file_abstractTransactionResponse_proto_depIdxs,
		MessageInfos:      file_abstractTransactionResponse_proto_msgTypes,
	}.Build()
	File_abstractTransactionResponse_proto = out.File
	file_abstractTransactionResponse_proto_rawDesc = nil
	file_abstractTransactionResponse_proto_goTypes = nil
	file_abstractTransactionResponse_proto_depIdxs = nil
}
