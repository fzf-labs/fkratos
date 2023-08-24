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

type RPCSysGrpc struct {
	grpcClientConn *grpc.ClientConn
}

func NewRPCSysGrpc(c *conf.Bootstrap, r registry.Discovery) *RPCSysGrpc {
	return &RPCSysGrpc{grpcClientConn: bootstrap.NewGrpcClient(context.Background(), r, c.Registry.Type, constant.RPCSys, c.Server.Grpc.GetTimeout())}
}

func NewSysAdminServiceClient(rpcSysGrpc *RPCSysGrpc) sysV1.AdminClient {
	return sysV1.NewAdminClient(rpcSysGrpc.grpcClientConn)
}
func NewSysAuthServiceClient(rpcSysGrpc *RPCSysGrpc) sysV1.AuthClient {
	return sysV1.NewAuthClient(rpcSysGrpc.grpcClientConn)
}
func NewSysDashboardServiceClient(rpcSysGrpc *RPCSysGrpc) sysV1.DashboardClient {
	return sysV1.NewDashboardClient(rpcSysGrpc.grpcClientConn)
}
func NewSysRoleServiceClient(rpcSysGrpc *RPCSysGrpc) sysV1.RoleClient {
	return sysV1.NewRoleClient(rpcSysGrpc.grpcClientConn)
}
func NewSysPermissionServiceClient(rpcSysGrpc *RPCSysGrpc) sysV1.PermissionClient {
	return sysV1.NewPermissionClient(rpcSysGrpc.grpcClientConn)
}
func NewSysAPIServiceClient(rpcSysGrpc *RPCSysGrpc) sysV1.APIClient {
	return sysV1.NewAPIClient(rpcSysGrpc.grpcClientConn)
}
func NewSysLogServiceClient(rpcSysGrpc *RPCSysGrpc) sysV1.LogClient {
	return sysV1.NewLogClient(rpcSysGrpc.grpcClientConn)
}
func NewSysJobServiceClient(rpcSysGrpc *RPCSysGrpc) sysV1.JobClient {
	return sysV1.NewJobClient(rpcSysGrpc.grpcClientConn)
}
func NewSysDeptServiceClient(rpcSysGrpc *RPCSysGrpc) sysV1.DeptClient {
	return sysV1.NewDeptClient(rpcSysGrpc.grpcClientConn)
}
