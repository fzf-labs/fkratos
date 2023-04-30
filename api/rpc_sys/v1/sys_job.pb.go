// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: rpc_sys/v1/sys_job.proto

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

type SysJobInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`               // 编号
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`           // 岗位名称
	Code      string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`           //岗位编码
	Remark    string `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`       //备注
	Status    int32  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`      // 0=禁用 1=开启
	Sort      int32  `protobuf:"varint,6,opt,name=sort,proto3" json:"sort,omitempty"`          // 排序值
	CreatedAt string `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"` // 创建时间
	UpdatedAt string `protobuf:"bytes,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"` // 更新时间
}

func (x *SysJobInfo) Reset() {
	*x = SysJobInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysJobInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysJobInfo) ProtoMessage() {}

func (x *SysJobInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysJobInfo.ProtoReflect.Descriptor instead.
func (*SysJobInfo) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_job_proto_rawDescGZIP(), []int{0}
}

func (x *SysJobInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SysJobInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SysJobInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SysJobInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *SysJobInfo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SysJobInfo) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *SysJobInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SysJobInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// 请求-岗位列表
type SysJobListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page        int32                 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`              //页码
	PageSize    int32                 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`      //页数
	QuickSearch string                `protobuf:"bytes,3,opt,name=quickSearch,proto3" json:"quickSearch,omitempty"` //快捷搜索
	Search      []*common.SearchParam `protobuf:"bytes,4,rep,name=search,proto3" json:"search,omitempty"`           //详细搜索
}

func (x *SysJobListReq) Reset() {
	*x = SysJobListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysJobListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysJobListReq) ProtoMessage() {}

func (x *SysJobListReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysJobListReq.ProtoReflect.Descriptor instead.
func (*SysJobListReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_job_proto_rawDescGZIP(), []int{1}
}

func (x *SysJobListReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SysJobListReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SysJobListReq) GetQuickSearch() string {
	if x != nil {
		return x.QuickSearch
	}
	return ""
}

func (x *SysJobListReq) GetSearch() []*common.SearchParam {
	if x != nil {
		return x.Search
	}
	return nil
}

// 响应-岗位列表
type SysJobListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List      []*SysJobInfo     `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`           //列表
	Paginator *common.Paginator `protobuf:"bytes,2,opt,name=paginator,proto3" json:"paginator,omitempty"` //分页
}

func (x *SysJobListReply) Reset() {
	*x = SysJobListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysJobListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysJobListReply) ProtoMessage() {}

func (x *SysJobListReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysJobListReply.ProtoReflect.Descriptor instead.
func (*SysJobListReply) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_job_proto_rawDescGZIP(), []int{2}
}

func (x *SysJobListReply) GetList() []*SysJobInfo {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *SysJobListReply) GetPaginator() *common.Paginator {
	if x != nil {
		return x.Paginator
	}
	return nil
}

// 请求-岗位信息
type SysJobInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SysJobInfoReq) Reset() {
	*x = SysJobInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_job_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysJobInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysJobInfoReq) ProtoMessage() {}

func (x *SysJobInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_job_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysJobInfoReq.ProtoReflect.Descriptor instead.
func (*SysJobInfoReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_job_proto_rawDescGZIP(), []int{3}
}

func (x *SysJobInfoReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// 响应-岗位信息
type SysJobInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *SysJobInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SysJobInfoReply) Reset() {
	*x = SysJobInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_job_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysJobInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysJobInfoReply) ProtoMessage() {}

func (x *SysJobInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_job_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysJobInfoReply.ProtoReflect.Descriptor instead.
func (*SysJobInfoReply) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_job_proto_rawDescGZIP(), []int{4}
}

func (x *SysJobInfoReply) GetInfo() *SysJobInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// 请求-岗位保存
type SysJobStoreReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`          // 编号
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`      // 岗位名称
	Code   string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`      //岗位编码
	Remark string `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`  //备注
	Status int32  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"` // 0=禁用 1=开启
	Sort   int32  `protobuf:"varint,6,opt,name=sort,proto3" json:"sort,omitempty"`     // 排序值
}

func (x *SysJobStoreReq) Reset() {
	*x = SysJobStoreReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_job_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysJobStoreReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysJobStoreReq) ProtoMessage() {}

