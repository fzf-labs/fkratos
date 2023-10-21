package service

import (
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewSysJobService(
	logger log.Logger,
	sysJobUseCase *biz.SysJobUseCase,
) *SysJobService {
	l := log.NewHelper(log.With(logger, "module", "service/sysJob"))
	return &SysJobService{
		log:           l,
		sysJobUseCase: sysJobUseCase,
	}
}

type SysJobService struct {
	pb.UnimplementedSysJobServer
	log           *log.Helper
	sysJobUseCase *biz.SysJobUseCase
}
