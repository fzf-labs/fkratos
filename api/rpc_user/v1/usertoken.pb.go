// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: user/v1/usertoken.proto

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

type CreateUsertokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateUsertokenRequest) Reset() {
	*x = CreateUsertokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_usertoken_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUsertokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUsertokenRequest) ProtoMessage() {}

func (x *CreateUsertokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_usertoken_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUsertokenRequest.ProtoReflect.Descriptor instead.
func (*CreateUsertokenRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_usertoken_proto_rawDescGZIP(), []int{0}
}

type CreateUsertokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateUsertokenReply) Reset() {
	*x = CreateUsertokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_usertoken_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUsertokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUsertokenReply) ProtoMessage() {}

func (x *CreateUsertokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_usertoken_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUsertokenReply.ProtoReflect.Descriptor instead.
func (*CreateUsertokenReply) Descriptor() ([]byte, []int) {
	return file_user_v1_usertoken_proto_rawDescGZIP(), []int{1}
}

type UpdateUsertokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUsertokenRequest) Reset() {
	*x = UpdateUsertokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_usertoken_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUsertokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUsertokenRequest) ProtoMessage() {}

func (x *UpdateUsertokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_usertoken_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUsertokenRequest.ProtoReflect.Descriptor instead.
func (*UpdateUsertokenRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_usertoken_proto_rawDescGZIP(), []int{2}
}

type UpdateUsertokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUsertokenReply) Reset() {
	*x = UpdateUsertokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_usertoken_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUsertokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUsertokenReply) ProtoMessage() {}

func (x *UpdateUsertokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_usertoken_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUsertokenReply.ProtoReflect.Descriptor instead.
func (*UpdateUsertokenReply) Descriptor() ([]byte, []int) {
	return file_user_v1_usertoken_proto_rawDescGZIP(), []int{3}
}

type DeleteUsertokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUsertokenRequest) Reset() {
	*x = DeleteUsertokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_usertoken_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUsertokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUsertokenRequest) ProtoMessage() {}

func (x *DeleteUsertokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_usertoken_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUsertokenRequest.ProtoReflect.Descriptor instead.
func (*DeleteUsertokenRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_usertoken_proto_rawDescGZIP(), []int{4}
}

type DeleteUsertokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUsertokenReply) Reset() {
	*x = DeleteUsertokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_usertoken_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUsertokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUsertokenReply) ProtoMessage() {}

func (x *DeleteUsertokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_usertoken_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUsertokenReply.ProtoReflect.Descriptor instead.
func (*DeleteUsertokenReply) Descriptor() ([]byte, []int) {
	return file_user_v1_usertoken_proto_rawDescGZIP(), []int{5}
}

type GetUsertokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUsertokenRequest) Reset() {
	*x = GetUsertokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_usertoken_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsertokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsertokenRequest) ProtoMessage() {}

func (x *GetUsertokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_usertoken_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsertokenRequest.ProtoReflect.Descriptor instead.
func (*GetUsertokenRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_usertoken_proto_rawDescGZIP(), []int{6}
}

type GetUsertokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUsertokenReply) Reset() {
	*x = GetUsertokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_usertoken_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsertokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsertokenReply) ProtoMessage() {}

func (x *GetUsertokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_usertoken_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsertokenReply.ProtoReflect.Descriptor instead.
func (*GetUsertokenReply) Descriptor() ([]byte, []int) {
	return file_user_v1_usertoken_proto_rawDescGZIP(), []int{7}
}

type ListUsertokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListUsertokenRequest) Reset() {
	*x = ListUsertokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_usertoken_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsertokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsertokenRequest) ProtoMessage() {}

func (x *ListUsertokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_usertoken_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsertokenRequest.ProtoReflect.Descriptor instead.
func (*ListUsertokenRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_usertoken_proto_rawDescGZIP(), []int{8}
}

type ListUsertokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListUsertokenReply) Reset() {
	*x = ListUsertokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_usertoken_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsertokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsertokenReply) ProtoMessage() {}

