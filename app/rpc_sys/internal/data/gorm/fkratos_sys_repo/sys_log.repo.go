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
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ ISysLogRepo = (*SysLogRepo)(nil)

var (
	// 缓存管理器
	cacheKeySysLogManage = cachekey.NewKeyManage("SysLogRepo")
	// 只针对唯一索引做缓存
	CacheSysLogByID = cacheKeySysLogManage.AddKey("CacheSysLogByID", time.Hour*24, "CacheSysLogByID")
)

type (
	ISysLogRepo interface {
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *fkratos_sys_model.SysLog) error
		// UpdateOne 更新一条数据
		UpdateOne(ctx context.Context, data *fkratos_sys_model.SysLog) error
		// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
		DeleteOneCacheByID(ctx context.Context, ID string) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
		DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error
		// DeleteUniqueIndexCache 删除唯一索引存在的缓存
		DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_sys_model.SysLog) error
		// FindMultiByTenantIDAdminID 根据TenantIDAdminID查询多条数据
		FindMultiByTenantIDAdminID(ctx context.Context, tenantID string, adminID string) ([]*fkratos_sys_model.SysLog, error)
		// FindOneByID 根据ID查询一条数据并设置缓存
		FindOneByID(ctx context.Context, ID string) (*fkratos_sys_model.SysLog, error)
		// FindMultiByIDS 根据IDS查询多条数据并设置缓存
		FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_sys_model.SysLog, error)
	}

	SysLogRepo struct {
		db    *gorm.DB
		redis *redis.Client
	}
)

func NewSysLogRepo(db *gorm.DB, redis *redis.Client) *SysLogRepo {
	return &SysLogRepo{db: db, redis: redis}
}

// CreateOne 创建一条数据
func (r *SysLogRepo) CreateOne(ctx context.Context, data *fkratos_sys_model.SysLog) error {
	dao := fkratos_sys_dao.Use(r.db).SysLog
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne 更新一条数据
func (r *SysLogRepo) UpdateOne(ctx context.Context, data *fkratos_sys_model.SysLog) error {
	dao := fkratos_sys_dao.Use(r.db).SysLog
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = r.DeleteUniqueIndexCache(ctx, []*fkratos_sys_model.SysLog{data})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (r *SysLogRepo) DeleteOneCacheByID(ctx context.Context, ID string) error {
	dao := fkratos_sys_dao.Use(r.db).SysLog
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
	err = r.DeleteUniqueIndexCache(ctx, []*fkratos_sys_model.SysLog{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (r *SysLogRepo) DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error {
	dao := fkratos_sys_dao.Use(r.db).SysLog
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
func (r *SysLogRepo) DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_sys_model.SysLog) error {
	var err error
	cacheSysLogByID := CacheSysLogByID.NewSingleKey(r.redis)

	for _, v := range data {
		err = cacheSysLogByID.SingleCacheDel(ctx, cacheSysLogByID.BuildKey(v.ID))
		if err != nil {
			return err
		}

	}
	return nil
}

// FindMultiByTenantIDAdminID 根据TenantIDAdminID查询多条数据
func (r *SysLogRepo) FindMultiByTenantIDAdminID(ctx context.Context, tenantID string, adminID string) ([]*fkratos_sys_model.SysLog, error) {
	dao := fkratos_sys_dao.Use(r.db).SysLog
	result, err := dao.WithContext(ctx).Where(dao.TenantID.Eq(tenantID), dao.AdminID.Eq(adminID)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindOneByID 根据ID查询一条数据并设置缓存
func (r *SysLogRepo) FindOneByID(ctx context.Context, ID string) (*fkratos_sys_model.SysLog, error) {
	resp := new(fkratos_sys_model.SysLog)
	cache := CacheSysLogByID.NewSingleKey(r.redis)
	cacheValue, err := cache.SingleCache(ctx, ID, func() (string, error) {
		dao := fkratos_sys_dao.Use(r.db).SysLog
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
func (r *SysLogRepo) FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_sys_model.SysLog, error) {
	resp := make([]*fkratos_sys_model.SysLog, 0)
	cacheKey := CacheSysLogByID.NewBatchKey(r.redis)
	cacheValue, err := cacheKey.BatchKeyCache(ctx, IDS, func() (map[string]string, error) {
		dao := fkratos_sys_dao.Use(r.db).SysLog
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
		tmp := new(fkratos_sys_model.SysLog)
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}
