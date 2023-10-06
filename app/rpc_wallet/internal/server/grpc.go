package server

import (
	v1 "fkratos/api/rpc_wallet/v1"
	"fkratos/app/rpc_wallet/internal/service"
	"fkratos/internal/bootstrap"
	"fkratos/internal/bootstrap/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Bootstrap, logger log.Logger, payService *service.PayService, walletService *service.WalletService) *grpc.Server {
	srv := bootstrap.NewGrpcServer(c, logger)
	v1.RegisterPayServer(srv, payService)
	v1.RegisterWalletServer(srv, walletService)
	return srv
}
