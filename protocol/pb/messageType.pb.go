// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: messageType.proto

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
type MessageTypeProto int32

const (
	MessageTypeProto_TYPE_GLOBAL_PRESERVED    MessageTypeProto = 0
	MessageTypeProto_TYPE_GLOBAL_BEGIN        MessageTypeProto = 1
	MessageTypeProto_TYPE_GLOBAL_BEGIN_RESULT MessageTypeProto = 2
	//*
	// The constant TYPE_GLOBAL_COMMIT.
	MessageTypeProto_TYPE_GLOBAL_COMMIT MessageTypeProto = 7
	//*
	// The constant TYPE_GLOBAL_COMMIT_RESULT.
	MessageTypeProto_TYPE_GLOBAL_COMMIT_RESULT MessageTypeProto = 8
	//*
	// The constant TYPE_GLOBAL_ROLLBACK.
	MessageTypeProto_TYPE_GLOBAL_ROLLBACK MessageTypeProto = 9
	//*
	// The constant TYPE_GLOBAL_ROLLBACK_RESULT.
	MessageTypeProto_TYPE_GLOBAL_ROLLBACK_RESULT MessageTypeProto = 10
	//*
	// The constant TYPE_GLOBAL_STATUS.
	MessageTypeProto_TYPE_GLOBAL_STATUS MessageTypeProto = 15
	//*
	// The constant TYPE_GLOBAL_STATUS_RESULT.
	MessageTypeProto_TYPE_GLOBAL_STATUS_RESULT MessageTypeProto = 16
	//*
	// The constant TYPE_GLOBAL_REPORT.
	MessageTypeProto_TYPE_GLOBAL_REPORT MessageTypeProto = 17
	//*
	// The constant TYPE_GLOBAL_REPORT_RESULT.
	MessageTypeProto_TYPE_GLOBAL_REPORT_RESULT MessageTypeProto = 18
	//*
	// The constant TYPE_GLOBAL_LOCK_QUERY.
	MessageTypeProto_TYPE_GLOBAL_LOCK_QUERY MessageTypeProto = 21
	//*
	// The constant TYPE_GLOBAL_LOCK_QUERY_RESULT.
	MessageTypeProto_TYPE_GLOBAL_LOCK_QUERY_RESULT MessageTypeProto = 22
	//*
	// The constant TYPE_BRANCH_COMMIT.
	MessageTypeProto_TYPE_BRANCH_COMMIT MessageTypeProto = 3
	//*
	// The constant TYPE_BRANCH_COMMIT_RESULT.
	MessageTypeProto_TYPE_BRANCH_COMMIT_RESULT MessageTypeProto = 4
	//*
	// The constant TYPE_BRANCH_ROLLBACK.
	MessageTypeProto_TYPE_BRANCH_ROLLBACK MessageTypeProto = 5
	//*
	// The constant TYPE_BRANCH_ROLLBACK_RESULT.
	MessageTypeProto_TYPE_BRANCH_ROLLBACK_RESULT MessageTypeProto = 6
	//*
	// The constant TYPE_BRANCH_REGISTER.
	MessageTypeProto_TYPE_BRANCH_REGISTER MessageTypeProto = 11
	//*
	// The constant TYPE_BRANCH_REGISTER_RESULT.
	MessageTypeProto_TYPE_BRANCH_REGISTER_RESULT MessageTypeProto = 12
	//*
	// The constant TYPE_BRANCH_STATUS_REPORT.
	MessageTypeProto_TYPE_BRANCH_STATUS_REPORT MessageTypeProto = 13
	//*
	// The constant TYPE_BRANCH_STATUS_REPORT_RESULT.
	MessageTypeProto_TYPE_BRANCH_STATUS_REPORT_RESULT MessageTypeProto = 14
	//*
	// The constant TYPE_SEATA_MERGE.
	MessageTypeProto_TYPE_SEATA_MERGE MessageTypeProto = 59
	//*
	// The constant TYPE_SEATA_MERGE_RESULT.
	MessageTypeProto_TYPE_SEATA_MERGE_RESULT MessageTypeProto = 60
	//*
	// The constant TYPE_REG_CLT.
	MessageTypeProto_TYPE_REG_CLT MessageTypeProto = 101
	//*
	// The constant TYPE_REG_CLT_RESULT.
	MessageTypeProto_TYPE_REG_CLT_RESULT MessageTypeProto = 102
	//*
	// The constant TYPE_REG_RM.
	MessageTypeProto_TYPE_REG_RM MessageTypeProto = 103
	//*
	// The constant TYPE_REG_RM_RESULT.
	MessageTypeProto_TYPE_REG_RM_RESULT MessageTypeProto = 104
	//*
	// The constant TYPE_UNDO_LOG_DELETE.
	MessageTypeProto_TYPE_UNDO_LOG_DELETE MessageTypeProto = 111
)

