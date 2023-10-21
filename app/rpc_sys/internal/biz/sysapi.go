package biz

import (
	"context"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"

	"github.com/go-kratos/kratos/v2/log"
)

type SysAPIRepo interface {
	fkratos_sys_repo.ISysAPIRepo
	GetAPIIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
	SysAPIStore(ctx context.Context, model *fkratos_sys_model.SysAPI) (*fkratos_sys_model.SysAPI, error)
}

func NewSysAPIUseCase(
	logger log.Logger,
	sysAPIRepo SysAPIRepo,
) *SysAPIUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/sysAPI"))
	return &SysAPIUseCase{
		log:        l,
		sysAPIRepo: sysAPIRepo,
	}
}

type SysAPIUseCase struct {
	log        *log.Helper
	sysAPIRepo SysAPIRepo
}
