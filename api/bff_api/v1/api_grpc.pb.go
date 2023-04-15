// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: bff_api/v1/api.proto

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
	Api_CreateApi_FullMethodName = "/api.Api.v1.Api/CreateApi"
	Api_UpdateApi_FullMethodName = "/api.Api.v1.Api/UpdateApi"
	Api_DeleteApi_FullMethodName = "/api.Api.v1.Api/DeleteApi"
	Api_GetApi_FullMethodName    = "/api.Api.v1.Api/GetApi"
	Api_ListApi_FullMethodName   = "/api.Api.v1.Api/ListApi"
)

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	// 创建用户
	CreateApi(ctx context.Context, in *CreateApiReq, opts ...grpc.CallOption) (*CreateApiReply, error)
	// 更新用户
	UpdateApi(ctx context.Context, in *UpdateApiReq, opts ...grpc.CallOption) (*UpdateApiReply, error)
	// 删除用户
	DeleteApi(ctx context.Context, in *DeleteApiReq, opts ...grpc.CallOption) (*DeleteApiReply, error)
	// 获取单个用户
	GetApi(ctx context.Context, in *GetApiReq, opts ...grpc.CallOption) (*GetApiReply, error)
	// 获取用户列表
	ListApi(ctx context.Context, in *ListApiReq, opts ...grpc.CallOption) (*ListApiReply, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) CreateApi(ctx context.Context, in *CreateApiReq, opts ...grpc.CallOption) (*CreateApiReply, error) {
	out := new(CreateApiReply)
	err := c.cc.Invoke(ctx, Api_CreateApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) UpdateApi(ctx context.Context, in *UpdateApiReq, opts ...grpc.CallOption) (*UpdateApiReply, error) {
	out := new(UpdateApiReply)
	err := c.cc.Invoke(ctx, Api_UpdateApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) DeleteApi(ctx context.Context, in *DeleteApiReq, opts ...grpc.CallOption) (*DeleteApiReply, error) {
	out := new(DeleteApiReply)
	err := c.cc.Invoke(ctx, Api_DeleteApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetApi(ctx context.Context, in *GetApiReq, opts ...grpc.CallOption) (*GetApiReply, error) {
	out := new(GetApiReply)
	err := c.cc.Invoke(ctx, Api_GetApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ListApi(ctx context.Context, in *ListApiReq, opts ...grpc.CallOption) (*ListApiReply, error) {
	out := new(ListApiReply)
	err := c.cc.Invoke(ctx, Api_ListApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	// 创建用户
	CreateApi(context.Context, *CreateApiReq) (*CreateApiReply, error)
	// 更新用户
	UpdateApi(context.Context, *UpdateApiReq) (*UpdateApiReply, error)
	// 删除用户
	DeleteApi(context.Context, *DeleteApiReq) (*DeleteApiReply, error)
	// 获取单个用户
	GetApi(context.Context, *GetApiReq) (*GetApiReply, error)
	// 获取用户列表
	ListApi(context.Context, *ListApiReq) (*ListApiReply, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) CreateApi(context.Context, *CreateApiReq) (*CreateApiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApi not implemented")
}
func (UnimplementedApiServer) UpdateApi(context.Context, *UpdateApiReq) (*UpdateApiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApi not implemented")
}
func (UnimplementedApiServer) DeleteApi(context.Context, *DeleteApiReq) (*DeleteApiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApi not implemented")
}
func (UnimplementedApiServer) GetApi(context.Context, *GetApiReq) (*GetApiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApi not implemented")
}
func (UnimplementedApiServer) ListApi(context.Context, *ListApiReq) (*ListApiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApi not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_CreateApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApiReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).CreateApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_CreateApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).CreateApi(ctx, req.(*CreateApiReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_UpdateApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApiReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).UpdateApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_UpdateApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).UpdateApi(ctx, req.(*UpdateApiReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_DeleteApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApiReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).DeleteApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_DeleteApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).DeleteApi(ctx, req.(*DeleteApiReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApiReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_GetApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetApi(ctx, req.(*GetApiReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ListApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApiReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ListApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_ListApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ListApi(ctx, req.(*ListApiReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Api.v1.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApi",
			Handler:    _Api_CreateApi_Handler,
		},
		{
			MethodName: "UpdateApi",
			Handler:    _Api_UpdateApi_Handler,
		},
		{
			MethodName: "DeleteApi",
			Handler:    _Api_DeleteApi_Handler,
		},
		{
			MethodName: "GetApi",
			Handler:    _Api_GetApi_Handler,
		},
		{
			MethodName: "ListApi",
			Handler:    _Api_ListApi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bff_api/v1/api.proto",
}
