// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: registry.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 注册发现中心
type Registry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       string               `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Consul     *Registry_Consul     `protobuf:"bytes,2,opt,name=consul,proto3" json:"consul,omitempty"`         // Consul
	Etcd       *Registry_Etcd       `protobuf:"bytes,3,opt,name=etcd,proto3" json:"etcd,omitempty"`             // Etcd
	Nacos      *Registry_Nacos      `protobuf:"bytes,4,opt,name=nacos,proto3" json:"nacos,omitempty"`           // Nacos
	Polaris    *Registry_Polaris    `protobuf:"bytes,5,opt,name=polaris,proto3" json:"polaris,omitempty"`       // Polaris
	Kubernetes *Registry_Kubernetes `protobuf:"bytes,6,opt,name=kubernetes,proto3" json:"kubernetes,omitempty"` // Kubernetes
}

func (x *Registry) Reset() {
	*x = Registry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry) ProtoMessage() {}

func (x *Registry) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry.ProtoReflect.Descriptor instead.
func (*Registry) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{0}
}

func (x *Registry) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Registry) GetConsul() *Registry_Consul {
	if x != nil {
		return x.Consul
	}
	return nil
}

func (x *Registry) GetEtcd() *Registry_Etcd {
	if x != nil {
		return x.Etcd
	}
	return nil
}

func (x *Registry) GetNacos() *Registry_Nacos {
	if x != nil {
		return x.Nacos
	}
	return nil
}

func (x *Registry) GetPolaris() *Registry_Polaris {
	if x != nil {
		return x.Polaris
	}
	return nil
}

func (x *Registry) GetKubernetes() *Registry_Kubernetes {
	if x != nil {
		return x.Kubernetes
	}
	return nil
}

// Consul
type Registry_Consul struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scheme      string `protobuf:"bytes,1,opt,name=scheme,proto3" json:"scheme,omitempty"`            // 网络样式
	Address     string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`          // 服务端地址
	HealthCheck bool   `protobuf:"varint,3,opt,name=healthCheck,proto3" json:"healthCheck,omitempty"` // 健康检查
}

func (x *Registry_Consul) Reset() {
	*x = Registry_Consul{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry_Consul) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry_Consul) ProtoMessage() {}

func (x *Registry_Consul) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry_Consul.ProtoReflect.Descriptor instead.
func (*Registry_Consul) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Registry_Consul) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *Registry_Consul) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Registry_Consul) GetHealthCheck() bool {
	if x != nil {
		return x.HealthCheck
	}
	return false
}

