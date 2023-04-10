// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: bff/admin/v1/admin.proto

package v1

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

// 请求-创建用户
type CreateAdminReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAdminReq) Reset() {
	*x = CreateAdminReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_admin_v1_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdminReq) ProtoMessage() {}

func (x *CreateAdminReq) ProtoReflect() protoreflect.Message {
	mi := &file_bff_admin_v1_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdminReq.ProtoReflect.Descriptor instead.
func (*CreateAdminReq) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{0}
}

// 响应-创建用户
type CreateAdminReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAdminReply) Reset() {
	*x = CreateAdminReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_admin_v1_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAdminReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdminReply) ProtoMessage() {}

func (x *CreateAdminReply) ProtoReflect() protoreflect.Message {
	mi := &file_bff_admin_v1_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdminReply.ProtoReflect.Descriptor instead.
func (*CreateAdminReply) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{1}
}

// 请求-更新用户
type UpdateAdminReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAdminReq) Reset() {
	*x = UpdateAdminReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_admin_v1_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdminReq) ProtoMessage() {}

func (x *UpdateAdminReq) ProtoReflect() protoreflect.Message {
	mi := &file_bff_admin_v1_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAdminReq.ProtoReflect.Descriptor instead.
func (*UpdateAdminReq) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{2}
}

// 响应-更新用户
type UpdateAdminReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAdminReply) Reset() {
	*x = UpdateAdminReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_admin_v1_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAdminReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdminReply) ProtoMessage() {}

func (x *UpdateAdminReply) ProtoReflect() protoreflect.Message {
	mi := &file_bff_admin_v1_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAdminReply.ProtoReflect.Descriptor instead.
func (*UpdateAdminReply) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{3}
}

// 请求-删除用户
type DeleteAdminReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAdminReq) Reset() {
	*x = DeleteAdminReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_admin_v1_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAdminReq) ProtoMessage() {}

func (x *DeleteAdminReq) ProtoReflect() protoreflect.Message {
	mi := &file_bff_admin_v1_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAdminReq.ProtoReflect.Descriptor instead.
func (*DeleteAdminReq) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{4}
}

// 响应-删除用户
type DeleteAdminReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAdminReply) Reset() {
	*x = DeleteAdminReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_admin_v1_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAdminReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAdminReply) ProtoMessage() {}

func (x *DeleteAdminReply) ProtoReflect() protoreflect.Message {
	mi := &file_bff_admin_v1_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAdminReply.ProtoReflect.Descriptor instead.
func (*DeleteAdminReply) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{5}
}

// 请求-获取单个用户
type GetAdminReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAdminReq) Reset() {
	*x = GetAdminReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_admin_v1_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminReq) ProtoMessage() {}

func (x *GetAdminReq) ProtoReflect() protoreflect.Message {
	mi := &file_bff_admin_v1_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminReq.ProtoReflect.Descriptor instead.
func (*GetAdminReq) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{6}
}

// 响应-获取单个用户
type GetAdminReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAdminReply) Reset() {
	*x = GetAdminReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_admin_v1_admin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminReply) ProtoMessage() {}

func (x *GetAdminReply) ProtoReflect() protoreflect.Message {
	mi := &file_bff_admin_v1_admin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminReply.ProtoReflect.Descriptor instead.
func (*GetAdminReply) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{7}
}

// 请求-获取用户列表
type ListAdminReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAdminReq) Reset() {
	*x = ListAdminReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_admin_v1_admin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAdminReq) ProtoMessage() {}

func (x *ListAdminReq) ProtoReflect() protoreflect.Message {
	mi := &file_bff_admin_v1_admin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAdminReq.ProtoReflect.Descriptor instead.
func (*ListAdminReq) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{8}
}

// 响应-获取用户列表
type ListAdminReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAdminReply) Reset() {
	*x = ListAdminReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_admin_v1_admin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAdminReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAdminReply) ProtoMessage() {}

func (x *ListAdminReply) ProtoReflect() protoreflect.Message {
	mi := &file_bff_admin_v1_admin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAdminReply.ProtoReflect.Descriptor instead.
func (*ListAdminReply) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{9}
}

var File_bff_admin_v1_admin_proto protoreflect.FileDescriptor

