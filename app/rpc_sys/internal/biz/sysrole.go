package biz

import (
	"context"
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"

	"github.com/go-kratos/kratos/v2/log"
)

type SysRoleRepo interface {
	fkratos_sys_repo.ISysRoleRepo
	GetRoleIDToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
	SysRoleList(ctx context.Context) ([]*fkratos_sys_model.SysRole, error)
	SysRoleStore(ctx context.Context, req *pb.SysRoleStoreReq) (*fkratos_sys_model.SysRole, error)
}

func NewSysRoleUseCase(
	logger log.Logger,
	sysRoleRepo SysRoleRepo,
) *SysRoleUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/sysRole"))
	return &SysRoleUseCase{
		log:         l,
		sysRoleRepo: sysRoleRepo,
	}
}

type SysRoleUseCase struct {
	log         *log.Helper
	sysRoleRepo SysRoleRepo
}
