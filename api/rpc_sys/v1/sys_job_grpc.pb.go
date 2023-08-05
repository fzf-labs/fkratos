// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: rpc_sys/v1/sys_job.proto

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
	Job_SysJobList_FullMethodName  = "/api.rpc_sys.v1.Job/SysJobList"
	Job_SysJobInfo_FullMethodName  = "/api.rpc_sys.v1.Job/SysJobInfo"
	Job_SysJobStore_FullMethodName = "/api.rpc_sys.v1.Job/SysJobStore"
	Job_SysJobDel_FullMethodName   = "/api.rpc_sys.v1.Job/SysJobDel"
)

// JobClient is the client API for Job service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobClient interface {
	// 岗位-列表
	SysJobList(ctx context.Context, in *paginator.PaginatorReq, opts ...grpc.CallOption) (*SysJobListReply, error)
	// 岗位-单个岗位信息
	SysJobInfo(ctx context.Context, in *SysJobInfoReq, opts ...grpc.CallOption) (*SysJobInfoReply, error)
	// 岗位-保存
	SysJobStore(ctx context.Context, in *SysJobStoreReq, opts ...grpc.CallOption) (*SysJobStoreReply, error)
	// 岗位-删除
	SysJobDel(ctx context.Context, in *SysJobDelReq, opts ...grpc.CallOption) (*SysJobDelReply, error)
}

type jobClient struct {
	cc grpc.ClientConnInterface
}

func NewJobClient(cc grpc.ClientConnInterface) JobClient {
	return &jobClient{cc}
}

func (c *jobClient) SysJobList(ctx context.Context, in *paginator.PaginatorReq, opts ...grpc.CallOption) (*SysJobListReply, error) {
	out := new(SysJobListReply)
	err := c.cc.Invoke(ctx, Job_SysJobList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) SysJobInfo(ctx context.Context, in *SysJobInfoReq, opts ...grpc.CallOption) (*SysJobInfoReply, error) {
	out := new(SysJobInfoReply)
	err := c.cc.Invoke(ctx, Job_SysJobInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) SysJobStore(ctx context.Context, in *SysJobStoreReq, opts ...grpc.CallOption) (*SysJobStoreReply, error) {
	out := new(SysJobStoreReply)
	err := c.cc.Invoke(ctx, Job_SysJobStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) SysJobDel(ctx context.Context, in *SysJobDelReq, opts ...grpc.CallOption) (*SysJobDelReply, error) {
	out := new(SysJobDelReply)
	err := c.cc.Invoke(ctx, Job_SysJobDel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServer is the server API for Job service.
// All implementations must embed UnimplementedJobServer
// for forward compatibility
type JobServer interface {
	// 岗位-列表
	SysJobList(context.Context, *paginator.PaginatorReq) (*SysJobListReply, error)
	// 岗位-单个岗位信息
	SysJobInfo(context.Context, *SysJobInfoReq) (*SysJobInfoReply, error)
	// 岗位-保存
	SysJobStore(context.Context, *SysJobStoreReq) (*SysJobStoreReply, error)
	// 岗位-删除
	SysJobDel(context.Context, *SysJobDelReq) (*SysJobDelReply, error)
	mustEmbedUnimplementedJobServer()
}

// UnimplementedJobServer must be embedded to have forward compatible implementations.
type UnimplementedJobServer struct {
}

func (UnimplementedJobServer) SysJobList(context.Context, *paginator.PaginatorReq) (*SysJobListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysJobList not implemented")
}
func (UnimplementedJobServer) SysJobInfo(context.Context, *SysJobInfoReq) (*SysJobInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysJobInfo not implemented")
}
func (UnimplementedJobServer) SysJobStore(context.Context, *SysJobStoreReq) (*SysJobStoreReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysJobStore not implemented")
}
func (UnimplementedJobServer) SysJobDel(context.Context, *SysJobDelReq) (*SysJobDelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysJobDel not implemented")
}
func (UnimplementedJobServer) mustEmbedUnimplementedJobServer() {}

// UnsafeJobServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobServer will
// result in compilation errors.
type UnsafeJobServer interface {
	mustEmbedUnimplementedJobServer()
}

func RegisterJobServer(s grpc.ServiceRegistrar, srv JobServer) {
	s.RegisterService(&Job_ServiceDesc, srv)
}

func _Job_SysJobList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(paginator.PaginatorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).SysJobList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_SysJobList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).SysJobList(ctx, req.(*paginator.PaginatorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_SysJobInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysJobInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).SysJobInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_SysJobInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).SysJobInfo(ctx, req.(*SysJobInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_SysJobStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysJobStoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).SysJobStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_SysJobStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).SysJobStore(ctx, req.(*SysJobStoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_SysJobDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysJobDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).SysJobDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_SysJobDel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).SysJobDel(ctx, req.(*SysJobDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Job_ServiceDesc is the grpc.ServiceDesc for Job service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Job_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.rpc_sys.v1.Job",
	HandlerType: (*JobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SysJobList",
			Handler:    _Job_SysJobList_Handler,
		},
		{
			MethodName: "SysJobInfo",
			Handler:    _Job_SysJobInfo_Handler,
		},
		{
			MethodName: "SysJobStore",
			Handler:    _Job_SysJobStore_Handler,
		},
		{
			MethodName: "SysJobDel",
			Handler:    _Job_SysJobDel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_sys/v1/sys_job.proto",
}
