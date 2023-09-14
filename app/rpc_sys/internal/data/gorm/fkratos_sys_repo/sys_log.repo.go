// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.

package fkratos_sys_repo

import (
	"context"
	"encoding/json"
	"errors"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_dao"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"

	"github.com/fzf-labs/fpkg/orm"
	"gorm.io/gorm"
)

var _ ISysLogRepo = (*SysLogRepo)(nil)

var (
	cacheSysLogByIDPrefix = "DBCache:fkratos_sys:SysLogByID"
)

type (
	ISysLogRepo interface {
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *fkratos_sys_model.SysLog) error
		// CreateOneByTx 创建一条数据(事务)
		CreateOneByTx(ctx context.Context, tx *fkratos_sys_dao.Query, data *fkratos_sys_model.SysLog) error
		// CreateBatch 批量创建数据
		CreateBatch(ctx context.Context, data []*fkratos_sys_model.SysLog, batchSize int) error
		// UpdateOne 更新一条数据
		UpdateOne(ctx context.Context, data *fkratos_sys_model.SysLog) error
		// UpdateOne 更新一条数据(事务)
		UpdateOneByTx(ctx context.Context, tx *fkratos_sys_dao.Query, data *fkratos_sys_model.SysLog) error
		// FindOneCacheByID 根据ID查询一条数据并设置缓存
		FindOneCacheByID(ctx context.Context, ID string) (*fkratos_sys_model.SysLog, error)
		// FindOneByID 根据ID查询一条数据
		FindOneByID(ctx context.Context, ID string) (*fkratos_sys_model.SysLog, error)
		// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
		FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*fkratos_sys_model.SysLog, error)
		// FindMultiByIDS 根据IDS查询多条数据
		FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_sys_model.SysLog, error)
		// FindMultiByPaginator 查询分页数据(通用)
		FindMultiByPaginator(ctx context.Context, params *orm.PaginatorParams) ([]*fkratos_sys_model.SysLog, int64, error)
		// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
		DeleteOneCacheByID(ctx context.Context, ID string) error
		// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
		DeleteOneCacheByIDTx(ctx context.Context, tx *fkratos_sys_dao.Query, ID string) error
		// DeleteOneByID 根据ID删除一条数据
		DeleteOneByID(ctx context.Context, ID string) error
		// DeleteOneByID 根据ID删除一条数据
		DeleteOneByIDTx(ctx context.Context, tx *fkratos_sys_dao.Query, ID string) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
		DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
		DeleteMultiCacheByIDSTx(ctx context.Context, tx *fkratos_sys_dao.Query, IDS []string) error
		// DeleteMultiByIDS 根据IDS删除多条数据
		DeleteMultiByIDS(ctx context.Context, IDS []string) error
		// DeleteMultiByIDS 根据IDS删除多条数据
		DeleteMultiByIDSTx(ctx context.Context, tx *fkratos_sys_dao.Query, IDS []string) error
		// DeleteUniqueIndexCache 删除唯一索引存在的缓存
		DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_sys_model.SysLog) error
	}
	ISysLogCache interface {
		Key(fields ...any) string
		Fetch(ctx context.Context, key string, fn func() (string, error)) (string, error)
		FetchBatch(ctx context.Context, keys []string, fn func(miss []string) (map[string]string, error)) (map[string]string, error)
		Del(ctx context.Context, key string) error
		DelBatch(ctx context.Context, keys []string) error
	}
	SysLogRepo struct {
		db    *gorm.DB
		cache ISysLogCache
	}
)

func NewSysLogRepo(db *gorm.DB, cache ISysLogCache) *SysLogRepo {
	return &SysLogRepo{
		db:    db,
		cache: cache,
	}
}

