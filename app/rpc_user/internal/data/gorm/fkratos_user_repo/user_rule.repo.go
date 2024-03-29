// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.

package fkratos_user_repo

import (
	"context"
	"encoding/json"
	"errors"
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_dao"
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_model"

	"github.com/fzf-labs/fpkg/orm"
	"github.com/fzf-labs/fpkg/orm/gen/cache"
	"gorm.io/gorm"
)

var _ IUserRuleRepo = (*UserRuleRepo)(nil)

var (
	cacheUserRuleByIDPrefix = "DBCache:fkratos_user:UserRuleByID"
)

type (
	IUserRuleRepo interface {
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *fkratos_user_model.UserRule) error
		// CreateOneByTx 创建一条数据(事务)
		CreateOneByTx(ctx context.Context, tx *fkratos_user_dao.Query, data *fkratos_user_model.UserRule) error
		// UpsertOne Upsert一条数据
		UpsertOne(ctx context.Context, data *fkratos_user_model.UserRule) error
		// UpsertOneByTx Upsert一条数据(事务)
		UpsertOneByTx(ctx context.Context, tx *fkratos_user_dao.Query, data *fkratos_user_model.UserRule) error
		// CreateBatch 批量创建数据
		CreateBatch(ctx context.Context, data []*fkratos_user_model.UserRule, batchSize int) error
		// UpdateOne 更新一条数据
		UpdateOne(ctx context.Context, data *fkratos_user_model.UserRule) error
		// UpdateOne 更新一条数据(事务)
		UpdateOneByTx(ctx context.Context, tx *fkratos_user_dao.Query, data *fkratos_user_model.UserRule) error
		// UpdateOneWithZero 更新一条数据,包含零值
		UpdateOneWithZero(ctx context.Context, data *fkratos_user_model.UserRule) error
		// UpdateOneWithZero 更新一条数据,包含零值(事务)
		UpdateOneWithZeroByTx(ctx context.Context, tx *fkratos_user_dao.Query, data *fkratos_user_model.UserRule) error
		// FindOneCacheByID 根据ID查询一条数据并设置缓存
		FindOneCacheByID(ctx context.Context, ID string) (*fkratos_user_model.UserRule, error)
		// FindOneByID 根据ID查询一条数据
		FindOneByID(ctx context.Context, ID string) (*fkratos_user_model.UserRule, error)
		// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
		FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*fkratos_user_model.UserRule, error)
		// FindMultiByIDS 根据IDS查询多条数据
		FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_user_model.UserRule, error)
		// FindMultiByPid 根据pid查询多条数据
		FindMultiByPid(ctx context.Context, pid string) ([]*fkratos_user_model.UserRule, error)
		// FindMultiByPids 根据pids查询多条数据
		FindMultiByPids(ctx context.Context, pids []string) ([]*fkratos_user_model.UserRule, error)
		// FindMultiByPaginator 查询分页数据(通用)
		FindMultiByPaginator(ctx context.Context, paginatorReq *orm.PaginatorReq) ([]*fkratos_user_model.UserRule, *orm.PaginatorReply, error)
		// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
		DeleteOneCacheByID(ctx context.Context, ID string) error
		// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
		DeleteOneCacheByIDTx(ctx context.Context, tx *fkratos_user_dao.Query, ID string) error
		// DeleteOneByID 根据ID删除一条数据
		DeleteOneByID(ctx context.Context, ID string) error
		// DeleteOneByID 根据ID删除一条数据
		DeleteOneByIDTx(ctx context.Context, tx *fkratos_user_dao.Query, ID string) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
		DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
		DeleteMultiCacheByIDSTx(ctx context.Context, tx *fkratos_user_dao.Query, IDS []string) error
		// DeleteMultiByIDS 根据IDS删除多条数据
		DeleteMultiByIDS(ctx context.Context, IDS []string) error
		// DeleteMultiByIDS 根据IDS删除多条数据
		DeleteMultiByIDSTx(ctx context.Context, tx *fkratos_user_dao.Query, IDS []string) error
		// DeleteUniqueIndexCache 删除唯一索引存在的缓存
		DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_user_model.UserRule) error
	}
	UserRuleRepo struct {
		db    *gorm.DB
		cache cache.IDBCache
	}
)

