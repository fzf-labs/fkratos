package service

import (
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewSysRoleService(
	logger log.Logger,
	sysRoleUseCase *biz.SysRoleUseCase,
) *SysRoleService {
	l := log.NewHelper(log.With(logger, "module", "service/sysRole"))
	return &SysRoleService{
		log:            l,
		sysRoleUseCase: sysRoleUseCase,
	}
}

type SysRoleService struct {
	pb.UnimplementedSysRoleServer
	log            *log.Helper
	sysRoleUseCase *biz.SysRoleUseCase
}
