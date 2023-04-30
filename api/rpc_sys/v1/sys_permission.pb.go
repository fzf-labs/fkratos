// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: rpc_sys/v1/sys_permission.proto

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

type SysPermMenu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                 // ID
	Pid       string         `protobuf:"bytes,2,opt,name=pid,proto3" json:"pid,omitempty"`               // 上级菜单
	Type      string         `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`             // 类型:menu_dir=菜单目录,menu=菜单项,button=页面按钮
	Title     string         `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`           // 标题
	Name      string         `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`             // 规则名称
	Path      string         `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`             // 路由路径
	Icon      string         `protobuf:"bytes,7,opt,name=icon,proto3" json:"icon,omitempty"`             // 图标
	MenuType  string         `protobuf:"bytes,8,opt,name=menuType,proto3" json:"menuType,omitempty"`     // 菜单类型:tab=选项卡,link=链接,iframe=Iframe
	Url       string         `protobuf:"bytes,9,opt,name=url,proto3" json:"url,omitempty"`               // Url
	Component string         `protobuf:"bytes,10,opt,name=component,proto3" json:"component,omitempty"`  // 组件路径
	Keepalive int32          `protobuf:"varint,11,opt,name=keepalive,proto3" json:"keepalive,omitempty"` // 缓存:0=关闭,1=开启
	Extend    string         `protobuf:"bytes,12,opt,name=extend,proto3" json:"extend,omitempty"`        // 扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单
	Remark    string         `protobuf:"bytes,13,opt,name=remark,proto3" json:"remark,omitempty"`        // 备注
	Sort      int32          `protobuf:"varint,14,opt,name=sort,proto3" json:"sort,omitempty"`           // 权重(排序)
	Status    int32          `protobuf:"varint,15,opt,name=status,proto3" json:"status,omitempty"`       // 状态:0=禁用,1=启用
	CreatedAt string         `protobuf:"bytes,16,opt,name=createdAt,proto3" json:"createdAt,omitempty"`  // 创建时间
	UpdatedAt string         `protobuf:"bytes,17,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`  // 更新时间
	Children  []*SysPermMenu `protobuf:"bytes,18,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *SysPermMenu) Reset() {
	*x = SysPermMenu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysPermMenu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysPermMenu) ProtoMessage() {}

func (x *SysPermMenu) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysPermMenu.ProtoReflect.Descriptor instead.
func (*SysPermMenu) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_permission_proto_rawDescGZIP(), []int{0}
}

func (x *SysPermMenu) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SysPermMenu) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

func (x *SysPermMenu) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SysPermMenu) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SysPermMenu) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SysPermMenu) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *SysPermMenu) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *SysPermMenu) GetMenuType() string {
	if x != nil {
		return x.MenuType
	}
	return ""
}

func (x *SysPermMenu) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SysPermMenu) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *SysPermMenu) GetKeepalive() int32 {
	if x != nil {
		return x.Keepalive
	}
	return 0
}

func (x *SysPermMenu) GetExtend() string {
	if x != nil {
		return x.Extend
	}
	return ""
}

func (x *SysPermMenu) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *SysPermMenu) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *SysPermMenu) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SysPermMenu) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SysPermMenu) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *SysPermMenu) GetChildren() []*SysPermMenu {
	if x != nil {
		return x.Children
	}
	return nil
}

type SysPermissionListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SysPermissionListReq) Reset() {
	*x = SysPermissionListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysPermissionListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysPermissionListReq) ProtoMessage() {}

func (x *SysPermissionListReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysPermissionListReq.ProtoReflect.Descriptor instead.
func (*SysPermissionListReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_permission_proto_rawDescGZIP(), []int{1}
}

type SysPermissionListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*SysPermMenu `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *SysPermissionListResp) Reset() {
	*x = SysPermissionListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysPermissionListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysPermissionListResp) ProtoMessage() {}

func (x *SysPermissionListResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysPermissionListResp.ProtoReflect.Descriptor instead.
func (*SysPermissionListResp) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_permission_proto_rawDescGZIP(), []int{2}
}

func (x *SysPermissionListResp) GetList() []*SysPermMenu {
	if x != nil {
		return x.List
	}
	return nil
}

type SysPermissionInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SysPermissionInfoReq) Reset() {
	*x = SysPermissionInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysPermissionInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysPermissionInfoReq) ProtoMessage() {}

