package service

import (
	pb "fkratos/api/bff_admin/v1"
	"fkratos/app/bff_admin/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewSysService(
	logger log.Logger,
	sysUseCase *biz.SysUseCase,
) *SysService {
	l := log.NewHelper(log.With(logger, "module", "service/sys"))
	return &SysService{
		log:        l,
		sysUseCase: sysUseCase,
	}
}

type SysService struct {
	pb.UnimplementedSysServer
	log        *log.Helper
	sysUseCase *biz.SysUseCase
}
