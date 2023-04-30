// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: logger.proto

package conf

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

// 日志
type Logger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string          `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Zap     *Logger_Zap     `protobuf:"bytes,2,opt,name=zap,proto3" json:"zap,omitempty"`
	Logrus  *Logger_Logrus  `protobuf:"bytes,3,opt,name=logrus,proto3" json:"logrus,omitempty"`
	Aliyun  *Logger_Aliyun  `protobuf:"bytes,4,opt,name=aliyun,proto3" json:"aliyun,omitempty"`
	Tencent *Logger_Tencent `protobuf:"bytes,5,opt,name=tencent,proto3" json:"tencent,omitempty"`
}

func (x *Logger) Reset() {
	*x = Logger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger) ProtoMessage() {}

func (x *Logger) ProtoReflect() protoreflect.Message {
	mi := &file_logger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger.ProtoReflect.Descriptor instead.
func (*Logger) Descriptor() ([]byte, []int) {
	return file_logger_proto_rawDescGZIP(), []int{0}
}

func (x *Logger) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Logger) GetZap() *Logger_Zap {
	if x != nil {
		return x.Zap
	}
	return nil
}

func (x *Logger) GetLogrus() *Logger_Logrus {
	if x != nil {
		return x.Logrus
	}
	return nil
}

func (x *Logger) GetAliyun() *Logger_Aliyun {
	if x != nil {
		return x.Aliyun
	}
	return nil
}

func (x *Logger) GetTencent() *Logger_Tencent {
	if x != nil {
		return x.Tencent
	}
	return nil
}

// Zap
type Logger_Zap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename   string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`      //
	Level      string `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`            //
	MaxSize    int32  `protobuf:"varint,3,opt,name=maxSize,proto3" json:"maxSize,omitempty"`       //
	MaxAge     int32  `protobuf:"varint,4,opt,name=maxAge,proto3" json:"maxAge,omitempty"`         //
	MaxBackups int32  `protobuf:"varint,5,opt,name=maxBackups,proto3" json:"maxBackups,omitempty"` //
}

func (x *Logger_Zap) Reset() {
	*x = Logger_Zap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logger_Zap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger_Zap) ProtoMessage() {}

func (x *Logger_Zap) ProtoReflect() protoreflect.Message {
	mi := &file_logger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger_Zap.ProtoReflect.Descriptor instead.
func (*Logger_Zap) Descriptor() ([]byte, []int) {
	return file_logger_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Logger_Zap) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Logger_Zap) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Logger_Zap) GetMaxSize() int32 {
	if x != nil {
		return x.MaxSize
	}
	return 0
}

func (x *Logger_Zap) GetMaxAge() int32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

func (x *Logger_Zap) GetMaxBackups() int32 {
	if x != nil {
		return x.MaxBackups
	}
	return 0
}

// logrus
type Logger_Logrus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level            string `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`                        // 日志等级
	Formatter        string `protobuf:"bytes,2,opt,name=formatter,proto3" json:"formatter,omitempty"`                // 输出格式：text, json.
	TimestampFormat  string `protobuf:"bytes,3,opt,name=timestampFormat,proto3" json:"timestampFormat,omitempty"`    // 定义时间戳格式，例如："2006-01-02 15:04:05"
	DisableColors    bool   `protobuf:"varint,4,opt,name=disableColors,proto3" json:"disableColors,omitempty"`       // 不需要彩色日志
	DisableTimestamp bool   `protobuf:"varint,5,opt,name=disableTimestamp,proto3" json:"disableTimestamp,omitempty"` // 不需要时间戳
}

func (x *Logger_Logrus) Reset() {
	*x = Logger_Logrus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logger_Logrus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger_Logrus) ProtoMessage() {}

func (x *Logger_Logrus) ProtoReflect() protoreflect.Message {
	mi := &file_logger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger_Logrus.ProtoReflect.Descriptor instead.
func (*Logger_Logrus) Descriptor() ([]byte, []int) {
	return file_logger_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Logger_Logrus) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Logger_Logrus) GetFormatter() string {
	if x != nil {
		return x.Formatter
	}
	return ""
}

func (x *Logger_Logrus) GetTimestampFormat() string {
	if x != nil {
		return x.TimestampFormat
	}
	return ""
}

func (x *Logger_Logrus) GetDisableColors() bool {
	if x != nil {
		return x.DisableColors
	}
	return false
}

func (x *Logger_Logrus) GetDisableTimestamp() bool {
	if x != nil {
		return x.DisableTimestamp
	}
	return false
}

// 阿里云
type Logger_Aliyun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint     string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`         // 公网接入地址
	Project      string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`           //
	AccessKey    string `protobuf:"bytes,3,opt,name=accessKey,proto3" json:"accessKey,omitempty"`       // 访问密钥ID
	AccessSecret string `protobuf:"bytes,4,opt,name=accessSecret,proto3" json:"accessSecret,omitempty"` // 访问密钥
}

func (x *Logger_Aliyun) Reset() {
	*x = Logger_Aliyun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logger_Aliyun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger_Aliyun) ProtoMessage() {}

