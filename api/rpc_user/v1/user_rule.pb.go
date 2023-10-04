// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: rpc_user/v1/user_rule.proto

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

type UserRuleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserRuleInfo) Reset() {
	*x = UserRuleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_v1_user_rule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRuleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRuleInfo) ProtoMessage() {}

func (x *UserRuleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_v1_user_rule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRuleInfo.ProtoReflect.Descriptor instead.
func (*UserRuleInfo) Descriptor() ([]byte, []int) {
	return file_rpc_user_v1_user_rule_proto_rawDescGZIP(), []int{0}
}

type UserRuleListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paginator *paginator.PaginatorReq `protobuf:"bytes,1,opt,name=paginator,proto3" json:"paginator,omitempty"` //分页
}

func (x *UserRuleListReq) Reset() {
	*x = UserRuleListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_v1_user_rule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRuleListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRuleListReq) ProtoMessage() {}

func (x *UserRuleListReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_v1_user_rule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRuleListReq.ProtoReflect.Descriptor instead.
func (*UserRuleListReq) Descriptor() ([]byte, []int) {
	return file_rpc_user_v1_user_rule_proto_rawDescGZIP(), []int{1}
}

func (x *UserRuleListReq) GetPaginator() *paginator.PaginatorReq {
	if x != nil {
		return x.Paginator
	}
	return nil
}

type UserRuleListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paginator *paginator.PaginatorReply `protobuf:"bytes,1,opt,name=paginator,proto3" json:"paginator,omitempty"` //分页
	List      []*UserRuleInfo           `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`           //列表
}

func (x *UserRuleListReply) Reset() {
	*x = UserRuleListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_v1_user_rule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRuleListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRuleListReply) ProtoMessage() {}

func (x *UserRuleListReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_v1_user_rule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRuleListReply.ProtoReflect.Descriptor instead.
func (*UserRuleListReply) Descriptor() ([]byte, []int) {
	return file_rpc_user_v1_user_rule_proto_rawDescGZIP(), []int{2}
}

func (x *UserRuleListReply) GetPaginator() *paginator.PaginatorReply {
	if x != nil {
		return x.Paginator
	}
	return nil
}

func (x *UserRuleListReply) GetList() []*UserRuleInfo {
	if x != nil {
		return x.List
	}
	return nil
}

type UserRuleInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UID int64 `protobuf:"varint,2,opt,name=UID,proto3" json:"UID,omitempty"` // uid
}

func (x *UserRuleInfoReq) Reset() {
	*x = UserRuleInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_v1_user_rule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRuleInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRuleInfoReq) ProtoMessage() {}

func (x *UserRuleInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_v1_user_rule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRuleInfoReq.ProtoReflect.Descriptor instead.
func (*UserRuleInfoReq) Descriptor() ([]byte, []int) {
	return file_rpc_user_v1_user_rule_proto_rawDescGZIP(), []int{3}
}

func (x *UserRuleInfoReq) GetUID() int64 {
	if x != nil {
		return x.UID
	}
	return 0
}

type UserRuleInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *UserRuleInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"` //用户规则信息
}

func (x *UserRuleInfoReply) Reset() {
	*x = UserRuleInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_v1_user_rule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRuleInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRuleInfoReply) ProtoMessage() {}

func (x *UserRuleInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_v1_user_rule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRuleInfoReply.ProtoReflect.Descriptor instead.
func (*UserRuleInfoReply) Descriptor() ([]byte, []int) {
	return file_rpc_user_v1_user_rule_proto_rawDescGZIP(), []int{4}
}

func (x *UserRuleInfoReply) GetInfo() *UserRuleInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type UserRuleStoreReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserRuleStoreReq) Reset() {
	*x = UserRuleStoreReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_v1_user_rule_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRuleStoreReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRuleStoreReq) ProtoMessage() {}

