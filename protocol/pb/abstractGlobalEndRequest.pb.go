// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: abstractGlobalEndRequest.proto

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
type AbstractGlobalEndRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbstractTransactionRequest *AbstractTransactionRequestProto `protobuf:"bytes,1,opt,name=abstractTransactionRequest,proto3" json:"abstractTransactionRequest,omitempty"`
	Xid                        string                           `protobuf:"bytes,2,opt,name=xid,proto3" json:"xid,omitempty"`
	ExtraData                  string                           `protobuf:"bytes,3,opt,name=extraData,proto3" json:"extraData,omitempty"`
}

func (x *AbstractGlobalEndRequestProto) Reset() {
	*x = AbstractGlobalEndRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_abstractGlobalEndRequest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbstractGlobalEndRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbstractGlobalEndRequestProto) ProtoMessage() {}

func (x *AbstractGlobalEndRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_abstractGlobalEndRequest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbstractGlobalEndRequestProto.ProtoReflect.Descriptor instead.
func (*AbstractGlobalEndRequestProto) Descriptor() ([]byte, []int) {
	return file_abstractGlobalEndRequest_proto_rawDescGZIP(), []int{0}
}

func (x *AbstractGlobalEndRequestProto) GetAbstractTransactionRequest() *AbstractTransactionRequestProto {
	if x != nil {
		return x.AbstractTransactionRequest
	}
	return nil
}

func (x *AbstractGlobalEndRequestProto) GetXid() string {
	if x != nil {
		return x.Xid
	}
	return ""
}

func (x *AbstractGlobalEndRequestProto) GetExtraData() string {
	if x != nil {
		return x.ExtraData
	}
	return ""
}

var File_abstractGlobalEndRequest_proto protoreflect.FileDescriptor

var file_abstractGlobalEndRequest_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1a, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x20, 0x61, 0x62,
	0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc,
	0x01, 0x0a, 0x1d, 0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x7b, 0x0a, 0x1a, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x52, 0x1a, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x78, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x78, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x42, 0x49, 0x0a,
	0x26, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x42, 0x18, 0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x50, 0x01, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_abstractGlobalEndRequest_proto_rawDescOnce sync.Once
	file_abstractGlobalEndRequest_proto_rawDescData = file_abstractGlobalEndRequest_proto_rawDesc
)

func file_abstractGlobalEndRequest_proto_rawDescGZIP() []byte {
	file_abstractGlobalEndRequest_proto_rawDescOnce.Do(func() {
		file_abstractGlobalEndRequest_proto_rawDescData = protoimpl.X.CompressGZIP(file_abstractGlobalEndRequest_proto_rawDescData)
	})
	return file_abstractGlobalEndRequest_proto_rawDescData
}

var file_abstractGlobalEndRequest_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_abstractGlobalEndRequest_proto_goTypes = []interface{}{
	(*AbstractGlobalEndRequestProto)(nil),   // 0: io.seata.protocol.protobuf.AbstractGlobalEndRequestProto
	(*AbstractTransactionRequestProto)(nil), // 1: io.seata.protocol.protobuf.AbstractTransactionRequestProto
}
var file_abstractGlobalEndRequest_proto_depIdxs = []int32{
	1, // 0: io.seata.protocol.protobuf.AbstractGlobalEndRequestProto.abstractTransactionRequest:type_name -> io.seata.protocol.protobuf.AbstractTransactionRequestProto
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_abstractGlobalEndRequest_proto_init() }
func file_abstractGlobalEndRequest_proto_init() {
	if File_abstractGlobalEndRequest_proto != nil {
		return
	}
	file_abstractTransactionRequest_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_abstractGlobalEndRequest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbstractGlobalEndRequestProto); i {
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
			RawDescriptor: file_abstractGlobalEndRequest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_abstractGlobalEndRequest_proto_goTypes,
		DependencyIndexes: file_abstractGlobalEndRequest_proto_depIdxs,
		MessageInfos:      file_abstractGlobalEndRequest_proto_msgTypes,
	}.Build()
	File_abstractGlobalEndRequest_proto = out.File
	file_abstractGlobalEndRequest_proto_rawDesc = nil
	file_abstractGlobalEndRequest_proto_goTypes = nil
	file_abstractGlobalEndRequest_proto_depIdxs = nil
}