// Enum value maps for MessageTypeProto.
var (
	MessageTypeProto_name = map[int32]string{
		0:   "TYPE_GLOBAL_PRESERVED",
		1:   "TYPE_GLOBAL_BEGIN",
		2:   "TYPE_GLOBAL_BEGIN_RESULT",
		7:   "TYPE_GLOBAL_COMMIT",
		8:   "TYPE_GLOBAL_COMMIT_RESULT",
		9:   "TYPE_GLOBAL_ROLLBACK",
		10:  "TYPE_GLOBAL_ROLLBACK_RESULT",
		15:  "TYPE_GLOBAL_STATUS",
		16:  "TYPE_GLOBAL_STATUS_RESULT",
		17:  "TYPE_GLOBAL_REPORT",
		18:  "TYPE_GLOBAL_REPORT_RESULT",
		21:  "TYPE_GLOBAL_LOCK_QUERY",
		22:  "TYPE_GLOBAL_LOCK_QUERY_RESULT",
		3:   "TYPE_BRANCH_COMMIT",
		4:   "TYPE_BRANCH_COMMIT_RESULT",
		5:   "TYPE_BRANCH_ROLLBACK",
		6:   "TYPE_BRANCH_ROLLBACK_RESULT",
		11:  "TYPE_BRANCH_REGISTER",
		12:  "TYPE_BRANCH_REGISTER_RESULT",
		13:  "TYPE_BRANCH_STATUS_REPORT",
		14:  "TYPE_BRANCH_STATUS_REPORT_RESULT",
		59:  "TYPE_SEATA_MERGE",
		60:  "TYPE_SEATA_MERGE_RESULT",
		101: "TYPE_REG_CLT",
		102: "TYPE_REG_CLT_RESULT",
		103: "TYPE_REG_RM",
		104: "TYPE_REG_RM_RESULT",
		111: "TYPE_UNDO_LOG_DELETE",
	}
	MessageTypeProto_value = map[string]int32{
		"TYPE_GLOBAL_PRESERVED":            0,
		"TYPE_GLOBAL_BEGIN":                1,
		"TYPE_GLOBAL_BEGIN_RESULT":         2,
		"TYPE_GLOBAL_COMMIT":               7,
		"TYPE_GLOBAL_COMMIT_RESULT":        8,
		"TYPE_GLOBAL_ROLLBACK":             9,
		"TYPE_GLOBAL_ROLLBACK_RESULT":      10,
		"TYPE_GLOBAL_STATUS":               15,
		"TYPE_GLOBAL_STATUS_RESULT":        16,
		"TYPE_GLOBAL_REPORT":               17,
		"TYPE_GLOBAL_REPORT_RESULT":        18,
		"TYPE_GLOBAL_LOCK_QUERY":           21,
		"TYPE_GLOBAL_LOCK_QUERY_RESULT":    22,
		"TYPE_BRANCH_COMMIT":               3,
		"TYPE_BRANCH_COMMIT_RESULT":        4,
		"TYPE_BRANCH_ROLLBACK":             5,
		"TYPE_BRANCH_ROLLBACK_RESULT":      6,
		"TYPE_BRANCH_REGISTER":             11,
		"TYPE_BRANCH_REGISTER_RESULT":      12,
		"TYPE_BRANCH_STATUS_REPORT":        13,
		"TYPE_BRANCH_STATUS_REPORT_RESULT": 14,
		"TYPE_SEATA_MERGE":                 59,
		"TYPE_SEATA_MERGE_RESULT":          60,
		"TYPE_REG_CLT":                     101,
		"TYPE_REG_CLT_RESULT":              102,
		"TYPE_REG_RM":                      103,
		"TYPE_REG_RM_RESULT":               104,
		"TYPE_UNDO_LOG_DELETE":             111,
	}
)

