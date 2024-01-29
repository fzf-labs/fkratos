package server

import (
	v1 "fkratos/api/rpc_common/v1"
	"fkratos/app/rpc_common/internal/service"

	conf "github.com/fzf-labs/fkratos-contrib/api/conf/v1"

	"github.com/fzf-labs/fkratos-contrib/bootstrap"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Bootstrap, logger log.Logger, sensitiveWordService *service.SensitiveWordService) *grpc.Server {
	srv := bootstrap.NewGrpcServer(c, logger)
	v1.RegisterSensitiveWordServer(srv, sensitiveWordService)
	return srv
}
