// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: rpc_common/v1/sensitive_word.proto

package v1

import (
	"context"
	"fkratos/api/paginator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SensitiveWord_SensitiveWordList_FullMethodName  = "/api.rpc_common.v1.SensitiveWord/SensitiveWordList"
	SensitiveWord_SensitiveWordStore_FullMethodName = "/api.rpc_common.v1.SensitiveWord/SensitiveWordStore"
	SensitiveWord_SensitiveWordDel_FullMethodName   = "/api.rpc_common.v1.SensitiveWord/SensitiveWordDel"
	SensitiveWord_SensitiveWordCheck_FullMethodName = "/api.rpc_common.v1.SensitiveWord/SensitiveWordCheck"
)

// SensitiveWordClient is the client API for SensitiveWord service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SensitiveWordClient interface {
	// 敏感词-列表
	SensitiveWordList(ctx context.Context, in *paginator.PaginatorReq, opts ...grpc.CallOption) (*SensitiveWordListReply, error)
	// 敏感词-保存
	SensitiveWordStore(ctx context.Context, in *SensitiveWordStoreReq, opts ...grpc.CallOption) (*SensitiveWordStoreReply, error)
	// 敏感词-删除
	SensitiveWordDel(ctx context.Context, in *SensitiveWordDelReq, opts ...grpc.CallOption) (*SensitiveWordDelReply, error)
	// 敏感词-检测
	SensitiveWordCheck(ctx context.Context, in *SensitiveWordCheckReq, opts ...grpc.CallOption) (*SensitiveWordCheckResp, error)
}

type sensitiveWordClient struct {
	cc grpc.ClientConnInterface
}

func NewSensitiveWordClient(cc grpc.ClientConnInterface) SensitiveWordClient {
	return &sensitiveWordClient{cc}
}

func (c *sensitiveWordClient) SensitiveWordList(ctx context.Context, in *paginator.PaginatorReq, opts ...grpc.CallOption) (*SensitiveWordListReply, error) {
	out := new(SensitiveWordListReply)
	err := c.cc.Invoke(ctx, SensitiveWord_SensitiveWordList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensitiveWordClient) SensitiveWordStore(ctx context.Context, in *SensitiveWordStoreReq, opts ...grpc.CallOption) (*SensitiveWordStoreReply, error) {
	out := new(SensitiveWordStoreReply)
	err := c.cc.Invoke(ctx, SensitiveWord_SensitiveWordStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensitiveWordClient) SensitiveWordDel(ctx context.Context, in *SensitiveWordDelReq, opts ...grpc.CallOption) (*SensitiveWordDelReply, error) {
	out := new(SensitiveWordDelReply)
	err := c.cc.Invoke(ctx, SensitiveWord_SensitiveWordDel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensitiveWordClient) SensitiveWordCheck(ctx context.Context, in *SensitiveWordCheckReq, opts ...grpc.CallOption) (*SensitiveWordCheckResp, error) {
	out := new(SensitiveWordCheckResp)
	err := c.cc.Invoke(ctx, SensitiveWord_SensitiveWordCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SensitiveWordServer is the server API for SensitiveWord service.
// All implementations must embed UnimplementedSensitiveWordServer
// for forward compatibility
type SensitiveWordServer interface {
	// 敏感词-列表
	SensitiveWordList(context.Context, *paginator.PaginatorReq) (*SensitiveWordListReply, error)
	// 敏感词-保存
	SensitiveWordStore(context.Context, *SensitiveWordStoreReq) (*SensitiveWordStoreReply, error)
	// 敏感词-删除
	SensitiveWordDel(context.Context, *SensitiveWordDelReq) (*SensitiveWordDelReply, error)
	// 敏感词-检测
	SensitiveWordCheck(context.Context, *SensitiveWordCheckReq) (*SensitiveWordCheckResp, error)
	mustEmbedUnimplementedSensitiveWordServer()
}

// UnimplementedSensitiveWordServer must be embedded to have forward compatible implementations.
type UnimplementedSensitiveWordServer struct {
}

func (UnimplementedSensitiveWordServer) SensitiveWordList(context.Context, *paginator.PaginatorReq) (*SensitiveWordListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SensitiveWordList not implemented")
}
func (UnimplementedSensitiveWordServer) SensitiveWordStore(context.Context, *SensitiveWordStoreReq) (*SensitiveWordStoreReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SensitiveWordStore not implemented")
}
func (UnimplementedSensitiveWordServer) SensitiveWordDel(context.Context, *SensitiveWordDelReq) (*SensitiveWordDelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SensitiveWordDel not implemented")
}
func (UnimplementedSensitiveWordServer) SensitiveWordCheck(context.Context, *SensitiveWordCheckReq) (*SensitiveWordCheckResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SensitiveWordCheck not implemented")
}
func (UnimplementedSensitiveWordServer) mustEmbedUnimplementedSensitiveWordServer() {}

// UnsafeSensitiveWordServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SensitiveWordServer will
// result in compilation errors.
type UnsafeSensitiveWordServer interface {
	mustEmbedUnimplementedSensitiveWordServer()
}

func RegisterSensitiveWordServer(s grpc.ServiceRegistrar, srv SensitiveWordServer) {
	s.RegisterService(&SensitiveWord_ServiceDesc, srv)
}

func _SensitiveWord_SensitiveWordList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(paginator.PaginatorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensitiveWordServer).SensitiveWordList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SensitiveWord_SensitiveWordList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensitiveWordServer).SensitiveWordList(ctx, req.(*paginator.PaginatorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SensitiveWord_SensitiveWordStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SensitiveWordStoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensitiveWordServer).SensitiveWordStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SensitiveWord_SensitiveWordStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensitiveWordServer).SensitiveWordStore(ctx, req.(*SensitiveWordStoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SensitiveWord_SensitiveWordDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SensitiveWordDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensitiveWordServer).SensitiveWordDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SensitiveWord_SensitiveWordDel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensitiveWordServer).SensitiveWordDel(ctx, req.(*SensitiveWordDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SensitiveWord_SensitiveWordCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SensitiveWordCheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensitiveWordServer).SensitiveWordCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SensitiveWord_SensitiveWordCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensitiveWordServer).SensitiveWordCheck(ctx, req.(*SensitiveWordCheckReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SensitiveWord_ServiceDesc is the grpc.ServiceDesc for SensitiveWord service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SensitiveWord_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.rpc_common.v1.SensitiveWord",
	HandlerType: (*SensitiveWordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SensitiveWordList",
			Handler:    _SensitiveWord_SensitiveWordList_Handler,
		},
		{
			MethodName: "SensitiveWordStore",
			Handler:    _SensitiveWord_SensitiveWordStore_Handler,
		},
		{
			MethodName: "SensitiveWordDel",
			Handler:    _SensitiveWord_SensitiveWordDel_Handler,
		},
		{
			MethodName: "SensitiveWordCheck",
			Handler:    _SensitiveWord_SensitiveWordCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_common/v1/sensitive_word.proto",
}
