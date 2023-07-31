// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.

package fkratos_user_repo

import (
	"context"
	"encoding/json"
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_dao"
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_model"
	"time"

	"github.com/fzf-labs/fpkg/cache/cachekey"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var _ IUserRepo = (*userRepo)(nil)

var (
	// 缓存管理器
	cacheKeyUserManage = cachekey.NewKeyManage("userRepo")
	// 只针对唯一索引做缓存
	CacheUserByUsername = cacheKeyUserManage.AddKey("CacheUserByUsername", time.Hour*24, "CacheUserByUsername")
	CacheUserByID       = cacheKeyUserManage.AddKey("CacheUserByID", time.Hour*24, "CacheUserByID")
	CacheUserByPhone    = cacheKeyUserManage.AddKey("CacheUserByPhone", time.Hour*24, "CacheUserByPhone")
	CacheUserByEmail    = cacheKeyUserManage.AddKey("CacheUserByEmail", time.Hour*24, "CacheUserByEmail")
)

type (
	IUserRepo interface {
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *fkratos_user_model.User) error
		// UpdateOne 更新一条数据
		UpdateOne(ctx context.Context, data *fkratos_user_model.User) error
		// DeleteOneCacheByUsername 根据username删除一条数据并清理缓存
		DeleteOneCacheByUsername(ctx context.Context, username string) error
		// DeleteMultiCacheByUsernames 根据Usernames删除多条数据并清理缓存
		DeleteMultiCacheByUsernames(ctx context.Context, usernames []string) error
		// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
		DeleteOneCacheByID(ctx context.Context, ID string) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
		DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error
		// DeleteOneCacheByPhone 根据phone删除一条数据并清理缓存
		DeleteOneCacheByPhone(ctx context.Context, phone string) error
		// DeleteMultiCacheByPhones 根据Phones删除多条数据并清理缓存
		DeleteMultiCacheByPhones(ctx context.Context, phones []string) error
		// DeleteOneCacheByEmail 根据email删除一条数据并清理缓存
		DeleteOneCacheByEmail(ctx context.Context, email string) error
		// DeleteMultiCacheByEmails 根据Emails删除多条数据并清理缓存
		DeleteMultiCacheByEmails(ctx context.Context, emails []string) error
		// DeleteUniqueIndexCache 删除唯一索引存在的缓存
		DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_user_model.User) error
		// FindOneByUsername 根据username查询一条数据并设置缓存
		FindOneByUsername(ctx context.Context, username string) (*fkratos_user_model.User, error)
		// FindMultiByUsernames 根据usernames查询多条数据并设置缓存
		FindMultiByUsernames(ctx context.Context, usernames []string) ([]*fkratos_user_model.User, error)
		// FindMultiByStatuses 根据statuses查询多条数据
		FindMultiByStatuses(ctx context.Context, statuses []int16) ([]*fkratos_user_model.User, error)
		// FindOneByID 根据ID查询一条数据并设置缓存
		FindOneByID(ctx context.Context, ID string) (*fkratos_user_model.User, error)
		// FindMultiByIDS 根据IDS查询多条数据并设置缓存
		FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_user_model.User, error)
		// FindOneByPhone 根据phone查询一条数据并设置缓存
		FindOneByPhone(ctx context.Context, phone string) (*fkratos_user_model.User, error)
		// FindMultiByPhones 根据phones查询多条数据并设置缓存
		FindMultiByPhones(ctx context.Context, phones []string) ([]*fkratos_user_model.User, error)
		// FindOneByEmail 根据email查询一条数据并设置缓存
		FindOneByEmail(ctx context.Context, email string) (*fkratos_user_model.User, error)
		// FindMultiByEmails 根据emails查询多条数据并设置缓存
		FindMultiByEmails(ctx context.Context, emails []string) ([]*fkratos_user_model.User, error)
	}

	userRepo struct {
		db    *gorm.DB
		redis *redis.Client
	}
)

func NewUserRepo(db *gorm.DB, redis *redis.Client) IUserRepo {
	return &userRepo{db: db, redis: redis}
}

