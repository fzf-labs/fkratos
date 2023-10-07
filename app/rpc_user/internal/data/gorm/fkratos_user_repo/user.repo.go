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

var _ IUserRepo = (*UserRepo)(nil)

var (
	cacheUserByUsernamePrefix = "DBCache:fkratos_user:UserByUsername"
	cacheUserByUIDPrefix      = "DBCache:fkratos_user:UserByUID"
	cacheUserByIDPrefix       = "DBCache:fkratos_user:UserByID"
)

type (
	IUserRepo interface {
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *fkratos_user_model.User) error
		// CreateOneByTx 创建一条数据(事务)
		CreateOneByTx(ctx context.Context, tx *fkratos_user_dao.Query, data *fkratos_user_model.User) error
		// UpsertOne Upsert一条数据
		UpsertOne(ctx context.Context, data *fkratos_user_model.User) error
		// UpsertOneByTx Upsert一条数据(事务)
		UpsertOneByTx(ctx context.Context, tx *fkratos_user_dao.Query, data *fkratos_user_model.User) error
		// CreateBatch 批量创建数据
		CreateBatch(ctx context.Context, data []*fkratos_user_model.User, batchSize int) error
		// UpdateOne 更新一条数据
		UpdateOne(ctx context.Context, data *fkratos_user_model.User) error
		// UpdateOne 更新一条数据(事务)
		UpdateOneByTx(ctx context.Context, tx *fkratos_user_dao.Query, data *fkratos_user_model.User) error
		// UpdateOneWithZero 更新一条数据,包含零值
		UpdateOneWithZero(ctx context.Context, data *fkratos_user_model.User) error
		// UpdateOneWithZero 更新一条数据,包含零值(事务)
		UpdateOneWithZeroByTx(ctx context.Context, tx *fkratos_user_dao.Query, data *fkratos_user_model.User) error
		// FindOneCacheByUsername 根据username查询一条数据并设置缓存
		FindOneCacheByUsername(ctx context.Context, username string) (*fkratos_user_model.User, error)
		// FindOneByUsername 根据username查询一条数据
		FindOneByUsername(ctx context.Context, username string) (*fkratos_user_model.User, error)
		// FindMultiCacheByUsernames 根据usernames查询多条数据并设置缓存
		FindMultiCacheByUsernames(ctx context.Context, usernames []string) ([]*fkratos_user_model.User, error)
		// FindMultiByUsernames 根据usernames查询多条数据
		FindMultiByUsernames(ctx context.Context, usernames []string) ([]*fkratos_user_model.User, error)
		// FindMultiByUserGroupID 根据userGroupID查询多条数据
		FindMultiByUserGroupID(ctx context.Context, userGroupID string) ([]*fkratos_user_model.User, error)
		// FindMultiByUserGroupIDS 根据userGroupIDS查询多条数据
		FindMultiByUserGroupIDS(ctx context.Context, userGroupIDS []string) ([]*fkratos_user_model.User, error)
		// FindOneCacheByUID 根据UID查询一条数据并设置缓存
		FindOneCacheByUID(ctx context.Context, UID string) (*fkratos_user_model.User, error)
		// FindOneByUID 根据UID查询一条数据
		FindOneByUID(ctx context.Context, UID string) (*fkratos_user_model.User, error)
		// FindMultiCacheByUIDS 根据UIDS查询多条数据并设置缓存
		FindMultiCacheByUIDS(ctx context.Context, UIDS []string) ([]*fkratos_user_model.User, error)
		// FindMultiByUIDS 根据UIDS查询多条数据
		FindMultiByUIDS(ctx context.Context, UIDS []string) ([]*fkratos_user_model.User, error)
		// FindOneCacheByID 根据ID查询一条数据并设置缓存
		FindOneCacheByID(ctx context.Context, ID string) (*fkratos_user_model.User, error)
		// FindOneByID 根据ID查询一条数据
		FindOneByID(ctx context.Context, ID string) (*fkratos_user_model.User, error)
		// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
		FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*fkratos_user_model.User, error)
		// FindMultiByIDS 根据IDS查询多条数据
		FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_user_model.User, error)
		// FindMultiByPhone 根据phone查询多条数据
		FindMultiByPhone(ctx context.Context, phone string) ([]*fkratos_user_model.User, error)
		// FindMultiByPhones 根据phones查询多条数据
		FindMultiByPhones(ctx context.Context, phones []string) ([]*fkratos_user_model.User, error)
		// FindMultiByEmail 根据email查询多条数据
		FindMultiByEmail(ctx context.Context, email string) ([]*fkratos_user_model.User, error)
		// FindMultiByEmails 根据emails查询多条数据
		FindMultiByEmails(ctx context.Context, emails []string) ([]*fkratos_user_model.User, error)
		// FindMultiByPaginator 查询分页数据(通用)
		FindMultiByPaginator(ctx context.Context, paginatorReq *orm.PaginatorReq) ([]*fkratos_user_model.User, *orm.PaginatorReply, error)
		// DeleteOneCacheByUsername 根据username删除一条数据并清理缓存
		DeleteOneCacheByUsername(ctx context.Context, username string) error
		// DeleteOneCacheByUsername 根据username删除一条数据并清理缓存
		DeleteOneCacheByUsernameTx(ctx context.Context, tx *fkratos_user_dao.Query, username string) error
		// DeleteOneByUsername 根据username删除一条数据
		DeleteOneByUsername(ctx context.Context, username string) error
		// DeleteOneByUsername 根据username删除一条数据
		DeleteOneByUsernameTx(ctx context.Context, tx *fkratos_user_dao.Query, username string) error
		// DeleteMultiCacheByUsernames 根据Usernames删除多条数据并清理缓存
		DeleteMultiCacheByUsernames(ctx context.Context, usernames []string) error
		// DeleteMultiCacheByUsernames 根据Usernames删除多条数据并清理缓存
		DeleteMultiCacheByUsernamesTx(ctx context.Context, tx *fkratos_user_dao.Query, usernames []string) error
		// DeleteMultiByUsernames 根据Usernames删除多条数据
		DeleteMultiByUsernames(ctx context.Context, usernames []string) error
		// DeleteMultiByUsernames 根据Usernames删除多条数据
		DeleteMultiByUsernamesTx(ctx context.Context, tx *fkratos_user_dao.Query, usernames []string) error
		// DeleteOneCacheByUID 根据UID删除一条数据并清理缓存
		DeleteOneCacheByUID(ctx context.Context, UID string) error
		// DeleteOneCacheByUID 根据UID删除一条数据并清理缓存
		DeleteOneCacheByUIDTx(ctx context.Context, tx *fkratos_user_dao.Query, UID string) error
		// DeleteOneByUID 根据UID删除一条数据
		DeleteOneByUID(ctx context.Context, UID string) error
		// DeleteOneByUID 根据UID删除一条数据
		DeleteOneByUIDTx(ctx context.Context, tx *fkratos_user_dao.Query, UID string) error
		// DeleteMultiCacheByUIDS 根据UIDS删除多条数据并清理缓存
		DeleteMultiCacheByUIDS(ctx context.Context, UIDS []string) error
		// DeleteMultiCacheByUIDS 根据UIDS删除多条数据并清理缓存
		DeleteMultiCacheByUIDSTx(ctx context.Context, tx *fkratos_user_dao.Query, UIDS []string) error
		// DeleteMultiByUIDS 根据UIDS删除多条数据
		DeleteMultiByUIDS(ctx context.Context, UIDS []string) error
		// DeleteMultiByUIDS 根据UIDS删除多条数据
		DeleteMultiByUIDSTx(ctx context.Context, tx *fkratos_user_dao.Query, UIDS []string) error
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
		DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_user_model.User) error
	}
	UserRepo struct {
		db    *gorm.DB
		cache cache.IDBCache
	}
)

