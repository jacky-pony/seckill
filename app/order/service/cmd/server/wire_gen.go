// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/peter-wow/seckill/app/order/service/internal/biz"
	"github.com/peter-wow/seckill/app/order/service/internal/conf"
	"github.com/peter-wow/seckill/app/order/service/internal/data"
	"github.com/peter-wow/seckill/app/order/service/internal/queue/receiver"
	"github.com/peter-wow/seckill/app/order/service/internal/queue/sender"
	"github.com/peter-wow/seckill/app/order/service/internal/server"
	"github.com/peter-wow/seckill/app/order/service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	registry := server.NewRegistrar()
	dataData, cleanup, err := data.NewData(confData, logger, registry)
	if err != nil {
		return nil, nil, err
	}
	orderRepo := data.NewOrderRepo(dataData, logger)
	orderUsecase := biz.NewOrderUsecase(orderRepo, logger)
	seckillOrderRepo := data.NewSeckillOrderRepo(dataData, logger)
	seckillOrderUsecase := biz.NewSeckillOrderUsecase(seckillOrderRepo, logger)
	queue := sender.NewQueue()
	orderQueueRepo := sender.NewOrderQueueRepo(queue, logger)
	orderQueueUsecase := biz.NewOrderQueueUsecase(orderQueueRepo, logger)
	seckillGoodsRepo := data.NewSeckillGoodsRepo(dataData, logger)
	seckillGoodsUsecase := biz.NewSeckillGoodsUsecase(seckillGoodsRepo, logger)
	receiverReceiver := receiver.NewQueue()
	orderQueueReceiverRepo := receiver.NewOrderQueueReceiverRepo(receiverReceiver, logger)
	orderQueueReceiverUsecase := biz.NewOrderQueueReceiverUsecase(orderQueueReceiverRepo, logger)
	orderService := service.NewOrderService(orderUsecase, seckillOrderUsecase, orderQueueUsecase, seckillGoodsUsecase, orderQueueReceiverUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, orderService, logger)
	grpcServer := server.NewGRPCServer(confServer, orderService, logger)
	app := newApp(logger, httpServer, grpcServer, registry)
	return app, func() {
		cleanup()
	}, nil
}