func (x *SysJobStoreReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_job_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysJobStoreReq.ProtoReflect.Descriptor instead.
func (*SysJobStoreReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_job_proto_rawDescGZIP(), []int{5}
}

func (x *SysJobStoreReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SysJobStoreReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SysJobStoreReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SysJobStoreReq) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *SysJobStoreReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SysJobStoreReq) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

// 响应-岗位保存
type SysJobStoreReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SysJobStoreReply) Reset() {
	*x = SysJobStoreReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_job_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysJobStoreReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysJobStoreReply) ProtoMessage() {}

func (x *SysJobStoreReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_job_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysJobStoreReply.ProtoReflect.Descriptor instead.
func (*SysJobStoreReply) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_job_proto_rawDescGZIP(), []int{6}
}

// 请求-岗位删除
type SysJobDelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *SysJobDelReq) Reset() {
	*x = SysJobDelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_job_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysJobDelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysJobDelReq) ProtoMessage() {}

func (x *SysJobDelReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_job_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysJobDelReq.ProtoReflect.Descriptor instead.
func (*SysJobDelReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_job_proto_rawDescGZIP(), []int{7}
}

func (x *SysJobDelReq) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

// 响应-岗位删除
type SysJobDelReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SysJobDelReply) Reset() {
	*x = SysJobDelReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_job_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysJobDelReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysJobDelReply) ProtoMessage() {}

func (x *SysJobDelReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_job_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysJobDelReply.ProtoReflect.Descriptor instead.
func (*SysJobDelReply) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_job_proto_rawDescGZIP(), []int{8}
}

var File_rpc_sys_v1_sys_job_proto protoreflect.FileDescriptor

var file_rpc_sys_v1_sys_job_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73,
	0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc4, 0x01, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x4a, 0x6f,
	0x62, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x71, 0x75, 0x69, 0x63,
	0x6b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x71,
	0x75, 0x69, 0x63, 0x6b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x72, 0x0a, 0x0f, 0x53, 0x79, 0x73, 0x4a, 0x6f,
	0x62, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x4a, 0x6f, 0x62,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72,
	0x52, 0x09, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x1f, 0x0a, 0x0d, 0x53,
	0x79, 0x73, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x0f,
	0x53, 0x79, 0x73, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x2e, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x79, 0x73, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22,
	0x8c, 0x01, 0x0a, 0x0e, 0x53, 0x79, 0x73, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x12,
	0x0a, 0x10, 0x53, 0x79, 0x73, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x20, 0x0a, 0x0c, 0x53, 0x79, 0x73, 0x4a, 0x6f, 0x62, 0x44, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x79, 0x73, 0x4a, 0x6f, 0x62, 0x44, 0x65,
	0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xbd, 0x02, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x4c,
	0x0a, 0x0a, 0x53, 0x79, 0x73, 0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79,
	0x73, 0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73,
	0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4c, 0x0a, 0x0a,
	0x53, 0x79, 0x73, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x4a,
	0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x4a, 0x6f,
	0x62, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4f, 0x0a, 0x0b, 0x53, 0x79,
	0x73, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x4a, 0x6f,
	0x62, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x4a, 0x6f,
	0x62, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x49, 0x0a, 0x09, 0x53,
	0x79, 0x73, 0x4a, 0x6f, 0x62, 0x44, 0x65, 0x6c, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x4a, 0x6f, 0x62,
	0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x4a, 0x6f, 0x62, 0x44, 0x65,
	0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x1b, 0x5a, 0x19, 0x66, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_sys_v1_sys_job_proto_rawDescOnce sync.Once
	file_rpc_sys_v1_sys_job_proto_rawDescData = file_rpc_sys_v1_sys_job_proto_rawDesc
)

func file_rpc_sys_v1_sys_job_proto_rawDescGZIP() []byte {
	file_rpc_sys_v1_sys_job_proto_rawDescOnce.Do(func() {
		file_rpc_sys_v1_sys_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_sys_v1_sys_job_proto_rawDescData)
	})
	return file_rpc_sys_v1_sys_job_proto_rawDescData
}

