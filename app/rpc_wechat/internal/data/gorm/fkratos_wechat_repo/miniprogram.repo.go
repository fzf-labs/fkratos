// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.

package fkratos_wechat_repo

import (
	"context"
	"encoding/json"
	"errors"
	"fkratos/app/rpc_wechat/internal/data/gorm/fkratos_wechat_dao"
	"fkratos/app/rpc_wechat/internal/data/gorm/fkratos_wechat_model"

	"github.com/fzf-labs/fpkg/orm"
	"github.com/fzf-labs/fpkg/orm/gen/cache"
	"gorm.io/gorm"
)

var _ IMiniprogramRepo = (*MiniprogramRepo)(nil)

var (
	cacheMiniprogramByIDPrefix = "DBCache:fkratos_wechat:MiniprogramByID"
)

type (
	IMiniprogramRepo interface {
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *fkratos_wechat_model.Miniprogram) error
		// CreateOneByTx 创建一条数据(事务)
		CreateOneByTx(ctx context.Context, tx *fkratos_wechat_dao.Query, data *fkratos_wechat_model.Miniprogram) error
		// UpsertOne Upsert一条数据
		UpsertOne(ctx context.Context, data *fkratos_wechat_model.Miniprogram) error
		// UpsertOneByTx Upsert一条数据(事务)
		UpsertOneByTx(ctx context.Context, tx *fkratos_wechat_dao.Query, data *fkratos_wechat_model.Miniprogram) error
		// CreateBatch 批量创建数据
		CreateBatch(ctx context.Context, data []*fkratos_wechat_model.Miniprogram, batchSize int) error
		// UpdateOne 更新一条数据
		UpdateOne(ctx context.Context, data *fkratos_wechat_model.Miniprogram) error
		// UpdateOne 更新一条数据(事务)
		UpdateOneByTx(ctx context.Context, tx *fkratos_wechat_dao.Query, data *fkratos_wechat_model.Miniprogram) error
		// UpdateOneWithZero 更新一条数据,包含零值
		UpdateOneWithZero(ctx context.Context, data *fkratos_wechat_model.Miniprogram) error
		// UpdateOneWithZero 更新一条数据,包含零值(事务)
		UpdateOneWithZeroByTx(ctx context.Context, tx *fkratos_wechat_dao.Query, data *fkratos_wechat_model.Miniprogram) error
		// FindOneCacheByID 根据ID查询一条数据并设置缓存
		FindOneCacheByID(ctx context.Context, ID string) (*fkratos_wechat_model.Miniprogram, error)
		// FindOneByID 根据ID查询一条数据
		FindOneByID(ctx context.Context, ID string) (*fkratos_wechat_model.Miniprogram, error)
		// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
		FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*fkratos_wechat_model.Miniprogram, error)
		// FindMultiByIDS 根据IDS查询多条数据
		FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_wechat_model.Miniprogram, error)
		// FindMultiByPaginator 查询分页数据(通用)
		FindMultiByPaginator(ctx context.Context, paginatorReq *orm.PaginatorReq) ([]*fkratos_wechat_model.Miniprogram, *orm.PaginatorReply, error)
		// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
		DeleteOneCacheByID(ctx context.Context, ID string) error
		// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
		DeleteOneCacheByIDTx(ctx context.Context, tx *fkratos_wechat_dao.Query, ID string) error
		// DeleteOneByID 根据ID删除一条数据
		DeleteOneByID(ctx context.Context, ID string) error
		// DeleteOneByID 根据ID删除一条数据
		DeleteOneByIDTx(ctx context.Context, tx *fkratos_wechat_dao.Query, ID string) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
		DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
		DeleteMultiCacheByIDSTx(ctx context.Context, tx *fkratos_wechat_dao.Query, IDS []string) error
		// DeleteMultiByIDS 根据IDS删除多条数据
		DeleteMultiByIDS(ctx context.Context, IDS []string) error
		// DeleteMultiByIDS 根据IDS删除多条数据
		DeleteMultiByIDSTx(ctx context.Context, tx *fkratos_wechat_dao.Query, IDS []string) error
		// DeleteUniqueIndexCache 删除唯一索引存在的缓存
		DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_wechat_model.Miniprogram) error
	}
	MiniprogramRepo struct {
		db    *gorm.DB
		cache cache.IDBCache
	}
)

