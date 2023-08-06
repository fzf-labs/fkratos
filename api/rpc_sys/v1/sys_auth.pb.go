// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: rpc_sys/v1/sys_auth.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// 请求-验证码
type SysAuthLoginCaptchaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SysAuthLoginCaptchaReq) Reset() {
	*x = SysAuthLoginCaptchaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysAuthLoginCaptchaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysAuthLoginCaptchaReq) ProtoMessage() {}

func (x *SysAuthLoginCaptchaReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysAuthLoginCaptchaReq.ProtoReflect.Descriptor instead.
func (*SysAuthLoginCaptchaReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_auth_proto_rawDescGZIP(), []int{0}
}

// 响应-验证码
type SysAuthLoginCaptchaReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaptchaId  string `protobuf:"bytes,1,opt,name=captchaId,proto3" json:"captchaId,omitempty"`   //验证码ID
	CaptchaImg string `protobuf:"bytes,2,opt,name=captchaImg,proto3" json:"captchaImg,omitempty"` //验证码图片
}

func (x *SysAuthLoginCaptchaReply) Reset() {
	*x = SysAuthLoginCaptchaReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysAuthLoginCaptchaReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysAuthLoginCaptchaReply) ProtoMessage() {}

func (x *SysAuthLoginCaptchaReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysAuthLoginCaptchaReply.ProtoReflect.Descriptor instead.
func (*SysAuthLoginCaptchaReply) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_auth_proto_rawDescGZIP(), []int{1}
}

func (x *SysAuthLoginCaptchaReply) GetCaptchaId() string {
	if x != nil {
		return x.CaptchaId
	}
	return ""
}

func (x *SysAuthLoginCaptchaReply) GetCaptchaImg() string {
	if x != nil {
		return x.CaptchaImg
	}
	return ""
}

// 请求-登录
type SysAuthLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaptchaId  string `protobuf:"bytes,1,opt,name=captchaId,proto3" json:"captchaId,omitempty"`   //验证码id
	VerifyCode string `protobuf:"bytes,2,opt,name=verifyCode,proto3" json:"verifyCode,omitempty"` //验证码
	Username   string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`     //账号
	Password   string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`     //密码
}

func (x *SysAuthLoginReq) Reset() {
	*x = SysAuthLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysAuthLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysAuthLoginReq) ProtoMessage() {}