var file_bff_admin_v1_admin_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x66, 0x66, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10,
	0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0d, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0e, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xf9, 0x02, 0x0a,
	0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x4b, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x4b, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x4b, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12,
	0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x42, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x45, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x29, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x17, 0x66, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bff_admin_v1_admin_proto_rawDescOnce sync.Once
	file_bff_admin_v1_admin_proto_rawDescData = file_bff_admin_v1_admin_proto_rawDesc
)

func file_bff_admin_v1_admin_proto_rawDescGZIP() []byte {
	file_bff_admin_v1_admin_proto_rawDescOnce.Do(func() {
		file_bff_admin_v1_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_bff_admin_v1_admin_proto_rawDescData)
	})
	return file_bff_admin_v1_admin_proto_rawDescData
}

var file_bff_admin_v1_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_bff_admin_v1_admin_proto_goTypes = []interface{}{
	(*CreateAdminReq)(nil),   // 0: api.Admin.v1.CreateAdminReq
	(*CreateAdminReply)(nil), // 1: api.Admin.v1.CreateAdminReply
	(*UpdateAdminReq)(nil),   // 2: api.Admin.v1.UpdateAdminReq
	(*UpdateAdminReply)(nil), // 3: api.Admin.v1.UpdateAdminReply
	(*DeleteAdminReq)(nil),   // 4: api.Admin.v1.DeleteAdminReq
	(*DeleteAdminReply)(nil), // 5: api.Admin.v1.DeleteAdminReply
	(*GetAdminReq)(nil),      // 6: api.Admin.v1.GetAdminReq
	(*GetAdminReply)(nil),    // 7: api.Admin.v1.GetAdminReply
	(*ListAdminReq)(nil),     // 8: api.Admin.v1.ListAdminReq
	(*ListAdminReply)(nil),   // 9: api.Admin.v1.ListAdminReply
}
var file_bff_admin_v1_admin_proto_depIdxs = []int32{
	0, // 0: api.Admin.v1.Admin.CreateAdmin:input_type -> api.Admin.v1.CreateAdminReq
	2, // 1: api.Admin.v1.Admin.UpdateAdmin:input_type -> api.Admin.v1.UpdateAdminReq
	4, // 2: api.Admin.v1.Admin.DeleteAdmin:input_type -> api.Admin.v1.DeleteAdminReq
	6, // 3: api.Admin.v1.Admin.GetAdmin:input_type -> api.Admin.v1.GetAdminReq
	8, // 4: api.Admin.v1.Admin.ListAdmin:input_type -> api.Admin.v1.ListAdminReq
	1, // 5: api.Admin.v1.Admin.CreateAdmin:output_type -> api.Admin.v1.CreateAdminReply
	3, // 6: api.Admin.v1.Admin.UpdateAdmin:output_type -> api.Admin.v1.UpdateAdminReply
	5, // 7: api.Admin.v1.Admin.DeleteAdmin:output_type -> api.Admin.v1.DeleteAdminReply
	7, // 8: api.Admin.v1.Admin.GetAdmin:output_type -> api.Admin.v1.GetAdminReply
	9, // 9: api.Admin.v1.Admin.ListAdmin:output_type -> api.Admin.v1.ListAdminReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bff_admin_v1_admin_proto_init() }
func file_bff_admin_v1_admin_proto_init() {
	if File_bff_admin_v1_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bff_admin_v1_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAdminReq); i {
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
		file_bff_admin_v1_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAdminReply); i {
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
		file_bff_admin_v1_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAdminReq); i {
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
		file_bff_admin_v1_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAdminReply); i {
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
		file_bff_admin_v1_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAdminReq); i {
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
		file_bff_admin_v1_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAdminReply); i {
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
		file_bff_admin_v1_admin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdminReq); i {
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
		file_bff_admin_v1_admin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdminReply); i {
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
		file_bff_admin_v1_admin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAdminReq); i {
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
		file_bff_admin_v1_admin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAdminReply); i {
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
			RawDescriptor: file_bff_admin_v1_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bff_admin_v1_admin_proto_goTypes,
		DependencyIndexes: file_bff_admin_v1_admin_proto_depIdxs,
		MessageInfos:      file_bff_admin_v1_admin_proto_msgTypes,
	}.Build()
	File_bff_admin_v1_admin_proto = out.File
	file_bff_admin_v1_admin_proto_rawDesc = nil
	file_bff_admin_v1_admin_proto_goTypes = nil
	file_bff_admin_v1_admin_proto_depIdxs = nil
}
