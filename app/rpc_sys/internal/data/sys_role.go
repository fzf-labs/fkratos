package data

import (
	"context"
	"fkratos/app/rpc_sys/internal/biz"
	"fkratos/app/rpc_sys/internal/data/gorm/rpc_sys_dao"
	"fkratos/internal/bootstrap/conf"
	"fkratos/internal/errorx"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.SysRoleRepo = (*SysRoleRepo)(nil)

func NewSysRoleRepo(data *Data, logger log.Logger) biz.SysRoleRepo {
	l := log.NewHelper(log.With(logger, "module", "auth/repo/auth-service"))
	return &SysRoleRepo{
		data: data,
		log:  l,
	}
}

type SysRoleRepo struct {
	config *conf.Bootstrap
	data   *Data
	log    *log.Helper
}

func (s *SysRoleRepo) GetRoleIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error) {
	resp := make(map[string]string)
	dao := rpc_sys_dao.Use(s.data.gorm).SysRole
	results, err := dao.WithContext(ctx).Where(dao.ID.In(ids...)).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithCause(err)
	}
	for _, v := range results {
		resp[v.ID] = v.Name
	}
	return resp, nil
}
