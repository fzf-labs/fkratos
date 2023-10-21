package service

import (
	pb "fkratos/api/rpc_wallet/v1"
	"fkratos/app/rpc_wallet/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewWalletService(
	logger log.Logger,
	walletUseCase *biz.WalletUseCase,
) *WalletService {
	l := log.NewHelper(log.With(logger, "module", "service/wallet"))
	return &WalletService{
		log:           l,
		walletUseCase: walletUseCase,
	}
}

type WalletService struct {
	pb.UnimplementedWalletServer
	log           *log.Helper
	walletUseCase *biz.WalletUseCase
}
