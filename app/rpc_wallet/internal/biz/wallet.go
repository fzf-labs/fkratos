package biz

import "github.com/go-kratos/kratos/v2/log"

type WalletRepo interface {
}

func NewWalletUseCase(
	logger log.Logger,
	walletRepo WalletRepo,
) *WalletUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/wallet"))
	return &WalletUseCase{
		log:        l,
		walletRepo: walletRepo,
	}
}

type WalletUseCase struct {
	log        *log.Helper
	walletRepo WalletRepo
}