func (x *UserRuleStoreReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_v1_user_rule_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRuleStoreReq.ProtoReflect.Descriptor instead.
func (*UserRuleStoreReq) Descriptor() ([]byte, []int) {
	return file_rpc_user_v1_user_rule_proto_rawDescGZIP(), []int{5}
}

type UserRuleStoreReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"` // Id
}

func (x *UserRuleStoreReply) Reset() {
	*x = UserRuleStoreReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_v1_user_rule_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRuleStoreReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRuleStoreReply) ProtoMessage() {}

func (x *UserRuleStoreReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_v1_user_rule_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRuleStoreReply.ProtoReflect.Descriptor instead.
func (*UserRuleStoreReply) Descriptor() ([]byte, []int) {
	return file_rpc_user_v1_user_rule_proto_rawDescGZIP(), []int{6}
}

func (x *UserRuleStoreReply) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type UserRuleDelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"` //ID集合
}

func (x *UserRuleDelReq) Reset() {
	*x = UserRuleDelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_v1_user_rule_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRuleDelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRuleDelReq) ProtoMessage() {}

func (x *UserRuleDelReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_v1_user_rule_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRuleDelReq.ProtoReflect.Descriptor instead.
func (*UserRuleDelReq) Descriptor() ([]byte, []int) {
	return file_rpc_user_v1_user_rule_proto_rawDescGZIP(), []int{7}
}

func (x *UserRuleDelReq) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type UserRuleDelReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserRuleDelReply) Reset() {
	*x = UserRuleDelReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_v1_user_rule_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRuleDelReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRuleDelReply) ProtoMessage() {}

func (x *UserRuleDelReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_v1_user_rule_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRuleDelReply.ProtoReflect.Descriptor instead.
func (*UserRuleDelReply) Descriptor() ([]byte, []int) {
	return file_rpc_user_v1_user_rule_proto_rawDescGZIP(), []int{8}
}

var File_rpc_user_v1_user_rule_proto protoreflect.FileDescriptor

var file_rpc_user_v1_user_rule_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x19,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a, 0x0c, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x48, 0x0a, 0x0f, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x35, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x52, 0x09, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x22, 0x7f, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x09, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x31, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x23, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x55, 0x49, 0x44, 0x22, 0x46, 0x0a, 0x11, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x31,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x22, 0x24, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x52, 0x75, 0x6c,
	0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x22, 0x0a, 0x0e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22,
	0x12, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x32, 0xe2, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65,
	0x12, 0x54, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x54, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x75,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x75, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x57, 0x0a, 0x0d,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x21, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x51, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x75, 0x6c,
	0x65, 0x44, 0x65, 0x6c, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x44,
	0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65,
	0x44, 0x65, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x1c, 0x5a, 0x1a, 0x66, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_user_v1_user_rule_proto_rawDescOnce sync.Once
	file_rpc_user_v1_user_rule_proto_rawDescData = file_rpc_user_v1_user_rule_proto_rawDesc
)

func file_rpc_user_v1_user_rule_proto_rawDescGZIP() []byte {
	file_rpc_user_v1_user_rule_proto_rawDescOnce.Do(func() {
		file_rpc_user_v1_user_rule_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_user_v1_user_rule_proto_rawDescData)
	})
	return file_rpc_user_v1_user_rule_proto_rawDescData
}

