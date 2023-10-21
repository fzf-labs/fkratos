package biz

import (
	"context"
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"

	"github.com/fzf-labs/fpkg/jwt"
	"github.com/go-kratos/kratos/v2/log"
)

type SysAdminRepo interface {
	fkratos_sys_repo.ISysAdminRepo
	GenerateJwTToken(ctx context.Context, kv map[string]interface{}) (*jwt.Token, error)
	ClearJwTToken(ctx context.Context, jwtUID string) error
	SysManageStore(ctx context.Context, req *pb.SysManageStoreReq) (*fkratos_sys_model.SysAdmin, error)
	GetAdminIDToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
	SysAuthJwtTokenCheck(ctx context.Context, token string) (string, error)
}

func NewSysAdminUseCase(
	logger log.Logger,
	sysAdminRepo SysAdminRepo,
	sysRoleRepo SysRoleRepo,
	sysJobRepo SysJobRepo,
	sysDeptRepo SysDeptRepo,
	sysPermissionRepo SysPermissionRepo,
) *SysAdminUseCase {
	l := log.NewHelper(log.With(logger, "module", "rpc_user/biz/admin"))
	return &SysAdminUseCase{
		log:               l,
		sysAdminRepo:      sysAdminRepo,
		sysRoleRepo:       sysRoleRepo,
		sysJobRepo:        sysJobRepo,
		sysDeptRepo:       sysDeptRepo,
		sysPermissionRepo: sysPermissionRepo,
	}
}

type SysAdminUseCase struct {
	log               *log.Helper
	sysAdminRepo      SysAdminRepo
	sysRoleRepo       SysRoleRepo
	sysJobRepo        SysJobRepo
	sysDeptRepo       SysDeptRepo
	sysPermissionRepo SysPermissionRepo
}
