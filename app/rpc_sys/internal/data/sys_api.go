package data

import (
	"context"
	"fkratos/app/rpc_sys/internal/biz"
	"fkratos/app/rpc_sys/internal/data/gorm/rpc_sys_dao"
	"fkratos/app/rpc_sys/internal/data/gorm/rpc_sys_model"
	"fkratos/internal/errorx"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.SysApiRepo = (*SysApiRepo)(nil)

func NewSysApiRepo(data *Data, logger log.Logger) biz.SysApiRepo {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/data/sys_log"))
	return &SysApiRepo{
		data: data,
		log:  l,
	}
}

type SysApiRepo struct {
	data *Data
	log  *log.Helper
}

func (s *SysApiRepo) SysApiDel(ctx context.Context, ids []string) error {
	sysAPIDao := rpc_sys_dao.Use(s.data.gorm).SysAPI
	_, err := sysAPIDao.WithContext(ctx).Where(sysAPIDao.ID.In(ids...)).Delete()
	if err != nil {
		return errorx.DataSqlErr.WithCause(err)
	}
	return nil
}

func (s *SysApiRepo) SysApiStore(ctx context.Context, model *rpc_sys_model.SysAPI) (*rpc_sys_model.SysAPI, error) {
	sysAPIDao := rpc_sys_dao.Use(s.data.gorm).SysAPI
	err := sysAPIDao.WithContext(ctx).Create(model)
	if err != nil {
		return nil, errorx.DataSqlErr.WithCause(err)
	}
	return nil, err
}

func (s *SysApiRepo) SysApiList(ctx context.Context, permissionId string) ([]*rpc_sys_model.SysAPI, error) {
	sysAPIDao := rpc_sys_dao.Use(s.data.gorm).SysAPI
	sysAPIS, err := sysAPIDao.WithContext(ctx).Where(sysAPIDao.PermissionID.Eq(permissionId)).Find()
	if err != nil {
		return nil, errorx.DataSqlErr.WithCause(err)
	}
	return sysAPIS, nil
}

func (s *SysApiRepo) GetApiIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error) {
	resp := make(map[string]string)
	sysAPIDao := rpc_sys_dao.Use(s.data.gorm).SysAPI
	sysAPIS, err := sysAPIDao.WithContext(ctx).Where(sysAPIDao.ID.In(ids...)).Find()
	if err != nil {
		return nil, errorx.DataSqlErr.WithCause(err)
	}
	for _, v := range sysAPIS {
		resp[v.ID] = v.Desc
	}
	return resp, nil
}
