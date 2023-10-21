package service

import (
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewSysAdminService(
	logger log.Logger,
	sysAdminUseCase *biz.SysAdminUseCase,
) *SysAdminService {
	l := log.NewHelper(log.With(logger, "module", "service/sysAdmin"))
	return &SysAdminService{
		log:             l,
		sysAdminUseCase: sysAdminUseCase,
	}
}

type SysAdminService struct {
	pb.UnimplementedSysAdminServer
	log             *log.Helper
	sysAdminUseCase *biz.SysAdminUseCase
}
