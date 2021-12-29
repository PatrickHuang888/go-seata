// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: abstractBranchEndRequest.proto

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
type AbstractBranchEndRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbstractTransactionRequest *AbstractTransactionRequestProto `protobuf:"bytes,1,opt,name=abstractTransactionRequest,proto3" json:"abstractTransactionRequest,omitempty"`
	Xid                        string                           `protobuf:"bytes,2,opt,name=xid,proto3" json:"xid,omitempty"`
	//*
	// The Branch id.
	BranchId int64 `protobuf:"varint,3,opt,name=branchId,proto3" json:"branchId,omitempty"`
	//*
	// The Branch type.
	BranchType BranchTypeProto `protobuf:"varint,4,opt,name=branchType,proto3,enum=io.seata.protocol.protobuf.BranchTypeProto" json:"branchType,omitempty"`
	//*
	// The Resource id.
	ResourceId string `protobuf:"bytes,5,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
	//*
	// The Application data.
	ApplicationData string `protobuf:"bytes,6,opt,name=applicationData,proto3" json:"applicationData,omitempty"`
}

func (x *AbstractBranchEndRequestProto) Reset() {
	*x = AbstractBranchEndRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_abstractBranchEndRequest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbstractBranchEndRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbstractBranchEndRequestProto) ProtoMessage() {}

func (x *AbstractBranchEndRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_abstractBranchEndRequest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbstractBranchEndRequestProto.ProtoReflect.Descriptor instead.
func (*AbstractBranchEndRequestProto) Descriptor() ([]byte, []int) {
	return file_abstractBranchEndRequest_proto_rawDescGZIP(), []int{0}
}

func (x *AbstractBranchEndRequestProto) GetAbstractTransactionRequest() *AbstractTransactionRequestProto {
	if x != nil {
		return x.AbstractTransactionRequest
	}
	return nil
}

func (x *AbstractBranchEndRequestProto) GetXid() string {
	if x != nil {
		return x.Xid
	}
	return ""
}

func (x *AbstractBranchEndRequestProto) GetBranchId() int64 {
	if x != nil {
		return x.BranchId
	}
	return 0
}

func (x *AbstractBranchEndRequestProto) GetBranchType() BranchTypeProto {
	if x != nil {
		return x.BranchType
	}
	return BranchTypeProto_AT
}

func (x *AbstractBranchEndRequestProto) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *AbstractBranchEndRequestProto) GetApplicationData() string {
	if x != nil {
		return x.ApplicationData
	}
	return ""
}

var File_abstractBranchEndRequest_proto protoreflect.FileDescriptor

var file_abstractBranchEndRequest_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1a, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x20, 0x61, 0x62,
	0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe1, 0x02, 0x0a, 0x1d, 0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x7b, 0x0a, 0x1a, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x1a, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x78, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x78, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x4b, 0x0a,
	0x0a, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2b, 0x2e, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0a,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x42, 0x49, 0x0a, 0x26, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x61,
	0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x42, 0x18,
	0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x45, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x01, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_abstractBranchEndRequest_proto_rawDescOnce sync.Once
	file_abstractBranchEndRequest_proto_rawDescData = file_abstractBranchEndRequest_proto_rawDesc
)

func file_abstractBranchEndRequest_proto_rawDescGZIP() []byte {
	file_abstractBranchEndRequest_proto_rawDescOnce.Do(func() {
		file_abstractBranchEndRequest_proto_rawDescData = protoimpl.X.CompressGZIP(file_abstractBranchEndRequest_proto_rawDescData)
	})
	return file_abstractBranchEndRequest_proto_rawDescData
}

var file_abstractBranchEndRequest_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_abstractBranchEndRequest_proto_goTypes = []interface{}{
	(*AbstractBranchEndRequestProto)(nil),   // 0: io.seata.protocol.protobuf.AbstractBranchEndRequestProto
	(*AbstractTransactionRequestProto)(nil), // 1: io.seata.protocol.protobuf.AbstractTransactionRequestProto
	(BranchTypeProto)(0),                    // 2: io.seata.protocol.protobuf.BranchTypeProto
}
var file_abstractBranchEndRequest_proto_depIdxs = []int32{
	1, // 0: io.seata.protocol.protobuf.AbstractBranchEndRequestProto.abstractTransactionRequest:type_name -> io.seata.protocol.protobuf.AbstractTransactionRequestProto
	2, // 1: io.seata.protocol.protobuf.AbstractBranchEndRequestProto.branchType:type_name -> io.seata.protocol.protobuf.BranchTypeProto
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_abstractBranchEndRequest_proto_init() }
func file_abstractBranchEndRequest_proto_init() {
	if File_abstractBranchEndRequest_proto != nil {
		return
	}
	file_abstractTransactionRequest_proto_init()
	file_branchType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_abstractBranchEndRequest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbstractBranchEndRequestProto); i {
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
			RawDescriptor: file_abstractBranchEndRequest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_abstractBranchEndRequest_proto_goTypes,
		DependencyIndexes: file_abstractBranchEndRequest_proto_depIdxs,
		MessageInfos:      file_abstractBranchEndRequest_proto_msgTypes,
	}.Build()
	File_abstractBranchEndRequest_proto = out.File
	file_abstractBranchEndRequest_proto_rawDesc = nil
	file_abstractBranchEndRequest_proto_goTypes = nil
	file_abstractBranchEndRequest_proto_depIdxs = nil
}