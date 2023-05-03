package main

import (
	"fkratos/internal/bootstrap"
	"fkratos/internal/service"
	kAsynq "fkratos/pkg/asynq"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "go.uber.org/automaxprocs"
)

var (
	Service = bootstrap.NewService(
		service.RpcSys,
		"1.0.0",
		"",
	)
)

func newApp(logger log.Logger, registry registry.Registrar, gs *grpc.Server, as *kAsynq.Server) *kratos.App {
	return kratos.New(
		kratos.ID(Service.GetInstanceId()),
		kratos.Name(Service.Name),
		kratos.Version(Service.Version),
		kratos.Metadata(Service.Metadata),
		kratos.Logger(logger),
		kratos.Registrar(registry),
		kratos.Server(
			gs,
			as,
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
