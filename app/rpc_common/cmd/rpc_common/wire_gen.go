// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"fkratos/app/rpc_common/internal/biz"
	"fkratos/app/rpc_common/internal/data"
	"fkratos/app/rpc_common/internal/data/gorm/fkratos_common_repo"
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
	client := bootstrap.NewRueidis(confBootstrap, logger)
	dataData, cleanup, err := data.NewData(confBootstrap, logger, db, client)
	if err != nil {
		return nil, nil, err
	}
	idbCache := data.NewDBCache(client)
	sensitiveWordRepo := fkratos_common_repo.NewSensitiveWordRepo(db, idbCache)
	bizSensitiveWordRepo := data.NewSensitiveWordRepo(logger, dataData, sensitiveWordRepo)
	sensitiveWordUseCase := biz.NewSensitiveWordUseCase(logger, bizSensitiveWordRepo)
	sensitiveWordService := service.NewSensitiveWordService(logger, sensitiveWordUseCase)
	grpcServer := server.NewGRPCServer(confBootstrap, logger, sensitiveWordService)
	app := newApp(logger, registrar, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
