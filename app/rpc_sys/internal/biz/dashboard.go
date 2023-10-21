package biz

import "github.com/go-kratos/kratos/v2/log"

type DashboardRepo interface {
}

func NewDashboardUseCase(
	logger log.Logger,
) *DashboardUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/dashboard"))
	return &DashboardUseCase{
		log: l,
	}
}

type DashboardUseCase struct {
	log *log.Helper
}
