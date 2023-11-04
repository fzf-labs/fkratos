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
	"github.com/fzf-labs/fpkg/orm/gen/cache"
	"gorm.io/gorm"
)

var _ ISysAPIRepo = (*SysAPIRepo)(nil)

var (
	cacheSysAPIByIDPrefix = "DBCache:fkratos_sys:SysAPIByID"
)

type (
	ISysAPIRepo interface {
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *fkratos_sys_model.SysAPI) error
		// CreateOneByTx 创建一条数据(事务)
		CreateOneByTx(ctx context.Context, tx *fkratos_sys_dao.Query, data *fkratos_sys_model.SysAPI) error
		// UpsertOne Upsert一条数据
		UpsertOne(ctx context.Context, data *fkratos_sys_model.SysAPI) error
		// UpsertOneByTx Upsert一条数据(事务)
		UpsertOneByTx(ctx context.Context, tx *fkratos_sys_dao.Query, data *fkratos_sys_model.SysAPI) error
		// CreateBatch 批量创建数据
		CreateBatch(ctx context.Context, data []*fkratos_sys_model.SysAPI, batchSize int) error
		// UpdateOne 更新一条数据
		UpdateOne(ctx context.Context, data *fkratos_sys_model.SysAPI) error
		// UpdateOne 更新一条数据(事务)
		UpdateOneByTx(ctx context.Context, tx *fkratos_sys_dao.Query, data *fkratos_sys_model.SysAPI) error
		// UpdateOneWithZero 更新一条数据,包含零值
		UpdateOneWithZero(ctx context.Context, data *fkratos_sys_model.SysAPI) error
		// UpdateOneWithZero 更新一条数据,包含零值(事务)
		UpdateOneWithZeroByTx(ctx context.Context, tx *fkratos_sys_dao.Query, data *fkratos_sys_model.SysAPI) error
		// FindOneCacheByID 根据ID查询一条数据并设置缓存
		FindOneCacheByID(ctx context.Context, ID string) (*fkratos_sys_model.SysAPI, error)
		// FindOneByID 根据ID查询一条数据
		FindOneByID(ctx context.Context, ID string) (*fkratos_sys_model.SysAPI, error)
		// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
		FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*fkratos_sys_model.SysAPI, error)
		// FindMultiByIDS 根据IDS查询多条数据
		FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_sys_model.SysAPI, error)
		// FindMultiByPermissionID 根据permissionID查询多条数据
		FindMultiByPermissionID(ctx context.Context, permissionID string) ([]*fkratos_sys_model.SysAPI, error)
		// FindMultiByPermissionIDS 根据permissionIDS查询多条数据
		FindMultiByPermissionIDS(ctx context.Context, permissionIDS []string) ([]*fkratos_sys_model.SysAPI, error)
		// FindMultiByPaginator 查询分页数据(通用)
		FindMultiByPaginator(ctx context.Context, paginatorReq *orm.PaginatorReq) ([]*fkratos_sys_model.SysAPI, *orm.PaginatorReply, error)
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
		DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_sys_model.SysAPI) error
	}
	SysAPIRepo struct {
		db    *gorm.DB
		cache cache.IDBCache
	}
)

func NewSysAPIRepo(db *gorm.DB, cache cache.IDBCache) *SysAPIRepo {
	return &SysAPIRepo{
		db:    db,
		cache: cache,
	}
}