var file_rpc_user_v1_user_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rpc_user_v1_user_rule_proto_goTypes = []interface{}{
	(*UserRuleInfo)(nil),             // 0: api.rpc_user.v1.UserRuleInfo
	(*UserRuleListReq)(nil),          // 1: api.rpc_user.v1.UserRuleListReq
	(*UserRuleListReply)(nil),        // 2: api.rpc_user.v1.UserRuleListReply
	(*UserRuleInfoReq)(nil),          // 3: api.rpc_user.v1.UserRuleInfoReq
	(*UserRuleInfoReply)(nil),        // 4: api.rpc_user.v1.UserRuleInfoReply
	(*UserRuleStoreReq)(nil),         // 5: api.rpc_user.v1.UserRuleStoreReq
	(*UserRuleStoreReply)(nil),       // 6: api.rpc_user.v1.UserRuleStoreReply
	(*UserRuleDelReq)(nil),           // 7: api.rpc_user.v1.UserRuleDelReq
	(*UserRuleDelReply)(nil),         // 8: api.rpc_user.v1.UserRuleDelReply
	(*paginator.PaginatorReq)(nil),   // 9: paginator.PaginatorReq
	(*paginator.PaginatorReply)(nil), // 10: paginator.PaginatorReply
}
var file_rpc_user_v1_user_rule_proto_depIdxs = []int32{
	9,  // 0: api.rpc_user.v1.UserRuleListReq.paginator:type_name -> paginator.PaginatorReq
	10, // 1: api.rpc_user.v1.UserRuleListReply.paginator:type_name -> paginator.PaginatorReply
	0,  // 2: api.rpc_user.v1.UserRuleListReply.list:type_name -> api.rpc_user.v1.UserRuleInfo
	0,  // 3: api.rpc_user.v1.UserRuleInfoReply.info:type_name -> api.rpc_user.v1.UserRuleInfo
	1,  // 4: api.rpc_user.v1.UserRule.UserRuleList:input_type -> api.rpc_user.v1.UserRuleListReq
	3,  // 5: api.rpc_user.v1.UserRule.UserRuleInfo:input_type -> api.rpc_user.v1.UserRuleInfoReq
	5,  // 6: api.rpc_user.v1.UserRule.UserRuleStore:input_type -> api.rpc_user.v1.UserRuleStoreReq
	7,  // 7: api.rpc_user.v1.UserRule.UserRuleDel:input_type -> api.rpc_user.v1.UserRuleDelReq
	2,  // 8: api.rpc_user.v1.UserRule.UserRuleList:output_type -> api.rpc_user.v1.UserRuleListReply
	4,  // 9: api.rpc_user.v1.UserRule.UserRuleInfo:output_type -> api.rpc_user.v1.UserRuleInfoReply
	6,  // 10: api.rpc_user.v1.UserRule.UserRuleStore:output_type -> api.rpc_user.v1.UserRuleStoreReply
	8,  // 11: api.rpc_user.v1.UserRule.UserRuleDel:output_type -> api.rpc_user.v1.UserRuleDelReply
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_rpc_user_v1_user_rule_proto_init() }
func file_rpc_user_v1_user_rule_proto_init() {
	if File_rpc_user_v1_user_rule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_user_v1_user_rule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRuleInfo); i {
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
		file_rpc_user_v1_user_rule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRuleListReq); i {
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
		file_rpc_user_v1_user_rule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRuleListReply); i {
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
		file_rpc_user_v1_user_rule_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRuleInfoReq); i {
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
		file_rpc_user_v1_user_rule_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRuleInfoReply); i {
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
		file_rpc_user_v1_user_rule_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRuleStoreReq); i {
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
		file_rpc_user_v1_user_rule_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRuleStoreReply); i {
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
		file_rpc_user_v1_user_rule_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRuleDelReq); i {
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
		file_rpc_user_v1_user_rule_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRuleDelReply); i {
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
			RawDescriptor: file_rpc_user_v1_user_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_user_v1_user_rule_proto_goTypes,
		DependencyIndexes: file_rpc_user_v1_user_rule_proto_depIdxs,
		MessageInfos:      file_rpc_user_v1_user_rule_proto_msgTypes,
	}.Build()
	File_rpc_user_v1_user_rule_proto = out.File
	file_rpc_user_v1_user_rule_proto_rawDesc = nil
	file_rpc_user_v1_user_rule_proto_goTypes = nil
	file_rpc_user_v1_user_rule_proto_depIdxs = nil
}
