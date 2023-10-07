// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.

package fkratos_common_repo

import (
	"context"
	"encoding/json"
	"errors"
	"fkratos/app/rpc_common/internal/data/gorm/fkratos_common_dao"
	"fkratos/app/rpc_common/internal/data/gorm/fkratos_common_model"

	"github.com/fzf-labs/fpkg/orm"
	"github.com/fzf-labs/fpkg/orm/gen/cache"
	"gorm.io/gorm"
)

var _ IDictTypeRepo = (*DictTypeRepo)(nil)

var (
	cacheDictTypeByIDPrefix = "DBCache:fkratos_common:DictTypeByID"
)

type (
	IDictTypeRepo interface {
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *fkratos_common_model.DictType) error
		// CreateOneByTx 创建一条数据(事务)
		CreateOneByTx(ctx context.Context, tx *fkratos_common_dao.Query, data *fkratos_common_model.DictType) error
		// CreateBatch 批量创建数据
		CreateBatch(ctx context.Context, data []*fkratos_common_model.DictType, batchSize int) error
		// UpdateOne 更新一条数据
		UpdateOne(ctx context.Context, data *fkratos_common_model.DictType) error
		// UpdateOne 更新一条数据(事务)
		UpdateOneByTx(ctx context.Context, tx *fkratos_common_dao.Query, data *fkratos_common_model.DictType) error
		// FindOneCacheByID 根据ID查询一条数据并设置缓存
		FindOneCacheByID(ctx context.Context, ID string) (*fkratos_common_model.DictType, error)
		// FindOneByID 根据ID查询一条数据
		FindOneByID(ctx context.Context, ID string) (*fkratos_common_model.DictType, error)
		// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
		FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*fkratos_common_model.DictType, error)
		// FindMultiByIDS 根据IDS查询多条数据
		FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_common_model.DictType, error)
		// FindMultiByPaginator 查询分页数据(通用)
		FindMultiByPaginator(ctx context.Context, paginatorReq *orm.PaginatorReq) ([]*fkratos_common_model.DictType, *orm.PaginatorReply, error)
		// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
		DeleteOneCacheByID(ctx context.Context, ID string) error
		// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
		DeleteOneCacheByIDTx(ctx context.Context, tx *fkratos_common_dao.Query, ID string) error
		// DeleteOneByID 根据ID删除一条数据
		DeleteOneByID(ctx context.Context, ID string) error
		// DeleteOneByID 根据ID删除一条数据
		DeleteOneByIDTx(ctx context.Context, tx *fkratos_common_dao.Query, ID string) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
		DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
		DeleteMultiCacheByIDSTx(ctx context.Context, tx *fkratos_common_dao.Query, IDS []string) error
		// DeleteMultiByIDS 根据IDS删除多条数据
		DeleteMultiByIDS(ctx context.Context, IDS []string) error
		// DeleteMultiByIDS 根据IDS删除多条数据
		DeleteMultiByIDSTx(ctx context.Context, tx *fkratos_common_dao.Query, IDS []string) error
		// DeleteUniqueIndexCache 删除唯一索引存在的缓存
		DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_common_model.DictType) error
	}
	DictTypeRepo struct {
		db    *gorm.DB
		cache cache.IDBCache
	}
)

func NewDictTypeRepo(db *gorm.DB, cache cache.IDBCache) *DictTypeRepo {
	return &DictTypeRepo{
		db:    db,
		cache: cache,
	}
}

