package data

import (
	"context"
	"fkratos/app/rpc_sys/internal/biz"
	"fkratos/app/rpc_sys/internal/data/gorm/rpc_sys_dao"
	"fkratos/app/rpc_sys/internal/data/gorm/rpc_sys_model"

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

func (s *SysApiRepo) SysDeptList(ctx context.Context, permissionId string) ([]*rpc_sys_model.SysAPI, error) {
	sysAPIDao := rpc_sys_dao.Use(s.data.gorm).SysAPI
	sysAPIS, err := sysAPIDao.WithContext(ctx).Where(sysAPIDao.PermissionID.Eq(permissionId)).Find()
	if err != nil {
		return nil, err
	}
	return sysAPIS, nil
}

func (s *SysApiRepo) GetApiIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error) {
	resp := make(map[string]string)
	sysAPIDao := rpc_sys_dao.Use(s.data.gorm).SysAPI
	sysAPIS, err := sysAPIDao.WithContext(ctx).Where(sysAPIDao.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	for _, v := range sysAPIS {
		resp[v.ID] = v.Desc
	}
	return resp, nil
}
