// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: bff_admin/v1/admin.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type AdminInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdminInfo) Reset() {
	*x = AdminInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_admin_v1_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminInfo) ProtoMessage() {}

func (x *AdminInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AdminInfo.ProtoReflect.Descriptor instead.
func (*AdminInfo) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{0}
}

// 请求-创建用户
type CreateAdminReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminId  string `protobuf:"bytes,1,opt,name=adminId,proto3" json:"adminId,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"` //昵称
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`       //邮件
	Mobile   string `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`     //手机号
	Motto    string `protobuf:"bytes,5,opt,name=motto,proto3" json:"motto,omitempty"`       //个性签名
	Password string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"` //密码
}

func (x *CreateAdminReq) Reset() {
	*x = CreateAdminReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_admin_v1_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdminReq) ProtoMessage() {}

func (x *CreateAdminReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateAdminReq.ProtoReflect.Descriptor instead.
func (*CreateAdminReq) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAdminReq) GetAdminId() string {
	if x != nil {
		return x.AdminId
	}
	return ""
}

func (x *CreateAdminReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *CreateAdminReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateAdminReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *CreateAdminReq) GetMotto() string {
	if x != nil {
		return x.Motto
	}
	return ""
}

func (x *CreateAdminReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
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
		mi := &file_bff_admin_v1_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAdminReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdminReply) ProtoMessage() {}

func (x *CreateAdminReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateAdminReply.ProtoReflect.Descriptor instead.
func (*CreateAdminReply) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{2}
}

// 请求-更新用户
type UpdateAdminReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminInfo *AdminInfo `protobuf:"bytes,1,opt,name=adminInfo,proto3" json:"adminInfo,omitempty"`
}

func (x *UpdateAdminReq) Reset() {
	*x = UpdateAdminReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_admin_v1_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdminReq) ProtoMessage() {}

func (x *UpdateAdminReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateAdminReq.ProtoReflect.Descriptor instead.
func (*UpdateAdminReq) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAdminReq) GetAdminInfo() *AdminInfo {
	if x != nil {
		return x.AdminInfo
	}
	return nil
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
		mi := &file_bff_admin_v1_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAdminReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdminReply) ProtoMessage() {}

func (x *UpdateAdminReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateAdminReply.ProtoReflect.Descriptor instead.
func (*UpdateAdminReply) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{4}
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
		mi := &file_bff_admin_v1_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAdminReq) ProtoMessage() {}

func (x *DeleteAdminReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteAdminReq.ProtoReflect.Descriptor instead.
func (*DeleteAdminReq) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{5}
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
		mi := &file_bff_admin_v1_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAdminReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAdminReply) ProtoMessage() {}

func (x *DeleteAdminReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteAdminReply.ProtoReflect.Descriptor instead.
func (*DeleteAdminReply) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{6}
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
		mi := &file_bff_admin_v1_admin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminReq) ProtoMessage() {}

func (x *GetAdminReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetAdminReq.ProtoReflect.Descriptor instead.
func (*GetAdminReq) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{7}
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
		mi := &file_bff_admin_v1_admin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminReply) ProtoMessage() {}

func (x *GetAdminReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetAdminReply.ProtoReflect.Descriptor instead.
func (*GetAdminReply) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{8}
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
		mi := &file_bff_admin_v1_admin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAdminReq) ProtoMessage() {}