func (x *SysPermissionInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysPermissionInfoReq.ProtoReflect.Descriptor instead.
func (*SysPermissionInfoReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_permission_proto_rawDescGZIP(), []int{3}
}

func (x *SysPermissionInfoReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SysPermissionInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *SysPermMenu `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SysPermissionInfoResp) Reset() {
	*x = SysPermissionInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysPermissionInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysPermissionInfoResp) ProtoMessage() {}

func (x *SysPermissionInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysPermissionInfoResp.ProtoReflect.Descriptor instead.
func (*SysPermissionInfoResp) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_permission_proto_rawDescGZIP(), []int{4}
}

func (x *SysPermissionInfoResp) GetInfo() *SysPermMenu {
	if x != nil {
		return x.Info
	}
	return nil
}

type SysPermissionStoreReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                 // ID
	Pid       string `protobuf:"bytes,2,opt,name=pid,proto3" json:"pid,omitempty"`               // 上级菜单
	Type      string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`             // 类型:menu_dir=菜单目录,menu=菜单项,button=页面按钮
	Title     string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`           // 标题
	Name      string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`             // 规则名称
	Path      string `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`             // 路由路径
	Icon      string `protobuf:"bytes,7,opt,name=icon,proto3" json:"icon,omitempty"`             // 图标
	MenuType  string `protobuf:"bytes,8,opt,name=menuType,proto3" json:"menuType,omitempty"`     // 菜单类型:tab=选项卡,link=链接,iframe=Iframe
	Url       string `protobuf:"bytes,9,opt,name=url,proto3" json:"url,omitempty"`               // Url
	Component string `protobuf:"bytes,10,opt,name=component,proto3" json:"component,omitempty"`  // 组件路径
	Keepalive int32  `protobuf:"varint,11,opt,name=keepalive,proto3" json:"keepalive,omitempty"` // 缓存:0=关闭,1=开启
	Extend    string `protobuf:"bytes,12,opt,name=extend,proto3" json:"extend,omitempty"`        // 扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单
	Remark    string `protobuf:"bytes,13,opt,name=remark,proto3" json:"remark,omitempty"`        // 备注
	Sort      int32  `protobuf:"varint,14,opt,name=sort,proto3" json:"sort,omitempty"`           // 权重(排序)
	Status    int32  `protobuf:"varint,15,opt,name=status,proto3" json:"status,omitempty"`       // 状态:0=禁用,1=启用
}

func (x *SysPermissionStoreReq) Reset() {
	*x = SysPermissionStoreReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysPermissionStoreReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysPermissionStoreReq) ProtoMessage() {}

func (x *SysPermissionStoreReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysPermissionStoreReq.ProtoReflect.Descriptor instead.
func (*SysPermissionStoreReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_permission_proto_rawDescGZIP(), []int{5}
}

func (x *SysPermissionStoreReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SysPermissionStoreReq) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

func (x *SysPermissionStoreReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SysPermissionStoreReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SysPermissionStoreReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SysPermissionStoreReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *SysPermissionStoreReq) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *SysPermissionStoreReq) GetMenuType() string {
	if x != nil {
		return x.MenuType
	}
	return ""
}

func (x *SysPermissionStoreReq) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SysPermissionStoreReq) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *SysPermissionStoreReq) GetKeepalive() int32 {
	if x != nil {
		return x.Keepalive
	}
	return 0
}

func (x *SysPermissionStoreReq) GetExtend() string {
	if x != nil {
		return x.Extend
	}
	return ""
}

func (x *SysPermissionStoreReq) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *SysPermissionStoreReq) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *SysPermissionStoreReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type SysPermissionStoreResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SysPermissionStoreResp) Reset() {
	*x = SysPermissionStoreResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysPermissionStoreResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysPermissionStoreResp) ProtoMessage() {}

func (x *SysPermissionStoreResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysPermissionStoreResp.ProtoReflect.Descriptor instead.
func (*SysPermissionStoreResp) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_permission_proto_rawDescGZIP(), []int{6}
}

type SysPermissionDelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *SysPermissionDelReq) Reset() {
	*x = SysPermissionDelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysPermissionDelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysPermissionDelReq) ProtoMessage() {}

func (x *SysPermissionDelReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysPermissionDelReq.ProtoReflect.Descriptor instead.
func (*SysPermissionDelReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_permission_proto_rawDescGZIP(), []int{7}
}

func (x *SysPermissionDelReq) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type SysPermissionDelResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SysPermissionDelResp) Reset() {
	*x = SysPermissionDelResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysPermissionDelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysPermissionDelResp) ProtoMessage() {}

