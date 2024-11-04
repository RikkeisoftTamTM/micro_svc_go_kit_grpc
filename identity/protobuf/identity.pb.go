// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: identity.proto

package protobuf

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type InsertReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *InsertReq) Reset() {
	*x = InsertReq{}
	mi := &file_identity_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InsertReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertReq) ProtoMessage() {}

func (x *InsertReq) ProtoReflect() protoreflect.Message {
	mi := &file_identity_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertReq.ProtoReflect.Descriptor instead.
func (*InsertReq) Descriptor() ([]byte, []int) {
	return file_identity_proto_rawDescGZIP(), []int{0}
}

func (x *InsertReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *InsertReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type InsertRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *InsertRes) Reset() {
	*x = InsertRes{}
	mi := &file_identity_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InsertRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertRes) ProtoMessage() {}

func (x *InsertRes) ProtoReflect() protoreflect.Message {
	mi := &file_identity_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertRes.ProtoReflect.Descriptor instead.
func (*InsertRes) Descriptor() ([]byte, []int) {
	return file_identity_proto_rawDescGZIP(), []int{1}
}

func (x *InsertRes) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	mi := &file_identity_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_identity_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_identity_proto_rawDescGZIP(), []int{2}
}

func (x *GetReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *GetRes) Reset() {
	*x = GetRes{}
	mi := &file_identity_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRes) ProtoMessage() {}

func (x *GetRes) ProtoReflect() protoreflect.Message {
	mi := &file_identity_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRes.ProtoReflect.Descriptor instead.
func (*GetRes) Descriptor() ([]byte, []int) {
	return file_identity_proto_rawDescGZIP(), []int{3}
}

func (x *GetRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetRes) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetRes) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_identity_proto protoreflect.FileDescriptor

var file_identity_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x43, 0x0a, 0x09, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x21, 0x0a, 0x09, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x24, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x50, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x32, 0x82, 0x01, 0x0a, 0x0f, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x3a, 0x01, 0x2a, 0x22, 0x07, 0x2f,
	0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x22, 0x0c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x06, 0x12, 0x04, 0x2f, 0x67, 0x65, 0x74, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_identity_proto_rawDescOnce sync.Once
	file_identity_proto_rawDescData = file_identity_proto_rawDesc
)

func file_identity_proto_rawDescGZIP() []byte {
	file_identity_proto_rawDescOnce.Do(func() {
		file_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_identity_proto_rawDescData)
	})
	return file_identity_proto_rawDescData
}

var file_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_identity_proto_goTypes = []any{
	(*InsertReq)(nil), // 0: pb.InsertReq
	(*InsertRes)(nil), // 1: pb.InsertRes
	(*GetReq)(nil),    // 2: pb.GetReq
	(*GetRes)(nil),    // 3: pb.GetRes
}
var file_identity_proto_depIdxs = []int32{
	0, // 0: pb.IdentityService.InsertUser:input_type -> pb.InsertReq
	2, // 1: pb.IdentityService.GetUser:input_type -> pb.GetReq
	1, // 2: pb.IdentityService.InsertUser:output_type -> pb.InsertRes
	3, // 3: pb.IdentityService.GetUser:output_type -> pb.GetRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_identity_proto_init() }
func file_identity_proto_init() {
	if File_identity_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_identity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_identity_proto_goTypes,
		DependencyIndexes: file_identity_proto_depIdxs,
		MessageInfos:      file_identity_proto_msgTypes,
	}.Build()
	File_identity_proto = out.File
	file_identity_proto_rawDesc = nil
	file_identity_proto_goTypes = nil
	file_identity_proto_depIdxs = nil
}
