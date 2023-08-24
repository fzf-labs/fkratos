package main

import (
	"fkratos/internal/bootstrap"
	"fkratos/internal/constant"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
)

var (
	Service = bootstrap.NewService(
		constant.BffAPI,
		"1.0.0",
		"",
	)
)

func newApp(logger log.Logger, r registry.Registrar, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(Service.GetInstanceID()),
		kratos.Name(Service.Name),
		kratos.Version(Service.Version),
		kratos.Metadata(Service.Metadata),
		kratos.Logger(logger),
		kratos.Server(
			hs,
		),
		kratos.Registrar(r),
	)
}

func main() {
	cfg, logger, req, dis := bootstrap.Bootstrap(Service)
	app, cleanup, err := wireApp(cfg, logger, req, dis)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
