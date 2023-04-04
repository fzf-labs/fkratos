// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.19.4
// source: demo/v1/demo.proto

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

const OperationDemoCreateDemo = "/api.demo.v1.Demo/CreateDemo"
const OperationDemoDeleteDemo = "/api.demo.v1.Demo/DeleteDemo"
const OperationDemoGetDemo = "/api.demo.v1.Demo/GetDemo"
const OperationDemoListDemo = "/api.demo.v1.Demo/ListDemo"
const OperationDemoUpdateDemo = "/api.demo.v1.Demo/UpdateDemo"

type DemoHTTPServer interface {
	CreateDemo(context.Context, *CreateDemoRequest) (*CreateDemoReply, error)
	DeleteDemo(context.Context, *DeleteDemoRequest) (*DeleteDemoReply, error)
	GetDemo(context.Context, *GetDemoRequest) (*GetDemoReply, error)
	ListDemo(context.Context, *ListDemoRequest) (*ListDemoReply, error)
	UpdateDemo(context.Context, *UpdateDemoRequest) (*UpdateDemoReply, error)
}

func RegisterDemoHTTPServer(s *http.Server, srv DemoHTTPServer) {
	r := s.Route("/")
	r.POST("/demo/v1/create", _Demo_CreateDemo0_HTTP_Handler(srv))
	r.PUT("/demo/v1/create", _Demo_UpdateDemo0_HTTP_Handler(srv))
	r.DELETE("/demo/v1/delete", _Demo_DeleteDemo0_HTTP_Handler(srv))
	r.GET("/demo/v1/info", _Demo_GetDemo0_HTTP_Handler(srv))
	r.GET("/demo/v1/list", _Demo_ListDemo0_HTTP_Handler(srv))
}

func _Demo_CreateDemo0_HTTP_Handler(srv DemoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateDemoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDemoCreateDemo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateDemo(ctx, req.(*CreateDemoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateDemoReply)
		return ctx.Result(200, reply)
	}
}

func _Demo_UpdateDemo0_HTTP_Handler(srv DemoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateDemoRequest
		if err := ctx.Bind(&in.Demo); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDemoUpdateDemo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDemo(ctx, req.(*UpdateDemoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateDemoReply)
		return ctx.Result(200, reply)
	}
}

func _Demo_DeleteDemo0_HTTP_Handler(srv DemoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteDemoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDemoDeleteDemo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteDemo(ctx, req.(*DeleteDemoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteDemoReply)
		return ctx.Result(200, reply)
	}
}

func _Demo_GetDemo0_HTTP_Handler(srv DemoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDemoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDemoGetDemo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDemo(ctx, req.(*GetDemoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetDemoReply)
		return ctx.Result(200, reply)
	}
}

func _Demo_ListDemo0_HTTP_Handler(srv DemoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListDemoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDemoListDemo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDemo(ctx, req.(*ListDemoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDemoReply)
		return ctx.Result(200, reply)
	}
}

type DemoHTTPClient interface {
	CreateDemo(ctx context.Context, req *CreateDemoRequest, opts ...http.CallOption) (rsp *CreateDemoReply, err error)
	DeleteDemo(ctx context.Context, req *DeleteDemoRequest, opts ...http.CallOption) (rsp *DeleteDemoReply, err error)
	GetDemo(ctx context.Context, req *GetDemoRequest, opts ...http.CallOption) (rsp *GetDemoReply, err error)
	ListDemo(ctx context.Context, req *ListDemoRequest, opts ...http.CallOption) (rsp *ListDemoReply, err error)
	UpdateDemo(ctx context.Context, req *UpdateDemoRequest, opts ...http.CallOption) (rsp *UpdateDemoReply, err error)
}

type DemoHTTPClientImpl struct {
	cc *http.Client
}

func NewDemoHTTPClient(client *http.Client) DemoHTTPClient {
	return &DemoHTTPClientImpl{client}
}

func (c *DemoHTTPClientImpl) CreateDemo(ctx context.Context, in *CreateDemoRequest, opts ...http.CallOption) (*CreateDemoReply, error) {
	var out CreateDemoReply
	pattern := "/demo/v1/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDemoCreateDemo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DemoHTTPClientImpl) DeleteDemo(ctx context.Context, in *DeleteDemoRequest, opts ...http.CallOption) (*DeleteDemoReply, error) {
	var out DeleteDemoReply
	pattern := "/demo/v1/delete"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDemoDeleteDemo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DemoHTTPClientImpl) GetDemo(ctx context.Context, in *GetDemoRequest, opts ...http.CallOption) (*GetDemoReply, error) {
	var out GetDemoReply
	pattern := "/demo/v1/info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDemoGetDemo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DemoHTTPClientImpl) ListDemo(ctx context.Context, in *ListDemoRequest, opts ...http.CallOption) (*ListDemoReply, error) {
	var out ListDemoReply
	pattern := "/demo/v1/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDemoListDemo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DemoHTTPClientImpl) UpdateDemo(ctx context.Context, in *UpdateDemoRequest, opts ...http.CallOption) (*UpdateDemoReply, error) {
	var out UpdateDemoReply
	pattern := "/demo/v1/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDemoUpdateDemo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Demo, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
