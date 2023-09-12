package data

import (
	"context"
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_dao"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"
	"fkratos/internal/errorx"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.SysPermissionRepo = (*SysPermissionRepo)(nil)

func NewSysPermissionRepo(data *Data, logger log.Logger) biz.SysPermissionRepo {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/data/sys_permission"))
	return &SysPermissionRepo{
		data:              data,
		log:               l,
		SysPermissionRepo: fkratos_sys_repo.NewSysPermissionRepo(data.gorm, data.dbCache),
	}
}

type SysPermissionRepo struct {
	data *Data
	log  *log.Helper
	*fkratos_sys_repo.SysPermissionRepo
}

func (s *SysPermissionRepo) SysPermissionByStatus(ctx context.Context, status int16) ([]*fkratos_sys_model.SysPermission, error) {
	sysPermissionDao := fkratos_sys_dao.Use(s.data.gorm).SysPermission
	res, err := sysPermissionDao.WithContext(ctx).Where(sysPermissionDao.Status.Eq(status)).Find()
	if err != nil {
		return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return res, nil
}

func (s *SysPermissionRepo) SysPermissionByIdsAndStatus(ctx context.Context, ids []string, status int16) ([]*fkratos_sys_model.SysPermission, error) {
	sysPermissionDao := fkratos_sys_dao.Use(s.data.gorm).SysPermission
	res, err := sysPermissionDao.WithContext(ctx).Where(sysPermissionDao.ID.In(ids...), sysPermissionDao.Status.Eq(status)).Find()
	if err != nil {
		return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return res, nil
}

func (s *SysPermissionRepo) SysPermissionUpdateStatus(ctx context.Context, id string, status int32) error {
	sysPermMenuDao := fkratos_sys_dao.Use(s.data.gorm).SysPermission
	_, err := sysPermMenuDao.WithContext(ctx).Where(sysPermMenuDao.ID.Eq(id)).UpdateColumn(sysPermMenuDao.Status, status)
	if err != nil {
		return errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return nil
}

func (s *SysPermissionRepo) SysPermissionList(ctx context.Context) ([]*fkratos_sys_model.SysPermission, error) {
	sysPermMenuDao := fkratos_sys_dao.Use(s.data.gorm).SysPermission
	sysPermMenus, err := sysPermMenuDao.WithContext(ctx).Find()
	if err != nil {
		return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return sysPermMenus, nil
}

func (s *SysPermissionRepo) SysPermissionInfoByID(ctx context.Context, id string) (*fkratos_sys_model.SysPermission, error) {
	sysPermMenuDao := fkratos_sys_dao.Use(s.data.gorm).SysPermission
	sysPermMenu, err := sysPermMenuDao.WithContext(ctx).Where(sysPermMenuDao.ID.Eq(id)).First()
	if err != nil {
		return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return sysPermMenu, nil
}

func (s *SysPermissionRepo) SysPermissionDelByIds(ctx context.Context, ids []string) error {
	sysPermMenuDao := fkratos_sys_dao.Use(s.data.gorm).SysPermission
	_, err := sysPermMenuDao.WithContext(ctx).Where(sysPermMenuDao.ID.In(ids...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (s *SysPermissionRepo) SysPermissionStore(ctx context.Context, req *v1.SysPermissionStoreReq) (*fkratos_sys_model.SysPermission, error) {
	sysPermMenuDao := fkratos_sys_dao.Use(s.data.gorm).SysPermission
	model := &fkratos_sys_model.SysPermission{
		Pid:       req.Pid,
		Type:      req.Type,
		Title:     req.Title,
		Name:      req.Name,
		Path:      req.Path,
		Icon:      req.Icon,
		MenuType:  req.MenuType,
		URL:       req.Url,
		Component: req.Component,
		Extend:    req.Extend,
		Remark:    req.Remark,
		Sort:      int64(req.Sort),
		Status:    int16(req.Status),
	}
	if req.Id != "" {
		_, err := sysPermMenuDao.WithContext(ctx).Where(sysPermMenuDao.ID.Eq(req.Id)).Select(
			sysPermMenuDao.Pid,
			sysPermMenuDao.Type,
			sysPermMenuDao.Pid,
			sysPermMenuDao.Title,
			sysPermMenuDao.Name,
			sysPermMenuDao.Path,
			sysPermMenuDao.Icon,
			sysPermMenuDao.MenuType,
			sysPermMenuDao.URL,
			sysPermMenuDao.Component,
			sysPermMenuDao.Extend,
			sysPermMenuDao.Remark,
			sysPermMenuDao.Sort,
			sysPermMenuDao.Status,
		).Updates(model)
		if err != nil {
			return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
	} else {
		err := sysPermMenuDao.WithContext(ctx).Create(model)
		if err != nil {
			return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
	}
	return model, nil
}
