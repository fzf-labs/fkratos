package server

import (
	"fkratos/api/rpc_user/v1"
	"fkratos/app/rpc_user/internal/service"
	"fkratos/internal/bootstrap/conf"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/tracing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Bootstrap, logger log.Logger, userService *service.UserService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
		),
	}
	if c.Server.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Server.Grpc.Network))
	}
	if c.Server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Server.Grpc.Addr))
	}
	if c.Server.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Server.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterUserServer(srv, userService)
	return srv
}
