// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: rpc_sys/v1/sys_log.proto

package v1

import (
	common "fkratos/api/common"
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

type SysLogInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                // 编号
	AdminID   string `protobuf:"bytes,2,opt,name=adminID,proto3" json:"adminID,omitempty"`      // 管理员ID
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`    // 账号
	Ip        string `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`                // ip
	Uri       string `protobuf:"bytes,5,opt,name=uri,proto3" json:"uri,omitempty"`              // 请求路径
	UriDesc   string `protobuf:"bytes,6,opt,name=uriDesc,proto3" json:"uriDesc,omitempty"`      // 请求描述
	Useragent string `protobuf:"bytes,7,opt,name=useragent,proto3" json:"useragent,omitempty"`  // 浏览器标识
	Req       string `protobuf:"bytes,8,opt,name=req,proto3" json:"req,omitempty"`              // 请求数据
	Resp      string `protobuf:"bytes,9,opt,name=resp,proto3" json:"resp,omitempty"`            // 响应数据
	CreatedAt string `protobuf:"bytes,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"` // 创建时间
}

func (x *SysLogInfo) Reset() {
	*x = SysLogInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysLogInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysLogInfo) ProtoMessage() {}

func (x *SysLogInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysLogInfo.ProtoReflect.Descriptor instead.
func (*SysLogInfo) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_log_proto_rawDescGZIP(), []int{0}
}

func (x *SysLogInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SysLogInfo) GetAdminID() string {
	if x != nil {
		return x.AdminID
	}
	return ""
}

func (x *SysLogInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SysLogInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *SysLogInfo) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *SysLogInfo) GetUriDesc() string {
	if x != nil {
		return x.UriDesc
	}
	return ""
}

func (x *SysLogInfo) GetUseragent() string {
	if x != nil {
		return x.Useragent
	}
	return ""
}

func (x *SysLogInfo) GetReq() string {
	if x != nil {
		return x.Req
	}
	return ""
}

func (x *SysLogInfo) GetResp() string {
	if x != nil {
		return x.Resp
	}
	return ""
}

func (x *SysLogInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

// 请求-日志列表
type SysLogListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`         //页码
	PageSize int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"` //页数
}

func (x *SysLogListReq) Reset() {
	*x = SysLogListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysLogListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysLogListReq) ProtoMessage() {}

func (x *SysLogListReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysLogListReq.ProtoReflect.Descriptor instead.
func (*SysLogListReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_log_proto_rawDescGZIP(), []int{1}
}

func (x *SysLogListReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SysLogListReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// 响应-日志列表
type SysLogListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List      []*SysLogInfo     `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`           //管理员列表
	Paginator *common.Paginator `protobuf:"bytes,2,opt,name=paginator,proto3" json:"paginator,omitempty"` //分页
}

func (x *SysLogListResp) Reset() {
	*x = SysLogListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysLogListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysLogListResp) ProtoMessage() {}

func (x *SysLogListResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysLogListResp.ProtoReflect.Descriptor instead.
func (*SysLogListResp) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_log_proto_rawDescGZIP(), []int{2}
}

func (x *SysLogListResp) GetList() []*SysLogInfo {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *SysLogListResp) GetPaginator() *common.Paginator {
	if x != nil {
		return x.Paginator
	}
	return nil
}

// 请求-单条日志
type SysLogInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SysLogInfoReq) Reset() {
	*x = SysLogInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysLogInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysLogInfoReq) ProtoMessage() {}

