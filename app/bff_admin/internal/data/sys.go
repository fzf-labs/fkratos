package data

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/bootstrap"
	"fkratos/internal/bootstrap/conf"
	"fkratos/internal/constant"

	"github.com/go-kratos/kratos/v2/registry"
	"google.golang.org/grpc"
)

type RpcSysGrpc struct {
	grpcClientConn *grpc.ClientConn
}

func NewRpcSysGrpc(c *conf.Bootstrap, r registry.Discovery) *RpcSysGrpc {
	return &RpcSysGrpc{grpcClientConn: bootstrap.NewGrpcClient(context.Background(), r, c.Registry.Type, constant.RpcSys, c.Server.Grpc.GetTimeout())}
}

func NewSysAdminServiceClient(rpcSysGrpc *RpcSysGrpc) sysV1.AdminClient {
	return sysV1.NewAdminClient(rpcSysGrpc.grpcClientConn)
}
func NewSysAuthServiceClient(rpcSysGrpc *RpcSysGrpc) sysV1.AuthClient {
	return sysV1.NewAuthClient(rpcSysGrpc.grpcClientConn)
}
func NewSysDashboardServiceClient(rpcSysGrpc *RpcSysGrpc) sysV1.DashboardClient {
	return sysV1.NewDashboardClient(rpcSysGrpc.grpcClientConn)
}
func NewSysRoleServiceClient(rpcSysGrpc *RpcSysGrpc) sysV1.RoleClient {
	return sysV1.NewRoleClient(rpcSysGrpc.grpcClientConn)
}
func NewSysPermissionServiceClient(rpcSysGrpc *RpcSysGrpc) sysV1.PermissionClient {
	return sysV1.NewPermissionClient(rpcSysGrpc.grpcClientConn)
}
func NewSysApiServiceClient(rpcSysGrpc *RpcSysGrpc) sysV1.ApiClient {
	return sysV1.NewApiClient(rpcSysGrpc.grpcClientConn)
}
func NewSysLogServiceClient(rpcSysGrpc *RpcSysGrpc) sysV1.LogClient {
	return sysV1.NewLogClient(rpcSysGrpc.grpcClientConn)
}
func NewSysJobServiceClient(rpcSysGrpc *RpcSysGrpc) sysV1.JobClient {
	return sysV1.NewJobClient(rpcSysGrpc.grpcClientConn)
}
func NewSysDeptServiceClient(rpcSysGrpc *RpcSysGrpc) sysV1.DeptClient {
	return sysV1.NewDeptClient(rpcSysGrpc.grpcClientConn)
}
