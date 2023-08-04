// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.

package fkratos_sys_repo

import (
	"context"
	"encoding/json"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_dao"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"time"

	"github.com/fzf-labs/fpkg/cache/cachekey"
	"github.com/fzf-labs/fpkg/conv"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ ISysPermissionRepo = (*SysPermissionRepo)(nil)

var (
	// 缓存管理器
	cacheKeySysPermissionManage = cachekey.NewKeyManage("SysPermissionRepo")
	// 只针对唯一索引做缓存
	CacheSysPermissionByID = cacheKeySysPermissionManage.AddKey("CacheSysPermissionByID", time.Hour*24, "CacheSysPermissionByID")
)

type (
	ISysPermissionRepo interface {
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *fkratos_sys_model.SysPermission) error
		// UpdateOne 更新一条数据
		UpdateOne(ctx context.Context, data *fkratos_sys_model.SysPermission) error
		// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
		DeleteOneCacheByID(ctx context.Context, ID string) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
		DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error
		// DeleteUniqueIndexCache 删除唯一索引存在的缓存
		DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_sys_model.SysPermission) error
		// FindOneCacheByID 根据ID查询一条数据并设置缓存
		FindOneCacheByID(ctx context.Context, ID string) (*fkratos_sys_model.SysPermission, error)
		// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
		FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*fkratos_sys_model.SysPermission, error)
	}

	SysPermissionRepo struct {
		db    *gorm.DB
		redis *redis.Client
	}
)

func NewSysPermissionRepo(db *gorm.DB, redis *redis.Client) *SysPermissionRepo {
	return &SysPermissionRepo{db: db, redis: redis}
}

// CreateOne 创建一条数据
func (r *SysPermissionRepo) CreateOne(ctx context.Context, data *fkratos_sys_model.SysPermission) error {
	dao := fkratos_sys_dao.Use(r.db).SysPermission
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne 更新一条数据
func (r *SysPermissionRepo) UpdateOne(ctx context.Context, data *fkratos_sys_model.SysPermission) error {
	dao := fkratos_sys_dao.Use(r.db).SysPermission
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = r.DeleteUniqueIndexCache(ctx, []*fkratos_sys_model.SysPermission{data})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (r *SysPermissionRepo) DeleteOneCacheByID(ctx context.Context, ID string) error {
	dao := fkratos_sys_dao.Use(r.db).SysPermission
	first, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if first == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	err = r.DeleteUniqueIndexCache(ctx, []*fkratos_sys_model.SysPermission{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (r *SysPermissionRepo) DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error {
	dao := fkratos_sys_dao.Use(r.db).SysPermission
	list, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return err
	}
	if len(list) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	err = r.DeleteUniqueIndexCache(ctx, list)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUniqueIndexCache 删除唯一索引存在的缓存
func (r *SysPermissionRepo) DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_sys_model.SysPermission) error {
	var err error
	cacheSysPermissionByID := CacheSysPermissionByID.NewSingleKey(r.redis)

	for _, v := range data {
		err = cacheSysPermissionByID.SingleCacheDel(ctx, cacheSysPermissionByID.BuildKey(v.ID))
		if err != nil {
			return err
		}

	}
	return nil
}

// FindOneCacheByID 根据ID查询一条数据并设置缓存
func (r *SysPermissionRepo) FindOneCacheByID(ctx context.Context, ID string) (*fkratos_sys_model.SysPermission, error) {
	resp := new(fkratos_sys_model.SysPermission)
	cache := CacheSysPermissionByID.NewSingleKey(r.redis)
	cacheValue, err := cache.SingleCache(ctx, conv.String(ID), func() (string, error) {
		dao := fkratos_sys_dao.Use(r.db).SysPermission
		result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
		if err != nil && err != gorm.ErrRecordNotFound {
			return "", err
		}
		marshal, err := json.Marshal(result)
		if err != nil {
			return "", err
		}
		return string(marshal), nil
	})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(cacheValue), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
func (r *SysPermissionRepo) FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*fkratos_sys_model.SysPermission, error) {
	resp := make([]*fkratos_sys_model.SysPermission, 0)
	cacheKey := CacheSysPermissionByID.NewBatchKey(r.redis)
	batchKeys := make([]string, 0)
	for _, v := range IDS {
		batchKeys = append(batchKeys, conv.String(v))
	}
	cacheValue, err := cacheKey.BatchKeyCache(ctx, batchKeys, func() (map[string]string, error) {
		dao := fkratos_sys_dao.Use(r.db).SysPermission
		result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
		value := make(map[string]string)
		for _, v := range result {
			marshal, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			value[conv.String(v.ID)] = string(marshal)
		}
		return value, nil
	})
	if err != nil {
		return nil, err
	}
	for _, v := range cacheValue {
		tmp := new(fkratos_sys_model.SysPermission)
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}
