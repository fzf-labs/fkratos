// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"fkratos/app/rpc_wechat/internal/biz"
	"fkratos/app/rpc_wechat/internal/data"
	"fkratos/app/rpc_wechat/internal/server"
	"fkratos/app/rpc_wechat/internal/service"
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
	idbCache := data.NewDBCache(client)
	dataData, cleanup, err := data.NewData(confBootstrap, logger, db, client, idbCache)
	if err != nil {
		return nil, nil, err
	}
	officialAccountRepo := data.NewOfficialAccountRepo(logger, dataData)
	officialAccountUseCase := biz.NewOfficialAccountUseCase(logger, officialAccountRepo)
	officialAccountService := service.NewOfficialAccountService(logger, officialAccountUseCase)
	miniProgramRepo := data.NewMiniProgramRepo(logger, dataData)
	miniProgramUseCase := biz.NewMiniProgramUseCase(logger, miniProgramRepo)
	miniProgramService := service.NewMiniProgramService(logger, miniProgramUseCase)
	grpcServer := server.NewGRPCServer(confBootstrap, logger, officialAccountService, miniProgramService)
	app := newApp(logger, registrar, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