func (x MessageTypeProto) Enum() *MessageTypeProto {
	p := new(MessageTypeProto)
	*p = x
	return p
}

func (x MessageTypeProto) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageTypeProto) Descriptor() protoreflect.EnumDescriptor {
	return file_messageType_proto_enumTypes[0].Descriptor()
}

func (MessageTypeProto) Type() protoreflect.EnumType {
	return &file_messageType_proto_enumTypes[0]
}

func (x MessageTypeProto) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageTypeProto.Descriptor instead.
func (MessageTypeProto) EnumDescriptor() ([]byte, []int) {
	return file_messageType_proto_rawDescGZIP(), []int{0}
}

type TestMessageType int32

const (
	TestMessageType_Test     TestMessageType = 0
	TestMessageType_Timeout  TestMessageType = 1
	TestMessageType_Deadline TestMessageType = 2
	TestMessageType_Cancel   TestMessageType = 3
)

// Enum value maps for TestMessageType.
var (
	TestMessageType_name = map[int32]string{
		0: "Test",
		1: "Timeout",
		2: "Deadline",
		3: "Cancel",
	}
	TestMessageType_value = map[string]int32{
		"Test":     0,
		"Timeout":  1,
		"Deadline": 2,
		"Cancel":   3,
	}
)

func (x TestMessageType) Enum() *TestMessageType {
	p := new(TestMessageType)
	*p = x
	return p
}

func (x TestMessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestMessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_messageType_proto_enumTypes[1].Descriptor()
}

func (TestMessageType) Type() protoreflect.EnumType {
	return &file_messageType_proto_enumTypes[1]
}

