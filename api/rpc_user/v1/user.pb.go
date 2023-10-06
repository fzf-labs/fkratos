// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: rpc_user/v1/user.proto

package v1

import (
	paginator "fkratos/api/paginator"
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

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                // Id
	Uid       string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`              // uid
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`    // 用户名
	Phone     string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`          // 手机
	Email     string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`          // 邮箱
	Nickname  string `protobuf:"bytes,6,opt,name=nickname,proto3" json:"nickname,omitempty"`    // 昵称
	Sex       int32  `protobuf:"varint,7,opt,name=sex,proto3" json:"sex,omitempty"`             // 性别（0未知 1男 2女）
	Avatar    string `protobuf:"bytes,8,opt,name=avatar,proto3" json:"avatar,omitempty"`        // 头像
	Profile   string `protobuf:"bytes,9,opt,name=profile,proto3" json:"profile,omitempty"`      // 简介
	Other     string `protobuf:"bytes,10,opt,name=other,proto3" json:"other,omitempty"`         // 其他
	Status    int32  `protobuf:"varint,11,opt,name=status,proto3" json:"status,omitempty"`      // 状态
	CreatedAt string `protobuf:"bytes,12,opt,name=createdAt,proto3" json:"createdAt,omitempty"` // 创建时间
	UpdatedAt string `protobuf:"bytes,13,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"` // 更新时间
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_rpc_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserInfo) GetSex() int32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *UserInfo) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserInfo) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *UserInfo) GetOther() string {
	if x != nil {
		return x.Other
	}
	return ""
}

func (x *UserInfo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UserInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UserListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paginator *paginator.PaginatorReq `protobuf:"bytes,1,opt,name=paginator,proto3" json:"paginator,omitempty"` //分页
}

func (x *UserListReq) Reset() {
	*x = UserListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListReq) ProtoMessage() {}

func (x *UserListReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListReq.ProtoReflect.Descriptor instead.
func (*UserListReq) Descriptor() ([]byte, []int) {
	return file_rpc_user_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserListReq) GetPaginator() *paginator.PaginatorReq {
	if x != nil {
		return x.Paginator
	}
	return nil
}

type UserListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paginator *paginator.PaginatorReply `protobuf:"bytes,1,opt,name=paginator,proto3" json:"paginator,omitempty"` //分页
	List      []*UserInfo               `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`           //列表
}

func (x *UserListReply) Reset() {
	*x = UserListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListReply) ProtoMessage() {}

func (x *UserListReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListReply.ProtoReflect.Descriptor instead.
func (*UserListReply) Descriptor() ([]byte, []int) {
	return file_rpc_user_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserListReply) GetPaginator() *paginator.PaginatorReply {
	if x != nil {
		return x.Paginator
	}
	return nil
}

func (x *UserListReply) GetList() []*UserInfo {
	if x != nil {
		return x.List
	}
	return nil
}

type UserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"` // uid
}

func (x *UserInfoReq) Reset() {
	*x = UserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_v1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoReq) ProtoMessage() {}

func (x *UserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_v1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoReq.ProtoReflect.Descriptor instead.
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return file_rpc_user_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserInfoReq) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type UserInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *UserInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"` //用户信息
}

func (x *UserInfoReply) Reset() {
	*x = UserInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_v1_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoReply) ProtoMessage() {}

func (x *UserInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_v1_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoReply.ProtoReflect.Descriptor instead.
func (*UserInfoReply) Descriptor() ([]byte, []int) {
	return file_rpc_user_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserInfoReply) GetInfo() *UserInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type UserStoreReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                    // Id
	Username    string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`        // 用户名
	Phone       string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`              // 手机
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`              // 邮箱
	Password    string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`        // 密码
	Nickname    string `protobuf:"bytes,6,opt,name=nickname,proto3" json:"nickname,omitempty"`        // 昵称
	Sex         int32  `protobuf:"varint,7,opt,name=sex,proto3" json:"sex,omitempty"`                 // 性别（0未知 1男 2女）
	Avatar      string `protobuf:"bytes,8,opt,name=avatar,proto3" json:"avatar,omitempty"`            // 头像
	Profile     string `protobuf:"bytes,9,opt,name=profile,proto3" json:"profile,omitempty"`          // 简介
	Status      int32  `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`          // 状态
	UserGroupID string `protobuf:"bytes,11,opt,name=userGroupID,proto3" json:"userGroupID,omitempty"` // 分组ID
}

func (x *UserStoreReq) Reset() {
	*x = UserStoreReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_v1_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStoreReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStoreReq) ProtoMessage() {}

func (x *UserStoreReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_v1_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStoreReq.ProtoReflect.Descriptor instead.
func (*UserStoreReq) Descriptor() ([]byte, []int) {
	return file_rpc_user_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserStoreReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserStoreReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserStoreReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserStoreReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserStoreReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserStoreReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserStoreReq) GetSex() int32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *UserStoreReq) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserStoreReq) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *UserStoreReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UserStoreReq) GetUserGroupID() string {
	if x != nil {
		return x.UserGroupID
	}
	return ""
}

type UserStoreReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // Id
}

func (x *UserStoreReply) Reset() {
	*x = UserStoreReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_v1_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStoreReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStoreReply) ProtoMessage() {}

