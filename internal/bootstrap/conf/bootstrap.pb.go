// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: bootstrap.proto

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

// 引导信息
type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string        `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Server      *Server       `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	Data        *Data         `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Trace       *Tracer       `protobuf:"bytes,4,opt,name=trace,proto3" json:"trace,omitempty"`
	Logger      *Logger       `protobuf:"bytes,5,opt,name=logger,proto3" json:"logger,omitempty"`
	Registry    *Registry     `protobuf:"bytes,6,opt,name=registry,proto3" json:"registry,omitempty"`
	Config      *RemoteConfig `protobuf:"bytes,7,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootstrap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_bootstrap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_bootstrap_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetTrace() *Tracer {
	if x != nil {
		return x.Trace
	}
	return nil
}

func (x *Bootstrap) GetLogger() *Logger {
	if x != nil {
		return x.Logger
	}
	return nil
}

func (x *Bootstrap) GetRegistry() *Registry {
	if x != nil {
		return x.Registry
	}
	return nil
}

func (x *Bootstrap) GetConfig() *RemoteConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_bootstrap_proto protoreflect.FileDescriptor

var file_bootstrap_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x1a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0a, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x02, 0x0a, 0x09,
	0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x6f,
	0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x72, 0x52, 0x05, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x06, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61,
	0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x34, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42,
	0x1d, 0x5a, 0x1b, 0x66, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x73,
	0x74, 0x72, 0x61, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bootstrap_proto_rawDescOnce sync.Once
	file_bootstrap_proto_rawDescData = file_bootstrap_proto_rawDesc
)

func file_bootstrap_proto_rawDescGZIP() []byte {
	file_bootstrap_proto_rawDescOnce.Do(func() {
		file_bootstrap_proto_rawDescData = protoimpl.X.CompressGZIP(file_bootstrap_proto_rawDescData)
	})
	return file_bootstrap_proto_rawDescData
}

var file_bootstrap_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bootstrap_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),    // 0: bootstrap.conf.Bootstrap
	(*Server)(nil),       // 1: bootstrap.conf.Server
	(*Data)(nil),         // 2: bootstrap.conf.Data
	(*Tracer)(nil),       // 3: bootstrap.conf.Tracer
	(*Logger)(nil),       // 4: bootstrap.conf.Logger
	(*Registry)(nil),     // 5: bootstrap.conf.Registry
	(*RemoteConfig)(nil), // 6: bootstrap.conf.RemoteConfig
}
var file_bootstrap_proto_depIdxs = []int32{
	1, // 0: bootstrap.conf.Bootstrap.server:type_name -> bootstrap.conf.Server
	2, // 1: bootstrap.conf.Bootstrap.data:type_name -> bootstrap.conf.Data
	3, // 2: bootstrap.conf.Bootstrap.trace:type_name -> bootstrap.conf.Tracer
	4, // 3: bootstrap.conf.Bootstrap.logger:type_name -> bootstrap.conf.Logger
	5, // 4: bootstrap.conf.Bootstrap.registry:type_name -> bootstrap.conf.Registry
	6, // 5: bootstrap.conf.Bootstrap.config:type_name -> bootstrap.conf.RemoteConfig
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_bootstrap_proto_init() }
func file_bootstrap_proto_init() {
	if File_bootstrap_proto != nil {
		return
	}
	file_server_proto_init()
	file_data_proto_init()
	file_tracer_proto_init()
	file_logger_proto_init()
	file_registry_proto_init()
	file_remoteconfig_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bootstrap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
			RawDescriptor: file_bootstrap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bootstrap_proto_goTypes,
		DependencyIndexes: file_bootstrap_proto_depIdxs,
		MessageInfos:      file_bootstrap_proto_msgTypes,
	}.Build()
	File_bootstrap_proto = out.File
	file_bootstrap_proto_rawDesc = nil
	file_bootstrap_proto_goTypes = nil
	file_bootstrap_proto_depIdxs = nil
}
