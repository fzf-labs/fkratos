// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: rpc_wallet/v1/wallet.proto

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
	Wallet_WalletStore_FullMethodName = "/api.rpc_wallet.v1.Wallet/WalletStore"
	Wallet_WalletDel_FullMethodName   = "/api.rpc_wallet.v1.Wallet/WalletDel"
	Wallet_WalletOne_FullMethodName   = "/api.rpc_wallet.v1.Wallet/WalletOne"
	Wallet_WalletList_FullMethodName  = "/api.rpc_wallet.v1.Wallet/WalletList"
)

// WalletClient is the client API for Wallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletClient interface {
	// 钱包-创建一条数据
	WalletStore(ctx context.Context, in *WalletStoreReq, opts ...grpc.CallOption) (*WalletStoreReply, error)
	// 钱包-删除多条数据
	WalletDel(ctx context.Context, in *WalletDelReq, opts ...grpc.CallOption) (*WalletDelReply, error)
	// 钱包-单条数据查询
	WalletOne(ctx context.Context, in *WalletOneReq, opts ...grpc.CallOption) (*WalletOneReply, error)
	// 钱包-列表数据查询
	WalletList(ctx context.Context, in *WalletListReq, opts ...grpc.CallOption) (*WalletListReply, error)
}

type walletClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletClient(cc grpc.ClientConnInterface) WalletClient {
	return &walletClient{cc}
}

func (c *walletClient) WalletStore(ctx context.Context, in *WalletStoreReq, opts ...grpc.CallOption) (*WalletStoreReply, error) {
	out := new(WalletStoreReply)
	err := c.cc.Invoke(ctx, Wallet_WalletStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) WalletDel(ctx context.Context, in *WalletDelReq, opts ...grpc.CallOption) (*WalletDelReply, error) {
	out := new(WalletDelReply)
	err := c.cc.Invoke(ctx, Wallet_WalletDel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) WalletOne(ctx context.Context, in *WalletOneReq, opts ...grpc.CallOption) (*WalletOneReply, error) {
	out := new(WalletOneReply)
	err := c.cc.Invoke(ctx, Wallet_WalletOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) WalletList(ctx context.Context, in *WalletListReq, opts ...grpc.CallOption) (*WalletListReply, error) {
	out := new(WalletListReply)
	err := c.cc.Invoke(ctx, Wallet_WalletList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServer is the server API for Wallet service.
// All implementations must embed UnimplementedWalletServer
// for forward compatibility
type WalletServer interface {
	// 钱包-创建一条数据
	WalletStore(context.Context, *WalletStoreReq) (*WalletStoreReply, error)
	// 钱包-删除多条数据
	WalletDel(context.Context, *WalletDelReq) (*WalletDelReply, error)
	// 钱包-单条数据查询
	WalletOne(context.Context, *WalletOneReq) (*WalletOneReply, error)
	// 钱包-列表数据查询
	WalletList(context.Context, *WalletListReq) (*WalletListReply, error)
	mustEmbedUnimplementedWalletServer()
}

// UnimplementedWalletServer must be embedded to have forward compatible implementations.
type UnimplementedWalletServer struct {
}

func (UnimplementedWalletServer) WalletStore(context.Context, *WalletStoreReq) (*WalletStoreReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WalletStore not implemented")
}
func (UnimplementedWalletServer) WalletDel(context.Context, *WalletDelReq) (*WalletDelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WalletDel not implemented")
}
func (UnimplementedWalletServer) WalletOne(context.Context, *WalletOneReq) (*WalletOneReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WalletOne not implemented")
}
func (UnimplementedWalletServer) WalletList(context.Context, *WalletListReq) (*WalletListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WalletList not implemented")
}
func (UnimplementedWalletServer) mustEmbedUnimplementedWalletServer() {}

// UnsafeWalletServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServer will
// result in compilation errors.
type UnsafeWalletServer interface {
	mustEmbedUnimplementedWalletServer()
}

func RegisterWalletServer(s grpc.ServiceRegistrar, srv WalletServer) {
	s.RegisterService(&Wallet_ServiceDesc, srv)
}

func _Wallet_WalletStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletStoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).WalletStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_WalletStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).WalletStore(ctx, req.(*WalletStoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_WalletDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).WalletDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_WalletDel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).WalletDel(ctx, req.(*WalletDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_WalletOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletOneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).WalletOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_WalletOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).WalletOne(ctx, req.(*WalletOneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_WalletList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).WalletList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_WalletList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).WalletList(ctx, req.(*WalletListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Wallet_ServiceDesc is the grpc.ServiceDesc for Wallet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wallet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.rpc_wallet.v1.Wallet",
	HandlerType: (*WalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WalletStore",
			Handler:    _Wallet_WalletStore_Handler,
		},
		{
			MethodName: "WalletDel",
			Handler:    _Wallet_WalletDel_Handler,
		},
		{
			MethodName: "WalletOne",
			Handler:    _Wallet_WalletOne_Handler,
		},
		{
			MethodName: "WalletList",
			Handler:    _Wallet_WalletList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_wallet/v1/wallet.proto",
}
