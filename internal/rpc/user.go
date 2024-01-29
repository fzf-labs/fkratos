package rpc

import (
	"context"
	userV1 "fkratos/api/rpc_user/v1"
	"fkratos/internal/constant"

	conf "github.com/fzf-labs/fkratos-contrib/api/conf/v1"
	"github.com/fzf-labs/fkratos-contrib/bootstrap"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"google.golang.org/grpc"
)

type UserGrpc struct {
	grpcClientConn *grpc.ClientConn
}

func NewRPCUserGrpc(c *conf.Bootstrap, logger log.Logger, r registry.Discovery) *UserGrpc {
	return &UserGrpc{grpcClientConn: bootstrap.NewGrpcClient(context.Background(), c, logger, constant.RPCUser, r)}
}
func NewUserServiceClient(rpcSysGrpc *UserGrpc) userV1.UserClient {
	return userV1.NewUserClient(rpcSysGrpc.grpcClientConn)
}
