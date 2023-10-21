package service

import (
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewSysAuthService(
	logger log.Logger,
	sysAuthUseCase *biz.SysAuthUseCase,
) *SysAuthService {
	l := log.NewHelper(log.With(logger, "module", "service/sysAuth"))
	return &SysAuthService{
		log:            l,
		sysAuthUseCase: sysAuthUseCase,
	}
}

type SysAuthService struct {
	pb.UnimplementedSysAuthServer
	log            *log.Helper
	sysAuthUseCase *biz.SysAuthUseCase
}
