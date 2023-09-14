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
	"gorm.io/gorm"
)

var _ IUserRepo = (*UserRepo)(nil)

var (
	cacheUserByUsernamePrefix = "DBCache:fkratos_user:UserByUsername"
	cacheUserByIDPrefix       = "DBCache:fkratos_user:UserByID"
	cacheUserByPhonePrefix    = "DBCache:fkratos_user:UserByPhone"
	cacheUserByEmailPrefix    = "DBCache:fkratos_user:UserByEmail"
)

type (
	IUserRepo interface {
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *fkratos_user_model.User) error
		// CreateOneByTx 创建一条数据(事务)
		CreateOneByTx(ctx context.Context, tx *fkratos_user_dao.Query, data *fkratos_user_model.User) error
		// CreateBatch 批量创建数据
		CreateBatch(ctx context.Context, data []*fkratos_user_model.User, batchSize int) error
		// UpdateOne 更新一条数据
		UpdateOne(ctx context.Context, data *fkratos_user_model.User) error
		// UpdateOne 更新一条数据(事务)
		UpdateOneByTx(ctx context.Context, tx *fkratos_user_dao.Query, data *fkratos_user_model.User) error
		// FindOneCacheByUsername 根据username查询一条数据并设置缓存
		FindOneCacheByUsername(ctx context.Context, username string) (*fkratos_user_model.User, error)
		// FindOneByUsername 根据username查询一条数据
		FindOneByUsername(ctx context.Context, username string) (*fkratos_user_model.User, error)
		// FindMultiCacheByUsernames 根据usernames查询多条数据并设置缓存
		FindMultiCacheByUsernames(ctx context.Context, usernames []string) ([]*fkratos_user_model.User, error)
		// FindMultiByUsernames 根据usernames查询多条数据
		FindMultiByUsernames(ctx context.Context, usernames []string) ([]*fkratos_user_model.User, error)
		// FindMultiByStatus 根据status查询多条数据
		FindMultiByStatus(ctx context.Context, status int16) ([]*fkratos_user_model.User, error)
		// FindMultiByStatuses 根据statuses查询多条数据
		FindMultiByStatuses(ctx context.Context, statuses []int16) ([]*fkratos_user_model.User, error)
		// FindOneCacheByID 根据ID查询一条数据并设置缓存
		FindOneCacheByID(ctx context.Context, ID string) (*fkratos_user_model.User, error)
		// FindOneByID 根据ID查询一条数据
		FindOneByID(ctx context.Context, ID string) (*fkratos_user_model.User, error)
		// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
		FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*fkratos_user_model.User, error)
		// FindMultiByIDS 根据IDS查询多条数据
		FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_user_model.User, error)
		// FindOneCacheByPhone 根据phone查询一条数据并设置缓存
		FindOneCacheByPhone(ctx context.Context, phone string) (*fkratos_user_model.User, error)
		// FindOneByPhone 根据phone查询一条数据
		FindOneByPhone(ctx context.Context, phone string) (*fkratos_user_model.User, error)
		// FindMultiCacheByPhones 根据phones查询多条数据并设置缓存
		FindMultiCacheByPhones(ctx context.Context, phones []string) ([]*fkratos_user_model.User, error)
		// FindMultiByPhones 根据phones查询多条数据
		FindMultiByPhones(ctx context.Context, phones []string) ([]*fkratos_user_model.User, error)
		// FindOneCacheByEmail 根据email查询一条数据并设置缓存
		FindOneCacheByEmail(ctx context.Context, email string) (*fkratos_user_model.User, error)
		// FindOneByEmail 根据email查询一条数据
		FindOneByEmail(ctx context.Context, email string) (*fkratos_user_model.User, error)
		// FindMultiCacheByEmails 根据emails查询多条数据并设置缓存
		FindMultiCacheByEmails(ctx context.Context, emails []string) ([]*fkratos_user_model.User, error)
		// FindMultiByEmails 根据emails查询多条数据
		FindMultiByEmails(ctx context.Context, emails []string) ([]*fkratos_user_model.User, error)
		// FindMultiByPaginator 查询分页数据(通用)
		FindMultiByPaginator(ctx context.Context, params *orm.PaginatorParams) ([]*fkratos_user_model.User, int64, error)
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
		// DeleteOneCacheByPhone 根据phone删除一条数据并清理缓存
		DeleteOneCacheByPhone(ctx context.Context, phone string) error
		// DeleteOneCacheByPhone 根据phone删除一条数据并清理缓存
		DeleteOneCacheByPhoneTx(ctx context.Context, tx *fkratos_user_dao.Query, phone string) error
		// DeleteOneByPhone 根据phone删除一条数据
		DeleteOneByPhone(ctx context.Context, phone string) error
		// DeleteOneByPhone 根据phone删除一条数据
		DeleteOneByPhoneTx(ctx context.Context, tx *fkratos_user_dao.Query, phone string) error
		// DeleteMultiCacheByPhones 根据Phones删除多条数据并清理缓存
		DeleteMultiCacheByPhones(ctx context.Context, phones []string) error
		// DeleteMultiCacheByPhones 根据Phones删除多条数据并清理缓存
		DeleteMultiCacheByPhonesTx(ctx context.Context, tx *fkratos_user_dao.Query, phones []string) error
		// DeleteMultiByPhones 根据Phones删除多条数据
		DeleteMultiByPhones(ctx context.Context, phones []string) error
		// DeleteMultiByPhones 根据Phones删除多条数据
		DeleteMultiByPhonesTx(ctx context.Context, tx *fkratos_user_dao.Query, phones []string) error
		// DeleteOneCacheByEmail 根据email删除一条数据并清理缓存
		DeleteOneCacheByEmail(ctx context.Context, email string) error
		// DeleteOneCacheByEmail 根据email删除一条数据并清理缓存
		DeleteOneCacheByEmailTx(ctx context.Context, tx *fkratos_user_dao.Query, email string) error
		// DeleteOneByEmail 根据email删除一条数据
		DeleteOneByEmail(ctx context.Context, email string) error
		// DeleteOneByEmail 根据email删除一条数据
		DeleteOneByEmailTx(ctx context.Context, tx *fkratos_user_dao.Query, email string) error
		// DeleteMultiCacheByEmails 根据Emails删除多条数据并清理缓存
		DeleteMultiCacheByEmails(ctx context.Context, emails []string) error
		// DeleteMultiCacheByEmails 根据Emails删除多条数据并清理缓存
		DeleteMultiCacheByEmailsTx(ctx context.Context, tx *fkratos_user_dao.Query, emails []string) error
		// DeleteMultiByEmails 根据Emails删除多条数据
		DeleteMultiByEmails(ctx context.Context, emails []string) error
		// DeleteMultiByEmails 根据Emails删除多条数据
		DeleteMultiByEmailsTx(ctx context.Context, tx *fkratos_user_dao.Query, emails []string) error
		// DeleteUniqueIndexCache 删除唯一索引存在的缓存
		DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_user_model.User) error
	}
	IUserCache interface {
		Key(fields ...any) string
		Fetch(ctx context.Context, key string, fn func() (string, error)) (string, error)
		FetchBatch(ctx context.Context, keys []string, fn func(miss []string) (map[string]string, error)) (map[string]string, error)
		Del(ctx context.Context, key string) error
		DelBatch(ctx context.Context, keys []string) error
	}
	UserRepo struct {
		db    *gorm.DB
		cache IUserCache
	}
)