func (x TestMessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestMessageType.Descriptor instead.
func (TestMessageType) EnumDescriptor() ([]byte, []int) {
	return file_messageType_proto_rawDescGZIP(), []int{1}
}

var File_messageType_proto protoreflect.FileDescriptor

var file_messageType_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2a,
	0x94, 0x06, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x4c, 0x4f,
	0x42, 0x41, 0x4c, 0x5f, 0x50, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x15, 0x0a, 0x11, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x5f, 0x42,
	0x45, 0x47, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47,
	0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x5f, 0x42, 0x45, 0x47, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55,
	0x4c, 0x54, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x4c, 0x4f,
	0x42, 0x41, 0x4c, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x10, 0x07, 0x12, 0x1d, 0x0a, 0x19,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x5f, 0x43, 0x4f, 0x4d, 0x4d,
	0x49, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x08, 0x12, 0x18, 0x0a, 0x14, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x5f, 0x52, 0x4f, 0x4c, 0x4c, 0x42,
	0x41, 0x43, 0x4b, 0x10, 0x09, 0x12, 0x1f, 0x0a, 0x1b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x4c,
	0x4f, 0x42, 0x41, 0x4c, 0x5f, 0x52, 0x4f, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x52, 0x45,
	0x53, 0x55, 0x4c, 0x54, 0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47,
	0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x0f, 0x12, 0x1d,
	0x0a, 0x19, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x10, 0x12, 0x16, 0x0a,
	0x12, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x5f, 0x52, 0x45, 0x50,
	0x4f, 0x52, 0x54, 0x10, 0x11, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x4c,
	0x4f, 0x42, 0x41, 0x4c, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x55,
	0x4c, 0x54, 0x10, 0x12, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x4c, 0x4f,
	0x42, 0x41, 0x4c, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x10, 0x15,
	0x12, 0x21, 0x0a, 0x1d, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x5f,
	0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c,
	0x54, 0x10, 0x16, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x52, 0x41, 0x4e,
	0x43, 0x48, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x42, 0x52, 0x41, 0x4e, 0x43, 0x48, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49,
	0x54, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x42, 0x52, 0x41, 0x4e, 0x43, 0x48, 0x5f, 0x52, 0x4f, 0x4c, 0x4c, 0x42, 0x41,
	0x43, 0x4b, 0x10, 0x05, 0x12, 0x1f, 0x0a, 0x1b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x52, 0x41,
	0x4e, 0x43, 0x48, 0x5f, 0x52, 0x4f, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x10, 0x06, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x52,
	0x41, 0x4e, 0x43, 0x48, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x0b, 0x12,
	0x1f, 0x0a, 0x1b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x52, 0x41, 0x4e, 0x43, 0x48, 0x5f, 0x52,
	0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x0c,
	0x12, 0x1d, 0x0a, 0x19, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x52, 0x41, 0x4e, 0x43, 0x48, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x0d, 0x12,
	0x24, 0x0a, 0x20, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x52, 0x41, 0x4e, 0x43, 0x48, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x10, 0x0e, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45,
	0x41, 0x54, 0x41, 0x5f, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x10, 0x3b, 0x12, 0x1b, 0x0a, 0x17, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x41, 0x54, 0x41, 0x5f, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x5f,
	0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x3c, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x52, 0x45, 0x47, 0x5f, 0x43, 0x4c, 0x54, 0x10, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x52, 0x45, 0x47, 0x5f, 0x43, 0x4c, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c,
	0x54, 0x10, 0x66, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x47, 0x5f,
	0x52, 0x4d, 0x10, 0x67, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x47,
	0x5f, 0x52, 0x4d, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x68, 0x12, 0x18, 0x0a, 0x14,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x4f, 0x5f, 0x4c, 0x4f, 0x47, 0x5f, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x10, 0x6f, 0x2a, 0x42, 0x0a, 0x0f, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x65, 0x73,
	0x74, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x10, 0x03, 0x42, 0x3c, 0x0a, 0x26, 0x69, 0x6f,
	0x2e, 0x73, 0x65, 0x61, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x50, 0x01, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messageType_proto_rawDescOnce sync.Once
	file_messageType_proto_rawDescData = file_messageType_proto_rawDesc
)

func file_messageType_proto_rawDescGZIP() []byte {
	file_messageType_proto_rawDescOnce.Do(func() {
		file_messageType_proto_rawDescData = protoimpl.X.CompressGZIP(file_messageType_proto_rawDescData)
	})
	return file_messageType_proto_rawDescData
}

var file_messageType_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_messageType_proto_goTypes = []interface{}{
	(MessageTypeProto)(0), // 0: io.seata.protocol.protobuf.MessageTypeProto
	(TestMessageType)(0),  // 1: io.seata.protocol.protobuf.TestMessageType
}
var file_messageType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_messageType_proto_init() }
func file_messageType_proto_init() {
	if File_messageType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messageType_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messageType_proto_goTypes,
		DependencyIndexes: file_messageType_proto_depIdxs,
		EnumInfos:         file_messageType_proto_enumTypes,
	}.Build()
	File_messageType_proto = out.File
	file_messageType_proto_rawDesc = nil
	file_messageType_proto_goTypes = nil
	file_messageType_proto_depIdxs = nil
}
