package main

import (
	bootstrap2 "fkratos/internal/bootstrap"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "go.uber.org/automaxprocs"
)

var (
	Service = bootstrap2.NewService(
		"rpc.user",
		"1.0.0",
		"",
	)
)

func newApp(logger log.Logger, registry registry.Registrar, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(Service.GetInstanceId()),
		kratos.Name(Service.Name),
		kratos.Version(Service.Version),
		kratos.Metadata(Service.Metadata),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		kratos.Registrar(registry),
	)
}

func main() {
	cfg, logger, registry := bootstrap2.Bootstrap(Service)
	app, cleanup, err := wireApp(cfg, logger, registry)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