// CreateOne 创建一条数据
func (r *userRepo) CreateOne(ctx context.Context, data *fkratos_user_model.User) error {
	dao := fkratos_user_dao.Use(r.db).User
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne 更新一条数据
func (r *userRepo) UpdateOne(ctx context.Context, data *fkratos_user_model.User) error {
	dao := fkratos_user_dao.Use(r.db).User
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = r.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{data})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByUsername 根据username删除一条数据并清理缓存
func (r *userRepo) DeleteOneCacheByUsername(ctx context.Context, username string) error {
	dao := fkratos_user_dao.Use(r.db).User
	first, err := dao.WithContext(ctx).Where(dao.Username.Eq(username)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if first == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Username.Eq(username)).Delete()
	if err != nil {
		return err
	}
	err = r.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByUsernames 根据usernames删除多条数据并清理缓存
func (r *userRepo) DeleteMultiCacheByUsernames(ctx context.Context, usernames []string) error {
	dao := fkratos_user_dao.Use(r.db).User
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
	err = r.DeleteUniqueIndexCache(ctx, list)
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (r *userRepo) DeleteOneCacheByID(ctx context.Context, ID string) error {
	dao := fkratos_user_dao.Use(r.db).User
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
	err = r.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (r *userRepo) DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error {
	dao := fkratos_user_dao.Use(r.db).User
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

// DeleteOneCacheByPhone 根据phone删除一条数据并清理缓存
func (r *userRepo) DeleteOneCacheByPhone(ctx context.Context, phone string) error {
	dao := fkratos_user_dao.Use(r.db).User
	first, err := dao.WithContext(ctx).Where(dao.Phone.Eq(phone)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if first == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Phone.Eq(phone)).Delete()
	if err != nil {
		return err
	}
	err = r.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByPhones 根据phones删除多条数据并清理缓存
func (r *userRepo) DeleteMultiCacheByPhones(ctx context.Context, phones []string) error {
	dao := fkratos_user_dao.Use(r.db).User
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
	err = r.DeleteUniqueIndexCache(ctx, list)
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByEmail 根据email删除一条数据并清理缓存
func (r *userRepo) DeleteOneCacheByEmail(ctx context.Context, email string) error {
	dao := fkratos_user_dao.Use(r.db).User
	first, err := dao.WithContext(ctx).Where(dao.Email.Eq(email)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if first == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Email.Eq(email)).Delete()
	if err != nil {
		return err
	}
	err = r.DeleteUniqueIndexCache(ctx, []*fkratos_user_model.User{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByEmails 根据emails删除多条数据并清理缓存
func (r *userRepo) DeleteMultiCacheByEmails(ctx context.Context, emails []string) error {
	dao := fkratos_user_dao.Use(r.db).User
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
	err = r.DeleteUniqueIndexCache(ctx, list)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUniqueIndexCache 删除唯一索引存在的缓存
func (r *userRepo) DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_user_model.User) error {
	var err error
	cacheUserByUsername := CacheUserByUsername.NewSingleKey(r.redis)
	cacheUserByID := CacheUserByID.NewSingleKey(r.redis)
	cacheUserByPhone := CacheUserByPhone.NewSingleKey(r.redis)
	cacheUserByEmail := CacheUserByEmail.NewSingleKey(r.redis)

	for _, v := range data {
		err = cacheUserByUsername.SingleCacheDel(ctx, cacheUserByUsername.BuildKey(v.Username))
		if err != nil {
			return err
		}
		err = cacheUserByID.SingleCacheDel(ctx, cacheUserByID.BuildKey(v.ID))
		if err != nil {
			return err
		}
		err = cacheUserByPhone.SingleCacheDel(ctx, cacheUserByPhone.BuildKey(v.Phone))
		if err != nil {
			return err
		}
		err = cacheUserByEmail.SingleCacheDel(ctx, cacheUserByEmail.BuildKey(v.Email))
		if err != nil {
			return err
		}

	}
	return nil
}

// FindOneByUsername 根据username查询一条数据并设置缓存
func (r *userRepo) FindOneByUsername(ctx context.Context, username string) (*fkratos_user_model.User, error) {
	resp := new(fkratos_user_model.User)
	cache := CacheUserByUsername.NewSingleKey(r.redis)
	cacheValue, err := cache.SingleCache(ctx, username, func() (string, error) {
		dao := fkratos_user_dao.Use(r.db).User
		result, err := dao.WithContext(ctx).Where(dao.Username.Eq(username)).First()
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

// FindMultiByUsernames 根据usernames查询多条数据并设置缓存
func (r *userRepo) FindMultiByUsernames(ctx context.Context, usernames []string) ([]*fkratos_user_model.User, error) {
	resp := make([]*fkratos_user_model.User, 0)
	cacheKey := CacheUserByUsername.NewBatchKey(r.redis)
	cacheValue, err := cacheKey.BatchKeyCache(ctx, usernames, func() (map[string]string, error) {
		dao := fkratos_user_dao.Use(r.db).User
		result, err := dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Find()
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
		value := make(map[string]string)
		for _, v := range result {
			marshal, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			value[v.Username] = string(marshal)
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

// FindMultiByStatuses 根据statuses查询多条数据
func (r *userRepo) FindMultiByStatuses(ctx context.Context, statuses []int16) ([]*fkratos_user_model.User, error) {
	dao := fkratos_user_dao.Use(r.db).User
	result, err := dao.WithContext(ctx).Where(dao.Status.In(statuses...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindOneByID 根据ID查询一条数据并设置缓存
func (r *userRepo) FindOneByID(ctx context.Context, ID string) (*fkratos_user_model.User, error) {
	resp := new(fkratos_user_model.User)
	cache := CacheUserByID.NewSingleKey(r.redis)
	cacheValue, err := cache.SingleCache(ctx, ID, func() (string, error) {
		dao := fkratos_user_dao.Use(r.db).User
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
func (r *userRepo) FindMultiByIDS(ctx context.Context, IDS []string) ([]*fkratos_user_model.User, error) {
	resp := make([]*fkratos_user_model.User, 0)
	cacheKey := CacheUserByID.NewBatchKey(r.redis)
	cacheValue, err := cacheKey.BatchKeyCache(ctx, IDS, func() (map[string]string, error) {
		dao := fkratos_user_dao.Use(r.db).User
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
		tmp := new(fkratos_user_model.User)
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}

// FindOneByPhone 根据phone查询一条数据并设置缓存
func (r *userRepo) FindOneByPhone(ctx context.Context, phone string) (*fkratos_user_model.User, error) {
	resp := new(fkratos_user_model.User)
	cache := CacheUserByPhone.NewSingleKey(r.redis)
	cacheValue, err := cache.SingleCache(ctx, phone, func() (string, error) {
		dao := fkratos_user_dao.Use(r.db).User
		result, err := dao.WithContext(ctx).Where(dao.Phone.Eq(phone)).First()
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

// FindMultiByPhones 根据phones查询多条数据并设置缓存
func (r *userRepo) FindMultiByPhones(ctx context.Context, phones []string) ([]*fkratos_user_model.User, error) {
	resp := make([]*fkratos_user_model.User, 0)
	cacheKey := CacheUserByPhone.NewBatchKey(r.redis)
	cacheValue, err := cacheKey.BatchKeyCache(ctx, phones, func() (map[string]string, error) {
		dao := fkratos_user_dao.Use(r.db).User
		result, err := dao.WithContext(ctx).Where(dao.Phone.In(phones...)).Find()
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
		value := make(map[string]string)
		for _, v := range result {
			marshal, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			value[v.Phone] = string(marshal)
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

// FindOneByEmail 根据email查询一条数据并设置缓存
func (r *userRepo) FindOneByEmail(ctx context.Context, email string) (*fkratos_user_model.User, error) {
	resp := new(fkratos_user_model.User)
	cache := CacheUserByEmail.NewSingleKey(r.redis)
	cacheValue, err := cache.SingleCache(ctx, email, func() (string, error) {
		dao := fkratos_user_dao.Use(r.db).User
		result, err := dao.WithContext(ctx).Where(dao.Email.Eq(email)).First()
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

// FindMultiByEmails 根据emails查询多条数据并设置缓存
func (r *userRepo) FindMultiByEmails(ctx context.Context, emails []string) ([]*fkratos_user_model.User, error) {
	resp := make([]*fkratos_user_model.User, 0)
	cacheKey := CacheUserByEmail.NewBatchKey(r.redis)
	cacheValue, err := cacheKey.BatchKeyCache(ctx, emails, func() (map[string]string, error) {
		dao := fkratos_user_dao.Use(r.db).User
		result, err := dao.WithContext(ctx).Where(dao.Email.In(emails...)).Find()
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
		value := make(map[string]string)
		for _, v := range result {
			marshal, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			value[v.Email] = string(marshal)
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