func NewUserRepo(db *gorm.DB, cache IUserCache) *UserRepo {
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

// DeleteOneCacheByUsername 根据username删除一条数据并清理缓存
func (u *UserRepo) DeleteOneCacheByUsername(ctx context.Context, username string) error {
	dao := fkratos_user_dao.Use(u.db).User
	first, err := dao.WithContext(ctx).Where(dao.Username.Eq(username)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if first == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Username.Eq(username)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByUsername 根据username删除一条数据并清理缓存
func (u *UserRepo) DeleteOneCacheByUsernameTx(ctx context.Context, tx *fkratos_user_dao.Query, username string) error {
	dao := tx.User
	first, err := dao.WithContext(ctx).Where(dao.Username.Eq(username)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if first == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Username.Eq(username)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{first})
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
	list, err := dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Find()
	if err != nil {
		return err
	}
	if len(list) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, list)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByUsernames 根据usernames删除多条数据并清理缓存
func (u *UserRepo) DeleteMultiCacheByUsernamesTx(ctx context.Context, tx *fkratos_user_dao.Query, usernames []string) error {
	dao := tx.User
	list, err := dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Find()
	if err != nil {
		return err
	}
	if len(list) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, list)
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

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (u *UserRepo) DeleteOneCacheByID(ctx context.Context, ID string) error {
	dao := fkratos_user_dao.Use(u.db).User
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
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (u *UserRepo) DeleteOneCacheByIDTx(ctx context.Context, tx *fkratos_user_dao.Query, ID string) error {
	dao := tx.User
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
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{first})
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
	err = u.DeleteUniqueIndexCache(ctx, list)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (u *UserRepo) DeleteMultiCacheByIDSTx(ctx context.Context, tx *fkratos_user_dao.Query, IDS []string) error {
	dao := tx.User
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
	err = u.DeleteUniqueIndexCache(ctx, list)
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

// DeleteOneCacheByPhone 根据phone删除一条数据并清理缓存
func (u *UserRepo) DeleteOneCacheByPhone(ctx context.Context, phone string) error {
	dao := fkratos_user_dao.Use(u.db).User
	first, err := dao.WithContext(ctx).Where(dao.Phone.Eq(phone)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if first == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Phone.Eq(phone)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByPhone 根据phone删除一条数据并清理缓存
func (u *UserRepo) DeleteOneCacheByPhoneTx(ctx context.Context, tx *fkratos_user_dao.Query, phone string) error {
	dao := tx.User
	first, err := dao.WithContext(ctx).Where(dao.Phone.Eq(phone)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if first == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Phone.Eq(phone)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByPhone 根据phone删除一条数据
func (u *UserRepo) DeleteOneByPhone(ctx context.Context, phone string) error {
	dao := fkratos_user_dao.Use(u.db).User
	_, err := dao.WithContext(ctx).Where(dao.Phone.Eq(phone)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByPhone 根据phone删除一条数据
func (u *UserRepo) DeleteOneByPhoneTx(ctx context.Context, tx *fkratos_user_dao.Query, phone string) error {
	dao := tx.User
	_, err := dao.WithContext(ctx).Where(dao.Phone.Eq(phone)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByPhones 根据phones删除多条数据并清理缓存
func (u *UserRepo) DeleteMultiCacheByPhones(ctx context.Context, phones []string) error {
	dao := fkratos_user_dao.Use(u.db).User
	list, err := dao.WithContext(ctx).Where(dao.Phone.In(phones...)).Find()
	if err != nil {
		return err
	}
	if len(list) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Phone.In(phones...)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, list)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByPhones 根据phones删除多条数据并清理缓存
func (u *UserRepo) DeleteMultiCacheByPhonesTx(ctx context.Context, tx *fkratos_user_dao.Query, phones []string) error {
	dao := tx.User
	list, err := dao.WithContext(ctx).Where(dao.Phone.In(phones...)).Find()
	if err != nil {
		return err
	}
	if len(list) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Phone.In(phones...)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, list)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByPhones 根据phones删除多条数据
func (u *UserRepo) DeleteMultiByPhones(ctx context.Context, phones []string) error {
	dao := fkratos_user_dao.Use(u.db).User
	_, err := dao.WithContext(ctx).Where(dao.Phone.In(phones...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByPhones 根据phones删除多条数据
func (u *UserRepo) DeleteMultiByPhonesTx(ctx context.Context, tx *fkratos_user_dao.Query, phones []string) error {
	dao := tx.User
	_, err := dao.WithContext(ctx).Where(dao.Phone.In(phones...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByEmail 根据email删除一条数据并清理缓存
func (u *UserRepo) DeleteOneCacheByEmail(ctx context.Context, email string) error {
	dao := fkratos_user_dao.Use(u.db).User
	first, err := dao.WithContext(ctx).Where(dao.Email.Eq(email)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if first == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Email.Eq(email)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByEmail 根据email删除一条数据并清理缓存
func (u *UserRepo) DeleteOneCacheByEmailTx(ctx context.Context, tx *fkratos_user_dao.Query, email string) error {
	dao := tx.User
	first, err := dao.WithContext(ctx).Where(dao.Email.Eq(email)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if first == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Email.Eq(email)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByEmail 根据email删除一条数据
func (u *UserRepo) DeleteOneByEmail(ctx context.Context, email string) error {
	dao := fkratos_user_dao.Use(u.db).User
	_, err := dao.WithContext(ctx).Where(dao.Email.Eq(email)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByEmail 根据email删除一条数据
func (u *UserRepo) DeleteOneByEmailTx(ctx context.Context, tx *fkratos_user_dao.Query, email string) error {
	dao := tx.User
	_, err := dao.WithContext(ctx).Where(dao.Email.Eq(email)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByEmails 根据emails删除多条数据并清理缓存
func (u *UserRepo) DeleteMultiCacheByEmails(ctx context.Context, emails []string) error {
	dao := fkratos_user_dao.Use(u.db).User
	list, err := dao.WithContext(ctx).Where(dao.Email.In(emails...)).Find()
	if err != nil {
		return err
	}
	if len(list) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Email.In(emails...)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, list)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByEmails 根据emails删除多条数据并清理缓存
func (u *UserRepo) DeleteMultiCacheByEmailsTx(ctx context.Context, tx *fkratos_user_dao.Query, emails []string) error {
	dao := tx.User
	list, err := dao.WithContext(ctx).Where(dao.Email.In(emails...)).Find()
	if err != nil {
		return err
	}
	if len(list) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Email.In(emails...)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, list)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByEmails 根据emails删除多条数据
func (u *UserRepo) DeleteMultiByEmails(ctx context.Context, emails []string) error {
	dao := fkratos_user_dao.Use(u.db).User
	_, err := dao.WithContext(ctx).Where(dao.Email.In(emails...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByEmails 根据emails删除多条数据
func (u *UserRepo) DeleteMultiByEmailsTx(ctx context.Context, tx *fkratos_user_dao.Query, emails []string) error {
	dao := tx.User
	_, err := dao.WithContext(ctx).Where(dao.Email.In(emails...)).Delete()
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
		keys = append(keys, u.cache.Key(cacheUserByIDPrefix, v.ID))
		keys = append(keys, u.cache.Key(cacheUserByPhonePrefix, v.Phone))
		keys = append(keys, u.cache.Key(cacheUserByEmailPrefix, v.Email))

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
	key := u.cache.Key(cacheUserByUsernamePrefix, username)
	cacheValue, err := u.cache.Fetch(ctx, key, func() (string, error) {
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
	keys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range usernames {
		key := u.cache.Key(cacheUserByUsernamePrefix, v)
		keys = append(keys, key)
		keyToParam[key] = v
	}
	cacheValue, err := u.cache.FetchBatch(ctx, keys, func(miss []string) (map[string]string, error) {
		params := make([]string, 0)
		for _, v := range miss {
			params = append(params, keyToParam[v])
		}
		dao := fkratos_user_dao.Use(u.db).User
		result, err := dao.WithContext(ctx).Where(dao.Username.In(params...)).Find()
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

// FindMultiByStatus 根据status查询多条数据
func (u *UserRepo) FindMultiByStatus(ctx context.Context, status int16) ([]*fkratos_user_model.User, error) {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.Status.Eq(status)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByStatuses 根据statuses查询多条数据
func (u *UserRepo) FindMultiByStatuses(ctx context.Context, statuses []int16) ([]*fkratos_user_model.User, error) {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.Status.In(statuses...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindOneCacheByID 根据ID查询一条数据并设置缓存
func (u *UserRepo) FindOneCacheByID(ctx context.Context, ID string) (*fkratos_user_model.User, error) {
	resp := new(fkratos_user_model.User)
	key := u.cache.Key(cacheUserByIDPrefix, ID)
	cacheValue, err := u.cache.Fetch(ctx, key, func() (string, error) {
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
	keys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range IDS {
		key := u.cache.Key(cacheUserByIDPrefix, v)
		keys = append(keys, key)
		keyToParam[key] = v
	}
	cacheValue, err := u.cache.FetchBatch(ctx, keys, func(miss []string) (map[string]string, error) {
		params := make([]string, 0)
		for _, v := range miss {
			params = append(params, keyToParam[v])
		}
		dao := fkratos_user_dao.Use(u.db).User
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

// FindOneCacheByPhone 根据phone查询一条数据并设置缓存
func (u *UserRepo) FindOneCacheByPhone(ctx context.Context, phone string) (*fkratos_user_model.User, error) {
	resp := new(fkratos_user_model.User)
	key := u.cache.Key(cacheUserByPhonePrefix, phone)
	cacheValue, err := u.cache.Fetch(ctx, key, func() (string, error) {
		dao := fkratos_user_dao.Use(u.db).User
		result, err := dao.WithContext(ctx).Where(dao.Phone.Eq(phone)).First()
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

// FindOneByPhone 根据phone查询一条数据
func (u *UserRepo) FindOneByPhone(ctx context.Context, phone string) (*fkratos_user_model.User, error) {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.Phone.Eq(phone)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByPhones 根据phones查询多条数据并设置缓存
func (u *UserRepo) FindMultiCacheByPhones(ctx context.Context, phones []string) ([]*fkratos_user_model.User, error) {
	resp := make([]*fkratos_user_model.User, 0)
	keys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range phones {
		key := u.cache.Key(cacheUserByPhonePrefix, v)
		keys = append(keys, key)
		keyToParam[key] = v
	}
	cacheValue, err := u.cache.FetchBatch(ctx, keys, func(miss []string) (map[string]string, error) {
		params := make([]string, 0)
		for _, v := range miss {
			params = append(params, keyToParam[v])
		}
		dao := fkratos_user_dao.Use(u.db).User
		result, err := dao.WithContext(ctx).Where(dao.Phone.In(params...)).Find()
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
			value[u.cache.Key(cacheUserByPhonePrefix, v.Phone)] = string(marshal)
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

// FindMultiByPhones 根据phones查询多条数据
func (u *UserRepo) FindMultiByPhones(ctx context.Context, phones []string) ([]*fkratos_user_model.User, error) {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.Phone.In(phones...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindOneCacheByEmail 根据email查询一条数据并设置缓存
func (u *UserRepo) FindOneCacheByEmail(ctx context.Context, email string) (*fkratos_user_model.User, error) {
	resp := new(fkratos_user_model.User)
	key := u.cache.Key(cacheUserByEmailPrefix, email)
	cacheValue, err := u.cache.Fetch(ctx, key, func() (string, error) {
		dao := fkratos_user_dao.Use(u.db).User
		result, err := dao.WithContext(ctx).Where(dao.Email.Eq(email)).First()
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

// FindOneByEmail 根据email查询一条数据
func (u *UserRepo) FindOneByEmail(ctx context.Context, email string) (*fkratos_user_model.User, error) {
	dao := fkratos_user_dao.Use(u.db).User
	result, err := dao.WithContext(ctx).Where(dao.Email.Eq(email)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByEmails 根据emails查询多条数据并设置缓存
func (u *UserRepo) FindMultiCacheByEmails(ctx context.Context, emails []string) ([]*fkratos_user_model.User, error) {
	resp := make([]*fkratos_user_model.User, 0)
	keys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range emails {
		key := u.cache.Key(cacheUserByEmailPrefix, v)
		keys = append(keys, key)
		keyToParam[key] = v
	}
	cacheValue, err := u.cache.FetchBatch(ctx, keys, func(miss []string) (map[string]string, error) {
		params := make([]string, 0)
		for _, v := range miss {
			params = append(params, keyToParam[v])
		}
		dao := fkratos_user_dao.Use(u.db).User
		result, err := dao.WithContext(ctx).Where(dao.Email.In(params...)).Find()
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
			value[u.cache.Key(cacheUserByEmailPrefix, v.Email)] = string(marshal)
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
func (u *UserRepo) FindMultiByPaginator(ctx context.Context, params *orm.PaginatorParams) ([]*fkratos_user_model.User, int64, error) {
	result := make([]*fkratos_user_model.User, 0)
	var total int64
	queryStr, args, err := params.ConvertToGormConditions()
	if err != nil {
		return nil, 0, err
	}
	err = u.db.WithContext(ctx).Model(&fkratos_user_model.User{}).Select([]string{"id"}).Where(queryStr, args...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, total, nil
	}
	query := u.db.WithContext(ctx)
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