func NewUserRuleRepo(db *gorm.DB, cache cache.IDBCache) *UserRuleRepo {
	return &UserRuleRepo{
		db:    db,
		cache: cache,
	}
}

// CreateOne 创建一条数据
func (u *UserRuleRepo) CreateOne(ctx context.Context, data *fkratos_user_model.UserRule) error {
	dao := fkratos_user_dao.Use(u.db).UserRule
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateOneByTx 创建一条数据(事务)
func (u *UserRuleRepo) CreateOneByTx(ctx context.Context, tx *fkratos_user_dao.Query, data *fkratos_user_model.UserRule) error {
	dao := tx.UserRule
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOne Upsert一条数据
func (u *UserRuleRepo) UpsertOne(ctx context.Context, data *fkratos_user_model.UserRule) error {
	dao := fkratos_user_dao.Use(u.db).UserRule
	err := dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneByTx Upsert一条数据(事务)
func (u *UserRuleRepo) UpsertOneByTx(ctx context.Context, tx *fkratos_user_dao.Query, data *fkratos_user_model.UserRule) error {
	dao := tx.UserRule
	err := dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateBatch 批量创建数据
func (u *UserRuleRepo) CreateBatch(ctx context.Context, data []*fkratos_user_model.UserRule, batchSize int) error {
	dao := fkratos_user_dao.Use(u.db).UserRule
	err := dao.WithContext(ctx).CreateInBatches(data, batchSize)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne 更新一条数据
func (u *UserRuleRepo) UpdateOne(ctx context.Context, data *fkratos_user_model.UserRule) error {
	dao := fkratos_user_dao.Use(u.db).UserRule
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.UserRule{data})
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneByTx 更新一条数据(事务)
func (u *UserRuleRepo) UpdateOneByTx(ctx context.Context, tx *fkratos_user_dao.Query, data *fkratos_user_model.UserRule) error {
	dao := tx.UserRule
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.UserRule{data})
	if err != nil {
		return err
	}
	return err
}

// UpdateOneWithZero 更新一条数据,包含零值
func (u *UserRuleRepo) UpdateOneWithZero(ctx context.Context, data *fkratos_user_model.UserRule) error {
	dao := fkratos_user_dao.Use(u.db).UserRule
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Select(dao.ALL.WithTable("")).Updates(data)
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.UserRule{data})
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneWithZeroByTx 更新一条数据(事务),包含零值
func (u *UserRuleRepo) UpdateOneWithZeroByTx(ctx context.Context, tx *fkratos_user_dao.Query, data *fkratos_user_model.UserRule) error {
	dao := tx.UserRule
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Select(dao.ALL.WithTable("")).Updates(data)
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.UserRule{data})
	if err != nil {
		return err
	}
	return err
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (u *UserRuleRepo) DeleteOneCacheByID(ctx context.Context, ID string) error {
	dao := fkratos_user_dao.Use(u.db).UserRule
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
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.UserRule{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (u *UserRuleRepo) DeleteOneCacheByIDTx(ctx context.Context, tx *fkratos_user_dao.Query, ID string) error {
	dao := tx.UserRule
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
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.UserRule{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (u *UserRuleRepo) DeleteOneByID(ctx context.Context, ID string) error {
	dao := fkratos_user_dao.Use(u.db).UserRule
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (u *UserRuleRepo) DeleteOneByIDTx(ctx context.Context, tx *fkratos_user_dao.Query, ID string) error {
	dao := tx.UserRule
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (u *UserRuleRepo) DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error {
	dao := fkratos_user_dao.Use(u.db).UserRule
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
	err = u.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (u *UserRuleRepo) DeleteMultiCacheByIDSTx(ctx context.Context, tx *fkratos_user_dao.Query, IDS []string) error {
	dao := tx.UserRule
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
	err = u.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (u *UserRuleRepo) DeleteMultiByIDS(ctx context.Context, IDS []string) error {
	dao := fkratos_user_dao.Use(u.db).UserRule
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (u *UserRuleRepo) DeleteMultiByIDSTx(ctx context.Context, tx *fkratos_user_dao.Query, IDS []string) error {
	dao := tx.UserRule
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteUniqueIndexCache 删除唯一索引存在的缓存
func (u *UserRuleRepo) DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_user_model.UserRule) error {
	keys := make([]string, 0)
	for _, v := range data {
		keys = append(keys, u.cache.Key(cacheUserRuleByIDPrefix, v.ID))

	}
	err := u.cache.DelBatch(ctx, keys)
	if err != nil {
		return err
	}
	return nil
}

// FindOneCacheByID 根据ID查询一条数据并设置缓存
func (u *UserRuleRepo) FindOneCacheByID(ctx context.Context, ID string) (*fkratos_user_model.UserRule, error) {
	resp := new(fkratos_user_model.UserRule)
	cacheKey := u.cache.Key(cacheUserRuleByIDPrefix, ID)
	cacheValue, err := u.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := fkratos_user_dao.Use(u.db).UserRule
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
func (u *UserRuleRepo) FindOneByID(ctx context.Context, ID string) (*fkratos_user_model.UserRule, error) {
	dao := fkratos_user_dao.Use(u.db).UserRule
	result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
func (u *UserRuleRepo) FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*fkratos_user_model.UserRule, error) {
	resp := make([]*fkratos_user_model.UserRule, 0)
	cacheKeys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range IDS {
		cacheKey := u.cache.Key(cacheUserRuleByIDPrefix, v)
		cacheKeys = append(cacheKeys, cacheKey)
		keyToParam[cacheKey] = v
	}
	cacheValue, err := u.cache.FetchBatch(ctx, cacheKeys, func(miss []string) (map[string]string, error) {
		parameters := make([]string, 0)
		for _, v := range miss {
			parameters = append(parameters, keyToParam[v])
		}
		dao := fkratos_user_dao.Use(u.db).UserRule
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
			value[u.cache.Key(cacheUserRuleByIDPrefix, v.ID)] = string(marshal)
		}
		return value, nil
	})
	if err != nil {
		return nil, err
	}
	for _, v := range cacheValue {
		tmp := new(fkratos_user_model.UserRule)
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}

// FindMultiByIDS 根据IDS查询多条数据
func (u *UserRuleRepo) FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_user_model.UserRule, error) {
	dao := fkratos_user_dao.Use(u.db).UserRule
	result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByPid 根据pid查询多条数据
func (u *UserRuleRepo) FindMultiByPid(ctx context.Context, pid string) ([]*fkratos_user_model.UserRule, error) {
	dao := fkratos_user_dao.Use(u.db).UserRule
	result, err := dao.WithContext(ctx).Where(dao.Pid.Eq(pid)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByPids 根据pids查询多条数据
func (u *UserRuleRepo) FindMultiByPids(ctx context.Context, pids []string) ([]*fkratos_user_model.UserRule, error) {
	dao := fkratos_user_dao.Use(u.db).UserRule
	result, err := dao.WithContext(ctx).Where(dao.Pid.In(pids...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByPaginator 查询分页数据(通用)
func (u *UserRuleRepo) FindMultiByPaginator(ctx context.Context, paginatorReq *orm.PaginatorReq) ([]*fkratos_user_model.UserRule, *orm.PaginatorReply, error) {
	result := make([]*fkratos_user_model.UserRule, 0)
	var total int64
	whereExpressions, orderExpressions, err := paginatorReq.ConvertToGormExpression(fkratos_user_model.UserRule{})
	if err != nil {
		return result, nil, err
	}
	err = u.db.WithContext(ctx).Model(&fkratos_user_model.UserRule{}).Select([]string{"*"}).Clauses(whereExpressions...).Count(&total).Error
	if err != nil {
		return result, nil, err
	}
	if total == 0 {
		return result, nil, nil
	}
	paginatorReply := paginatorReq.ConvertToPage(int(total))
	err = u.db.WithContext(ctx).Model(&fkratos_user_model.UserRule{}).Limit(paginatorReply.Limit).Offset(paginatorReply.Offset).Clauses(whereExpressions...).Clauses(orderExpressions...).Find(&result).Error
	if err != nil {
		return result, nil, err
	}
	return result, paginatorReply, err
}