// Etcd
type Registry_Etcd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []string `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *Registry_Etcd) Reset() {
	*x = Registry_Etcd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry_Etcd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry_Etcd) ProtoMessage() {}

func (x *Registry_Etcd) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry_Etcd.ProtoReflect.Descriptor instead.
func (*Registry_Etcd) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Registry_Etcd) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

// Nacos
type Registry_Nacos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address              string               `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`                             // 服务端地址
	Port                 uint64               `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`                                  // 服务端端口
	NamespaceId          string               `protobuf:"bytes,3,opt,name=namespaceId,proto3" json:"namespaceId,omitempty"`                     //
	LogLevel             string               `protobuf:"bytes,4,opt,name=logLevel,proto3" json:"logLevel,omitempty"`                           // 日志等级
	CacheDir             string               `protobuf:"bytes,5,opt,name=cacheDir,proto3" json:"cacheDir,omitempty"`                           // 缓存目录
	LogDir               string               `protobuf:"bytes,6,opt,name=logDir,proto3" json:"logDir,omitempty"`                               // 日志目录
	UpdateThreadNum      int32                `protobuf:"varint,7,opt,name=updateThreadNum,proto3" json:"updateThreadNum,omitempty"`            // 更新服务的线程数
	Timeout              *durationpb.Duration `protobuf:"bytes,8,opt,name=timeout,proto3" json:"timeout,omitempty"`                             // http请求超时时间，单位: 毫秒
	BeatInterval         *durationpb.Duration `protobuf:"bytes,9,opt,name=beatInterval,proto3" json:"beatInterval,omitempty"`                   // 心跳间隔时间，单位: 毫秒
	NotLoadCacheAtStart  bool                 `protobuf:"varint,10,opt,name=notLoadCacheAtStart,proto3" json:"notLoadCacheAtStart,omitempty"`   // 在启动时不读取本地缓存数据，true: 不读取，false: 读取
	UpdateCacheWhenEmpty bool                 `protobuf:"varint,11,opt,name=updateCacheWhenEmpty,proto3" json:"updateCacheWhenEmpty,omitempty"` // 当服务列表为空时是否更新本地缓存，true: 更新,false: 不更新
}

func (x *Registry_Nacos) Reset() {
	*x = Registry_Nacos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry_Nacos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry_Nacos) ProtoMessage() {}

func (x *Registry_Nacos) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry_Nacos.ProtoReflect.Descriptor instead.
func (*Registry_Nacos) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Registry_Nacos) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Registry_Nacos) GetPort() uint64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Registry_Nacos) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *Registry_Nacos) GetLogLevel() string {
	if x != nil {
		return x.LogLevel
	}
	return ""
}

func (x *Registry_Nacos) GetCacheDir() string {
	if x != nil {
		return x.CacheDir
	}
	return ""
}

func (x *Registry_Nacos) GetLogDir() string {
	if x != nil {
		return x.LogDir
	}
	return ""
}

func (x *Registry_Nacos) GetUpdateThreadNum() int32 {
	if x != nil {
		return x.UpdateThreadNum
	}
	return 0
}

func (x *Registry_Nacos) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Registry_Nacos) GetBeatInterval() *durationpb.Duration {
	if x != nil {
		return x.BeatInterval
	}
	return nil
}

func (x *Registry_Nacos) GetNotLoadCacheAtStart() bool {
	if x != nil {
		return x.NotLoadCacheAtStart
	}
	return false
}

func (x *Registry_Nacos) GetUpdateCacheWhenEmpty() bool {
	if x != nil {
		return x.UpdateCacheWhenEmpty
	}
	return false
}

// Polaris
type Registry_Polaris struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address   string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`     // 服务端地址
	Port      int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`          // 服务端端口
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"` //命名空间
	Service   string `protobuf:"bytes,4,opt,name=service,proto3" json:"service,omitempty"`     //服务名
	Token     string `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Registry_Polaris) Reset() {
	*x = Registry_Polaris{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry_Polaris) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry_Polaris) ProtoMessage() {}

func (x *Registry_Polaris) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry_Polaris.ProtoReflect.Descriptor instead.
func (*Registry_Polaris) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Registry_Polaris) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Registry_Polaris) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Registry_Polaris) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Registry_Polaris) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Registry_Polaris) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// Kubernetes
type Registry_Kubernetes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Registry_Kubernetes) Reset() {
	*x = Registry_Kubernetes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry_Kubernetes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry_Kubernetes) ProtoMessage() {}

func (x *Registry_Kubernetes) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry_Kubernetes.ProtoReflect.Descriptor instead.
func (*Registry_Kubernetes) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{0, 4}
}

var File_registry_proto protoreflect.FileDescriptor

var file_registry_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x89, 0x08, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x37, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x31, 0x0a, 0x04, 0x65, 0x74,
	0x63, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x73,
	0x74, 0x72, 0x61, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x45, 0x74, 0x63, 0x64, 0x52, 0x04, 0x65, 0x74, 0x63, 0x64, 0x12, 0x34, 0x0a,
	0x05, 0x6e, 0x61, 0x63, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62,
	0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4e, 0x61, 0x63, 0x6f, 0x73, 0x52, 0x05, 0x6e, 0x61,
	0x63, 0x6f, 0x73, 0x12, 0x3a, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x50,
	0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x12,
	0x43, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x1a, 0x5c, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x1a, 0x24, 0x0a, 0x04, 0x45, 0x74, 0x63, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0xab, 0x03, 0x0a, 0x05, 0x4e, 0x61, 0x63,
	0x6f, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x63, 0x68, 0x65, 0x44, 0x69, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x61, 0x63, 0x68, 0x65, 0x44, 0x69, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f,
	0x67, 0x44, 0x69, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x44,
	0x69, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x4e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x33, 0x0a, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x62, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x62, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x12, 0x30, 0x0a, 0x13, 0x6e, 0x6f, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x41, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x6e,
	0x6f, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x63, 0x68, 0x65, 0x41, 0x74, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x57, 0x68, 0x65, 0x6e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x14, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x57, 0x68, 0x65,
	0x6e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x85, 0x01, 0x0a, 0x07, 0x50, 0x6f, 0x6c, 0x61, 0x72,
	0x69, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x0c,
	0x0a, 0x0a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x42, 0x1d, 0x5a, 0x1b,
	0x66, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61,
	0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_registry_proto_rawDescOnce sync.Once
	file_registry_proto_rawDescData = file_registry_proto_rawDesc
)

func file_registry_proto_rawDescGZIP() []byte {
	file_registry_proto_rawDescOnce.Do(func() {
		file_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_registry_proto_rawDescData)
	})
	return file_registry_proto_rawDescData
}

var file_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_registry_proto_goTypes = []interface{}{
	(*Registry)(nil),            // 0: bootstrap.conf.Registry
	(*Registry_Consul)(nil),     // 1: bootstrap.conf.Registry.Consul
	(*Registry_Etcd)(nil),       // 2: bootstrap.conf.Registry.Etcd
	(*Registry_Nacos)(nil),      // 3: bootstrap.conf.Registry.Nacos
	(*Registry_Polaris)(nil),    // 4: bootstrap.conf.Registry.Polaris
	(*Registry_Kubernetes)(nil), // 5: bootstrap.conf.Registry.Kubernetes
	(*durationpb.Duration)(nil), // 6: google.protobuf.Duration
}
var file_registry_proto_depIdxs = []int32{
	1, // 0: bootstrap.conf.Registry.consul:type_name -> bootstrap.conf.Registry.Consul
	2, // 1: bootstrap.conf.Registry.etcd:type_name -> bootstrap.conf.Registry.Etcd
	3, // 2: bootstrap.conf.Registry.nacos:type_name -> bootstrap.conf.Registry.Nacos
	4, // 3: bootstrap.conf.Registry.polaris:type_name -> bootstrap.conf.Registry.Polaris
	5, // 4: bootstrap.conf.Registry.kubernetes:type_name -> bootstrap.conf.Registry.Kubernetes
	6, // 5: bootstrap.conf.Registry.Nacos.timeout:type_name -> google.protobuf.Duration
	6, // 6: bootstrap.conf.Registry.Nacos.beatInterval:type_name -> google.protobuf.Duration
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_registry_proto_init() }
func file_registry_proto_init() {
	if File_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry); i {
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
		file_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry_Consul); i {
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
		file_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry_Etcd); i {
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
		file_registry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry_Nacos); i {
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
		file_registry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry_Polaris); i {
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
		file_registry_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry_Kubernetes); i {
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
			RawDescriptor: file_registry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_registry_proto_goTypes,
		DependencyIndexes: file_registry_proto_depIdxs,
		MessageInfos:      file_registry_proto_msgTypes,
	}.Build()
	File_registry_proto = out.File
	file_registry_proto_rawDesc = nil
	file_registry_proto_goTypes = nil
	file_registry_proto_depIdxs = nil
}
