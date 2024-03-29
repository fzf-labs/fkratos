// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.19.4
// source: rpc_wallet/v1/wallet.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationWalletWalletDel = "/api.rpc_wallet.v1.Wallet/WalletDel"
const OperationWalletWalletList = "/api.rpc_wallet.v1.Wallet/WalletList"
const OperationWalletWalletOne = "/api.rpc_wallet.v1.Wallet/WalletOne"
const OperationWalletWalletStore = "/api.rpc_wallet.v1.Wallet/WalletStore"

type WalletHTTPServer interface {
	// WalletDel钱包-删除多条数据
	WalletDel(context.Context, *WalletDelReq) (*WalletDelReply, error)
	// WalletList钱包-列表数据查询
	WalletList(context.Context, *WalletListReq) (*WalletListReply, error)
	// WalletOne钱包-单条数据查询
	WalletOne(context.Context, *WalletOneReq) (*WalletOneReply, error)
	// WalletStore钱包-创建一条数据
	WalletStore(context.Context, *WalletStoreReq) (*WalletStoreReply, error)
}

func RegisterWalletHTTPServer(s *http.Server, srv WalletHTTPServer) {
	r := s.Route("/")
	r.POST("/wallet/v1/wallet_store", _Wallet_WalletStore0_HTTP_Handler(srv))
	r.POST("/wallet/v1/wallet_del", _Wallet_WalletDel0_HTTP_Handler(srv))
	r.GET("/wallet/v1/wallet_info", _Wallet_WalletOne0_HTTP_Handler(srv))
	r.POST("/wallet/v1/wallet_list", _Wallet_WalletList0_HTTP_Handler(srv))
}

func _Wallet_WalletStore0_HTTP_Handler(srv WalletHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WalletStoreReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWalletWalletStore)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WalletStore(ctx, req.(*WalletStoreReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WalletStoreReply)
		return ctx.Result(200, reply)
	}
}

func _Wallet_WalletDel0_HTTP_Handler(srv WalletHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WalletDelReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWalletWalletDel)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WalletDel(ctx, req.(*WalletDelReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WalletDelReply)
		return ctx.Result(200, reply)
	}
}

func _Wallet_WalletOne0_HTTP_Handler(srv WalletHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WalletOneReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWalletWalletOne)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WalletOne(ctx, req.(*WalletOneReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WalletOneReply)
		return ctx.Result(200, reply)
	}
}

func _Wallet_WalletList0_HTTP_Handler(srv WalletHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WalletListReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWalletWalletList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WalletList(ctx, req.(*WalletListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WalletListReply)
		return ctx.Result(200, reply)
	}
}

type WalletHTTPClient interface {
	WalletDel(ctx context.Context, req *WalletDelReq, opts ...http.CallOption) (rsp *WalletDelReply, err error)
	WalletList(ctx context.Context, req *WalletListReq, opts ...http.CallOption) (rsp *WalletListReply, err error)
	WalletOne(ctx context.Context, req *WalletOneReq, opts ...http.CallOption) (rsp *WalletOneReply, err error)
	WalletStore(ctx context.Context, req *WalletStoreReq, opts ...http.CallOption) (rsp *WalletStoreReply, err error)
}

type WalletHTTPClientImpl struct {
	cc *http.Client
}

func NewWalletHTTPClient(client *http.Client) WalletHTTPClient {
	return &WalletHTTPClientImpl{client}
}

func (c *WalletHTTPClientImpl) WalletDel(ctx context.Context, in *WalletDelReq, opts ...http.CallOption) (*WalletDelReply, error) {
	var out WalletDelReply
	pattern := "/wallet/v1/wallet_del"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWalletWalletDel))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WalletHTTPClientImpl) WalletList(ctx context.Context, in *WalletListReq, opts ...http.CallOption) (*WalletListReply, error) {
	var out WalletListReply
	pattern := "/wallet/v1/wallet_list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWalletWalletList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WalletHTTPClientImpl) WalletOne(ctx context.Context, in *WalletOneReq, opts ...http.CallOption) (*WalletOneReply, error) {
	var out WalletOneReply
	pattern := "/wallet/v1/wallet_info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationWalletWalletOne))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WalletHTTPClientImpl) WalletStore(ctx context.Context, in *WalletStoreReq, opts ...http.CallOption) (*WalletStoreReply, error) {
	var out WalletStoreReply
	pattern := "/wallet/v1/wallet_store"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWalletWalletStore))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