func (x *SysLogInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysLogInfoReq.ProtoReflect.Descriptor instead.
func (*SysLogInfoReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_log_proto_rawDescGZIP(), []int{3}
}

func (x *SysLogInfoReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// 响应-单条日志
type SysLogInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *SysLogInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SysLogInfoResp) Reset() {
	*x = SysLogInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysLogInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysLogInfoResp) ProtoMessage() {}

func (x *SysLogInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysLogInfoResp.ProtoReflect.Descriptor instead.
func (*SysLogInfoResp) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_log_proto_rawDescGZIP(), []int{4}
}

func (x *SysLogInfoResp) GetInfo() *SysLogInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// 请求-日志存储
type SysLogStoreReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminID   string `protobuf:"bytes,1,opt,name=adminID,proto3" json:"adminID,omitempty"`     // 管理员ID
	Username  string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`   // 账号
	Ip        string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`               // ip
	Uri       string `protobuf:"bytes,4,opt,name=uri,proto3" json:"uri,omitempty"`             // 请求路径
	UriDesc   string `protobuf:"bytes,5,opt,name=uriDesc,proto3" json:"uriDesc,omitempty"`     // 请求描述
	Useragent string `protobuf:"bytes,6,opt,name=useragent,proto3" json:"useragent,omitempty"` // 浏览器标识
	Req       string `protobuf:"bytes,7,opt,name=req,proto3" json:"req,omitempty"`             // 请求数据
	Resp      string `protobuf:"bytes,8,opt,name=resp,proto3" json:"resp,omitempty"`           // 响应数据
}

func (x *SysLogStoreReq) Reset() {
	*x = SysLogStoreReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_log_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysLogStoreReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysLogStoreReq) ProtoMessage() {}

func (x *SysLogStoreReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_log_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysLogStoreReq.ProtoReflect.Descriptor instead.
func (*SysLogStoreReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_log_proto_rawDescGZIP(), []int{5}
}

func (x *SysLogStoreReq) GetAdminID() string {
	if x != nil {
		return x.AdminID
	}
	return ""
}

func (x *SysLogStoreReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SysLogStoreReq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *SysLogStoreReq) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *SysLogStoreReq) GetUriDesc() string {
	if x != nil {
		return x.UriDesc
	}
	return ""
}

func (x *SysLogStoreReq) GetUseragent() string {
	if x != nil {
		return x.Useragent
	}
	return ""
}

func (x *SysLogStoreReq) GetReq() string {
	if x != nil {
		return x.Req
	}
	return ""
}

func (x *SysLogStoreReq) GetResp() string {
	if x != nil {
		return x.Resp
	}
	return ""
}

// 响应-日志存储
type SysLogStoreResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *SysLogInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SysLogStoreResp) Reset() {
	*x = SysLogStoreResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_log_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysLogStoreResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysLogStoreResp) ProtoMessage() {}

func (x *SysLogStoreResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_log_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysLogStoreResp.ProtoReflect.Descriptor instead.
func (*SysLogStoreResp) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_log_proto_rawDescGZIP(), []int{6}
}

func (x *SysLogStoreResp) GetInfo() *SysLogInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_rpc_sys_v1_sys_log_proto protoreflect.FileDescriptor

var file_rpc_sys_v1_sys_log_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73,
	0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf0, 0x01, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x72, 0x69, 0x44, 0x65, 0x73,
	0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x72, 0x69, 0x44, 0x65, 0x73, 0x63,
	0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x3f, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x22, 0x71, 0x0a, 0x0e, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x1f, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x0e, 0x53, 0x79, 0x73, 0x4c, 0x6f,
	0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x67, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x0e, 0x53, 0x79,
	0x73, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x69, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x72, 0x69, 0x44, 0x65, 0x73, 0x63, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x72, 0x69, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1c,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x65, 0x71, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65,
	0x73, 0x70, 0x22, 0x41, 0x0a, 0x0f, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xe7, 0x01, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x43, 0x0a,
	0x0a, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x4b, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a,
	0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x4e, 0x0a, 0x0b, 0x53, 0x79, 0x73, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x1e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x79, 0x73, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x79, 0x73, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42,
	0x1b, 0x5a, 0x19, 0x66, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_sys_v1_sys_log_proto_rawDescOnce sync.Once
	file_rpc_sys_v1_sys_log_proto_rawDescData = file_rpc_sys_v1_sys_log_proto_rawDesc
)

