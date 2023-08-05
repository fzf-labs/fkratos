// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: rpc_sys/v1/sys_permission.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Permission_SysPermissionList_FullMethodName   = "/api.rpc_sys.v1.Permission/SysPermissionList"
	Permission_SysPermissionInfo_FullMethodName   = "/api.rpc_sys.v1.Permission/SysPermissionInfo"
	Permission_SysPermissionStore_FullMethodName  = "/api.rpc_sys.v1.Permission/SysPermissionStore"
	Permission_SysPermissionDel_FullMethodName    = "/api.rpc_sys.v1.Permission/SysPermissionDel"
	Permission_SysPermissionStatus_FullMethodName = "/api.rpc_sys.v1.Permission/SysPermissionStatus"
)

// PermissionClient is the client API for Permission service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionClient interface {
	// 权限-列表
	SysPermissionList(ctx context.Context, in *SysPermissionListReq, opts ...grpc.CallOption) (*SysPermissionListResp, error)
	// 权限-单个权限信息
	SysPermissionInfo(ctx context.Context, in *SysPermissionInfoReq, opts ...grpc.CallOption) (*SysPermissionInfoResp, error)
	// 权限-保存
	SysPermissionStore(ctx context.Context, in *SysPermissionStoreReq, opts ...grpc.CallOption) (*SysPermissionStoreResp, error)
	// 权限-删除
	SysPermissionDel(ctx context.Context, in *SysPermissionDelReq, opts ...grpc.CallOption) (*SysPermissionDelResp, error)
	// 权限-修改状态
	SysPermissionStatus(ctx context.Context, in *SysPermissionStatusReq, opts ...grpc.CallOption) (*SysPermissionStatusResp, error)
}

type permissionClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionClient(cc grpc.ClientConnInterface) PermissionClient {
	return &permissionClient{cc}
}

func (c *permissionClient) SysPermissionList(ctx context.Context, in *SysPermissionListReq, opts ...grpc.CallOption) (*SysPermissionListResp, error) {
	out := new(SysPermissionListResp)
	err := c.cc.Invoke(ctx, Permission_SysPermissionList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionClient) SysPermissionInfo(ctx context.Context, in *SysPermissionInfoReq, opts ...grpc.CallOption) (*SysPermissionInfoResp, error) {
	out := new(SysPermissionInfoResp)
	err := c.cc.Invoke(ctx, Permission_SysPermissionInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionClient) SysPermissionStore(ctx context.Context, in *SysPermissionStoreReq, opts ...grpc.CallOption) (*SysPermissionStoreResp, error) {
	out := new(SysPermissionStoreResp)
	err := c.cc.Invoke(ctx, Permission_SysPermissionStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionClient) SysPermissionDel(ctx context.Context, in *SysPermissionDelReq, opts ...grpc.CallOption) (*SysPermissionDelResp, error) {
	out := new(SysPermissionDelResp)
	err := c.cc.Invoke(ctx, Permission_SysPermissionDel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionClient) SysPermissionStatus(ctx context.Context, in *SysPermissionStatusReq, opts ...grpc.CallOption) (*SysPermissionStatusResp, error) {
	out := new(SysPermissionStatusResp)
	err := c.cc.Invoke(ctx, Permission_SysPermissionStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionServer is the server API for Permission service.
// All implementations must embed UnimplementedPermissionServer
// for forward compatibility
type PermissionServer interface {
	// 权限-列表
	SysPermissionList(context.Context, *SysPermissionListReq) (*SysPermissionListResp, error)
	// 权限-单个权限信息
	SysPermissionInfo(context.Context, *SysPermissionInfoReq) (*SysPermissionInfoResp, error)
	// 权限-保存
	SysPermissionStore(context.Context, *SysPermissionStoreReq) (*SysPermissionStoreResp, error)
	// 权限-删除
	SysPermissionDel(context.Context, *SysPermissionDelReq) (*SysPermissionDelResp, error)
	// 权限-修改状态
	SysPermissionStatus(context.Context, *SysPermissionStatusReq) (*SysPermissionStatusResp, error)
	mustEmbedUnimplementedPermissionServer()
}

// UnimplementedPermissionServer must be embedded to have forward compatible implementations.
type UnimplementedPermissionServer struct {
}

func (UnimplementedPermissionServer) SysPermissionList(context.Context, *SysPermissionListReq) (*SysPermissionListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysPermissionList not implemented")
}
func (UnimplementedPermissionServer) SysPermissionInfo(context.Context, *SysPermissionInfoReq) (*SysPermissionInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysPermissionInfo not implemented")
}
func (UnimplementedPermissionServer) SysPermissionStore(context.Context, *SysPermissionStoreReq) (*SysPermissionStoreResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysPermissionStore not implemented")
}
func (UnimplementedPermissionServer) SysPermissionDel(context.Context, *SysPermissionDelReq) (*SysPermissionDelResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysPermissionDel not implemented")
}
func (UnimplementedPermissionServer) SysPermissionStatus(context.Context, *SysPermissionStatusReq) (*SysPermissionStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysPermissionStatus not implemented")
}
func (UnimplementedPermissionServer) mustEmbedUnimplementedPermissionServer() {}

// UnsafePermissionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionServer will
// result in compilation errors.
type UnsafePermissionServer interface {
	mustEmbedUnimplementedPermissionServer()
}

func RegisterPermissionServer(s grpc.ServiceRegistrar, srv PermissionServer) {
	s.RegisterService(&Permission_ServiceDesc, srv)
}

func _Permission_SysPermissionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysPermissionListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServer).SysPermissionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Permission_SysPermissionList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServer).SysPermissionList(ctx, req.(*SysPermissionListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Permission_SysPermissionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysPermissionInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServer).SysPermissionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Permission_SysPermissionInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServer).SysPermissionInfo(ctx, req.(*SysPermissionInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Permission_SysPermissionStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysPermissionStoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServer).SysPermissionStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Permission_SysPermissionStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServer).SysPermissionStore(ctx, req.(*SysPermissionStoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Permission_SysPermissionDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysPermissionDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServer).SysPermissionDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Permission_SysPermissionDel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServer).SysPermissionDel(ctx, req.(*SysPermissionDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Permission_SysPermissionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysPermissionStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServer).SysPermissionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Permission_SysPermissionStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServer).SysPermissionStatus(ctx, req.(*SysPermissionStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Permission_ServiceDesc is the grpc.ServiceDesc for Permission service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Permission_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.rpc_sys.v1.Permission",
	HandlerType: (*PermissionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SysPermissionList",
			Handler:    _Permission_SysPermissionList_Handler,
		},
		{
			MethodName: "SysPermissionInfo",
			Handler:    _Permission_SysPermissionInfo_Handler,
		},
		{
			MethodName: "SysPermissionStore",
			Handler:    _Permission_SysPermissionStore_Handler,
		},
		{
			MethodName: "SysPermissionDel",
			Handler:    _Permission_SysPermissionDel_Handler,
		},
		{
			MethodName: "SysPermissionStatus",
			Handler:    _Permission_SysPermissionStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_sys/v1/sys_permission.proto",
}
