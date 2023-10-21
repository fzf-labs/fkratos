package service

import (
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewSysDeptService(
	logger log.Logger,
	sysDeptUseCase *biz.SysDeptUseCase,
) *SysDeptService {
	l := log.NewHelper(log.With(logger, "module", "service/sysDept"))
	return &SysDeptService{
		log:            l,
		sysDeptUseCase: sysDeptUseCase,
	}
}

type SysDeptService struct {
	pb.UnimplementedSysDeptServer
	log            *log.Helper
	sysDeptUseCase *biz.SysDeptUseCase
}
