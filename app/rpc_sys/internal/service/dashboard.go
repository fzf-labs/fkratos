package service

import (
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewDashboardService(
	logger log.Logger,
	dashboardUseCase *biz.DashboardUseCase,
) *DashboardService {
	l := log.NewHelper(log.With(logger, "module", "service/dashboard"))
	return &DashboardService{
		log:              l,
		dashboardUseCase: dashboardUseCase,
	}
}

type DashboardService struct {
	pb.UnimplementedDashboardServer
	log              *log.Helper
	dashboardUseCase *biz.DashboardUseCase
}