func (x *SysPermissionDelResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysPermissionDelResp.ProtoReflect.Descriptor instead.
func (*SysPermissionDelResp) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_permission_proto_rawDescGZIP(), []int{8}
}

type SysPermissionStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status int32  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"` // 状态:0=禁用,1=启用
}

func (x *SysPermissionStatusReq) Reset() {
	*x = SysPermissionStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysPermissionStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysPermissionStatusReq) ProtoMessage() {}

func (x *SysPermissionStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysPermissionStatusReq.ProtoReflect.Descriptor instead.
func (*SysPermissionStatusReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_permission_proto_rawDescGZIP(), []int{9}
}

func (x *SysPermissionStatusReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SysPermissionStatusReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type SysPermissionStatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SysPermissionStatusResp) Reset() {
	*x = SysPermissionStatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysPermissionStatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysPermissionStatusResp) ProtoMessage() {}

func (x *SysPermissionStatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_permission_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysPermissionStatusResp.ProtoReflect.Descriptor instead.
func (*SysPermissionStatusResp) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_permission_proto_rawDescGZIP(), []int{10}
}

var File_rpc_sys_v1_sys_permission_proto protoreflect.FileDescriptor

var file_rpc_sys_v1_sys_permission_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73,
	0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76,
	0x31, 0x22, 0xd0, 0x03, 0x0a, 0x0b, 0x53, 0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x4d, 0x65, 0x6e,
	0x75, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x70, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6e,
	0x75, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6e,
	0x75, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69,
	0x76, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c,
	0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c,
	0x64, 0x72, 0x65, 0x6e, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x22, 0x48, 0x0a, 0x15,
	0x53, 0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x4d, 0x65, 0x6e, 0x75,
	0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x53, 0x79, 0x73, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x48,
	0x0a, 0x15, 0x53, 0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f,
	0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x4d, 0x65,
	0x6e, 0x75, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xe5, 0x02, 0x0a, 0x15, 0x53, 0x79, 0x73,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x70, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65,
	0x6e, 0x75, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65,
	0x6e, 0x75, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c,
	0x69, 0x76, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x61,
	0x6c, 0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x18, 0x0a, 0x16, 0x53, 0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x27, 0x0a, 0x13, 0x53, 0x79,
	0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x40, 0x0a, 0x16, 0x53,
	0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x19, 0x0a,
	0x17, 0x53, 0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x32, 0xfc, 0x03, 0x0a, 0x0a, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x60, 0x0a, 0x11, 0x53, 0x79, 0x73, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79,
	0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x60, 0x0a, 0x11, 0x53, 0x79, 0x73,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73,
	0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x63, 0x0a, 0x12, 0x53,
	0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x5d, 0x0a, 0x10, 0x53, 0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x6c, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73,
	0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x66, 0x0a, 0x13, 0x53, 0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x27,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x79, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x1b, 0x5a, 0x19, 0x66, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_sys_v1_sys_permission_proto_rawDescOnce sync.Once
	file_rpc_sys_v1_sys_permission_proto_rawDescData = file_rpc_sys_v1_sys_permission_proto_rawDesc
)

func file_rpc_sys_v1_sys_permission_proto_rawDescGZIP() []byte {
	file_rpc_sys_v1_sys_permission_proto_rawDescOnce.Do(func() {
		file_rpc_sys_v1_sys_permission_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_sys_v1_sys_permission_proto_rawDescData)
	})
	return file_rpc_sys_v1_sys_permission_proto_rawDescData
}