func NewMiniprogramRepo(db *gorm.DB, cache cache.IDBCache) *MiniprogramRepo {
	return &MiniprogramRepo{
		db:    db,
		cache: cache,
	}
}

// CreateOne 创建一条数据
func (m *MiniprogramRepo) CreateOne(ctx context.Context, data *fkratos_wechat_model.Miniprogram) error {
	dao := fkratos_wechat_dao.Use(m.db).Miniprogram
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateOneByTx 创建一条数据(事务)
func (m *MiniprogramRepo) CreateOneByTx(ctx context.Context, tx *fkratos_wechat_dao.Query, data *fkratos_wechat_model.Miniprogram) error {
	dao := tx.Miniprogram
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOne Upsert一条数据
func (m *MiniprogramRepo) UpsertOne(ctx context.Context, data *fkratos_wechat_model.Miniprogram) error {
	dao := fkratos_wechat_dao.Use(m.db).Miniprogram
	err := dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneByTx Upsert一条数据(事务)
func (m *MiniprogramRepo) UpsertOneByTx(ctx context.Context, tx *fkratos_wechat_dao.Query, data *fkratos_wechat_model.Miniprogram) error {
	dao := tx.Miniprogram
	err := dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateBatch 批量创建数据
func (m *MiniprogramRepo) CreateBatch(ctx context.Context, data []*fkratos_wechat_model.Miniprogram, batchSize int) error {
	dao := fkratos_wechat_dao.Use(m.db).Miniprogram
	err := dao.WithContext(ctx).CreateInBatches(data, batchSize)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne 更新一条数据
func (m *MiniprogramRepo) UpdateOne(ctx context.Context, data *fkratos_wechat_model.Miniprogram) error {
	dao := fkratos_wechat_dao.Use(m.db).Miniprogram
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = m.DeleteUniqueIndexCache(ctx, []*fkratos_wechat_model.Miniprogram{data})
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneByTx 更新一条数据(事务)
func (m *MiniprogramRepo) UpdateOneByTx(ctx context.Context, tx *fkratos_wechat_dao.Query, data *fkratos_wechat_model.Miniprogram) error {
	dao := tx.Miniprogram
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = m.DeleteUniqueIndexCache(ctx, []*fkratos_wechat_model.Miniprogram{data})
	if err != nil {
		return err
	}
	return err
}

// UpdateOneWithZero 更新一条数据,包含零值
func (m *MiniprogramRepo) UpdateOneWithZero(ctx context.Context, data *fkratos_wechat_model.Miniprogram) error {
	dao := fkratos_wechat_dao.Use(m.db).Miniprogram
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Select(dao.ALL.WithTable("")).Updates(data)
	if err != nil {
		return err
	}
	err = m.DeleteUniqueIndexCache(ctx, []*fkratos_wechat_model.Miniprogram{data})
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneWithZeroByTx 更新一条数据(事务),包含零值
func (m *MiniprogramRepo) UpdateOneWithZeroByTx(ctx context.Context, tx *fkratos_wechat_dao.Query, data *fkratos_wechat_model.Miniprogram) error {
	dao := tx.Miniprogram
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Select(dao.ALL.WithTable("")).Updates(data)
	if err != nil {
		return err
	}
	err = m.DeleteUniqueIndexCache(ctx, []*fkratos_wechat_model.Miniprogram{data})
	if err != nil {
		return err
	}
	return err
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (m *MiniprogramRepo) DeleteOneCacheByID(ctx context.Context, ID string) error {
	dao := fkratos_wechat_dao.Use(m.db).Miniprogram
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
	err = m.DeleteUniqueIndexCache(ctx, []*fkratos_wechat_model.Miniprogram{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (m *MiniprogramRepo) DeleteOneCacheByIDTx(ctx context.Context, tx *fkratos_wechat_dao.Query, ID string) error {
	dao := tx.Miniprogram
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
	err = m.DeleteUniqueIndexCache(ctx, []*fkratos_wechat_model.Miniprogram{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (m *MiniprogramRepo) DeleteOneByID(ctx context.Context, ID string) error {
	dao := fkratos_wechat_dao.Use(m.db).Miniprogram
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (m *MiniprogramRepo) DeleteOneByIDTx(ctx context.Context, tx *fkratos_wechat_dao.Query, ID string) error {
	dao := tx.Miniprogram
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (m *MiniprogramRepo) DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error {
	dao := fkratos_wechat_dao.Use(m.db).Miniprogram
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
	err = m.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (m *MiniprogramRepo) DeleteMultiCacheByIDSTx(ctx context.Context, tx *fkratos_wechat_dao.Query, IDS []string) error {
	dao := tx.Miniprogram
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
	err = m.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (m *MiniprogramRepo) DeleteMultiByIDS(ctx context.Context, IDS []string) error {
	dao := fkratos_wechat_dao.Use(m.db).Miniprogram
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (m *MiniprogramRepo) DeleteMultiByIDSTx(ctx context.Context, tx *fkratos_wechat_dao.Query, IDS []string) error {
	dao := tx.Miniprogram
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteUniqueIndexCache 删除唯一索引存在的缓存
func (m *MiniprogramRepo) DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_wechat_model.Miniprogram) error {
	keys := make([]string, 0)
	for _, v := range data {
		keys = append(keys, m.cache.Key(cacheMiniprogramByIDPrefix, v.ID))

	}
	err := m.cache.DelBatch(ctx, keys)
	if err != nil {
		return err
	}
	return nil
}

// FindOneCacheByID 根据ID查询一条数据并设置缓存
func (m *MiniprogramRepo) FindOneCacheByID(ctx context.Context, ID string) (*fkratos_wechat_model.Miniprogram, error) {
	resp := new(fkratos_wechat_model.Miniprogram)
	cacheKey := m.cache.Key(cacheMiniprogramByIDPrefix, ID)
	cacheValue, err := m.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := fkratos_wechat_dao.Use(m.db).Miniprogram
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
func (m *MiniprogramRepo) FindOneByID(ctx context.Context, ID string) (*fkratos_wechat_model.Miniprogram, error) {
	dao := fkratos_wechat_dao.Use(m.db).Miniprogram
	result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
func (m *MiniprogramRepo) FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*fkratos_wechat_model.Miniprogram, error) {
	resp := make([]*fkratos_wechat_model.Miniprogram, 0)
	cacheKeys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range IDS {
		cacheKey := m.cache.Key(cacheMiniprogramByIDPrefix, v)
		cacheKeys = append(cacheKeys, cacheKey)
		keyToParam[cacheKey] = v
	}
	cacheValue, err := m.cache.FetchBatch(ctx, cacheKeys, func(miss []string) (map[string]string, error) {
		parameters := make([]string, 0)
		for _, v := range miss {
			parameters = append(parameters, keyToParam[v])
		}
		dao := fkratos_wechat_dao.Use(m.db).Miniprogram
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
			value[m.cache.Key(cacheMiniprogramByIDPrefix, v.ID)] = string(marshal)
		}
		return value, nil
	})
	if err != nil {
		return nil, err
	}
	for _, v := range cacheValue {
		tmp := new(fkratos_wechat_model.Miniprogram)
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}

// FindMultiByIDS 根据IDS查询多条数据
func (m *MiniprogramRepo) FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_wechat_model.Miniprogram, error) {
	dao := fkratos_wechat_dao.Use(m.db).Miniprogram
	result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByPaginator 查询分页数据(通用)
func (m *MiniprogramRepo) FindMultiByPaginator(ctx context.Context, paginatorReq *orm.PaginatorReq) ([]*fkratos_wechat_model.Miniprogram, *orm.PaginatorReply, error) {
	result := make([]*fkratos_wechat_model.Miniprogram, 0)
	var total int64
	whereExpressions, orderExpressions, err := paginatorReq.ConvertToGormExpression(fkratos_wechat_model.Miniprogram{})
	if err != nil {
		return nil, nil, err
	}
	err = m.db.WithContext(ctx).Model(&fkratos_wechat_model.Miniprogram{}).Select([]string{"*"}).Clauses(whereExpressions...).Count(&total).Error
	if err != nil {
		return result, nil, err
	}
	if total == 0 {
		return result, nil, nil
	}
	paginatorReply := paginatorReq.ConvertToPage(int(total))
	err = m.db.WithContext(ctx).Model(&fkratos_wechat_model.Miniprogram{}).Limit(paginatorReply.Limit).Offset(paginatorReply.Offset).Clauses(whereExpressions...).Clauses(orderExpressions...).Find(&result).Error
	if err != nil {
		return result, nil, err
	}
	return result, paginatorReply, err
}
