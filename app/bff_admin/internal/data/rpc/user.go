package rpc

import (
	"context"
	userV1 "fkratos/api/rpc_user/v1"
	"fkratos/internal/bootstrap"
	"fkratos/internal/bootstrap/conf"
	"fkratos/internal/constant"

	"github.com/go-kratos/kratos/v2/registry"
	"google.golang.org/grpc"
)

type UserGrpc struct {
	grpcClientConn *grpc.ClientConn
}

func NewRPCUserGrpc(c *conf.Bootstrap, r registry.Discovery) *UserGrpc {
	return &UserGrpc{grpcClientConn: bootstrap.NewGrpcClient(context.Background(), r, c.Registry.Type, constant.RPCUser, c.Server.Grpc.GetTimeout())}
}

func NewUserServiceClient(rpcSysGrpc *UserGrpc) userV1.UserClient {
	return userV1.NewUserClient(rpcSysGrpc.grpcClientConn)
}