var file_rpc_sys_v1_sys_permission_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_rpc_sys_v1_sys_permission_proto_goTypes = []interface{}{
	(*SysPermMenu)(nil),             // 0: api.rpc_sys.v1.SysPermMenu
	(*SysPermissionListReq)(nil),    // 1: api.rpc_sys.v1.SysPermissionListReq
	(*SysPermissionListResp)(nil),   // 2: api.rpc_sys.v1.SysPermissionListResp
	(*SysPermissionInfoReq)(nil),    // 3: api.rpc_sys.v1.SysPermissionInfoReq
	(*SysPermissionInfoResp)(nil),   // 4: api.rpc_sys.v1.SysPermissionInfoResp
	(*SysPermissionStoreReq)(nil),   // 5: api.rpc_sys.v1.SysPermissionStoreReq
	(*SysPermissionStoreResp)(nil),  // 6: api.rpc_sys.v1.SysPermissionStoreResp
	(*SysPermissionDelReq)(nil),     // 7: api.rpc_sys.v1.SysPermissionDelReq
	(*SysPermissionDelResp)(nil),    // 8: api.rpc_sys.v1.SysPermissionDelResp
	(*SysPermissionStatusReq)(nil),  // 9: api.rpc_sys.v1.SysPermissionStatusReq
	(*SysPermissionStatusResp)(nil), // 10: api.rpc_sys.v1.SysPermissionStatusResp
}
var file_rpc_sys_v1_sys_permission_proto_depIdxs = []int32{
	0,  // 0: api.rpc_sys.v1.SysPermMenu.children:type_name -> api.rpc_sys.v1.SysPermMenu
	0,  // 1: api.rpc_sys.v1.SysPermissionListResp.list:type_name -> api.rpc_sys.v1.SysPermMenu
	0,  // 2: api.rpc_sys.v1.SysPermissionInfoResp.info:type_name -> api.rpc_sys.v1.SysPermMenu
	1,  // 3: api.rpc_sys.v1.Permission.SysPermissionList:input_type -> api.rpc_sys.v1.SysPermissionListReq
	3,  // 4: api.rpc_sys.v1.Permission.SysPermissionInfo:input_type -> api.rpc_sys.v1.SysPermissionInfoReq
	5,  // 5: api.rpc_sys.v1.Permission.SysPermissionStore:input_type -> api.rpc_sys.v1.SysPermissionStoreReq
	7,  // 6: api.rpc_sys.v1.Permission.SysPermissionDel:input_type -> api.rpc_sys.v1.SysPermissionDelReq
	9,  // 7: api.rpc_sys.v1.Permission.SysPermissionStatus:input_type -> api.rpc_sys.v1.SysPermissionStatusReq
	2,  // 8: api.rpc_sys.v1.Permission.SysPermissionList:output_type -> api.rpc_sys.v1.SysPermissionListResp
	4,  // 9: api.rpc_sys.v1.Permission.SysPermissionInfo:output_type -> api.rpc_sys.v1.SysPermissionInfoResp
	6,  // 10: api.rpc_sys.v1.Permission.SysPermissionStore:output_type -> api.rpc_sys.v1.SysPermissionStoreResp
	8,  // 11: api.rpc_sys.v1.Permission.SysPermissionDel:output_type -> api.rpc_sys.v1.SysPermissionDelResp
	10, // 12: api.rpc_sys.v1.Permission.SysPermissionStatus:output_type -> api.rpc_sys.v1.SysPermissionStatusResp
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_rpc_sys_v1_sys_permission_proto_init() }
func file_rpc_sys_v1_sys_permission_proto_init() {
	if File_rpc_sys_v1_sys_permission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_sys_v1_sys_permission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysPermMenu); i {
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
		file_rpc_sys_v1_sys_permission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysPermissionListReq); i {
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
		file_rpc_sys_v1_sys_permission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysPermissionListResp); i {
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
		file_rpc_sys_v1_sys_permission_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysPermissionInfoReq); i {
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
		file_rpc_sys_v1_sys_permission_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysPermissionInfoResp); i {
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
		file_rpc_sys_v1_sys_permission_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysPermissionStoreReq); i {
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
		file_rpc_sys_v1_sys_permission_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysPermissionStoreResp); i {
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
		file_rpc_sys_v1_sys_permission_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysPermissionDelReq); i {
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
		file_rpc_sys_v1_sys_permission_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysPermissionDelResp); i {
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
		file_rpc_sys_v1_sys_permission_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysPermissionStatusReq); i {
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
		file_rpc_sys_v1_sys_permission_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysPermissionStatusResp); i {
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
			RawDescriptor: file_rpc_sys_v1_sys_permission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_sys_v1_sys_permission_proto_goTypes,
		DependencyIndexes: file_rpc_sys_v1_sys_permission_proto_depIdxs,
		MessageInfos:      file_rpc_sys_v1_sys_permission_proto_msgTypes,
	}.Build()
	File_rpc_sys_v1_sys_permission_proto = out.File
	file_rpc_sys_v1_sys_permission_proto_rawDesc = nil
	file_rpc_sys_v1_sys_permission_proto_goTypes = nil
	file_rpc_sys_v1_sys_permission_proto_depIdxs = nil
}
