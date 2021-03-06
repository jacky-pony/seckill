package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "github.com/peter-wow/seckill/api/order/service/v1"
	"github.com/peter-wow/seckill/app/order/service/internal/biz"
)



// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewOrderService)

type OrderService struct {
	pb.UnimplementedOrderServer

	uc *biz.OrderUsecase
	goods *biz.SeckillGoodsUsecase
	so *biz.SeckillOrderUsecase
	log *log.Helper
}

func NewOrderService(uc *biz.OrderUsecase, so *biz.SeckillOrderUsecase, goods *biz.SeckillGoodsUsecase, logger log.Logger) *OrderService  {
	return &OrderService{
		uc: uc,
		so: so,
		goods: goods,
		log: log.NewHelper(logger),
	}
}