func file_rpc_sys_v1_sys_log_proto_rawDescGZIP() []byte {
	file_rpc_sys_v1_sys_log_proto_rawDescOnce.Do(func() {
		file_rpc_sys_v1_sys_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_sys_v1_sys_log_proto_rawDescData)
	})
	return file_rpc_sys_v1_sys_log_proto_rawDescData
}

var file_rpc_sys_v1_sys_log_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rpc_sys_v1_sys_log_proto_goTypes = []interface{}{
	(*SysLogInfo)(nil),           // 0: api.rpc_sys.v1.SysLogInfo
	(*SysLogListReq)(nil),        // 1: api.rpc_sys.v1.SysLogListReq
	(*SysLogListResp)(nil),       // 2: api.rpc_sys.v1.SysLogListResp
	(*SysLogInfoReq)(nil),        // 3: api.rpc_sys.v1.SysLogInfoReq
	(*SysLogInfoResp)(nil),       // 4: api.rpc_sys.v1.SysLogInfoResp
	(*SysLogStoreReq)(nil),       // 5: api.rpc_sys.v1.SysLogStoreReq
	(*SysLogStoreResp)(nil),      // 6: api.rpc_sys.v1.SysLogStoreResp
	(*common.Paginator)(nil),     // 7: common.Paginator
	(*common.SearchListReq)(nil), // 8: common.SearchListReq
}
var file_rpc_sys_v1_sys_log_proto_depIdxs = []int32{
	0, // 0: api.rpc_sys.v1.SysLogListResp.list:type_name -> api.rpc_sys.v1.SysLogInfo
	7, // 1: api.rpc_sys.v1.SysLogListResp.paginator:type_name -> common.Paginator
	0, // 2: api.rpc_sys.v1.SysLogInfoResp.info:type_name -> api.rpc_sys.v1.SysLogInfo
	0, // 3: api.rpc_sys.v1.SysLogStoreResp.info:type_name -> api.rpc_sys.v1.SysLogInfo
	8, // 4: api.rpc_sys.v1.Log.SysLogList:input_type -> common.SearchListReq
	3, // 5: api.rpc_sys.v1.Log.SysLogInfo:input_type -> api.rpc_sys.v1.SysLogInfoReq
	5, // 6: api.rpc_sys.v1.Log.SysLogStore:input_type -> api.rpc_sys.v1.SysLogStoreReq
	2, // 7: api.rpc_sys.v1.Log.SysLogList:output_type -> api.rpc_sys.v1.SysLogListResp
	4, // 8: api.rpc_sys.v1.Log.SysLogInfo:output_type -> api.rpc_sys.v1.SysLogInfoResp
	6, // 9: api.rpc_sys.v1.Log.SysLogStore:output_type -> api.rpc_sys.v1.SysLogStoreResp
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rpc_sys_v1_sys_log_proto_init() }
func file_rpc_sys_v1_sys_log_proto_init() {
	if File_rpc_sys_v1_sys_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_sys_v1_sys_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysLogInfo); i {
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
		file_rpc_sys_v1_sys_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysLogListReq); i {
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
		file_rpc_sys_v1_sys_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysLogListResp); i {
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
		file_rpc_sys_v1_sys_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysLogInfoReq); i {
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
		file_rpc_sys_v1_sys_log_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysLogInfoResp); i {
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
		file_rpc_sys_v1_sys_log_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysLogStoreReq); i {
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
		file_rpc_sys_v1_sys_log_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysLogStoreResp); i {
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
			RawDescriptor: file_rpc_sys_v1_sys_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_sys_v1_sys_log_proto_goTypes,
		DependencyIndexes: file_rpc_sys_v1_sys_log_proto_depIdxs,
		MessageInfos:      file_rpc_sys_v1_sys_log_proto_msgTypes,
	}.Build()
	File_rpc_sys_v1_sys_log_proto = out.File
	file_rpc_sys_v1_sys_log_proto_rawDesc = nil
	file_rpc_sys_v1_sys_log_proto_goTypes = nil
	file_rpc_sys_v1_sys_log_proto_depIdxs = nil
}
