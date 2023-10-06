package service

import (
	pb "fkratos/api/rpc_wallet/v1"
	"fkratos/app/rpc_wallet/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type PayService struct {
	pb.UnimplementedPayServer
	log        *log.Helper
	payUseCase *biz.PayUseCase
}

func NewPayService(logger log.Logger, payUseCase *biz.PayUseCase) *PayService {
	l := log.NewHelper(log.With(logger, "module", "service/pay"))
	return &PayService{
		log:        l,
		payUseCase: payUseCase,
	}
}
