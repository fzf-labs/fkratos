package rpc

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/constant"

	conf "github.com/fzf-labs/fkratos-contrib/api/conf/v1"
	"github.com/fzf-labs/fkratos-contrib/bootstrap"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"google.golang.org/grpc"
)

type SysGrpc struct {
	grpcClientConn *grpc.ClientConn
}

func NewRPCSysGrpc(c *conf.Bootstrap, logger log.Logger, r registry.Discovery) *SysGrpc {
	return &SysGrpc{grpcClientConn: bootstrap.NewGrpcClient(context.Background(), c, logger, constant.RPCSys, r)}
}

func NewSysAdminServiceClient(rpcSysGrpc *SysGrpc) sysV1.SysAdminClient {
	return sysV1.NewSysAdminClient(rpcSysGrpc.grpcClientConn)
}
func NewSysAuthServiceClient(rpcSysGrpc *SysGrpc) sysV1.SysAuthClient {
	return sysV1.NewSysAuthClient(rpcSysGrpc.grpcClientConn)
}
func NewSysDashboardServiceClient(rpcSysGrpc *SysGrpc) sysV1.DashboardClient {
	return sysV1.NewDashboardClient(rpcSysGrpc.grpcClientConn)
}
func NewSysRoleServiceClient(rpcSysGrpc *SysGrpc) sysV1.SysRoleClient {
	return sysV1.NewSysRoleClient(rpcSysGrpc.grpcClientConn)
}
func NewSysPermissionServiceClient(rpcSysGrpc *SysGrpc) sysV1.SysPermissionClient {
	return sysV1.NewSysPermissionClient(rpcSysGrpc.grpcClientConn)
}
func NewSysAPIServiceClient(rpcSysGrpc *SysGrpc) sysV1.SysAPIClient {
	return sysV1.NewSysAPIClient(rpcSysGrpc.grpcClientConn)
}
func NewSysLogServiceClient(rpcSysGrpc *SysGrpc) sysV1.SysLogClient {
	return sysV1.NewSysLogClient(rpcSysGrpc.grpcClientConn)
}
func NewSysJobServiceClient(rpcSysGrpc *SysGrpc) sysV1.SysJobClient {
	return sysV1.NewSysJobClient(rpcSysGrpc.grpcClientConn)
}
func NewSysDeptServiceClient(rpcSysGrpc *SysGrpc) sysV1.SysDeptClient {
	return sysV1.NewSysDeptClient(rpcSysGrpc.grpcClientConn)
}
