package data

import (
	"fkratos/app/rpc_wallet/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.PayRepo = (*PayRepo)(nil)

func NewPayRepo(
	logger log.Logger,
	data *Data,
) biz.PayRepo {
	l := log.NewHelper(log.With(logger, "module", "data/pay"))
	return &PayRepo{
		log:  l,
		data: data,
	}
}

type PayRepo struct {
	log  *log.Helper
	data *Data
}
