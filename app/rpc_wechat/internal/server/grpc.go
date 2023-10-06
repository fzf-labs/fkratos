package server

import (
	v1 "fkratos/api/rpc_wechat/v1"
	"fkratos/app/rpc_wechat/internal/service"
	"fkratos/internal/bootstrap"
	"fkratos/internal/bootstrap/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Bootstrap, logger log.Logger, officialAccountService *service.OfficialAccountService, miniProgramService *service.MiniProgramService) *grpc.Server {
	srv := bootstrap.NewGrpcServer(c, logger)
	v1.RegisterOfficialAccountServer(srv, officialAccountService)
	v1.RegisterMiniProgramServer(srv, miniProgramService)
	return srv
}
