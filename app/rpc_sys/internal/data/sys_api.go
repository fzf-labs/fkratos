package data

import (
	"context"
	"fkratos/app/rpc_sys/internal/biz"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_dao"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"
	"fkratos/internal/errorx"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.SysApiRepo = (*SysApiRepo)(nil)

func NewSysApiRepo(data *Data, logger log.Logger) biz.SysApiRepo {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/data/sys_log"))
	return &SysApiRepo{
		data:       data,
		log:        l,
		SysAPIRepo: fkratos_sys_repo.NewSysAPIRepo(data.gorm, data.redis),
	}
}

type SysApiRepo struct {
	data *Data
	log  *log.Helper
	*fkratos_sys_repo.SysAPIRepo
}

func (s *SysApiRepo) SysApiStore(ctx context.Context, model *fkratos_sys_model.SysAPI) (*fkratos_sys_model.SysAPI, error) {
	sysAPIDao := fkratos_sys_dao.Use(s.data.gorm).SysAPI
	err := sysAPIDao.WithContext(ctx).Create(model)
	if err != nil {
		return nil, errorx.DataSqlErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return nil, err
}

func (s *SysApiRepo) GetApiIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error) {
	resp := make(map[string]string)
	sysAPIDao := fkratos_sys_dao.Use(s.data.gorm).SysAPI
	sysAPIS, err := sysAPIDao.WithContext(ctx).Where(sysAPIDao.ID.In(ids...)).Find()
	if err != nil {
		return nil, errorx.DataSqlErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	for _, v := range sysAPIS {
		resp[v.ID] = v.Desc
	}
	return resp, nil
}
