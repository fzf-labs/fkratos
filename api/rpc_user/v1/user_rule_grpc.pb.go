// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: rpc_user/v1/user_rule.proto

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
	UserRule_UserRuleList_FullMethodName  = "/api.rpc_user.v1.UserRule/UserRuleList"
	UserRule_UserRuleInfo_FullMethodName  = "/api.rpc_user.v1.UserRule/UserRuleInfo"
	UserRule_UserRuleStore_FullMethodName = "/api.rpc_user.v1.UserRule/UserRuleStore"
	UserRule_UserRuleDel_FullMethodName   = "/api.rpc_user.v1.UserRule/UserRuleDel"
	UserRule_UserRules_FullMethodName     = "/api.rpc_user.v1.UserRule/UserRules"
)

// UserRuleClient is the client API for UserRule service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserRuleClient interface {
	// 用户规则-列表
	UserRuleList(ctx context.Context, in *UserRuleListReq, opts ...grpc.CallOption) (*UserRuleListReply, error)
	// 用户规则-信息
	UserRuleInfo(ctx context.Context, in *UserRuleInfoReq, opts ...grpc.CallOption) (*UserRuleInfoReply, error)
	// 用户规则-保存
	UserRuleStore(ctx context.Context, in *UserRuleStoreReq, opts ...grpc.CallOption) (*UserRuleStoreReply, error)
	// 用户规则-删除
	UserRuleDel(ctx context.Context, in *UserRuleDelReq, opts ...grpc.CallOption) (*UserRuleDelReply, error)
	// 用户规则-用户所有规则
	UserRules(ctx context.Context, in *UserRulesReq, opts ...grpc.CallOption) (*UserRulesReply, error)
}

type userRuleClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRuleClient(cc grpc.ClientConnInterface) UserRuleClient {
	return &userRuleClient{cc}
}

func (c *userRuleClient) UserRuleList(ctx context.Context, in *UserRuleListReq, opts ...grpc.CallOption) (*UserRuleListReply, error) {
	out := new(UserRuleListReply)
	err := c.cc.Invoke(ctx, UserRule_UserRuleList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRuleClient) UserRuleInfo(ctx context.Context, in *UserRuleInfoReq, opts ...grpc.CallOption) (*UserRuleInfoReply, error) {
	out := new(UserRuleInfoReply)
	err := c.cc.Invoke(ctx, UserRule_UserRuleInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRuleClient) UserRuleStore(ctx context.Context, in *UserRuleStoreReq, opts ...grpc.CallOption) (*UserRuleStoreReply, error) {
	out := new(UserRuleStoreReply)
	err := c.cc.Invoke(ctx, UserRule_UserRuleStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRuleClient) UserRuleDel(ctx context.Context, in *UserRuleDelReq, opts ...grpc.CallOption) (*UserRuleDelReply, error) {
	out := new(UserRuleDelReply)
	err := c.cc.Invoke(ctx, UserRule_UserRuleDel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRuleClient) UserRules(ctx context.Context, in *UserRulesReq, opts ...grpc.CallOption) (*UserRulesReply, error) {
	out := new(UserRulesReply)
	err := c.cc.Invoke(ctx, UserRule_UserRules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRuleServer is the server API for UserRule service.
// All implementations must embed UnimplementedUserRuleServer
// for forward compatibility
type UserRuleServer interface {
	// 用户规则-列表
	UserRuleList(context.Context, *UserRuleListReq) (*UserRuleListReply, error)
	// 用户规则-信息
	UserRuleInfo(context.Context, *UserRuleInfoReq) (*UserRuleInfoReply, error)
	// 用户规则-保存
	UserRuleStore(context.Context, *UserRuleStoreReq) (*UserRuleStoreReply, error)
	// 用户规则-删除
	UserRuleDel(context.Context, *UserRuleDelReq) (*UserRuleDelReply, error)
	// 用户规则-用户所有规则
	UserRules(context.Context, *UserRulesReq) (*UserRulesReply, error)
	mustEmbedUnimplementedUserRuleServer()
}

// UnimplementedUserRuleServer must be embedded to have forward compatible implementations.
type UnimplementedUserRuleServer struct {
}

func (UnimplementedUserRuleServer) UserRuleList(context.Context, *UserRuleListReq) (*UserRuleListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRuleList not implemented")
}
func (UnimplementedUserRuleServer) UserRuleInfo(context.Context, *UserRuleInfoReq) (*UserRuleInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRuleInfo not implemented")
}
func (UnimplementedUserRuleServer) UserRuleStore(context.Context, *UserRuleStoreReq) (*UserRuleStoreReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRuleStore not implemented")
}
func (UnimplementedUserRuleServer) UserRuleDel(context.Context, *UserRuleDelReq) (*UserRuleDelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRuleDel not implemented")
}
func (UnimplementedUserRuleServer) UserRules(context.Context, *UserRulesReq) (*UserRulesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRules not implemented")
}
func (UnimplementedUserRuleServer) mustEmbedUnimplementedUserRuleServer() {}

// UnsafeUserRuleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserRuleServer will
// result in compilation errors.
type UnsafeUserRuleServer interface {
	mustEmbedUnimplementedUserRuleServer()
}

func RegisterUserRuleServer(s grpc.ServiceRegistrar, srv UserRuleServer) {
	s.RegisterService(&UserRule_ServiceDesc, srv)
}

func _UserRule_UserRuleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRuleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRuleServer).UserRuleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRule_UserRuleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRuleServer).UserRuleList(ctx, req.(*UserRuleListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRule_UserRuleInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRuleInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRuleServer).UserRuleInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRule_UserRuleInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRuleServer).UserRuleInfo(ctx, req.(*UserRuleInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRule_UserRuleStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRuleStoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRuleServer).UserRuleStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRule_UserRuleStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRuleServer).UserRuleStore(ctx, req.(*UserRuleStoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRule_UserRuleDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRuleDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRuleServer).UserRuleDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRule_UserRuleDel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRuleServer).UserRuleDel(ctx, req.(*UserRuleDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRule_UserRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRulesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRuleServer).UserRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRule_UserRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRuleServer).UserRules(ctx, req.(*UserRulesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserRule_ServiceDesc is the grpc.ServiceDesc for UserRule service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserRule_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.rpc_user.v1.UserRule",
	HandlerType: (*UserRuleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRuleList",
			Handler:    _UserRule_UserRuleList_Handler,
		},
		{
			MethodName: "UserRuleInfo",
			Handler:    _UserRule_UserRuleInfo_Handler,
		},
		{
			MethodName: "UserRuleStore",
			Handler:    _UserRule_UserRuleStore_Handler,
		},
		{
			MethodName: "UserRuleDel",
			Handler:    _UserRule_UserRuleDel_Handler,
		},
		{
			MethodName: "UserRules",
			Handler:    _UserRule_UserRules_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_user/v1/user_rule.proto",
}