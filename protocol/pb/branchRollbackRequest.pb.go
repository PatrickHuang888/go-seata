// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: branchRollbackRequest.proto

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
type BranchRollbackRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbstractBranchEndRequest *AbstractBranchEndRequestProto `protobuf:"bytes,1,opt,name=abstractBranchEndRequest,proto3" json:"abstractBranchEndRequest,omitempty"`
}

func (x *BranchRollbackRequestProto) Reset() {
	*x = BranchRollbackRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branchRollbackRequest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BranchRollbackRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchRollbackRequestProto) ProtoMessage() {}

func (x *BranchRollbackRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_branchRollbackRequest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchRollbackRequestProto.ProtoReflect.Descriptor instead.
func (*BranchRollbackRequestProto) Descriptor() ([]byte, []int) {
	return file_branchRollbackRequest_proto_rawDescGZIP(), []int{0}
}

func (x *BranchRollbackRequestProto) GetAbstractBranchEndRequest() *AbstractBranchEndRequestProto {
	if x != nil {
		return x.AbstractBranchEndRequest
	}
	return nil
}

var File_branchRollbackRequest_proto protoreflect.FileDescriptor

var file_branchRollbackRequest_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x69,
	0x6f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x1e, 0x61, 0x62, 0x73, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x1a, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x75, 0x0a, 0x18, 0x61, 0x62, 0x73, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x69, 0x6f, 0x2e,
	0x73, 0x65, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x18, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0x46, 0x0a, 0x26, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x42, 0x15, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x50, 0x01, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_branchRollbackRequest_proto_rawDescOnce sync.Once
	file_branchRollbackRequest_proto_rawDescData = file_branchRollbackRequest_proto_rawDesc
)

func file_branchRollbackRequest_proto_rawDescGZIP() []byte {
	file_branchRollbackRequest_proto_rawDescOnce.Do(func() {
		file_branchRollbackRequest_proto_rawDescData = protoimpl.X.CompressGZIP(file_branchRollbackRequest_proto_rawDescData)
	})
	return file_branchRollbackRequest_proto_rawDescData
}

var file_branchRollbackRequest_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_branchRollbackRequest_proto_goTypes = []interface{}{
	(*BranchRollbackRequestProto)(nil),    // 0: io.seata.protocol.protobuf.BranchRollbackRequestProto
	(*AbstractBranchEndRequestProto)(nil), // 1: io.seata.protocol.protobuf.AbstractBranchEndRequestProto
}
var file_branchRollbackRequest_proto_depIdxs = []int32{
	1, // 0: io.seata.protocol.protobuf.BranchRollbackRequestProto.abstractBranchEndRequest:type_name -> io.seata.protocol.protobuf.AbstractBranchEndRequestProto
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_branchRollbackRequest_proto_init() }
func file_branchRollbackRequest_proto_init() {
	if File_branchRollbackRequest_proto != nil {
		return
	}
	file_abstractBranchEndRequest_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_branchRollbackRequest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BranchRollbackRequestProto); i {
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
			RawDescriptor: file_branchRollbackRequest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_branchRollbackRequest_proto_goTypes,
		DependencyIndexes: file_branchRollbackRequest_proto_depIdxs,
		MessageInfos:      file_branchRollbackRequest_proto_msgTypes,
	}.Build()
	File_branchRollbackRequest_proto = out.File
	file_branchRollbackRequest_proto_rawDesc = nil
	file_branchRollbackRequest_proto_goTypes = nil
	file_branchRollbackRequest_proto_depIdxs = nil
}