var file_rpc_sys_v1_sys_job_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rpc_sys_v1_sys_job_proto_goTypes = []interface{}{
	(*SysJobInfo)(nil),         // 0: api.rpc_sys.v1.SysJobInfo
	(*SysJobListReq)(nil),      // 1: api.rpc_sys.v1.SysJobListReq
	(*SysJobListReply)(nil),    // 2: api.rpc_sys.v1.SysJobListReply
	(*SysJobInfoReq)(nil),      // 3: api.rpc_sys.v1.SysJobInfoReq
	(*SysJobInfoReply)(nil),    // 4: api.rpc_sys.v1.SysJobInfoReply
	(*SysJobStoreReq)(nil),     // 5: api.rpc_sys.v1.SysJobStoreReq
	(*SysJobStoreReply)(nil),   // 6: api.rpc_sys.v1.SysJobStoreReply
	(*SysJobDelReq)(nil),       // 7: api.rpc_sys.v1.SysJobDelReq
	(*SysJobDelReply)(nil),     // 8: api.rpc_sys.v1.SysJobDelReply
	(*common.SearchParam)(nil), // 9: common.SearchParam
	(*common.Paginator)(nil),   // 10: common.Paginator
}
var file_rpc_sys_v1_sys_job_proto_depIdxs = []int32{
	9,  // 0: api.rpc_sys.v1.SysJobListReq.search:type_name -> common.SearchParam
	0,  // 1: api.rpc_sys.v1.SysJobListReply.list:type_name -> api.rpc_sys.v1.SysJobInfo
	10, // 2: api.rpc_sys.v1.SysJobListReply.paginator:type_name -> common.Paginator
	0,  // 3: api.rpc_sys.v1.SysJobInfoReply.info:type_name -> api.rpc_sys.v1.SysJobInfo
	1,  // 4: api.rpc_sys.v1.Job.SysJobList:input_type -> api.rpc_sys.v1.SysJobListReq
	3,  // 5: api.rpc_sys.v1.Job.SysJobInfo:input_type -> api.rpc_sys.v1.SysJobInfoReq
	5,  // 6: api.rpc_sys.v1.Job.SysJobStore:input_type -> api.rpc_sys.v1.SysJobStoreReq
	7,  // 7: api.rpc_sys.v1.Job.SysJobDel:input_type -> api.rpc_sys.v1.SysJobDelReq
	2,  // 8: api.rpc_sys.v1.Job.SysJobList:output_type -> api.rpc_sys.v1.SysJobListReply
	4,  // 9: api.rpc_sys.v1.Job.SysJobInfo:output_type -> api.rpc_sys.v1.SysJobInfoReply
	6,  // 10: api.rpc_sys.v1.Job.SysJobStore:output_type -> api.rpc_sys.v1.SysJobStoreReply
	8,  // 11: api.rpc_sys.v1.Job.SysJobDel:output_type -> api.rpc_sys.v1.SysJobDelReply
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_rpc_sys_v1_sys_job_proto_init() }
func file_rpc_sys_v1_sys_job_proto_init() {
	if File_rpc_sys_v1_sys_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_sys_v1_sys_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysJobInfo); i {
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
		file_rpc_sys_v1_sys_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysJobListReq); i {
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
		file_rpc_sys_v1_sys_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysJobListReply); i {
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
		file_rpc_sys_v1_sys_job_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysJobInfoReq); i {
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
		file_rpc_sys_v1_sys_job_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysJobInfoReply); i {
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
		file_rpc_sys_v1_sys_job_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysJobStoreReq); i {
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
		file_rpc_sys_v1_sys_job_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysJobStoreReply); i {
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
		file_rpc_sys_v1_sys_job_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysJobDelReq); i {
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
		file_rpc_sys_v1_sys_job_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysJobDelReply); i {
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
			RawDescriptor: file_rpc_sys_v1_sys_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_sys_v1_sys_job_proto_goTypes,
		DependencyIndexes: file_rpc_sys_v1_sys_job_proto_depIdxs,
		MessageInfos:      file_rpc_sys_v1_sys_job_proto_msgTypes,
	}.Build()
	File_rpc_sys_v1_sys_job_proto = out.File
	file_rpc_sys_v1_sys_job_proto_rawDesc = nil
	file_rpc_sys_v1_sys_job_proto_goTypes = nil
	file_rpc_sys_v1_sys_job_proto_depIdxs = nil
}
