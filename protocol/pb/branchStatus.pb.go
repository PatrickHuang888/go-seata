// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: branchStatus.proto

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
type BranchStatusProto int32

const (
	// special for Unknown
	BranchStatusProto_BUnknown BranchStatusProto = 0
	// Registered to TC.
	BranchStatusProto_Registered BranchStatusProto = 1
	// Branch logic is successfully done at phase one.
	BranchStatusProto_PhaseOne_Done BranchStatusProto = 2
	// Branch logic is failed at phase one.
	BranchStatusProto_PhaseOne_Failed BranchStatusProto = 3
	// Branch logic is NOT reported for a timeout.
	BranchStatusProto_PhaseOne_Timeout BranchStatusProto = 4
	// Commit logic is successfully done at phase two.
	BranchStatusProto_PhaseTwo_Committed BranchStatusProto = 5
	// Commit logic is failed but retryable.
	BranchStatusProto_PhaseTwo_CommitFailed_Retryable BranchStatusProto = 6
	// Commit logic is failed and NOT retryable.
	BranchStatusProto_PhaseTwo_CommitFailed_Unretryable BranchStatusProto = 7
	// Rollback logic is successfully done at phase two.
	BranchStatusProto_PhaseTwo_Rollbacked BranchStatusProto = 8
	// Rollback logic is failed but retryable.
	BranchStatusProto_PhaseTwo_RollbackFailed_Retryable BranchStatusProto = 9
	// Rollback logic is failed but NOT retryable.
	BranchStatusProto_PhaseTwo_RollbackFailed_Unretryable BranchStatusProto = 10
)

// Enum value maps for BranchStatusProto.
var (
	BranchStatusProto_name = map[int32]string{
		0:  "BUnknown",
		1:  "Registered",
		2:  "PhaseOne_Done",
		3:  "PhaseOne_Failed",
		4:  "PhaseOne_Timeout",
		5:  "PhaseTwo_Committed",
		6:  "PhaseTwo_CommitFailed_Retryable",
		7:  "PhaseTwo_CommitFailed_Unretryable",
		8:  "PhaseTwo_Rollbacked",
		9:  "PhaseTwo_RollbackFailed_Retryable",
		10: "PhaseTwo_RollbackFailed_Unretryable",
	}
	BranchStatusProto_value = map[string]int32{
		"BUnknown":                            0,
		"Registered":                          1,
		"PhaseOne_Done":                       2,
		"PhaseOne_Failed":                     3,
		"PhaseOne_Timeout":                    4,
		"PhaseTwo_Committed":                  5,
		"PhaseTwo_CommitFailed_Retryable":     6,
		"PhaseTwo_CommitFailed_Unretryable":   7,
		"PhaseTwo_Rollbacked":                 8,
		"PhaseTwo_RollbackFailed_Retryable":   9,
		"PhaseTwo_RollbackFailed_Unretryable": 10,
	}
)

func (x BranchStatusProto) Enum() *BranchStatusProto {
	p := new(BranchStatusProto)
	*p = x
	return p
}

func (x BranchStatusProto) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BranchStatusProto) Descriptor() protoreflect.EnumDescriptor {
	return file_branchStatus_proto_enumTypes[0].Descriptor()
}

func (BranchStatusProto) Type() protoreflect.EnumType {
	return &file_branchStatus_proto_enumTypes[0]
}

func (x BranchStatusProto) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BranchStatusProto.Descriptor instead.
func (BranchStatusProto) EnumDescriptor() ([]byte, []int) {
	return file_branchStatus_proto_rawDescGZIP(), []int{0}
}

var File_branchStatus_proto protoreflect.FileDescriptor

var file_branchStatus_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2a, 0xbc, 0x02, 0x0a, 0x11, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x65, 0x64, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x68, 0x61, 0x73, 0x65, 0x4f, 0x6e, 0x65,
	0x5f, 0x44, 0x6f, 0x6e, 0x65, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x68, 0x61, 0x73, 0x65,
	0x4f, 0x6e, 0x65, 0x5f, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10,
	0x50, 0x68, 0x61, 0x73, 0x65, 0x4f, 0x6e, 0x65, 0x5f, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x68, 0x61, 0x73, 0x65, 0x54, 0x77, 0x6f, 0x5f, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x10, 0x05, 0x12, 0x23, 0x0a, 0x1f, 0x50, 0x68,
	0x61, 0x73, 0x65, 0x54, 0x77, 0x6f, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x46, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x5f, 0x52, 0x65, 0x74, 0x72, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x06, 0x12,
	0x25, 0x0a, 0x21, 0x50, 0x68, 0x61, 0x73, 0x65, 0x54, 0x77, 0x6f, 0x5f, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x55, 0x6e, 0x72, 0x65, 0x74, 0x72, 0x79,
	0x61, 0x62, 0x6c, 0x65, 0x10, 0x07, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x68, 0x61, 0x73, 0x65, 0x54,
	0x77, 0x6f, 0x5f, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x10, 0x08, 0x12,
	0x25, 0x0a, 0x21, 0x50, 0x68, 0x61, 0x73, 0x65, 0x54, 0x77, 0x6f, 0x5f, 0x52, 0x6f, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x52, 0x65, 0x74, 0x72, 0x79,
	0x61, 0x62, 0x6c, 0x65, 0x10, 0x09, 0x12, 0x27, 0x0a, 0x23, 0x50, 0x68, 0x61, 0x73, 0x65, 0x54,
	0x77, 0x6f, 0x5f, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x46, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x5f, 0x55, 0x6e, 0x72, 0x65, 0x74, 0x72, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x0a, 0x42,
	0x3d, 0x0a, 0x26, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x42, 0x0c, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x01, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_branchStatus_proto_rawDescOnce sync.Once
	file_branchStatus_proto_rawDescData = file_branchStatus_proto_rawDesc
)

func file_branchStatus_proto_rawDescGZIP() []byte {
	file_branchStatus_proto_rawDescOnce.Do(func() {
		file_branchStatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_branchStatus_proto_rawDescData)
	})
	return file_branchStatus_proto_rawDescData
}

var file_branchStatus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_branchStatus_proto_goTypes = []interface{}{
	(BranchStatusProto)(0), // 0: io.seata.protocol.protobuf.BranchStatusProto
}
var file_branchStatus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_branchStatus_proto_init() }
func file_branchStatus_proto_init() {
	if File_branchStatus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_branchStatus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_branchStatus_proto_goTypes,
		DependencyIndexes: file_branchStatus_proto_depIdxs,
		EnumInfos:         file_branchStatus_proto_enumTypes,
	}.Build()
	File_branchStatus_proto = out.File
	file_branchStatus_proto_rawDesc = nil
	file_branchStatus_proto_goTypes = nil
	file_branchStatus_proto_depIdxs = nil
}
