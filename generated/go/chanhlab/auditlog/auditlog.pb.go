// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: chanhlab/auditlog/auditlog.proto

package v1

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

// An AuditLog.
type AuditLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefId       string                 `protobuf:"bytes,1,opt,name=ref_id,json=refId,proto3" json:"ref_id,omitempty"`
	ActorId     string                 `protobuf:"bytes,2,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"`
	ObjectId    string                 `protobuf:"bytes,3,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	EventName   string                 `protobuf:"bytes,4,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
	Description string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Sources     map[string]string      `protobuf:"bytes,6,rep,name=sources,proto3" json:"sources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tags        map[string]string      `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
}

func (x *AuditLog) Reset() {
	*x = AuditLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chanhlab_auditlog_auditlog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLog) ProtoMessage() {}

func (x *AuditLog) ProtoReflect() protoreflect.Message {
	mi := &file_chanhlab_auditlog_auditlog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditLog.ProtoReflect.Descriptor instead.
func (*AuditLog) Descriptor() ([]byte, []int) {
	return file_chanhlab_auditlog_auditlog_proto_rawDescGZIP(), []int{0}
}

func (x *AuditLog) GetRefId() string {
	if x != nil {
		return x.RefId
	}
	return ""
}

func (x *AuditLog) GetActorId() string {
	if x != nil {
		return x.ActorId
	}
	return ""
}

func (x *AuditLog) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *AuditLog) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *AuditLog) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AuditLog) GetSources() map[string]string {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *AuditLog) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *AuditLog) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

var File_chanhlab_auditlog_auditlog_proto protoreflect.FileDescriptor

var file_chanhlab_auditlog_auditlog_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x68, 0x61, 0x6e, 0x68, 0x6c, 0x61, 0x62, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x63, 0x68, 0x61, 0x6e, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x03, 0x0a, 0x08, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x66, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x66, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x68, 0x6c,
	0x61, 0x62, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x3c, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x68,
	0x61, 0x6e, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68,
	0x61, 0x6e, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x68,
	0x6c, 0x61, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chanhlab_auditlog_auditlog_proto_rawDescOnce sync.Once
	file_chanhlab_auditlog_auditlog_proto_rawDescData = file_chanhlab_auditlog_auditlog_proto_rawDesc
)

func file_chanhlab_auditlog_auditlog_proto_rawDescGZIP() []byte {
	file_chanhlab_auditlog_auditlog_proto_rawDescOnce.Do(func() {
		file_chanhlab_auditlog_auditlog_proto_rawDescData = protoimpl.X.CompressGZIP(file_chanhlab_auditlog_auditlog_proto_rawDescData)
	})
	return file_chanhlab_auditlog_auditlog_proto_rawDescData
}

var file_chanhlab_auditlog_auditlog_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chanhlab_auditlog_auditlog_proto_goTypes = []interface{}{
	(*AuditLog)(nil),              // 0: chanhlab.auditlog.v1.AuditLog
	nil,                           // 1: chanhlab.auditlog.v1.AuditLog.SourcesEntry
	nil,                           // 2: chanhlab.auditlog.v1.AuditLog.TagsEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_chanhlab_auditlog_auditlog_proto_depIdxs = []int32{
	1, // 0: chanhlab.auditlog.v1.AuditLog.sources:type_name -> chanhlab.auditlog.v1.AuditLog.SourcesEntry
	2, // 1: chanhlab.auditlog.v1.AuditLog.tags:type_name -> chanhlab.auditlog.v1.AuditLog.TagsEntry
	3, // 2: chanhlab.auditlog.v1.AuditLog.created_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_chanhlab_auditlog_auditlog_proto_init() }
func file_chanhlab_auditlog_auditlog_proto_init() {
	if File_chanhlab_auditlog_auditlog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chanhlab_auditlog_auditlog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditLog); i {
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
			RawDescriptor: file_chanhlab_auditlog_auditlog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chanhlab_auditlog_auditlog_proto_goTypes,
		DependencyIndexes: file_chanhlab_auditlog_auditlog_proto_depIdxs,
		MessageInfos:      file_chanhlab_auditlog_auditlog_proto_msgTypes,
	}.Build()
	File_chanhlab_auditlog_auditlog_proto = out.File
	file_chanhlab_auditlog_auditlog_proto_rawDesc = nil
	file_chanhlab_auditlog_auditlog_proto_goTypes = nil
	file_chanhlab_auditlog_auditlog_proto_depIdxs = nil
}