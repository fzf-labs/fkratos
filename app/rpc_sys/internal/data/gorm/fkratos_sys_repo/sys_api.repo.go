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
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var _ ISysAPIRepo = (*SysAPIRepo)(nil)

var (
	// 缓存管理器
	cacheKeySysAPIManage = cachekey.NewKeyManage("SysAPIRepo")
	// 只针对唯一索引做缓存
	CacheSysAPIByID = cacheKeySysAPIManage.AddKey("CacheSysAPIByID", time.Hour*24, "CacheSysAPIByID")
)

type (
	ISysAPIRepo interface {
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *fkratos_sys_model.SysAPI) error
		// UpdateOne 更新一条数据
		UpdateOne(ctx context.Context, data *fkratos_sys_model.SysAPI) error
		// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
		DeleteOneCacheByID(ctx context.Context, ID string) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
		DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error
		// DeleteUniqueIndexCache 删除唯一索引存在的缓存
		DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_sys_model.SysAPI) error
		// FindMultiByTenantIDS 根据tenantIDS查询多条数据
		FindMultiByTenantIDS(ctx context.Context, tenantIDS []string) ([]*fkratos_sys_model.SysAPI, error)
		// FindOneByID 根据ID查询一条数据并设置缓存
		FindOneByID(ctx context.Context, ID string) (*fkratos_sys_model.SysAPI, error)
		// FindMultiByIDS 根据IDS查询多条数据并设置缓存
		FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_sys_model.SysAPI, error)
		// FindMultiByPermissionIDS 根据permissionIDS查询多条数据
		FindMultiByPermissionIDS(ctx context.Context, permissionIDS []string) ([]*fkratos_sys_model.SysAPI, error)
	}

	SysAPIRepo struct {
		db    *gorm.DB
		redis *redis.Client
	}
)

func NewSysAPIRepo(db *gorm.DB, redis *redis.Client) *SysAPIRepo {
	return &SysAPIRepo{db: db, redis: redis}
}

// CreateOne 创建一条数据
func (r *SysAPIRepo) CreateOne(ctx context.Context, data *fkratos_sys_model.SysAPI) error {
	dao := fkratos_sys_dao.Use(r.db).SysAPI
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne 更新一条数据
func (r *SysAPIRepo) UpdateOne(ctx context.Context, data *fkratos_sys_model.SysAPI) error {
	dao := fkratos_sys_dao.Use(r.db).SysAPI
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = r.DeleteUniqueIndexCache(ctx, []*fkratos_sys_model.SysAPI{data})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (r *SysAPIRepo) DeleteOneCacheByID(ctx context.Context, ID string) error {
	dao := fkratos_sys_dao.Use(r.db).SysAPI
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
	err = r.DeleteUniqueIndexCache(ctx, []*fkratos_sys_model.SysAPI{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (r *SysAPIRepo) DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error {
	dao := fkratos_sys_dao.Use(r.db).SysAPI
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
func (r *SysAPIRepo) DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_sys_model.SysAPI) error {
	var err error
	cacheSysAPIByID := CacheSysAPIByID.NewSingleKey(r.redis)

	for _, v := range data {
		err = cacheSysAPIByID.SingleCacheDel(ctx, cacheSysAPIByID.BuildKey(v.ID))
		if err != nil {
			return err
		}

	}
	return nil
}

// FindMultiByTenantIDS 根据tenantIDS查询多条数据
func (r *SysAPIRepo) FindMultiByTenantIDS(ctx context.Context, tenantIDS []string) ([]*fkratos_sys_model.SysAPI, error) {
	dao := fkratos_sys_dao.Use(r.db).SysAPI
	result, err := dao.WithContext(ctx).Where(dao.TenantID.In(tenantIDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindOneByID 根据ID查询一条数据并设置缓存
func (r *SysAPIRepo) FindOneByID(ctx context.Context, ID string) (*fkratos_sys_model.SysAPI, error) {
	resp := new(fkratos_sys_model.SysAPI)
	cache := CacheSysAPIByID.NewSingleKey(r.redis)
	cacheValue, err := cache.SingleCache(ctx, ID, func() (string, error) {
		dao := fkratos_sys_dao.Use(r.db).SysAPI
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

// FindMultiByIDS 根据IDS查询多条数据并设置缓存
func (r *SysAPIRepo) FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_sys_model.SysAPI, error) {
	resp := make([]*fkratos_sys_model.SysAPI, 0)
	cacheKey := CacheSysAPIByID.NewBatchKey(r.redis)
	cacheValue, err := cacheKey.BatchKeyCache(ctx, IDS, func() (map[string]string, error) {
		dao := fkratos_sys_dao.Use(r.db).SysAPI
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
			value[v.ID] = string(marshal)
		}
		return value, nil
	})
	if err != nil {
		return nil, err
	}
	for _, v := range cacheValue {
		tmp := new(fkratos_sys_model.SysAPI)
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}

// FindMultiByPermissionIDS 根据permissionIDS查询多条数据
func (r *SysAPIRepo) FindMultiByPermissionIDS(ctx context.Context, permissionIDS []string) ([]*fkratos_sys_model.SysAPI, error) {
	dao := fkratos_sys_dao.Use(r.db).SysAPI
	result, err := dao.WithContext(ctx).Where(dao.PermissionID.In(permissionIDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}