// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: rpc_sys/v1/sys_admin.proto

package v1

import (
	context "context"
	paginator "fkratos/api/paginator"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Admin_SysAdminInfo_FullMethodName           = "/api.rpc_sys.v1.Admin/SysAdminInfo"
	Admin_SysAdminInfoUpdate_FullMethodName     = "/api.rpc_sys.v1.Admin/SysAdminInfoUpdate"
	Admin_SysAdminGenerateAvatar_FullMethodName = "/api.rpc_sys.v1.Admin/SysAdminGenerateAvatar"
	Admin_SysAdminPermission_FullMethodName     = "/api.rpc_sys.v1.Admin/SysAdminPermission"
	Admin_SysManageList_FullMethodName          = "/api.rpc_sys.v1.Admin/SysManageList"
	Admin_SysManageInfo_FullMethodName          = "/api.rpc_sys.v1.Admin/SysManageInfo"
	Admin_SysManageStore_FullMethodName         = "/api.rpc_sys.v1.Admin/SysManageStore"
	Admin_SysManageDel_FullMethodName           = "/api.rpc_sys.v1.Admin/SysManageDel"
)

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	// 管理员信息
	SysAdminInfo(ctx context.Context, in *SysAdminInfoReq, opts ...grpc.CallOption) (*SysAdminInfoReply, error)
	// 管理员更新
	SysAdminInfoUpdate(ctx context.Context, in *SysAdminInfoUpdateReq, opts ...grpc.CallOption) (*SysAdminInfoUpdateReply, error)
	// 生成头像
	SysAdminGenerateAvatar(ctx context.Context, in *SysAdminGenerateAvatarReq, opts ...grpc.CallOption) (*SysAdminGenerateAvatarReply, error)
	// 查询权限
	SysAdminPermission(ctx context.Context, in *SysAdminPermissionReq, opts ...grpc.CallOption) (*SysAdminPermissionReply, error)
	// 管理员列表
	SysManageList(ctx context.Context, in *paginator.PaginatorReq, opts ...grpc.CallOption) (*SysManageListReply, error)
	// 单个管理员
	SysManageInfo(ctx context.Context, in *SysManageInfoReq, opts ...grpc.CallOption) (*SysManageInfoReply, error)
	// 保存管理员
	SysManageStore(ctx context.Context, in *SysManageStoreReq, opts ...grpc.CallOption) (*SysManageStoreReply, error)
	// 删除管理员
	SysManageDel(ctx context.Context, in *SysManageDelReq, opts ...grpc.CallOption) (*SysManageDelReply, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) SysAdminInfo(ctx context.Context, in *SysAdminInfoReq, opts ...grpc.CallOption) (*SysAdminInfoReply, error) {
	out := new(SysAdminInfoReply)
	err := c.cc.Invoke(ctx, Admin_SysAdminInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) SysAdminInfoUpdate(ctx context.Context, in *SysAdminInfoUpdateReq, opts ...grpc.CallOption) (*SysAdminInfoUpdateReply, error) {
	out := new(SysAdminInfoUpdateReply)
	err := c.cc.Invoke(ctx, Admin_SysAdminInfoUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) SysAdminGenerateAvatar(ctx context.Context, in *SysAdminGenerateAvatarReq, opts ...grpc.CallOption) (*SysAdminGenerateAvatarReply, error) {
	out := new(SysAdminGenerateAvatarReply)
	err := c.cc.Invoke(ctx, Admin_SysAdminGenerateAvatar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) SysAdminPermission(ctx context.Context, in *SysAdminPermissionReq, opts ...grpc.CallOption) (*SysAdminPermissionReply, error) {
	out := new(SysAdminPermissionReply)
	err := c.cc.Invoke(ctx, Admin_SysAdminPermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) SysManageList(ctx context.Context, in *paginator.PaginatorReq, opts ...grpc.CallOption) (*SysManageListReply, error) {
	out := new(SysManageListReply)
	err := c.cc.Invoke(ctx, Admin_SysManageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) SysManageInfo(ctx context.Context, in *SysManageInfoReq, opts ...grpc.CallOption) (*SysManageInfoReply, error) {
	out := new(SysManageInfoReply)
	err := c.cc.Invoke(ctx, Admin_SysManageInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) SysManageStore(ctx context.Context, in *SysManageStoreReq, opts ...grpc.CallOption) (*SysManageStoreReply, error) {
	out := new(SysManageStoreReply)
	err := c.cc.Invoke(ctx, Admin_SysManageStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) SysManageDel(ctx context.Context, in *SysManageDelReq, opts ...grpc.CallOption) (*SysManageDelReply, error) {
	out := new(SysManageDelReply)
	err := c.cc.Invoke(ctx, Admin_SysManageDel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	// 管理员信息
	SysAdminInfo(context.Context, *SysAdminInfoReq) (*SysAdminInfoReply, error)
	// 管理员更新
	SysAdminInfoUpdate(context.Context, *SysAdminInfoUpdateReq) (*SysAdminInfoUpdateReply, error)
	// 生成头像
	SysAdminGenerateAvatar(context.Context, *SysAdminGenerateAvatarReq) (*SysAdminGenerateAvatarReply, error)
	// 查询权限
	SysAdminPermission(context.Context, *SysAdminPermissionReq) (*SysAdminPermissionReply, error)
	// 管理员列表
	SysManageList(context.Context, *paginator.PaginatorReq) (*SysManageListReply, error)
	// 单个管理员
	SysManageInfo(context.Context, *SysManageInfoReq) (*SysManageInfoReply, error)
	// 保存管理员
	SysManageStore(context.Context, *SysManageStoreReq) (*SysManageStoreReply, error)
	// 删除管理员
	SysManageDel(context.Context, *SysManageDelReq) (*SysManageDelReply, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) SysAdminInfo(context.Context, *SysAdminInfoReq) (*SysAdminInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysAdminInfo not implemented")
}
func (UnimplementedAdminServer) SysAdminInfoUpdate(context.Context, *SysAdminInfoUpdateReq) (*SysAdminInfoUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysAdminInfoUpdate not implemented")
}
func (UnimplementedAdminServer) SysAdminGenerateAvatar(context.Context, *SysAdminGenerateAvatarReq) (*SysAdminGenerateAvatarReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysAdminGenerateAvatar not implemented")
}
func (UnimplementedAdminServer) SysAdminPermission(context.Context, *SysAdminPermissionReq) (*SysAdminPermissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysAdminPermission not implemented")
}
func (UnimplementedAdminServer) SysManageList(context.Context, *paginator.PaginatorReq) (*SysManageListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysManageList not implemented")
}
func (UnimplementedAdminServer) SysManageInfo(context.Context, *SysManageInfoReq) (*SysManageInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysManageInfo not implemented")
}
func (UnimplementedAdminServer) SysManageStore(context.Context, *SysManageStoreReq) (*SysManageStoreReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysManageStore not implemented")
}
func (UnimplementedAdminServer) SysManageDel(context.Context, *SysManageDelReq) (*SysManageDelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysManageDel not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_SysAdminInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysAdminInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).SysAdminInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_SysAdminInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).SysAdminInfo(ctx, req.(*SysAdminInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_SysAdminInfoUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysAdminInfoUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).SysAdminInfoUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_SysAdminInfoUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).SysAdminInfoUpdate(ctx, req.(*SysAdminInfoUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_SysAdminGenerateAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysAdminGenerateAvatarReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).SysAdminGenerateAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_SysAdminGenerateAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).SysAdminGenerateAvatar(ctx, req.(*SysAdminGenerateAvatarReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_SysAdminPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysAdminPermissionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).SysAdminPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_SysAdminPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).SysAdminPermission(ctx, req.(*SysAdminPermissionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_SysManageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(paginator.PaginatorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).SysManageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_SysManageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).SysManageList(ctx, req.(*paginator.PaginatorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_SysManageInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysManageInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).SysManageInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_SysManageInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).SysManageInfo(ctx, req.(*SysManageInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_SysManageStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysManageStoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).SysManageStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_SysManageStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).SysManageStore(ctx, req.(*SysManageStoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_SysManageDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysManageDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).SysManageDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_SysManageDel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).SysManageDel(ctx, req.(*SysManageDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.rpc_sys.v1.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SysAdminInfo",
			Handler:    _Admin_SysAdminInfo_Handler,
		},
		{
			MethodName: "SysAdminInfoUpdate",
			Handler:    _Admin_SysAdminInfoUpdate_Handler,
		},
		{
			MethodName: "SysAdminGenerateAvatar",
			Handler:    _Admin_SysAdminGenerateAvatar_Handler,
		},
		{
			MethodName: "SysAdminPermission",
			Handler:    _Admin_SysAdminPermission_Handler,
		},
		{
			MethodName: "SysManageList",
			Handler:    _Admin_SysManageList_Handler,
		},
		{
			MethodName: "SysManageInfo",
			Handler:    _Admin_SysManageInfo_Handler,
		},
		{
			MethodName: "SysManageStore",
			Handler:    _Admin_SysManageStore_Handler,
		},
		{
			MethodName: "SysManageDel",
			Handler:    _Admin_SysManageDel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_sys/v1/sys_admin.proto",
}
