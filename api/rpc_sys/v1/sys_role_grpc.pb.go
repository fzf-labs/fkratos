// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: rpc_sys/v1/sys_role.proto

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
	Role_SysRoleList_FullMethodName   = "/api.rpc_sys.v1.Role/SysRoleList"
	Role_SysRoleInfo_FullMethodName   = "/api.rpc_sys.v1.Role/SysRoleInfo"
	Role_SysRoleStore_FullMethodName  = "/api.rpc_sys.v1.Role/SysRoleStore"
	Role_SysRoleDel_FullMethodName    = "/api.rpc_sys.v1.Role/SysRoleDel"
	Role_SysRoleStatus_FullMethodName = "/api.rpc_sys.v1.Role/SysRoleStatus"
)

// RoleClient is the client API for Role service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleClient interface {
	// 权限列表
	SysRoleList(ctx context.Context, in *SysRoleListReq, opts ...grpc.CallOption) (*SysRoleListResp, error)
	// 单个权限
	SysRoleInfo(ctx context.Context, in *SysRoleInfoReq, opts ...grpc.CallOption) (*SysRoleInfoResp, error)
	// 保存权限
	SysRoleStore(ctx context.Context, in *SysRoleStoreReq, opts ...grpc.CallOption) (*SysRoleStoreResp, error)
	// 删除权限
	SysRoleDel(ctx context.Context, in *SysRoleDelReq, opts ...grpc.CallOption) (*SysRoleDelResp, error)
	// 修改权限状态
	SysRoleStatus(ctx context.Context, in *SysRoleStatusReq, opts ...grpc.CallOption) (*SysRoleStatusResp, error)
}

type roleClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleClient(cc grpc.ClientConnInterface) RoleClient {
	return &roleClient{cc}
}

func (c *roleClient) SysRoleList(ctx context.Context, in *SysRoleListReq, opts ...grpc.CallOption) (*SysRoleListResp, error) {
	out := new(SysRoleListResp)
	err := c.cc.Invoke(ctx, Role_SysRoleList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) SysRoleInfo(ctx context.Context, in *SysRoleInfoReq, opts ...grpc.CallOption) (*SysRoleInfoResp, error) {
	out := new(SysRoleInfoResp)
	err := c.cc.Invoke(ctx, Role_SysRoleInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) SysRoleStore(ctx context.Context, in *SysRoleStoreReq, opts ...grpc.CallOption) (*SysRoleStoreResp, error) {
	out := new(SysRoleStoreResp)
	err := c.cc.Invoke(ctx, Role_SysRoleStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) SysRoleDel(ctx context.Context, in *SysRoleDelReq, opts ...grpc.CallOption) (*SysRoleDelResp, error) {
	out := new(SysRoleDelResp)
	err := c.cc.Invoke(ctx, Role_SysRoleDel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) SysRoleStatus(ctx context.Context, in *SysRoleStatusReq, opts ...grpc.CallOption) (*SysRoleStatusResp, error) {
	out := new(SysRoleStatusResp)
	err := c.cc.Invoke(ctx, Role_SysRoleStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServer is the server API for Role service.
// All implementations must embed UnimplementedRoleServer
// for forward compatibility
type RoleServer interface {
	// 权限列表
	SysRoleList(context.Context, *SysRoleListReq) (*SysRoleListResp, error)
	// 单个权限
	SysRoleInfo(context.Context, *SysRoleInfoReq) (*SysRoleInfoResp, error)
	// 保存权限
	SysRoleStore(context.Context, *SysRoleStoreReq) (*SysRoleStoreResp, error)
	// 删除权限
	SysRoleDel(context.Context, *SysRoleDelReq) (*SysRoleDelResp, error)
	// 修改权限状态
	SysRoleStatus(context.Context, *SysRoleStatusReq) (*SysRoleStatusResp, error)
	mustEmbedUnimplementedRoleServer()
}

// UnimplementedRoleServer must be embedded to have forward compatible implementations.
type UnimplementedRoleServer struct {
}

func (UnimplementedRoleServer) SysRoleList(context.Context, *SysRoleListReq) (*SysRoleListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysRoleList not implemented")
}
func (UnimplementedRoleServer) SysRoleInfo(context.Context, *SysRoleInfoReq) (*SysRoleInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysRoleInfo not implemented")
}
func (UnimplementedRoleServer) SysRoleStore(context.Context, *SysRoleStoreReq) (*SysRoleStoreResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysRoleStore not implemented")
}
func (UnimplementedRoleServer) SysRoleDel(context.Context, *SysRoleDelReq) (*SysRoleDelResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysRoleDel not implemented")
}
func (UnimplementedRoleServer) SysRoleStatus(context.Context, *SysRoleStatusReq) (*SysRoleStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysRoleStatus not implemented")
}
func (UnimplementedRoleServer) mustEmbedUnimplementedRoleServer() {}

// UnsafeRoleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServer will
// result in compilation errors.
type UnsafeRoleServer interface {
	mustEmbedUnimplementedRoleServer()
}

func RegisterRoleServer(s grpc.ServiceRegistrar, srv RoleServer) {
	s.RegisterService(&Role_ServiceDesc, srv)
}

func _Role_SysRoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysRoleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).SysRoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_SysRoleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).SysRoleList(ctx, req.(*SysRoleListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_SysRoleInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysRoleInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).SysRoleInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_SysRoleInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).SysRoleInfo(ctx, req.(*SysRoleInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_SysRoleStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysRoleStoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).SysRoleStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_SysRoleStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).SysRoleStore(ctx, req.(*SysRoleStoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_SysRoleDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysRoleDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).SysRoleDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_SysRoleDel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).SysRoleDel(ctx, req.(*SysRoleDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_SysRoleStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysRoleStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).SysRoleStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_SysRoleStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).SysRoleStatus(ctx, req.(*SysRoleStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Role_ServiceDesc is the grpc.ServiceDesc for Role service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Role_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.rpc_sys.v1.Role",
	HandlerType: (*RoleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SysRoleList",
			Handler:    _Role_SysRoleList_Handler,
		},
		{
			MethodName: "SysRoleInfo",
			Handler:    _Role_SysRoleInfo_Handler,
		},
		{
			MethodName: "SysRoleStore",
			Handler:    _Role_SysRoleStore_Handler,
		},
		{
			MethodName: "SysRoleDel",
			Handler:    _Role_SysRoleDel_Handler,
		},
		{
			MethodName: "SysRoleStatus",
			Handler:    _Role_SysRoleStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_sys/v1/sys_role.proto",
}
