package biz

import "github.com/go-kratos/kratos/v2/log"

type PayRepo interface {
}

func NewPayUseCase(
	logger log.Logger,
	payRepo PayRepo,
) *PayUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/pay"))
	return &PayUseCase{
		log:     l,
		payRepo: payRepo,
	}
}

type PayUseCase struct {
	log     *log.Helper
	payRepo PayRepo
}
