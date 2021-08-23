// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.17.3
// source: educode/protobuf/submission_timestamp.proto

package educode_protobuf

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

type SubmissionTimestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeId string `protobuf:"bytes,1,opt,name=challengeId,proto3" json:"challengeId,omitempty"`
	User        string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Timestamp   int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SubmissionTimestamp) Reset() {
	*x = SubmissionTimestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_educode_protobuf_submission_timestamp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmissionTimestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmissionTimestamp) ProtoMessage() {}

func (x *SubmissionTimestamp) ProtoReflect() protoreflect.Message {
	mi := &file_educode_protobuf_submission_timestamp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmissionTimestamp.ProtoReflect.Descriptor instead.
func (*SubmissionTimestamp) Descriptor() ([]byte, []int) {
	return file_educode_protobuf_submission_timestamp_proto_rawDescGZIP(), []int{0}
}

func (x *SubmissionTimestamp) GetChallengeId() string {
	if x != nil {
		return x.ChallengeId
	}
	return ""
}

func (x *SubmissionTimestamp) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *SubmissionTimestamp) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_educode_protobuf_submission_timestamp_proto protoreflect.FileDescriptor

var file_educode_protobuf_submission_timestamp_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x64, 0x75, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65,
	0x64, 0x75, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22,
	0x69, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x1b, 0x0a, 0x17, 0x64, 0x65,
	0x2e, 0x68, 0x68, 0x75, 0x2e, 0x65, 0x64, 0x75, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_educode_protobuf_submission_timestamp_proto_rawDescOnce sync.Once
	file_educode_protobuf_submission_timestamp_proto_rawDescData = file_educode_protobuf_submission_timestamp_proto_rawDesc
)

func file_educode_protobuf_submission_timestamp_proto_rawDescGZIP() []byte {
	file_educode_protobuf_submission_timestamp_proto_rawDescOnce.Do(func() {
		file_educode_protobuf_submission_timestamp_proto_rawDescData = protoimpl.X.CompressGZIP(file_educode_protobuf_submission_timestamp_proto_rawDescData)
	})
	return file_educode_protobuf_submission_timestamp_proto_rawDescData
}

var file_educode_protobuf_submission_timestamp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_educode_protobuf_submission_timestamp_proto_goTypes = []interface{}{
	(*SubmissionTimestamp)(nil), // 0: educode.protobuf.SubmissionTimestamp
}
var file_educode_protobuf_submission_timestamp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_educode_protobuf_submission_timestamp_proto_init() }
func file_educode_protobuf_submission_timestamp_proto_init() {
	if File_educode_protobuf_submission_timestamp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_educode_protobuf_submission_timestamp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmissionTimestamp); i {
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
			RawDescriptor: file_educode_protobuf_submission_timestamp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_educode_protobuf_submission_timestamp_proto_goTypes,
		DependencyIndexes: file_educode_protobuf_submission_timestamp_proto_depIdxs,
		MessageInfos:      file_educode_protobuf_submission_timestamp_proto_msgTypes,
	}.Build()
	File_educode_protobuf_submission_timestamp_proto = out.File
	file_educode_protobuf_submission_timestamp_proto_rawDesc = nil
	file_educode_protobuf_submission_timestamp_proto_goTypes = nil
	file_educode_protobuf_submission_timestamp_proto_depIdxs = nil
}
