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

type RPCUserGrpc struct {
	grpcClientConn *grpc.ClientConn
}

func NewRPCUserGrpc(c *conf.Bootstrap, r registry.Discovery) *RPCUserGrpc {
	return &RPCUserGrpc{grpcClientConn: bootstrap.NewGrpcClient(context.Background(), r, c.Registry.Type, constant.RPCUser, c.Server.Grpc.GetTimeout())}
}

func NewUserServiceClient(rpcSysGrpc *RPCUserGrpc) userV1.UserClient {
	return userV1.NewUserClient(rpcSysGrpc.grpcClientConn)
}
