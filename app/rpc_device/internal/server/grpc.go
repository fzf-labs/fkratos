package server

import (
	v1 "fkratos/api/rpc_device/v1"
	"fkratos/app/rpc_device/internal/service"

	conf "github.com/fzf-labs/fkratos-contrib/api/conf/v1"

	"github.com/fzf-labs/fkratos-contrib/bootstrap"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Bootstrap, logger log.Logger, deviceService *service.DeviceService) *grpc.Server {
	srv := bootstrap.NewGrpcServer(c, logger)
	v1.RegisterDeviceServer(srv, deviceService)
	return srv
}
