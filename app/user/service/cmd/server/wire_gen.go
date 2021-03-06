// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/peter-wow/seckill/app/user/service/internal/biz"
	"github.com/peter-wow/seckill/app/user/service/internal/conf"
	"github.com/peter-wow/seckill/app/user/service/internal/data"
	"github.com/peter-wow/seckill/app/user/service/internal/server"
	"github.com/peter-wow/seckill/app/user/service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, logger)
	userService := service.NewUserService(userUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, userService, logger)
	grpcServer := server.NewGRPCServer(confServer, userService, logger)
	registry := server.NewRegistrar()
	app := newApp(logger, httpServer, grpcServer, registry)
	return app, func() {
		cleanup()
	}, nil
}
