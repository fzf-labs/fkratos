package data

import (
	"fkratos/app/rpc_wallet/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.WalletRepo = (*WalletRepo)(nil)

func NewWalletRepo(logger log.Logger, data *Data) biz.WalletRepo {
	l := log.NewHelper(log.With(logger, "module", "data/wallet"))
	return &WalletRepo{
		log:  l,
		data: data,
	}
}

type WalletRepo struct {
	log  *log.Helper
	data *Data
}