// CreateOne 创建一条数据
func (d *DictTypeRepo) CreateOne(ctx context.Context, data *fkratos_common_model.DictType) error {
	dao := fkratos_common_dao.Use(d.db).DictType
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateOneByTx 创建一条数据(事务)
func (d *DictTypeRepo) CreateOneByTx(ctx context.Context, tx *fkratos_common_dao.Query, data *fkratos_common_model.DictType) error {
	dao := tx.DictType
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateBatch 批量创建数据
func (d *DictTypeRepo) CreateBatch(ctx context.Context, data []*fkratos_common_model.DictType, batchSize int) error {
	dao := fkratos_common_dao.Use(d.db).DictType
	err := dao.WithContext(ctx).CreateInBatches(data, batchSize)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne 更新一条数据
func (d *DictTypeRepo) UpdateOne(ctx context.Context, data *fkratos_common_model.DictType) error {
	dao := fkratos_common_dao.Use(d.db).DictType
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = d.DeleteUniqueIndexCache(ctx, []*fkratos_common_model.DictType{data})
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneByTx 更新一条数据(事务)
func (d *DictTypeRepo) UpdateOneByTx(ctx context.Context, tx *fkratos_common_dao.Query, data *fkratos_common_model.DictType) error {
	dao := tx.DictType
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = d.DeleteUniqueIndexCache(ctx, []*fkratos_common_model.DictType{data})
	if err != nil {
		return err
	}
	return err
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (d *DictTypeRepo) DeleteOneCacheByID(ctx context.Context, ID string) error {
	dao := fkratos_common_dao.Use(d.db).DictType
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
	err = d.DeleteUniqueIndexCache(ctx, []*fkratos_common_model.DictType{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (d *DictTypeRepo) DeleteOneCacheByIDTx(ctx context.Context, tx *fkratos_common_dao.Query, ID string) error {
	dao := tx.DictType
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
	err = d.DeleteUniqueIndexCache(ctx, []*fkratos_common_model.DictType{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (d *DictTypeRepo) DeleteOneByID(ctx context.Context, ID string) error {
	dao := fkratos_common_dao.Use(d.db).DictType
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (d *DictTypeRepo) DeleteOneByIDTx(ctx context.Context, tx *fkratos_common_dao.Query, ID string) error {
	dao := tx.DictType
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (d *DictTypeRepo) DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error {
	dao := fkratos_common_dao.Use(d.db).DictType
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
	err = d.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (d *DictTypeRepo) DeleteMultiCacheByIDSTx(ctx context.Context, tx *fkratos_common_dao.Query, IDS []string) error {
	dao := tx.DictType
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
	err = d.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (d *DictTypeRepo) DeleteMultiByIDS(ctx context.Context, IDS []string) error {
	dao := fkratos_common_dao.Use(d.db).DictType
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (d *DictTypeRepo) DeleteMultiByIDSTx(ctx context.Context, tx *fkratos_common_dao.Query, IDS []string) error {
	dao := tx.DictType
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteUniqueIndexCache 删除唯一索引存在的缓存
func (d *DictTypeRepo) DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_common_model.DictType) error {
	keys := make([]string, 0)
	for _, v := range data {
		keys = append(keys, d.cache.Key(cacheDictTypeByIDPrefix, v.ID))

	}
	err := d.cache.DelBatch(ctx, keys)
	if err != nil {
		return err
	}
	return nil
}

// FindOneCacheByID 根据ID查询一条数据并设置缓存
func (d *DictTypeRepo) FindOneCacheByID(ctx context.Context, ID string) (*fkratos_common_model.DictType, error) {
	resp := new(fkratos_common_model.DictType)
	cacheKey := d.cache.Key(cacheDictTypeByIDPrefix, ID)
	cacheValue, err := d.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := fkratos_common_dao.Use(d.db).DictType
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
func (d *DictTypeRepo) FindOneByID(ctx context.Context, ID string) (*fkratos_common_model.DictType, error) {
	dao := fkratos_common_dao.Use(d.db).DictType
	result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
func (d *DictTypeRepo) FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*fkratos_common_model.DictType, error) {
	resp := make([]*fkratos_common_model.DictType, 0)
	cacheKeys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range IDS {
		cacheKey := d.cache.Key(cacheDictTypeByIDPrefix, v)
		cacheKeys = append(cacheKeys, cacheKey)
		keyToParam[cacheKey] = v
	}
	cacheValue, err := d.cache.FetchBatch(ctx, cacheKeys, func(miss []string) (map[string]string, error) {
		parameters := make([]string, 0)
		for _, v := range miss {
			parameters = append(parameters, keyToParam[v])
		}
		dao := fkratos_common_dao.Use(d.db).DictType
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
			value[d.cache.Key(cacheDictTypeByIDPrefix, v.ID)] = string(marshal)
		}
		return value, nil
	})
	if err != nil {
		return nil, err
	}
	for _, v := range cacheValue {
		tmp := new(fkratos_common_model.DictType)
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}

// FindMultiByIDS 根据IDS查询多条数据
func (d *DictTypeRepo) FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_common_model.DictType, error) {
	dao := fkratos_common_dao.Use(d.db).DictType
	result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByPaginator 查询分页数据(通用)
func (d *DictTypeRepo) FindMultiByPaginator(ctx context.Context, paginatorReq *orm.PaginatorReq) ([]*fkratos_common_model.DictType, *orm.PaginatorReply, error) {
	result := make([]*fkratos_common_model.DictType, 0)
	var total int64
	queryStr, args, err := paginatorReq.ConvertToGormConditions()
	if err != nil {
		return nil, nil, err
	}
	err = d.db.WithContext(ctx).Model(&fkratos_common_model.DictType{}).Select([]string{"id"}).Where(queryStr, args...).Count(&total).Error
	if err != nil {
		return nil, nil, err
	}
	if total == 0 {
		return nil, nil, nil
	}
	query := d.db.WithContext(ctx)
	order := paginatorReq.ConvertToOrder()
	if order != "" {
		query = query.Order(order)
	}
	paginatorReply := paginatorReq.ConvertToPage(int(total))
	err = query.Limit(paginatorReply.Limit).Offset(paginatorReply.Offset).Where(queryStr, args...).Find(&result).Error
	if err != nil {
		return nil, nil, err
	}
	return result, paginatorReply, err
}