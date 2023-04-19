package server

import (
	v1 "fkratos/api/rpc_user/v1"
	"fkratos/app/rpc_user/internal/service"
	"fkratos/internal/bootstrap"
	"fkratos/internal/bootstrap/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Bootstrap, logger log.Logger, userService *service.UserService) *grpc.Server {
	//创建grpc服务
	srv := bootstrap.NewGrpcServer(c, logging.Server(logger))

	//注册服务
	v1.RegisterUserServer(srv, userService)
	return srv
}
