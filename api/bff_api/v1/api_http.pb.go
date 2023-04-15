// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.19.4
// source: bff_api/v1/api.proto

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

const OperationApiCreateApi = "/api.Api.v1.Api/CreateApi"
const OperationApiDeleteApi = "/api.Api.v1.Api/DeleteApi"
const OperationApiGetApi = "/api.Api.v1.Api/GetApi"
const OperationApiListApi = "/api.Api.v1.Api/ListApi"
const OperationApiUpdateApi = "/api.Api.v1.Api/UpdateApi"

type ApiHTTPServer interface {
	// CreateApi创建用户
	CreateApi(context.Context, *CreateApiReq) (*CreateApiReply, error)
	// DeleteApi删除用户
	DeleteApi(context.Context, *DeleteApiReq) (*DeleteApiReply, error)
	// GetApi获取单个用户
	GetApi(context.Context, *GetApiReq) (*GetApiReply, error)
	// ListApi获取用户列表
	ListApi(context.Context, *ListApiReq) (*ListApiReply, error)
	// UpdateApi更新用户
	UpdateApi(context.Context, *UpdateApiReq) (*UpdateApiReply, error)
}

func RegisterApiHTTPServer(s *http.Server, srv ApiHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/createApi", _Api_CreateApi0_HTTP_Handler(srv))
	r.PUT("/api/v1/updateApi", _Api_UpdateApi0_HTTP_Handler(srv))
	r.DELETE("/api/v1/deleteApi", _Api_DeleteApi0_HTTP_Handler(srv))
	r.GET("/api/v1/getApi", _Api_GetApi0_HTTP_Handler(srv))
	r.GET("/api/v1/listApi", _Api_ListApi0_HTTP_Handler(srv))
}

func _Api_CreateApi0_HTTP_Handler(srv ApiHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateApiReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApiCreateApi)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateApi(ctx, req.(*CreateApiReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateApiReply)
		return ctx.Result(200, reply)
	}
}

func _Api_UpdateApi0_HTTP_Handler(srv ApiHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateApiReq
		if err := ctx.Bind(&in.ApiInfo); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApiUpdateApi)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateApi(ctx, req.(*UpdateApiReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateApiReply)
		return ctx.Result(200, reply)
	}
}

func _Api_DeleteApi0_HTTP_Handler(srv ApiHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteApiReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApiDeleteApi)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteApi(ctx, req.(*DeleteApiReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteApiReply)
		return ctx.Result(200, reply)
	}
}

func _Api_GetApi0_HTTP_Handler(srv ApiHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetApiReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApiGetApi)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetApi(ctx, req.(*GetApiReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetApiReply)
		return ctx.Result(200, reply)
	}
}

func _Api_ListApi0_HTTP_Handler(srv ApiHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListApiReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApiListApi)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListApi(ctx, req.(*ListApiReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListApiReply)
		return ctx.Result(200, reply)
	}
}

type ApiHTTPClient interface {
	CreateApi(ctx context.Context, req *CreateApiReq, opts ...http.CallOption) (rsp *CreateApiReply, err error)
	DeleteApi(ctx context.Context, req *DeleteApiReq, opts ...http.CallOption) (rsp *DeleteApiReply, err error)
	GetApi(ctx context.Context, req *GetApiReq, opts ...http.CallOption) (rsp *GetApiReply, err error)
	ListApi(ctx context.Context, req *ListApiReq, opts ...http.CallOption) (rsp *ListApiReply, err error)
	UpdateApi(ctx context.Context, req *UpdateApiReq, opts ...http.CallOption) (rsp *UpdateApiReply, err error)
}

type ApiHTTPClientImpl struct {
	cc *http.Client
}

func NewApiHTTPClient(client *http.Client) ApiHTTPClient {
	return &ApiHTTPClientImpl{client}
}

func (c *ApiHTTPClientImpl) CreateApi(ctx context.Context, in *CreateApiReq, opts ...http.CallOption) (*CreateApiReply, error) {
	var out CreateApiReply
	pattern := "/api/v1/createApi"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationApiCreateApi))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApiHTTPClientImpl) DeleteApi(ctx context.Context, in *DeleteApiReq, opts ...http.CallOption) (*DeleteApiReply, error) {
	var out DeleteApiReply
	pattern := "/api/v1/deleteApi"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationApiDeleteApi))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApiHTTPClientImpl) GetApi(ctx context.Context, in *GetApiReq, opts ...http.CallOption) (*GetApiReply, error) {
	var out GetApiReply
	pattern := "/api/v1/getApi"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationApiGetApi))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApiHTTPClientImpl) ListApi(ctx context.Context, in *ListApiReq, opts ...http.CallOption) (*ListApiReply, error) {
	var out ListApiReply
	pattern := "/api/v1/listApi"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationApiListApi))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApiHTTPClientImpl) UpdateApi(ctx context.Context, in *UpdateApiReq, opts ...http.CallOption) (*UpdateApiReply, error) {
	var out UpdateApiReply
	pattern := "/api/v1/updateApi"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationApiUpdateApi))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.ApiInfo, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