// CreateOne 创建一条数据
func (s *SysLogRepo) CreateOne(ctx context.Context, data *fkratos_sys_model.SysLog) error {
	dao := fkratos_sys_dao.Use(s.db).SysLog
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateOneByTx 创建一条数据(事务)
func (s *SysLogRepo) CreateOneByTx(ctx context.Context, tx *fkratos_sys_dao.Query, data *fkratos_sys_model.SysLog) error {
	dao := tx.SysLog
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateBatch 批量创建数据
func (s *SysLogRepo) CreateBatch(ctx context.Context, data []*fkratos_sys_model.SysLog, batchSize int) error {
	dao := fkratos_sys_dao.Use(s.db).SysLog
	err := dao.WithContext(ctx).CreateInBatches(data, batchSize)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne 更新一条数据
func (s *SysLogRepo) UpdateOne(ctx context.Context, data *fkratos_sys_model.SysLog) error {
	dao := fkratos_sys_dao.Use(s.db).SysLog
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = s.DeleteUniqueIndexCache(ctx, []*fkratos_sys_model.SysLog{data})
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneByTx 更新一条数据(事务)
func (s *SysLogRepo) UpdateOneByTx(ctx context.Context, tx *fkratos_sys_dao.Query, data *fkratos_sys_model.SysLog) error {
	dao := tx.SysLog
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = s.DeleteUniqueIndexCache(ctx, []*fkratos_sys_model.SysLog{data})
	if err != nil {
		return err
	}
	return err
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (s *SysLogRepo) DeleteOneCacheByID(ctx context.Context, ID string) error {
	dao := fkratos_sys_dao.Use(s.db).SysLog
	first, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if first == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	err = s.DeleteUniqueIndexCache(ctx, []*fkratos_sys_model.SysLog{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (s *SysLogRepo) DeleteOneCacheByIDTx(ctx context.Context, tx *fkratos_sys_dao.Query, ID string) error {
	dao := tx.SysLog
	first, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if first == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	err = s.DeleteUniqueIndexCache(ctx, []*fkratos_sys_model.SysLog{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (s *SysLogRepo) DeleteOneByID(ctx context.Context, ID string) error {
	dao := fkratos_sys_dao.Use(s.db).SysLog
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (s *SysLogRepo) DeleteOneByIDTx(ctx context.Context, tx *fkratos_sys_dao.Query, ID string) error {
	dao := tx.SysLog
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (s *SysLogRepo) DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error {
	dao := fkratos_sys_dao.Use(s.db).SysLog
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
	err = s.DeleteUniqueIndexCache(ctx, list)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (s *SysLogRepo) DeleteMultiCacheByIDSTx(ctx context.Context, tx *fkratos_sys_dao.Query, IDS []string) error {
	dao := tx.SysLog
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
	err = s.DeleteUniqueIndexCache(ctx, list)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (s *SysLogRepo) DeleteMultiByIDS(ctx context.Context, IDS []string) error {
	dao := fkratos_sys_dao.Use(s.db).SysLog
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (s *SysLogRepo) DeleteMultiByIDSTx(ctx context.Context, tx *fkratos_sys_dao.Query, IDS []string) error {
	dao := tx.SysLog
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteUniqueIndexCache 删除唯一索引存在的缓存
func (s *SysLogRepo) DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_sys_model.SysLog) error {
	keys := make([]string, 0)
	for _, v := range data {
		keys = append(keys, s.cache.Key(cacheSysLogByIDPrefix, v.ID))

	}
	err := s.cache.DelBatch(ctx, keys)
	if err != nil {
		return err
	}
	return nil
}

// FindOneCacheByID 根据ID查询一条数据并设置缓存
func (s *SysLogRepo) FindOneCacheByID(ctx context.Context, ID string) (*fkratos_sys_model.SysLog, error) {
	resp := new(fkratos_sys_model.SysLog)
	key := s.cache.Key(cacheSysLogByIDPrefix, ID)
	cacheValue, err := s.cache.Fetch(ctx, key, func() (string, error) {
		dao := fkratos_sys_dao.Use(s.db).SysLog
		result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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

// FindOneByID 根据ID查询一条数据
func (s *SysLogRepo) FindOneByID(ctx context.Context, ID string) (*fkratos_sys_model.SysLog, error) {
	dao := fkratos_sys_dao.Use(s.db).SysLog
	result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
func (s *SysLogRepo) FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*fkratos_sys_model.SysLog, error) {
	resp := make([]*fkratos_sys_model.SysLog, 0)
	keys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range IDS {
		key := s.cache.Key(cacheSysLogByIDPrefix, v)
		keys = append(keys, key)
		keyToParam[key] = v
	}
	cacheValue, err := s.cache.FetchBatch(ctx, keys, func(miss []string) (map[string]string, error) {
		params := make([]string, 0)
		for _, v := range miss {
			params = append(params, keyToParam[v])
		}
		dao := fkratos_sys_dao.Use(s.db).SysLog
		result, err := dao.WithContext(ctx).Where(dao.ID.In(params...)).Find()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		value := make(map[string]string)
		for _, v := range miss {
			value[v] = ""
		}
		for _, v := range result {
			marshal, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			value[s.cache.Key(cacheSysLogByIDPrefix, v.ID)] = string(marshal)
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

// FindMultiByIDS 根据IDS查询多条数据
func (s *SysLogRepo) FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_sys_model.SysLog, error) {
	dao := fkratos_sys_dao.Use(s.db).SysLog
	result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByPaginator 查询分页数据(通用)
func (s *SysLogRepo) FindMultiByPaginator(ctx context.Context, params *orm.PaginatorParams) ([]*fkratos_sys_model.SysLog, int64, error) {
	result := make([]*fkratos_sys_model.SysLog, 0)
	var total int64
	queryStr, args, err := params.ConvertToGormConditions()
	if err != nil {
		return nil, 0, err
	}
	err = s.db.WithContext(ctx).Model(&fkratos_sys_model.SysLog{}).Select([]string{"id"}).Where(queryStr, args...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, total, nil
	}
	query := s.db.WithContext(ctx)
	order := params.ConvertToOrder()
	if order != "" {
		query = query.Order(order)
	}
	limit, offset := params.ConvertToPage()
	err = query.Limit(limit).Offset(offset).Where(queryStr, args...).Find(&result).Error
	if err != nil {
		return nil, 0, err
	}
	return result, total, err
}
