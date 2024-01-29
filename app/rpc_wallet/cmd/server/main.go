package main

import (
	"fkratos/internal/constant"

	"github.com/fzf-labs/fkratos-contrib/bootstrap"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "go.uber.org/automaxprocs"
)

var (
	Service = bootstrap.NewService(
		constant.RPCWallet,
		"1.0.0",
		"",
	)
)

func newApp(logger log.Logger, r registry.Registrar, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(Service.GetInstanceID()),
		kratos.Name(Service.Name),
		kratos.Version(Service.Version),
		kratos.Metadata(Service.Metadata),
		kratos.Logger(logger),
		kratos.Registrar(r),
		kratos.Server(
			gs,
		),
	)
}

func main() {
	cfg, logger, reg, dis := bootstrap.Bootstrap(Service)
	app, cleanup, err := wireApp(cfg, logger, reg, dis)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
