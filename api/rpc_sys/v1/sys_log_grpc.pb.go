// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: rpc_sys/v1/sys_log.proto

package v1

import (
	context "context"
	common "fkratos/api/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Log_SysLogList_FullMethodName  = "/api.rpc_sys.v1.Log/SysLogList"
	Log_SysLogInfo_FullMethodName  = "/api.rpc_sys.v1.Log/SysLogInfo"
	Log_SysLogStore_FullMethodName = "/api.rpc_sys.v1.Log/SysLogStore"
)

// LogClient is the client API for Log service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogClient interface {
	// 日志列表
	SysLogList(ctx context.Context, in *common.SearchListReq, opts ...grpc.CallOption) (*SysLogListResp, error)
	// 单条日志
	SysLogInfo(ctx context.Context, in *SysLogInfoReq, opts ...grpc.CallOption) (*SysLogInfoResp, error)
	// 日志记录
	SysLogStore(ctx context.Context, in *SysLogStoreReq, opts ...grpc.CallOption) (*SysLogStoreResp, error)
}

type logClient struct {
	cc grpc.ClientConnInterface
}

func NewLogClient(cc grpc.ClientConnInterface) LogClient {
	return &logClient{cc}
}

func (c *logClient) SysLogList(ctx context.Context, in *common.SearchListReq, opts ...grpc.CallOption) (*SysLogListResp, error) {
	out := new(SysLogListResp)
	err := c.cc.Invoke(ctx, Log_SysLogList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) SysLogInfo(ctx context.Context, in *SysLogInfoReq, opts ...grpc.CallOption) (*SysLogInfoResp, error) {
	out := new(SysLogInfoResp)
	err := c.cc.Invoke(ctx, Log_SysLogInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) SysLogStore(ctx context.Context, in *SysLogStoreReq, opts ...grpc.CallOption) (*SysLogStoreResp, error) {
	out := new(SysLogStoreResp)
	err := c.cc.Invoke(ctx, Log_SysLogStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServer is the server API for Log service.
// All implementations must embed UnimplementedLogServer
// for forward compatibility
type LogServer interface {
	// 日志列表
	SysLogList(context.Context, *common.SearchListReq) (*SysLogListResp, error)
	// 单条日志
	SysLogInfo(context.Context, *SysLogInfoReq) (*SysLogInfoResp, error)
	// 日志记录
	SysLogStore(context.Context, *SysLogStoreReq) (*SysLogStoreResp, error)
	mustEmbedUnimplementedLogServer()
}

// UnimplementedLogServer must be embedded to have forward compatible implementations.
type UnimplementedLogServer struct {
}

func (UnimplementedLogServer) SysLogList(context.Context, *common.SearchListReq) (*SysLogListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysLogList not implemented")
}
func (UnimplementedLogServer) SysLogInfo(context.Context, *SysLogInfoReq) (*SysLogInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysLogInfo not implemented")
}
func (UnimplementedLogServer) SysLogStore(context.Context, *SysLogStoreReq) (*SysLogStoreResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysLogStore not implemented")
}
func (UnimplementedLogServer) mustEmbedUnimplementedLogServer() {}

// UnsafeLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServer will
// result in compilation errors.
type UnsafeLogServer interface {
	mustEmbedUnimplementedLogServer()
}

func RegisterLogServer(s grpc.ServiceRegistrar, srv LogServer) {
	s.RegisterService(&Log_ServiceDesc, srv)
}

func _Log_SysLogList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.SearchListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).SysLogList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_SysLogList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).SysLogList(ctx, req.(*common.SearchListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_SysLogInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysLogInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).SysLogInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_SysLogInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).SysLogInfo(ctx, req.(*SysLogInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_SysLogStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysLogStoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).SysLogStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_SysLogStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).SysLogStore(ctx, req.(*SysLogStoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Log_ServiceDesc is the grpc.ServiceDesc for Log service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Log_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.rpc_sys.v1.Log",
	HandlerType: (*LogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SysLogList",
			Handler:    _Log_SysLogList_Handler,
		},
		{
			MethodName: "SysLogInfo",
			Handler:    _Log_SysLogInfo_Handler,
		},
		{
			MethodName: "SysLogStore",
			Handler:    _Log_SysLogStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_sys/v1/sys_log.proto",
}