func (x *UserStoreReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_v1_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStoreReply.ProtoReflect.Descriptor instead.
func (*UserStoreReply) Descriptor() ([]byte, []int) {
	return file_rpc_user_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserStoreReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UserDelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"` //ID集合
}

func (x *UserDelReq) Reset() {
	*x = UserDelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_v1_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDelReq) ProtoMessage() {}

func (x *UserDelReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_v1_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDelReq.ProtoReflect.Descriptor instead.
func (*UserDelReq) Descriptor() ([]byte, []int) {
	return file_rpc_user_v1_user_proto_rawDescGZIP(), []int{7}
}

func (x *UserDelReq) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type UserDelReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserDelReply) Reset() {
	*x = UserDelReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_v1_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDelReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDelReply) ProtoMessage() {}

func (x *UserDelReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_v1_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDelReply.ProtoReflect.Descriptor instead.
func (*UserDelReply) Descriptor() ([]byte, []int) {
	return file_rpc_user_v1_user_proto_rawDescGZIP(), []int{8}
}

var File_rpc_user_v1_user_proto protoreflect.FileDescriptor

var file_rpc_user_v1_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70,
	0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x48, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x39, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x52, 0x09, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x22,
	0x7b, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x3b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x52, 0x09, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2d, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x1f, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x3e, 0x0a,
	0x0d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2d,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x9c, 0x02,
	0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x65, 0x78,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x73,
	0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x75, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x22, 0x20, 0x0a, 0x0e,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1e,
	0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x0e,
	0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xae,
	0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x48, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4b, 0x0a, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70,
	0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x45, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x6c, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x1c, 0x5a, 0x1a, 0x66, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_user_v1_user_proto_rawDescOnce sync.Once
	file_rpc_user_v1_user_proto_rawDescData = file_rpc_user_v1_user_proto_rawDesc
)

func file_rpc_user_v1_user_proto_rawDescGZIP() []byte {
	file_rpc_user_v1_user_proto_rawDescOnce.Do(func() {
		file_rpc_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_user_v1_user_proto_rawDescData)
	})
	return file_rpc_user_v1_user_proto_rawDescData
}

var file_rpc_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rpc_user_v1_user_proto_goTypes = []interface{}{
	(*UserInfo)(nil),                 // 0: api.rpc_user.v1.UserInfo
	(*UserListReq)(nil),              // 1: api.rpc_user.v1.UserListReq
	(*UserListReply)(nil),            // 2: api.rpc_user.v1.UserListReply
	(*UserInfoReq)(nil),              // 3: api.rpc_user.v1.UserInfoReq
	(*UserInfoReply)(nil),            // 4: api.rpc_user.v1.UserInfoReply
	(*UserStoreReq)(nil),             // 5: api.rpc_user.v1.UserStoreReq
	(*UserStoreReply)(nil),           // 6: api.rpc_user.v1.UserStoreReply
	(*UserDelReq)(nil),               // 7: api.rpc_user.v1.UserDelReq
	(*UserDelReply)(nil),             // 8: api.rpc_user.v1.UserDelReply
	(*paginator.PaginatorReq)(nil),   // 9: api.paginator.PaginatorReq
	(*paginator.PaginatorReply)(nil), // 10: api.paginator.PaginatorReply
}
var file_rpc_user_v1_user_proto_depIdxs = []int32{
	9,  // 0: api.rpc_user.v1.UserListReq.paginator:type_name -> api.paginator.PaginatorReq
	10, // 1: api.rpc_user.v1.UserListReply.paginator:type_name -> api.paginator.PaginatorReply
	0,  // 2: api.rpc_user.v1.UserListReply.list:type_name -> api.rpc_user.v1.UserInfo
	0,  // 3: api.rpc_user.v1.UserInfoReply.info:type_name -> api.rpc_user.v1.UserInfo
	1,  // 4: api.rpc_user.v1.User.UserList:input_type -> api.rpc_user.v1.UserListReq
	3,  // 5: api.rpc_user.v1.User.UserInfo:input_type -> api.rpc_user.v1.UserInfoReq
	5,  // 6: api.rpc_user.v1.User.UserStore:input_type -> api.rpc_user.v1.UserStoreReq
	7,  // 7: api.rpc_user.v1.User.UserDel:input_type -> api.rpc_user.v1.UserDelReq
	2,  // 8: api.rpc_user.v1.User.UserList:output_type -> api.rpc_user.v1.UserListReply
	4,  // 9: api.rpc_user.v1.User.UserInfo:output_type -> api.rpc_user.v1.UserInfoReply
	6,  // 10: api.rpc_user.v1.User.UserStore:output_type -> api.rpc_user.v1.UserStoreReply
	8,  // 11: api.rpc_user.v1.User.UserDel:output_type -> api.rpc_user.v1.UserDelReply
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_rpc_user_v1_user_proto_init() }
func file_rpc_user_v1_user_proto_init() {
	if File_rpc_user_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_user_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_rpc_user_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListReq); i {
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
		file_rpc_user_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListReply); i {
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
		file_rpc_user_v1_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoReq); i {
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
		file_rpc_user_v1_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoReply); i {
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
		file_rpc_user_v1_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserStoreReq); i {
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
		file_rpc_user_v1_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserStoreReply); i {
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
		file_rpc_user_v1_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDelReq); i {
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
		file_rpc_user_v1_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDelReply); i {
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
			RawDescriptor: file_rpc_user_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_user_v1_user_proto_goTypes,
		DependencyIndexes: file_rpc_user_v1_user_proto_depIdxs,
		MessageInfos:      file_rpc_user_v1_user_proto_msgTypes,
	}.Build()
	File_rpc_user_v1_user_proto = out.File
	file_rpc_user_v1_user_proto_rawDesc = nil
	file_rpc_user_v1_user_proto_goTypes = nil
	file_rpc_user_v1_user_proto_depIdxs = nil
}
