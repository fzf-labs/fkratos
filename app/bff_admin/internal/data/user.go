package data

import (
	"context"
	userV1 "fkratos/api/rpc_user/v1"
	"fkratos/internal/bootstrap"
	"fkratos/internal/bootstrap/conf"
	"fkratos/internal/constant"

	"github.com/go-kratos/kratos/v2/registry"
	"google.golang.org/grpc"
)

type RpcUserGrpc struct {
	grpcClientConn *grpc.ClientConn
}

func NewRpcUserGrpc(c *conf.Bootstrap, r registry.Discovery) *RpcUserGrpc {
	return &RpcUserGrpc{grpcClientConn: bootstrap.NewGrpcClient(context.Background(), r, c.Registry.Type, constant.RpcUser, c.Server.Grpc.GetTimeout())}
}

func NewUserServiceClient(rpcSysGrpc *RpcUserGrpc) userV1.UserClient {
	return userV1.NewUserClient(rpcSysGrpc.grpcClientConn)
}
