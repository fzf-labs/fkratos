// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"fkratos/app/rpc_common/internal/biz"
	"fkratos/app/rpc_common/internal/data"
	"fkratos/app/rpc_common/internal/server"
	"fkratos/app/rpc_common/internal/service"
	"fkratos/internal/bootstrap"
	"fkratos/internal/bootstrap/conf"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confBootstrap *conf.Bootstrap, logger log.Logger, registrar registry.Registrar, discovery registry.Discovery) (*kratos.App, func(), error) {
	db := bootstrap.NewGorm(confBootstrap, logger)
	client := bootstrap.NewRedis(confBootstrap, logger)
	dataData, cleanup, err := data.NewData(confBootstrap, logger, db, client)
	if err != nil {
		return nil, nil, err
	}
	sensitiveWordRepo := data.NewSensitiveWordRepo(dataData, logger)
	sensitiveWordUseCase := biz.NewSensitiveWordUseCase(logger, sensitiveWordRepo)
	sensitiveWordService := service.NewSensitiveWordService(logger, sensitiveWordUseCase)
	grpcServer := server.NewGRPCServer(confBootstrap, logger, sensitiveWordService)
	app := newApp(logger, registrar, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
