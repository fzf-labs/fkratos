package service

import (
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewSysPermissionService(
	logger log.Logger,
	sysPermissionUseCase *biz.SysPermissionUseCase,
) *SysPermissionService {
	l := log.NewHelper(log.With(logger, "module", "service/sysPermission"))
	return &SysPermissionService{
		log:                  l,
		sysPermissionUseCase: sysPermissionUseCase,
	}
}

type SysPermissionService struct {
	pb.UnimplementedSysPermissionServer
	log                  *log.Helper
	sysPermissionUseCase *biz.SysPermissionUseCase
}