func (x *Logger_Aliyun) ProtoReflect() protoreflect.Message {
	mi := &file_logger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger_Aliyun.ProtoReflect.Descriptor instead.
func (*Logger_Aliyun) Descriptor() ([]byte, []int) {
	return file_logger_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Logger_Aliyun) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Logger_Aliyun) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Logger_Aliyun) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *Logger_Aliyun) GetAccessSecret() string {
	if x != nil {
		return x.AccessSecret
	}
	return ""
}

// 腾讯
type Logger_Tencent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint     string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`         // 公网接入地址
	TopicId      string `protobuf:"bytes,2,opt,name=topicId,proto3" json:"topicId,omitempty"`           //
	AccessKey    string `protobuf:"bytes,3,opt,name=accessKey,proto3" json:"accessKey,omitempty"`       // 访问密钥ID
	AccessSecret string `protobuf:"bytes,4,opt,name=accessSecret,proto3" json:"accessSecret,omitempty"` // 访问密钥
}

func (x *Logger_Tencent) Reset() {
	*x = Logger_Tencent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logger_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logger_Tencent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger_Tencent) ProtoMessage() {}

func (x *Logger_Tencent) ProtoReflect() protoreflect.Message {
	mi := &file_logger_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger_Tencent.ProtoReflect.Descriptor instead.
func (*Logger_Tencent) Descriptor() ([]byte, []int) {
	return file_logger_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Logger_Tencent) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Logger_Tencent) GetTopicId() string {
	if x != nil {
		return x.TopicId
	}
	return ""
}

func (x *Logger_Tencent) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *Logger_Tencent) GetAccessSecret() string {
	if x != nil {
		return x.AccessSecret
	}
	return ""
}

var File_logger_proto protoreflect.FileDescriptor

var file_logger_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x22, 0xc0,
	0x06, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a,
	0x03, 0x7a, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6f, 0x6f,
	0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x5a, 0x61, 0x70, 0x52, 0x03, 0x7a, 0x61, 0x70, 0x12, 0x35, 0x0a, 0x06, 0x6c,
	0x6f, 0x67, 0x72, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x6f,
	0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x72, 0x75, 0x73, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x72,
	0x75, 0x73, 0x12, 0x35, 0x0a, 0x06, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x69, 0x79, 0x75,
	0x6e, 0x52, 0x06, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x12, 0x38, 0x0a, 0x07, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x6f, 0x6f,
	0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x54, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x6e, 0x74, 0x1a, 0x89, 0x01, 0x0a, 0x03, 0x5a, 0x61, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x6d, 0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x1a,
	0xb8, 0x01, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x72, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x12, 0x28,
	0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x73, 0x12, 0x2a,
	0x0a, 0x10, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x1a, 0x80, 0x01, 0x0a, 0x06, 0x41,
	0x6c, 0x69, 0x79, 0x75, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x1a, 0x81, 0x01,
	0x0a, 0x07, 0x54, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x42, 0x1d, 0x5a, 0x1b, 0x66, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x6f, 0x6f,
	0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logger_proto_rawDescOnce sync.Once
	file_logger_proto_rawDescData = file_logger_proto_rawDesc
)

func file_logger_proto_rawDescGZIP() []byte {
	file_logger_proto_rawDescOnce.Do(func() {
		file_logger_proto_rawDescData = protoimpl.X.CompressGZIP(file_logger_proto_rawDescData)
	})
	return file_logger_proto_rawDescData
}

var file_logger_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_logger_proto_goTypes = []interface{}{
	(*Logger)(nil),         // 0: bootstrap.conf.Logger
	(*Logger_Zap)(nil),     // 1: bootstrap.conf.Logger.Zap
	(*Logger_Logrus)(nil),  // 2: bootstrap.conf.Logger.Logrus
	(*Logger_Aliyun)(nil),  // 3: bootstrap.conf.Logger.Aliyun
	(*Logger_Tencent)(nil), // 4: bootstrap.conf.Logger.Tencent
}
var file_logger_proto_depIdxs = []int32{
	1, // 0: bootstrap.conf.Logger.zap:type_name -> bootstrap.conf.Logger.Zap
	2, // 1: bootstrap.conf.Logger.logrus:type_name -> bootstrap.conf.Logger.Logrus
	3, // 2: bootstrap.conf.Logger.aliyun:type_name -> bootstrap.conf.Logger.Aliyun
	4, // 3: bootstrap.conf.Logger.tencent:type_name -> bootstrap.conf.Logger.Tencent
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_logger_proto_init() }
func file_logger_proto_init() {
	if File_logger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logger); i {
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
		file_logger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logger_Zap); i {
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
		file_logger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logger_Logrus); i {
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
		file_logger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logger_Aliyun); i {
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
		file_logger_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logger_Tencent); i {
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
			RawDescriptor: file_logger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_logger_proto_goTypes,
		DependencyIndexes: file_logger_proto_depIdxs,
		MessageInfos:      file_logger_proto_msgTypes,
	}.Build()
	File_logger_proto = out.File
	file_logger_proto_rawDesc = nil
	file_logger_proto_goTypes = nil
	file_logger_proto_depIdxs = nil
}