func (x *ListAdminReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListAdminReq.ProtoReflect.Descriptor instead.
func (*ListAdminReq) Descriptor() ([]byte, []int) {
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{9}
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
		mi := &file_bff_admin_v1_admin_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAdminReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAdminReply) ProtoMessage() {}

func (x *ListAdminReply) ProtoReflect() protoreflect.Message {
	mi := &file_bff_admin_v1_admin_proto_msgTypes[10]
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
	return file_bff_admin_v1_admin_proto_rawDescGZIP(), []int{10}
}

var File_bff_admin_v1_admin_proto protoreflect.FileDescriptor

var file_bff_admin_v1_admin_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x66, 0x66, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x66, 0x66, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x0b, 0x0a, 0x09, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0xe9, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x07,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04,
	0x10, 0x00, 0x18, 0x32, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x72, 0x04, 0x10, 0x00, 0x18, 0x32, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x21, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x00, 0x18, 0x0f, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x6d, 0x6f, 0x74, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x00, 0x18, 0xff, 0x01, 0x52, 0x05, 0x6d,
	0x6f, 0x74, 0x74, 0x6f, 0x12, 0x26, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x00, 0x18,
	0x80, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x12, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x4b, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x39, 0x0a, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x66, 0x66, 0x5f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x12, 0x0a,
	0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0e, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xc5, 0x04, 0x0a, 0x05, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x12, 0x75, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x66, 0x66, 0x5f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x66, 0x66, 0x5f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a,
	0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x7d, 0x0a, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x66, 0x66, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x66, 0x66, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x15, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x72, 0x0a, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x66, 0x66, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x62, 0x66, 0x66, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x2a, 0x15, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x66, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x66, 0x66, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x66, 0x66, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x12, 0x12, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x6a, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x66, 0x66, 0x5f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x66, 0x66, 0x5f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x42, 0x1d, 0x5a, 0x1b, 0x66, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x62, 0x66, 0x66, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_bff_admin_v1_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_bff_admin_v1_admin_proto_goTypes = []interface{}{
	(*AdminInfo)(nil),        // 0: api.bff_admin.v1.AdminInfo
	(*CreateAdminReq)(nil),   // 1: api.bff_admin.v1.CreateAdminReq
	(*CreateAdminReply)(nil), // 2: api.bff_admin.v1.CreateAdminReply
	(*UpdateAdminReq)(nil),   // 3: api.bff_admin.v1.UpdateAdminReq
	(*UpdateAdminReply)(nil), // 4: api.bff_admin.v1.UpdateAdminReply
	(*DeleteAdminReq)(nil),   // 5: api.bff_admin.v1.DeleteAdminReq
	(*DeleteAdminReply)(nil), // 6: api.bff_admin.v1.DeleteAdminReply
	(*GetAdminReq)(nil),      // 7: api.bff_admin.v1.GetAdminReq
	(*GetAdminReply)(nil),    // 8: api.bff_admin.v1.GetAdminReply
	(*ListAdminReq)(nil),     // 9: api.bff_admin.v1.ListAdminReq
	(*ListAdminReply)(nil),   // 10: api.bff_admin.v1.ListAdminReply
}
var file_bff_admin_v1_admin_proto_depIdxs = []int32{
	0,  // 0: api.bff_admin.v1.UpdateAdminReq.adminInfo:type_name -> api.bff_admin.v1.AdminInfo
	1,  // 1: api.bff_admin.v1.Admin.CreateAdmin:input_type -> api.bff_admin.v1.CreateAdminReq
	3,  // 2: api.bff_admin.v1.Admin.UpdateAdmin:input_type -> api.bff_admin.v1.UpdateAdminReq
	5,  // 3: api.bff_admin.v1.Admin.DeleteAdmin:input_type -> api.bff_admin.v1.DeleteAdminReq
	7,  // 4: api.bff_admin.v1.Admin.GetAdmin:input_type -> api.bff_admin.v1.GetAdminReq
	9,  // 5: api.bff_admin.v1.Admin.ListAdmin:input_type -> api.bff_admin.v1.ListAdminReq
	2,  // 6: api.bff_admin.v1.Admin.CreateAdmin:output_type -> api.bff_admin.v1.CreateAdminReply
	4,  // 7: api.bff_admin.v1.Admin.UpdateAdmin:output_type -> api.bff_admin.v1.UpdateAdminReply
	6,  // 8: api.bff_admin.v1.Admin.DeleteAdmin:output_type -> api.bff_admin.v1.DeleteAdminReply
	8,  // 9: api.bff_admin.v1.Admin.GetAdmin:output_type -> api.bff_admin.v1.GetAdminReply
	10, // 10: api.bff_admin.v1.Admin.ListAdmin:output_type -> api.bff_admin.v1.ListAdminReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_bff_admin_v1_admin_proto_init() }
func file_bff_admin_v1_admin_proto_init() {
	if File_bff_admin_v1_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bff_admin_v1_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminInfo); i {
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
		file_bff_admin_v1_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_bff_admin_v1_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_bff_admin_v1_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_bff_admin_v1_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_bff_admin_v1_admin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_bff_admin_v1_admin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_bff_admin_v1_admin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_bff_admin_v1_admin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
		file_bff_admin_v1_admin_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   11,
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