func NewUserRepo(db *gorm.DB, cache cache.IDBCache) *UserRepo {
	return &UserRepo{
		db:    db,
		cache: cache,
	}
}

// CreateOne 创建一条数据
func (u *UserRepo) CreateOne(ctx context.Context, data *fkratos_user_model.User) error {
	dao := fkratos_user_dao.Use(u.db).User
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateOneByTx 创建一条数据(事务)
func (u *UserRepo) CreateOneByTx(ctx context.Context, tx *fkratos_user_dao.Query, data *fkratos_user_model.User) error {
	dao := tx.User
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOne Upsert一条数据
func (u *UserRepo) UpsertOne(ctx context.Context, data *fkratos_user_model.User) error {
	dao := fkratos_user_dao.Use(u.db).User
	err := dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneByTx Upsert一条数据(事务)
func (u *UserRepo) UpsertOneByTx(ctx context.Context, tx *fkratos_user_dao.Query, data *fkratos_user_model.User) error {
	dao := tx.User
	err := dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateBatch 批量创建数据
func (u *UserRepo) CreateBatch(ctx context.Context, data []*fkratos_user_model.User, batchSize int) error {
	dao := fkratos_user_dao.Use(u.db).User
	err := dao.WithContext(ctx).CreateInBatches(data, batchSize)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne 更新一条数据
func (u *UserRepo) UpdateOne(ctx context.Context, data *fkratos_user_model.User) error {
	dao := fkratos_user_dao.Use(u.db).User
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{data})
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneByTx 更新一条数据(事务)
func (u *UserRepo) UpdateOneByTx(ctx context.Context, tx *fkratos_user_dao.Query, data *fkratos_user_model.User) error {
	dao := tx.User
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{data})
	if err != nil {
		return err
	}
	return err
}

// UpdateOneWithZero 更新一条数据,包含零值
func (u *UserRepo) UpdateOneWithZero(ctx context.Context, data *fkratos_user_model.User) error {
	dao := fkratos_user_dao.Use(u.db).User
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Select(dao.ALL).Updates(data)
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{data})
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneWithZeroByTx 更新一条数据(事务),包含零值
func (u *UserRepo) UpdateOneWithZeroByTx(ctx context.Context, tx *fkratos_user_dao.Query, data *fkratos_user_model.User) error {
	dao := tx.User
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Select(dao.ALL).Updates(data)
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{data})
	if err != nil {
		return err
	}
	return err
}

// DeleteOneCacheByUsername 根据username删除一条数据并清理缓存
func (u *UserRepo) DeleteOneCacheByUsername(ctx context.Context, username string) error {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.Username.Eq(username)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if result == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Username.Eq(username)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByUsername 根据username删除一条数据并清理缓存
func (u *UserRepo) DeleteOneCacheByUsernameTx(ctx context.Context, tx *fkratos_user_dao.Query, username string) error {
	dao := tx.User
	result, err := dao.WithContext(ctx).Where(dao.Username.Eq(username)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if result == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Username.Eq(username)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByUsername 根据username删除一条数据
func (u *UserRepo) DeleteOneByUsername(ctx context.Context, username string) error {
	dao := fkratos_user_dao.Use(u.db).User
	_, err := dao.WithContext(ctx).Where(dao.Username.Eq(username)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByUsername 根据username删除一条数据
func (u *UserRepo) DeleteOneByUsernameTx(ctx context.Context, tx *fkratos_user_dao.Query, username string) error {
	dao := tx.User
	_, err := dao.WithContext(ctx).Where(dao.Username.Eq(username)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByUsernames 根据usernames删除多条数据并清理缓存
func (u *UserRepo) DeleteMultiCacheByUsernames(ctx context.Context, usernames []string) error {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByUsernames 根据usernames删除多条数据并清理缓存
func (u *UserRepo) DeleteMultiCacheByUsernamesTx(ctx context.Context, tx *fkratos_user_dao.Query, usernames []string) error {
	dao := tx.User
	result, err := dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByUsernames 根据usernames删除多条数据
func (u *UserRepo) DeleteMultiByUsernames(ctx context.Context, usernames []string) error {
	dao := fkratos_user_dao.Use(u.db).User
	_, err := dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByUsernames 根据usernames删除多条数据
func (u *UserRepo) DeleteMultiByUsernamesTx(ctx context.Context, tx *fkratos_user_dao.Query, usernames []string) error {
	dao := tx.User
	_, err := dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByUID 根据UID删除一条数据并清理缓存
func (u *UserRepo) DeleteOneCacheByUID(ctx context.Context, UID string) error {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.UID.Eq(UID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if result == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.UID.Eq(UID)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByUID 根据UID删除一条数据并清理缓存
func (u *UserRepo) DeleteOneCacheByUIDTx(ctx context.Context, tx *fkratos_user_dao.Query, UID string) error {
	dao := tx.User
	result, err := dao.WithContext(ctx).Where(dao.UID.Eq(UID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if result == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.UID.Eq(UID)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByUID 根据UID删除一条数据
func (u *UserRepo) DeleteOneByUID(ctx context.Context, UID string) error {
	dao := fkratos_user_dao.Use(u.db).User
	_, err := dao.WithContext(ctx).Where(dao.UID.Eq(UID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByUID 根据UID删除一条数据
func (u *UserRepo) DeleteOneByUIDTx(ctx context.Context, tx *fkratos_user_dao.Query, UID string) error {
	dao := tx.User
	_, err := dao.WithContext(ctx).Where(dao.UID.Eq(UID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByUIDS 根据UIDS删除多条数据并清理缓存
func (u *UserRepo) DeleteMultiCacheByUIDS(ctx context.Context, UIDS []string) error {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.UID.In(UIDS...)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.UID.In(UIDS...)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByUIDS 根据UIDS删除多条数据并清理缓存
func (u *UserRepo) DeleteMultiCacheByUIDSTx(ctx context.Context, tx *fkratos_user_dao.Query, UIDS []string) error {
	dao := tx.User
	result, err := dao.WithContext(ctx).Where(dao.UID.In(UIDS...)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.UID.In(UIDS...)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByUIDS 根据UIDS删除多条数据
func (u *UserRepo) DeleteMultiByUIDS(ctx context.Context, UIDS []string) error {
	dao := fkratos_user_dao.Use(u.db).User
	_, err := dao.WithContext(ctx).Where(dao.UID.In(UIDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByUIDS 根据UIDS删除多条数据
func (u *UserRepo) DeleteMultiByUIDSTx(ctx context.Context, tx *fkratos_user_dao.Query, UIDS []string) error {
	dao := tx.User
	_, err := dao.WithContext(ctx).Where(dao.UID.In(UIDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (u *UserRepo) DeleteOneCacheByID(ctx context.Context, ID string) error {
	dao := fkratos_user_dao.Use(u.db).User
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
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (u *UserRepo) DeleteOneCacheByIDTx(ctx context.Context, tx *fkratos_user_dao.Query, ID string) error {
	dao := tx.User
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
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (u *UserRepo) DeleteOneByID(ctx context.Context, ID string) error {
	dao := fkratos_user_dao.Use(u.db).User
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (u *UserRepo) DeleteOneByIDTx(ctx context.Context, tx *fkratos_user_dao.Query, ID string) error {
	dao := tx.User
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (u *UserRepo) DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error {
	dao := fkratos_user_dao.Use(u.db).User
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
func (u *UserRepo) DeleteMultiCacheByIDSTx(ctx context.Context, tx *fkratos_user_dao.Query, IDS []string) error {
	dao := tx.User
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
func (u *UserRepo) DeleteMultiByIDS(ctx context.Context, IDS []string) error {
	dao := fkratos_user_dao.Use(u.db).User
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (u *UserRepo) DeleteMultiByIDSTx(ctx context.Context, tx *fkratos_user_dao.Query, IDS []string) error {
	dao := tx.User
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteUniqueIndexCache 删除唯一索引存在的缓存
func (u *UserRepo) DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_user_model.User) error {
	keys := make([]string, 0)
	for _, v := range data {
		keys = append(keys, u.cache.Key(cacheUserByUsernamePrefix, v.Username))
		keys = append(keys, u.cache.Key(cacheUserByUIDPrefix, v.UID))
		keys = append(keys, u.cache.Key(cacheUserByIDPrefix, v.ID))

	}
	err := u.cache.DelBatch(ctx, keys)
	if err != nil {
		return err
	}
	return nil
}

// FindOneCacheByUsername 根据username查询一条数据并设置缓存
func (u *UserRepo) FindOneCacheByUsername(ctx context.Context, username string) (*fkratos_user_model.User, error) {
	resp := new(fkratos_user_model.User)
	cacheKey := u.cache.Key(cacheUserByUsernamePrefix, username)
	cacheValue, err := u.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := fkratos_user_dao.Use(u.db).User
		result, err := dao.WithContext(ctx).Where(dao.Username.Eq(username)).First()
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

// FindOneByUsername 根据username查询一条数据
func (u *UserRepo) FindOneByUsername(ctx context.Context, username string) (*fkratos_user_model.User, error) {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.Username.Eq(username)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByUsernames 根据usernames查询多条数据并设置缓存
func (u *UserRepo) FindMultiCacheByUsernames(ctx context.Context, usernames []string) ([]*fkratos_user_model.User, error) {
	resp := make([]*fkratos_user_model.User, 0)
	cacheKeys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range usernames {
		cacheKey := u.cache.Key(cacheUserByUsernamePrefix, v)
		cacheKeys = append(cacheKeys, cacheKey)
		keyToParam[cacheKey] = v
	}
	cacheValue, err := u.cache.FetchBatch(ctx, cacheKeys, func(miss []string) (map[string]string, error) {
		parameters := make([]string, 0)
		for _, v := range miss {
			parameters = append(parameters, keyToParam[v])
		}
		dao := fkratos_user_dao.Use(u.db).User
		result, err := dao.WithContext(ctx).Where(dao.Username.In(parameters...)).Find()
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
			value[u.cache.Key(cacheUserByUsernamePrefix, v.Username)] = string(marshal)
		}
		return value, nil
	})
	if err != nil {
		return nil, err
	}
	for _, v := range cacheValue {
		tmp := new(fkratos_user_model.User)
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}

// FindMultiByUsernames 根据usernames查询多条数据
func (u *UserRepo) FindMultiByUsernames(ctx context.Context, usernames []string) ([]*fkratos_user_model.User, error) {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByUserGroupID 根据userGroupID查询多条数据
func (u *UserRepo) FindMultiByUserGroupID(ctx context.Context, userGroupID string) ([]*fkratos_user_model.User, error) {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.UserGroupID.Eq(userGroupID)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByUserGroupIDS 根据userGroupIDS查询多条数据
func (u *UserRepo) FindMultiByUserGroupIDS(ctx context.Context, userGroupIDS []string) ([]*fkratos_user_model.User, error) {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.UserGroupID.In(userGroupIDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindOneCacheByUID 根据UID查询一条数据并设置缓存
func (u *UserRepo) FindOneCacheByUID(ctx context.Context, UID string) (*fkratos_user_model.User, error) {
	resp := new(fkratos_user_model.User)
	cacheKey := u.cache.Key(cacheUserByUIDPrefix, UID)
	cacheValue, err := u.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := fkratos_user_dao.Use(u.db).User
		result, err := dao.WithContext(ctx).Where(dao.UID.Eq(UID)).First()
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

// FindOneByUID 根据UID查询一条数据
func (u *UserRepo) FindOneByUID(ctx context.Context, UID string) (*fkratos_user_model.User, error) {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.UID.Eq(UID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByUIDS 根据UIDS查询多条数据并设置缓存
func (u *UserRepo) FindMultiCacheByUIDS(ctx context.Context, UIDS []string) ([]*fkratos_user_model.User, error) {
	resp := make([]*fkratos_user_model.User, 0)
	cacheKeys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range UIDS {
		cacheKey := u.cache.Key(cacheUserByUIDPrefix, v)
		cacheKeys = append(cacheKeys, cacheKey)
		keyToParam[cacheKey] = v
	}
	cacheValue, err := u.cache.FetchBatch(ctx, cacheKeys, func(miss []string) (map[string]string, error) {
		parameters := make([]string, 0)
		for _, v := range miss {
			parameters = append(parameters, keyToParam[v])
		}
		dao := fkratos_user_dao.Use(u.db).User
		result, err := dao.WithContext(ctx).Where(dao.UID.In(parameters...)).Find()
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
			value[u.cache.Key(cacheUserByUIDPrefix, v.UID)] = string(marshal)
		}
		return value, nil
	})
	if err != nil {
		return nil, err
	}
	for _, v := range cacheValue {
		tmp := new(fkratos_user_model.User)
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}

// FindMultiByUIDS 根据UIDS查询多条数据
func (u *UserRepo) FindMultiByUIDS(ctx context.Context, UIDS []string) ([]*fkratos_user_model.User, error) {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.UID.In(UIDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindOneCacheByID 根据ID查询一条数据并设置缓存
func (u *UserRepo) FindOneCacheByID(ctx context.Context, ID string) (*fkratos_user_model.User, error) {
	resp := new(fkratos_user_model.User)
	cacheKey := u.cache.Key(cacheUserByIDPrefix, ID)
	cacheValue, err := u.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := fkratos_user_dao.Use(u.db).User
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
func (u *UserRepo) FindOneByID(ctx context.Context, ID string) (*fkratos_user_model.User, error) {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
func (u *UserRepo) FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*fkratos_user_model.User, error) {
	resp := make([]*fkratos_user_model.User, 0)
	cacheKeys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range IDS {
		cacheKey := u.cache.Key(cacheUserByIDPrefix, v)
		cacheKeys = append(cacheKeys, cacheKey)
		keyToParam[cacheKey] = v
	}
	cacheValue, err := u.cache.FetchBatch(ctx, cacheKeys, func(miss []string) (map[string]string, error) {
		parameters := make([]string, 0)
		for _, v := range miss {
			parameters = append(parameters, keyToParam[v])
		}
		dao := fkratos_user_dao.Use(u.db).User
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
			value[u.cache.Key(cacheUserByIDPrefix, v.ID)] = string(marshal)
		}
		return value, nil
	})
	if err != nil {
		return nil, err
	}
	for _, v := range cacheValue {
		tmp := new(fkratos_user_model.User)
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}

// FindMultiByIDS 根据IDS查询多条数据
func (u *UserRepo) FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_user_model.User, error) {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByPhone 根据phone查询多条数据
func (u *UserRepo) FindMultiByPhone(ctx context.Context, phone string) ([]*fkratos_user_model.User, error) {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.Phone.Eq(phone)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByPhones 根据phones查询多条数据
func (u *UserRepo) FindMultiByPhones(ctx context.Context, phones []string) ([]*fkratos_user_model.User, error) {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.Phone.In(phones...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByEmail 根据email查询多条数据
func (u *UserRepo) FindMultiByEmail(ctx context.Context, email string) ([]*fkratos_user_model.User, error) {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.Email.Eq(email)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByEmails 根据emails查询多条数据
func (u *UserRepo) FindMultiByEmails(ctx context.Context, emails []string) ([]*fkratos_user_model.User, error) {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.Email.In(emails...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByPaginator 查询分页数据(通用)
func (u *UserRepo) FindMultiByPaginator(ctx context.Context, paginatorReq *orm.PaginatorReq) ([]*fkratos_user_model.User, *orm.PaginatorReply, error) {
	result := make([]*fkratos_user_model.User, 0)
	var total int64
	queryStr, args, err := paginatorReq.ConvertToGormConditions()
	if err != nil {
		return result, nil, err
	}
	err = u.db.WithContext(ctx).Model(&fkratos_user_model.User{}).Select([]string{"id"}).Where(queryStr, args...).Count(&total).Error
	if err != nil {
		return result, nil, err
	}
	if total == 0 {
		return result, nil, nil
	}
	query := u.db.WithContext(ctx)
	order := paginatorReq.ConvertToOrder()
	if order != "" {
		query = query.Order(order)
	}
	paginatorReply := paginatorReq.ConvertToPage(int(total))
	err = query.Limit(paginatorReply.Limit).Offset(paginatorReply.Offset).Where(queryStr, args...).Find(&result).Error
	if err != nil {
		return result, nil, err
	}
	return result, paginatorReply, err
}
