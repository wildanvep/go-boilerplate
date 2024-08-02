// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: outbox.proto

package outbox

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

type Outbox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//
	//	*Outbox_Profile
	//	*Outbox_Other
	Content isOutbox_Content `protobuf_oneof:"content"`
}

func (x *Outbox) Reset() {
	*x = Outbox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outbox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outbox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outbox) ProtoMessage() {}

func (x *Outbox) ProtoReflect() protoreflect.Message {
	mi := &file_outbox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outbox.ProtoReflect.Descriptor instead.
func (*Outbox) Descriptor() ([]byte, []int) {
	return file_outbox_proto_rawDescGZIP(), []int{0}
}

func (m *Outbox) GetContent() isOutbox_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *Outbox) GetProfile() *Profile {
	if x, ok := x.GetContent().(*Outbox_Profile); ok {
		return x.Profile
	}
	return nil
}

func (x *Outbox) GetOther() string {
	if x, ok := x.GetContent().(*Outbox_Other); ok {
		return x.Other
	}
	return ""
}

type isOutbox_Content interface {
	isOutbox_Content()
}

type Outbox_Profile struct {
	Profile *Profile `protobuf:"bytes,1,opt,name=profile,proto3,oneof"`
}

type Outbox_Other struct {
	Other string `protobuf:"bytes,2,opt,name=other,proto3,oneof"`
}

func (*Outbox_Profile) isOutbox_Content() {}

func (*Outbox_Other) isOutbox_Content() {}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       []byte                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TenantID []byte                 `protobuf:"bytes,2,opt,name=TenantID,proto3" json:"TenantID,omitempty"`
	NIN      string                 `protobuf:"bytes,3,opt,name=NIN,proto3" json:"NIN,omitempty"`
	Name     string                 `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Email    string                 `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone    string                 `protobuf:"bytes,6,opt,name=Phone,proto3" json:"Phone,omitempty"`
	DOB      *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=DOB,proto3" json:"DOB,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outbox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_outbox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_outbox_proto_rawDescGZIP(), []int{1}
}

func (x *Profile) GetID() []byte {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Profile) GetTenantID() []byte {
	if x != nil {
		return x.TenantID
	}
	return nil
}

func (x *Profile) GetNIN() string {
	if x != nil {
		return x.NIN
	}
	return ""
}

func (x *Profile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Profile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Profile) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Profile) GetDOB() *timestamppb.Timestamp {
	if x != nil {
		return x.DOB
	}
	return nil
}

var File_outbox_proto protoreflect.FileDescriptor

var file_outbox_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6f, 0x75, 0x74, 0x62, 0x6f, 0x78, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x62, 0x6f,
	0x78, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x78, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0xb5, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x49, 0x4e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4e, 0x49, 0x4e, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x44,
	0x4f, 0x42, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x44, 0x4f, 0x42, 0x42, 0x99, 0x01, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x2e, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x78, 0x42, 0x0b, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x78,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6c, 0x6b, 0x6f, 0x6d, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0x65,
	0x73, 0x69, 0x61, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x78, 0xa2,
	0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x78, 0xca, 0x02,
	0x06, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x78, 0xe2, 0x02, 0x12, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x78,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x4f,
	0x75, 0x74, 0x62, 0x6f, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_outbox_proto_rawDescOnce sync.Once
	file_outbox_proto_rawDescData = file_outbox_proto_rawDesc
)

func file_outbox_proto_rawDescGZIP() []byte {
	file_outbox_proto_rawDescOnce.Do(func() {
		file_outbox_proto_rawDescData = protoimpl.X.CompressGZIP(file_outbox_proto_rawDescData)
	})
	return file_outbox_proto_rawDescData
}

var file_outbox_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_outbox_proto_goTypes = []interface{}{
	(*Outbox)(nil),                // 0: outbox.Outbox
	(*Profile)(nil),               // 1: outbox.Profile
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_outbox_proto_depIdxs = []int32{
	1, // 0: outbox.Outbox.profile:type_name -> outbox.Profile
	2, // 1: outbox.Profile.DOB:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_outbox_proto_init() }
func file_outbox_proto_init() {
	if File_outbox_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_outbox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outbox); i {
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
		file_outbox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
	file_outbox_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Outbox_Profile)(nil),
		(*Outbox_Other)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_outbox_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_outbox_proto_goTypes,
		DependencyIndexes: file_outbox_proto_depIdxs,
		MessageInfos:      file_outbox_proto_msgTypes,
	}.Build()
	File_outbox_proto = out.File
	file_outbox_proto_rawDesc = nil
	file_outbox_proto_goTypes = nil
	file_outbox_proto_depIdxs = nil
}