// CreateOne 创建一条数据
func (s *SysAPIRepo) CreateOne(ctx context.Context, data *fkratos_sys_model.SysAPI) error {
	dao := fkratos_sys_dao.Use(s.db).SysAPI
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateOneByTx 创建一条数据(事务)
func (s *SysAPIRepo) CreateOneByTx(ctx context.Context, tx *fkratos_sys_dao.Query, data *fkratos_sys_model.SysAPI) error {
	dao := tx.SysAPI
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOne Upsert一条数据
func (s *SysAPIRepo) UpsertOne(ctx context.Context, data *fkratos_sys_model.SysAPI) error {
	dao := fkratos_sys_dao.Use(s.db).SysAPI
	err := dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneByTx Upsert一条数据(事务)
func (s *SysAPIRepo) UpsertOneByTx(ctx context.Context, tx *fkratos_sys_dao.Query, data *fkratos_sys_model.SysAPI) error {
	dao := tx.SysAPI
	err := dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateBatch 批量创建数据
func (s *SysAPIRepo) CreateBatch(ctx context.Context, data []*fkratos_sys_model.SysAPI, batchSize int) error {
	dao := fkratos_sys_dao.Use(s.db).SysAPI
	err := dao.WithContext(ctx).CreateInBatches(data, batchSize)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne 更新一条数据
func (s *SysAPIRepo) UpdateOne(ctx context.Context, data *fkratos_sys_model.SysAPI) error {
	dao := fkratos_sys_dao.Use(s.db).SysAPI
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = s.DeleteUniqueIndexCache(ctx, []*fkratos_sys_model.SysAPI{data})
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneByTx 更新一条数据(事务)
func (s *SysAPIRepo) UpdateOneByTx(ctx context.Context, tx *fkratos_sys_dao.Query, data *fkratos_sys_model.SysAPI) error {
	dao := tx.SysAPI
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = s.DeleteUniqueIndexCache(ctx, []*fkratos_sys_model.SysAPI{data})
	if err != nil {
		return err
	}
	return err
}

// UpdateOneWithZero 更新一条数据,包含零值
func (s *SysAPIRepo) UpdateOneWithZero(ctx context.Context, data *fkratos_sys_model.SysAPI) error {
	dao := fkratos_sys_dao.Use(s.db).SysAPI
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Select(dao.ALL.WithTable("")).Updates(data)
	if err != nil {
		return err
	}
	err = s.DeleteUniqueIndexCache(ctx, []*fkratos_sys_model.SysAPI{data})
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneWithZeroByTx 更新一条数据(事务),包含零值
func (s *SysAPIRepo) UpdateOneWithZeroByTx(ctx context.Context, tx *fkratos_sys_dao.Query, data *fkratos_sys_model.SysAPI) error {
	dao := tx.SysAPI
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Select(dao.ALL.WithTable("")).Updates(data)
	if err != nil {
		return err
	}
	err = s.DeleteUniqueIndexCache(ctx, []*fkratos_sys_model.SysAPI{data})
	if err != nil {
		return err
	}
	return err
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (s *SysAPIRepo) DeleteOneCacheByID(ctx context.Context, ID string) error {
	dao := fkratos_sys_dao.Use(s.db).SysAPI
	result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if result == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	err = s.DeleteUniqueIndexCache(ctx, []*fkratos_sys_model.SysAPI{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (s *SysAPIRepo) DeleteOneCacheByIDTx(ctx context.Context, tx *fkratos_sys_dao.Query, ID string) error {
	dao := tx.SysAPI
	result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if result == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	err = s.DeleteUniqueIndexCache(ctx, []*fkratos_sys_model.SysAPI{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (s *SysAPIRepo) DeleteOneByID(ctx context.Context, ID string) error {
	dao := fkratos_sys_dao.Use(s.db).SysAPI
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (s *SysAPIRepo) DeleteOneByIDTx(ctx context.Context, tx *fkratos_sys_dao.Query, ID string) error {
	dao := tx.SysAPI
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (s *SysAPIRepo) DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error {
	dao := fkratos_sys_dao.Use(s.db).SysAPI
	result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	err = s.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (s *SysAPIRepo) DeleteMultiCacheByIDSTx(ctx context.Context, tx *fkratos_sys_dao.Query, IDS []string) error {
	dao := tx.SysAPI
	result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	err = s.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (s *SysAPIRepo) DeleteMultiByIDS(ctx context.Context, IDS []string) error {
	dao := fkratos_sys_dao.Use(s.db).SysAPI
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (s *SysAPIRepo) DeleteMultiByIDSTx(ctx context.Context, tx *fkratos_sys_dao.Query, IDS []string) error {
	dao := tx.SysAPI
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteUniqueIndexCache 删除唯一索引存在的缓存
func (s *SysAPIRepo) DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_sys_model.SysAPI) error {
	keys := make([]string, 0)
	for _, v := range data {
		keys = append(keys, s.cache.Key(cacheSysAPIByIDPrefix, v.ID))

	}
	err := s.cache.DelBatch(ctx, keys)
	if err != nil {
		return err
	}
	return nil
}

// FindOneCacheByID 根据ID查询一条数据并设置缓存
func (s *SysAPIRepo) FindOneCacheByID(ctx context.Context, ID string) (*fkratos_sys_model.SysAPI, error) {
	resp := new(fkratos_sys_model.SysAPI)
	cacheKey := s.cache.Key(cacheSysAPIByIDPrefix, ID)
	cacheValue, err := s.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := fkratos_sys_dao.Use(s.db).SysAPI
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
func (s *SysAPIRepo) FindOneByID(ctx context.Context, ID string) (*fkratos_sys_model.SysAPI, error) {
	dao := fkratos_sys_dao.Use(s.db).SysAPI
	result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
func (s *SysAPIRepo) FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*fkratos_sys_model.SysAPI, error) {
	resp := make([]*fkratos_sys_model.SysAPI, 0)
	cacheKeys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range IDS {
		cacheKey := s.cache.Key(cacheSysAPIByIDPrefix, v)
		cacheKeys = append(cacheKeys, cacheKey)
		keyToParam[cacheKey] = v
	}
	cacheValue, err := s.cache.FetchBatch(ctx, cacheKeys, func(miss []string) (map[string]string, error) {
		parameters := make([]string, 0)
		for _, v := range miss {
			parameters = append(parameters, keyToParam[v])
		}
		dao := fkratos_sys_dao.Use(s.db).SysAPI
		result, err := dao.WithContext(ctx).Where(dao.ID.In(parameters...)).Find()
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
			value[s.cache.Key(cacheSysAPIByIDPrefix, v.ID)] = string(marshal)
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

// FindMultiByIDS 根据IDS查询多条数据
func (s *SysAPIRepo) FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_sys_model.SysAPI, error) {
	dao := fkratos_sys_dao.Use(s.db).SysAPI
	result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByPermissionID 根据permissionID查询多条数据
func (s *SysAPIRepo) FindMultiByPermissionID(ctx context.Context, permissionID string) ([]*fkratos_sys_model.SysAPI, error) {
	dao := fkratos_sys_dao.Use(s.db).SysAPI
	result, err := dao.WithContext(ctx).Where(dao.PermissionID.Eq(permissionID)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByPermissionIDS 根据permissionIDS查询多条数据
func (s *SysAPIRepo) FindMultiByPermissionIDS(ctx context.Context, permissionIDS []string) ([]*fkratos_sys_model.SysAPI, error) {
	dao := fkratos_sys_dao.Use(s.db).SysAPI
	result, err := dao.WithContext(ctx).Where(dao.PermissionID.In(permissionIDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByPaginator 查询分页数据(通用)
func (s *SysAPIRepo) FindMultiByPaginator(ctx context.Context, paginatorReq *orm.PaginatorReq) ([]*fkratos_sys_model.SysAPI, *orm.PaginatorReply, error) {
	result := make([]*fkratos_sys_model.SysAPI, 0)
	var total int64
	whereExpressions, orderExpressions, err := paginatorReq.ConvertToGormExpression(fkratos_sys_model.SysAPI{})
	if err != nil {
		return result, nil, err
	}
	err = s.db.WithContext(ctx).Model(&fkratos_sys_model.SysAPI{}).Select([]string{"*"}).Clauses(whereExpressions...).Count(&total).Error
	if err != nil {
		return result, nil, err
	}
	if total == 0 {
		return result, nil, nil
	}
	paginatorReply := paginatorReq.ConvertToPage(int(total))
	err = s.db.WithContext(ctx).Model(&fkratos_sys_model.SysAPI{}).Limit(paginatorReply.Limit).Offset(paginatorReply.Offset).Clauses(whereExpressions...).Clauses(orderExpressions...).Find(&result).Error
	if err != nil {
		return result, nil, err
	}
	return result, paginatorReply, err
}
