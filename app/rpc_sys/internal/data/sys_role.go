package data

import (
	"context"
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_dao"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/internal/errorx"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.SysRoleRepo = (*SysRoleRepo)(nil)

func NewSysRoleRepo(data *Data, logger log.Logger) biz.SysRoleRepo {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/data/sys_role"))
	return &SysRoleRepo{
		data: data,
		log:  l,
	}
}

type SysRoleRepo struct {
	data *Data
	log  *log.Helper
}

func (s *SysRoleRepo) SysRoleInfoByIds(ctx context.Context, ids []string) ([]*fkratos_sys_model.SysRole, error) {
	sysRoleDao := fkratos_sys_dao.Use(s.data.gorm).SysRole
	sysRoles, err := sysRoleDao.WithContext(ctx).Where(sysRoleDao.ID.In(ids...)).Find()
	if err != nil {
		return nil, errorx.DataSqlErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return sysRoles, nil
}

func (s *SysRoleRepo) SysRoleList(ctx context.Context) ([]*fkratos_sys_model.SysRole, error) {
	sysRoleDao := fkratos_sys_dao.Use(s.data.gorm).SysRole
	sysRoles, err := sysRoleDao.WithContext(ctx).Find()
	if err != nil {
		return nil, errorx.DataSqlErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return sysRoles, nil
}

func (s *SysRoleRepo) SysRoleInfoById(ctx context.Context, id string) (*fkratos_sys_model.SysRole, error) {
	sysRoleDao := fkratos_sys_dao.Use(s.data.gorm).SysRole
	sysRole, err := sysRoleDao.WithContext(ctx).Where(sysRoleDao.ID.Eq(id)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return sysRole, nil
}

func (s *SysRoleRepo) SysRoleDelByIds(ctx context.Context, ids []string) error {
	sysRoleDao := fkratos_sys_dao.Use(s.data.gorm).SysRole
	_, err := sysRoleDao.WithContext(ctx).Where(sysRoleDao.ID.In(ids...)).Delete()
	if err != nil {
		return errorx.DataSqlErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return nil
}

func (s *SysRoleRepo) SysRoleStore(ctx context.Context, req *v1.SysRoleStoreReq) (*fkratos_sys_model.SysRole, error) {
	sysRoleDao := fkratos_sys_dao.Use(s.data.gorm).SysRole
	toString := strings.Join(req.GetPermissionIds(), ",")
	model := &fkratos_sys_model.SysRole{
		Pid:           req.Pid,
		Name:          req.Name,
		PermissionIds: toString,
		Remark:        req.Remark,
		Status:        int16(req.Status),
	}
	if req.Id != "" {
		_, err := sysRoleDao.WithContext(ctx).Where(sysRoleDao.ID.Eq(req.Id)).Select(sysRoleDao.Pid, sysRoleDao.Name, sysRoleDao.PermissionIds, sysRoleDao.Remark, sysRoleDao.Status).Updates(model)
		if err != nil {
			return nil, errorx.DataSqlErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
	} else {
		err := sysRoleDao.WithContext(ctx).Create(model)
		if err != nil {
			return nil, errorx.DataSqlErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
	}

	return model, nil
}

func (s *SysRoleRepo) GetRoleIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error) {
	resp := make(map[string]string)
	dao := fkratos_sys_dao.Use(s.data.gorm).SysRole
	results, err := dao.WithContext(ctx).Where(dao.ID.In(ids...)).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	for _, v := range results {
		resp[v.ID] = v.Name
	}
	return resp, nil
}