func (x *ListUsertokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_usertoken_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsertokenReply.ProtoReflect.Descriptor instead.
func (*ListUsertokenReply) Descriptor() ([]byte, []int) {
	return file_user_v1_usertoken_proto_rawDescGZIP(), []int{9}
}

var File_user_v1_usertoken_proto protoreflect.FileDescriptor

var file_user_v1_usertoken_proto_rawDesc = []byte{
	0x0a, 0x17, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xc3, 0x03, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x59, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x59, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x59, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x50, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x53, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x27, 0x0a, 0x0b,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x16, 0x66,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_v1_usertoken_proto_rawDescOnce sync.Once
	file_user_v1_usertoken_proto_rawDescData = file_user_v1_usertoken_proto_rawDesc
)

func file_user_v1_usertoken_proto_rawDescGZIP() []byte {
	file_user_v1_usertoken_proto_rawDescOnce.Do(func() {
		file_user_v1_usertoken_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_v1_usertoken_proto_rawDescData)
	})
	return file_user_v1_usertoken_proto_rawDescData
}

var file_user_v1_usertoken_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_user_v1_usertoken_proto_goTypes = []interface{}{
	(*CreateUsertokenRequest)(nil), // 0: api.user.v1.CreateUsertokenRequest
	(*CreateUsertokenReply)(nil),   // 1: api.user.v1.CreateUsertokenReply
	(*UpdateUsertokenRequest)(nil), // 2: api.user.v1.UpdateUsertokenRequest
	(*UpdateUsertokenReply)(nil),   // 3: api.user.v1.UpdateUsertokenReply
	(*DeleteUsertokenRequest)(nil), // 4: api.user.v1.DeleteUsertokenRequest
	(*DeleteUsertokenReply)(nil),   // 5: api.user.v1.DeleteUsertokenReply
	(*GetUsertokenRequest)(nil),    // 6: api.user.v1.GetUsertokenRequest
	(*GetUsertokenReply)(nil),      // 7: api.user.v1.GetUsertokenReply
	(*ListUsertokenRequest)(nil),   // 8: api.user.v1.ListUsertokenRequest
	(*ListUsertokenReply)(nil),     // 9: api.user.v1.ListUsertokenReply
}
var file_user_v1_usertoken_proto_depIdxs = []int32{
	0, // 0: api.user.v1.Usertoken.CreateUsertoken:input_type -> api.user.v1.CreateUsertokenRequest
	2, // 1: api.user.v1.Usertoken.UpdateUsertoken:input_type -> api.user.v1.UpdateUsertokenRequest
	4, // 2: api.user.v1.Usertoken.DeleteUsertoken:input_type -> api.user.v1.DeleteUsertokenRequest
	6, // 3: api.user.v1.Usertoken.GetUsertoken:input_type -> api.user.v1.GetUsertokenRequest
	8, // 4: api.user.v1.Usertoken.ListUsertoken:input_type -> api.user.v1.ListUsertokenRequest
	1, // 5: api.user.v1.Usertoken.CreateUsertoken:output_type -> api.user.v1.CreateUsertokenReply
	3, // 6: api.user.v1.Usertoken.UpdateUsertoken:output_type -> api.user.v1.UpdateUsertokenReply
	5, // 7: api.user.v1.Usertoken.DeleteUsertoken:output_type -> api.user.v1.DeleteUsertokenReply
	7, // 8: api.user.v1.Usertoken.GetUsertoken:output_type -> api.user.v1.GetUsertokenReply
	9, // 9: api.user.v1.Usertoken.ListUsertoken:output_type -> api.user.v1.ListUsertokenReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_v1_usertoken_proto_init() }
func file_user_v1_usertoken_proto_init() {
	if File_user_v1_usertoken_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_v1_usertoken_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUsertokenRequest); i {
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
		file_user_v1_usertoken_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUsertokenReply); i {
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
		file_user_v1_usertoken_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUsertokenRequest); i {
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
		file_user_v1_usertoken_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUsertokenReply); i {
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
		file_user_v1_usertoken_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUsertokenRequest); i {
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
		file_user_v1_usertoken_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUsertokenReply); i {
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
		file_user_v1_usertoken_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsertokenRequest); i {
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
		file_user_v1_usertoken_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsertokenReply); i {
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
		file_user_v1_usertoken_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsertokenRequest); i {
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
		file_user_v1_usertoken_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsertokenReply); i {
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
			RawDescriptor: file_user_v1_usertoken_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_v1_usertoken_proto_goTypes,
		DependencyIndexes: file_user_v1_usertoken_proto_depIdxs,
		MessageInfos:      file_user_v1_usertoken_proto_msgTypes,
	}.Build()
	File_user_v1_usertoken_proto = out.File
	file_user_v1_usertoken_proto_rawDesc = nil
	file_user_v1_usertoken_proto_goTypes = nil
	file_user_v1_usertoken_proto_depIdxs = nil
}