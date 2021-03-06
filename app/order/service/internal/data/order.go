package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/peter-wow/seckill/app/order/service/internal/biz"
)

var _ biz.OrderRepo = (*orderRepo)(nil)

type orderRepo struct {
	data *Data
	log *log.Helper
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (o orderRepo) CreateOrder(ctx context.Context, order *biz.Order) error {

	res, err := o.data.db.Order.Create().SetGid(order.Gid).SetSn("2222").SetUID(333).Save(ctx)

	if err != nil {
		return err
	}

	o.log.Infof("order-create-result: %v", res)
	return nil
}