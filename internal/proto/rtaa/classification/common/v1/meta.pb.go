// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/rtaa/classification/common/v1/meta.proto

package commonv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Meta message holds metadata for any user request (comments or accounts classification)
// Use @ threshold to adjust the probability boundary (default 0.5)
// Use @ version to choose the version of the classifier (accepts v1, v2, default v1)
type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Source           string                 `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	ScreenName       string                 `protobuf:"bytes,3,opt,name=screen_name,json=screenName,proto3" json:"screen_name,omitempty"`
	CommentCreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=comment_created_at,json=commentCreatedAt,proto3" json:"comment_created_at,omitempty"`
	UserCreatedAt    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=user_created_at,json=userCreatedAt,proto3" json:"user_created_at,omitempty"`
	ClassifiedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=classified_at,json=classifiedAt,proto3" json:"classified_at,omitempty"`
	Link             string                 `protobuf:"bytes,7,opt,name=link,proto3" json:"link,omitempty"`
	Lang             string                 `protobuf:"bytes,8,opt,name=lang,proto3" json:"lang,omitempty"`
	Threshold        float64                `protobuf:"fixed64,9,opt,name=threshold,proto3" json:"threshold,omitempty"`
	Version          string                 `protobuf:"bytes,10,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rtaa_classification_common_v1_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rtaa_classification_common_v1_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_proto_rtaa_classification_common_v1_meta_proto_rawDescGZIP(), []int{0}
}

func (x *Meta) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Meta) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Meta) GetScreenName() string {
	if x != nil {
		return x.ScreenName
	}
	return ""
}

func (x *Meta) GetCommentCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CommentCreatedAt
	}
	return nil
}

func (x *Meta) GetUserCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UserCreatedAt
	}
	return nil
}

func (x *Meta) GetClassifiedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ClassifiedAt
	}
	return nil
}

func (x *Meta) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Meta) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *Meta) GetThreshold() float64 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *Meta) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_proto_rtaa_classification_common_v1_meta_proto protoreflect.FileDescriptor

var file_proto_rtaa_classification_common_v1_meta_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x74, 0x61, 0x61, 0x2f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x74, 0x61, 0x61, 0x2e, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x02, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x72, 0x65, 0x65,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x42, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3f, 0x0a, 0x0d, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x61, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0xb5, 0x02, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x74, 0x61, 0x61, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x42, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x76, 0x63,
	0x69, 0x6f, 0x2f, 0x72, 0x74, 0x61, 0x61, 0x2d, 0x37, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x74, 0x61, 0x61, 0x2f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31,
	0xa2, 0x02, 0x04, 0x50, 0x52, 0x43, 0x43, 0xaa, 0x02, 0x23, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x74, 0x61, 0x61, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x23,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x52, 0x74, 0x61, 0x61, 0x5c, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x52, 0x74, 0x61, 0x61,
	0x5c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x27, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x52,
	0x74, 0x61, 0x61, 0x3a, 0x3a, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rtaa_classification_common_v1_meta_proto_rawDescOnce sync.Once
	file_proto_rtaa_classification_common_v1_meta_proto_rawDescData = file_proto_rtaa_classification_common_v1_meta_proto_rawDesc
)

func file_proto_rtaa_classification_common_v1_meta_proto_rawDescGZIP() []byte {
	file_proto_rtaa_classification_common_v1_meta_proto_rawDescOnce.Do(func() {
		file_proto_rtaa_classification_common_v1_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rtaa_classification_common_v1_meta_proto_rawDescData)
	})
	return file_proto_rtaa_classification_common_v1_meta_proto_rawDescData
}

var file_proto_rtaa_classification_common_v1_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_rtaa_classification_common_v1_meta_proto_goTypes = []interface{}{
	(*Meta)(nil),                  // 0: proto.rtaa.classification.common.v1.Meta
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_proto_rtaa_classification_common_v1_meta_proto_depIdxs = []int32{
	1, // 0: proto.rtaa.classification.common.v1.Meta.comment_created_at:type_name -> google.protobuf.Timestamp
	1, // 1: proto.rtaa.classification.common.v1.Meta.user_created_at:type_name -> google.protobuf.Timestamp
	1, // 2: proto.rtaa.classification.common.v1.Meta.classified_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_rtaa_classification_common_v1_meta_proto_init() }
func file_proto_rtaa_classification_common_v1_meta_proto_init() {
	if File_proto_rtaa_classification_common_v1_meta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rtaa_classification_common_v1_meta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
			RawDescriptor: file_proto_rtaa_classification_common_v1_meta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_rtaa_classification_common_v1_meta_proto_goTypes,
		DependencyIndexes: file_proto_rtaa_classification_common_v1_meta_proto_depIdxs,
		MessageInfos:      file_proto_rtaa_classification_common_v1_meta_proto_msgTypes,
	}.Build()
	File_proto_rtaa_classification_common_v1_meta_proto = out.File
	file_proto_rtaa_classification_common_v1_meta_proto_rawDesc = nil
	file_proto_rtaa_classification_common_v1_meta_proto_goTypes = nil
	file_proto_rtaa_classification_common_v1_meta_proto_depIdxs = nil
}