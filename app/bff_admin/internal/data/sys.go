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
	return &RpcSysGrpc{grpcClientConn: bootstrap.NewGrpcClient(context.Background(), r, c.Registry.Type, constant.RpcUser, c.Server.Grpc.GetTimeout())}
}

func NewSysAdminServiceClient(rpcSysGrpc *RpcSysGrpc) sysV1.AdminClient {
	return sysV1.NewAdminClient(rpcSysGrpc.grpcClientConn)
}