func (x *SysAuthLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysAuthLoginReq.ProtoReflect.Descriptor instead.
func (*SysAuthLoginReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_auth_proto_rawDescGZIP(), []int{2}
}

func (x *SysAuthLoginReq) GetCaptchaId() string {
	if x != nil {
		return x.CaptchaId
	}
	return ""
}

func (x *SysAuthLoginReq) GetVerifyCode() string {
	if x != nil {
		return x.VerifyCode
	}
	return ""
}

func (x *SysAuthLoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SysAuthLoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// 响应-登录
type SysAuthLoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`          //token
	ExpiredAt int64  `protobuf:"varint,2,opt,name=expiredAt,proto3" json:"expiredAt,omitempty"` //过期时间
	RefreshAt int64  `protobuf:"varint,3,opt,name=refreshAt,proto3" json:"refreshAt,omitempty"` //刷新时间
}

func (x *SysAuthLoginReply) Reset() {
	*x = SysAuthLoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysAuthLoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysAuthLoginReply) ProtoMessage() {}

func (x *SysAuthLoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysAuthLoginReply.ProtoReflect.Descriptor instead.
func (*SysAuthLoginReply) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_auth_proto_rawDescGZIP(), []int{3}
}

func (x *SysAuthLoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SysAuthLoginReply) GetExpiredAt() int64 {
	if x != nil {
		return x.ExpiredAt
	}
	return 0
}

func (x *SysAuthLoginReply) GetRefreshAt() int64 {
	if x != nil {
		return x.RefreshAt
	}
	return 0
}

// 请求-退出
type SysAuthLogoutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminId string `protobuf:"bytes,1,opt,name=adminId,proto3" json:"adminId,omitempty"` //ID
}

func (x *SysAuthLogoutReq) Reset() {
	*x = SysAuthLogoutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysAuthLogoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysAuthLogoutReq) ProtoMessage() {}

func (x *SysAuthLogoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysAuthLogoutReq.ProtoReflect.Descriptor instead.
func (*SysAuthLogoutReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_auth_proto_rawDescGZIP(), []int{4}
}

func (x *SysAuthLogoutReq) GetAdminId() string {
	if x != nil {
		return x.AdminId
	}
	return ""
}

// 响应-退出
type SysAuthLogoutReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SysAuthLogoutReply) Reset() {
	*x = SysAuthLogoutReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysAuthLogoutReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysAuthLogoutReply) ProtoMessage() {}

func (x *SysAuthLogoutReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysAuthLogoutReply.ProtoReflect.Descriptor instead.
func (*SysAuthLogoutReply) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_auth_proto_rawDescGZIP(), []int{5}
}

// 请求-Auth-Token校验
type SysAuthJwtTokenCheckReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` //token
}

func (x *SysAuthJwtTokenCheckReq) Reset() {
	*x = SysAuthJwtTokenCheckReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysAuthJwtTokenCheckReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysAuthJwtTokenCheckReq) ProtoMessage() {}

func (x *SysAuthJwtTokenCheckReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysAuthJwtTokenCheckReq.ProtoReflect.Descriptor instead.
func (*SysAuthJwtTokenCheckReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_auth_proto_rawDescGZIP(), []int{6}
}

func (x *SysAuthJwtTokenCheckReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// 响应-Auth-Token校验
type SysAuthJwtTokenCheckReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminId string `protobuf:"bytes,1,opt,name=adminId,proto3" json:"adminId,omitempty"` //管理员Id
}

func (x *SysAuthJwtTokenCheckReply) Reset() {
	*x = SysAuthJwtTokenCheckReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_v1_sys_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysAuthJwtTokenCheckReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysAuthJwtTokenCheckReply) ProtoMessage() {}

func (x *SysAuthJwtTokenCheckReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_v1_sys_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysAuthJwtTokenCheckReply.ProtoReflect.Descriptor instead.
func (*SysAuthJwtTokenCheckReply) Descriptor() ([]byte, []int) {
	return file_rpc_sys_v1_sys_auth_proto_rawDescGZIP(), []int{7}
}

func (x *SysAuthJwtTokenCheckReply) GetAdminId() string {
	if x != nil {
		return x.AdminId
	}
	return ""
}

var File_rpc_sys_v1_sys_auth_proto protoreflect.FileDescriptor

var file_rpc_sys_v1_sys_auth_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x53, 0x79, 0x73, 0x41, 0x75, 0x74, 0x68, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71, 0x22, 0x58,
	0x0a, 0x18, 0x53, 0x79, 0x73, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x70, 0x74,
	0x63, 0x68, 0x61, 0x49, 0x6d, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x49, 0x6d, 0x67, 0x22, 0x87, 0x01, 0x0a, 0x0f, 0x53, 0x79, 0x73,
	0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x65, 0x0a, 0x11, 0x53, 0x79, 0x73, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x74, 0x22, 0x36, 0x0a, 0x10, 0x53, 0x79, 0x73,
	0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a,
	0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49,
	0x64, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x79, 0x73, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2f, 0x0a, 0x17, 0x53, 0x79, 0x73, 0x41, 0x75,
	0x74, 0x68, 0x4a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x35, 0x0a, 0x19, 0x53, 0x79, 0x73, 0x41,
	0x75, 0x74, 0x68, 0x4a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x32,
	0x88, 0x03, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x67, 0x0a, 0x13, 0x53, 0x79, 0x73, 0x41,
	0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12,
	0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x79, 0x73, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x75, 0x74, 0x68,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x52, 0x0a, 0x0c, 0x53, 0x79, 0x73, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x55, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x41, 0x75, 0x74, 0x68,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x75, 0x74, 0x68, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x75, 0x74,
	0x68, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x6c, 0x0a, 0x14,
	0x53, 0x79, 0x73, 0x41, 0x75, 0x74, 0x68, 0x4a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73,
	0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x75, 0x74, 0x68, 0x4a, 0x77, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x79, 0x73, 0x41, 0x75, 0x74, 0x68, 0x4a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x66, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x79,
	0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_sys_v1_sys_auth_proto_rawDescOnce sync.Once
	file_rpc_sys_v1_sys_auth_proto_rawDescData = file_rpc_sys_v1_sys_auth_proto_rawDesc
)

func file_rpc_sys_v1_sys_auth_proto_rawDescGZIP() []byte {
	file_rpc_sys_v1_sys_auth_proto_rawDescOnce.Do(func() {
		file_rpc_sys_v1_sys_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_sys_v1_sys_auth_proto_rawDescData)
	})
	return file_rpc_sys_v1_sys_auth_proto_rawDescData
}

var file_rpc_sys_v1_sys_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_rpc_sys_v1_sys_auth_proto_goTypes = []interface{}{
	(*SysAuthLoginCaptchaReq)(nil),    // 0: api.rpc_sys.v1.SysAuthLoginCaptchaReq
	(*SysAuthLoginCaptchaReply)(nil),  // 1: api.rpc_sys.v1.SysAuthLoginCaptchaReply
	(*SysAuthLoginReq)(nil),           // 2: api.rpc_sys.v1.SysAuthLoginReq
	(*SysAuthLoginReply)(nil),         // 3: api.rpc_sys.v1.SysAuthLoginReply
	(*SysAuthLogoutReq)(nil),          // 4: api.rpc_sys.v1.SysAuthLogoutReq
	(*SysAuthLogoutReply)(nil),        // 5: api.rpc_sys.v1.SysAuthLogoutReply
	(*SysAuthJwtTokenCheckReq)(nil),   // 6: api.rpc_sys.v1.SysAuthJwtTokenCheckReq
	(*SysAuthJwtTokenCheckReply)(nil), // 7: api.rpc_sys.v1.SysAuthJwtTokenCheckReply
}
var file_rpc_sys_v1_sys_auth_proto_depIdxs = []int32{
	0, // 0: api.rpc_sys.v1.Auth.SysAuthLoginCaptcha:input_type -> api.rpc_sys.v1.SysAuthLoginCaptchaReq
	2, // 1: api.rpc_sys.v1.Auth.SysAuthLogin:input_type -> api.rpc_sys.v1.SysAuthLoginReq
	4, // 2: api.rpc_sys.v1.Auth.SysAuthLogout:input_type -> api.rpc_sys.v1.SysAuthLogoutReq
	6, // 3: api.rpc_sys.v1.Auth.SysAuthJwtTokenCheck:input_type -> api.rpc_sys.v1.SysAuthJwtTokenCheckReq
	1, // 4: api.rpc_sys.v1.Auth.SysAuthLoginCaptcha:output_type -> api.rpc_sys.v1.SysAuthLoginCaptchaReply
	3, // 5: api.rpc_sys.v1.Auth.SysAuthLogin:output_type -> api.rpc_sys.v1.SysAuthLoginReply
	5, // 6: api.rpc_sys.v1.Auth.SysAuthLogout:output_type -> api.rpc_sys.v1.SysAuthLogoutReply
	7, // 7: api.rpc_sys.v1.Auth.SysAuthJwtTokenCheck:output_type -> api.rpc_sys.v1.SysAuthJwtTokenCheckReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_sys_v1_sys_auth_proto_init() }
func file_rpc_sys_v1_sys_auth_proto_init() {
	if File_rpc_sys_v1_sys_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_sys_v1_sys_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysAuthLoginCaptchaReq); i {
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
		file_rpc_sys_v1_sys_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysAuthLoginCaptchaReply); i {
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
		file_rpc_sys_v1_sys_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysAuthLoginReq); i {
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
		file_rpc_sys_v1_sys_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysAuthLoginReply); i {
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
		file_rpc_sys_v1_sys_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysAuthLogoutReq); i {
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
		file_rpc_sys_v1_sys_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysAuthLogoutReply); i {
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
		file_rpc_sys_v1_sys_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysAuthJwtTokenCheckReq); i {
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
		file_rpc_sys_v1_sys_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysAuthJwtTokenCheckReply); i {
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
			RawDescriptor: file_rpc_sys_v1_sys_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_sys_v1_sys_auth_proto_goTypes,
		DependencyIndexes: file_rpc_sys_v1_sys_auth_proto_depIdxs,
		MessageInfos:      file_rpc_sys_v1_sys_auth_proto_msgTypes,
	}.Build()
	File_rpc_sys_v1_sys_auth_proto = out.File
	file_rpc_sys_v1_sys_auth_proto_rawDesc = nil
	file_rpc_sys_v1_sys_auth_proto_goTypes = nil
	file_rpc_sys_v1_sys_auth_proto_depIdxs = nil
}
