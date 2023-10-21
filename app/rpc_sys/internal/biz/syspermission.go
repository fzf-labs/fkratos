package biz

import (
	"context"
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"

	"github.com/go-kratos/kratos/v2/log"
)

type SysPermissionRepo interface {
	fkratos_sys_repo.ISysPermissionRepo
	SysPermissionList(ctx context.Context) ([]*fkratos_sys_model.SysPermission, error)
	SysPermissionStore(ctx context.Context, req *pb.SysPermissionStoreReq) (*fkratos_sys_model.SysPermission, error)
	SysPermissionUpdateStatus(ctx context.Context, id string, status int32) error
	SysPermissionByIdsAndStatus(ctx context.Context, ids []string, status int16) ([]*fkratos_sys_model.SysPermission, error)
	SysPermissionByStatus(ctx context.Context, status int16) ([]*fkratos_sys_model.SysPermission, error)
}

func NewSysPermissionUseCase(
	logger log.Logger,
	sysPermissionRepo SysPermissionRepo,
) *SysPermissionUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/sysPermission"))
	return &SysPermissionUseCase{
		log:               l,
		sysPermissionRepo: sysPermissionRepo,
	}
}

type SysPermissionUseCase struct {
	log               *log.Helper
	sysPermissionRepo SysPermissionRepo
}
