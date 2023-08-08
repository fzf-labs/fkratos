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
	"time"

	"github.com/dtm-labs/rockscache"
	"github.com/fzf-labs/fpkg/cache/cachekey"
	"github.com/fzf-labs/fpkg/conv"
	"gorm.io/gorm"
)

var _ ISensitiveWordRepo = (*SensitiveWordRepo)(nil)

var (
	// 缓存管理器
	cacheKeySensitiveWordManage = cachekey.NewKeyManage("SensitiveWordRepo")
	// 只针对唯一索引做缓存
	CacheSensitiveWordByWord = cacheKeySensitiveWordManage.AddKey("CacheSensitiveWordByWord", time.Hour*24, "CacheSensitiveWordByWord")
	CacheSensitiveWordByID   = cacheKeySensitiveWordManage.AddKey("CacheSensitiveWordByID", time.Hour*24, "CacheSensitiveWordByID")
)

type (
	ISensitiveWordRepo interface {
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *fkratos_common_model.SensitiveWord) error
		// UpdateOne 更新一条数据
		UpdateOne(ctx context.Context, data *fkratos_common_model.SensitiveWord) error
		// DeleteOneCacheByWord 根据word删除一条数据并清理缓存
		DeleteOneCacheByWord(ctx context.Context, word string) error
		// DeleteMultiCacheByWords 根据Words删除多条数据并清理缓存
		DeleteMultiCacheByWords(ctx context.Context, words []string) error
		// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
		DeleteOneCacheByID(ctx context.Context, ID string) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
		DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error
		// DeleteUniqueIndexCache 删除唯一索引存在的缓存
		DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_common_model.SensitiveWord) error
		// FindOneCacheByWord 根据word查询一条数据并设置缓存
		FindOneCacheByWord(ctx context.Context, word string) (*fkratos_common_model.SensitiveWord, error)
		// FindMultiCacheByWords 根据words查询多条数据并设置缓存
		FindMultiCacheByWords(ctx context.Context, words []string) ([]*fkratos_common_model.SensitiveWord, error)
		// FindOneCacheByID 根据ID查询一条数据并设置缓存
		FindOneCacheByID(ctx context.Context, ID string) (*fkratos_common_model.SensitiveWord, error)
		// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
		FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*fkratos_common_model.SensitiveWord, error)
	}

	SensitiveWordRepo struct {
		db         *gorm.DB
		rockscache *rockscache.Client
	}
)

func NewSensitiveWordRepo(db *gorm.DB, rockscache *rockscache.Client) *SensitiveWordRepo {
	return &SensitiveWordRepo{
		db:         db,
		rockscache: rockscache,
	}
}

// CreateOne 创建一条数据
func (r *SensitiveWordRepo) CreateOne(ctx context.Context, data *fkratos_common_model.SensitiveWord) error {
	dao := fkratos_common_dao.Use(r.db).SensitiveWord
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne 更新一条数据
func (r *SensitiveWordRepo) UpdateOne(ctx context.Context, data *fkratos_common_model.SensitiveWord) error {
	dao := fkratos_common_dao.Use(r.db).SensitiveWord
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = r.DeleteUniqueIndexCache(ctx, []*fkratos_common_model.SensitiveWord{data})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByWord 根据word删除一条数据并清理缓存
func (r *SensitiveWordRepo) DeleteOneCacheByWord(ctx context.Context, word string) error {
	dao := fkratos_common_dao.Use(r.db).SensitiveWord
	first, err := dao.WithContext(ctx).Where(dao.Word.Eq(word)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if first == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Word.Eq(word)).Delete()
	if err != nil {
		return err
	}
	err = r.DeleteUniqueIndexCache(ctx, []*fkratos_common_model.SensitiveWord{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByWords 根据words删除多条数据并清理缓存
func (r *SensitiveWordRepo) DeleteMultiCacheByWords(ctx context.Context, words []string) error {
	dao := fkratos_common_dao.Use(r.db).SensitiveWord
	list, err := dao.WithContext(ctx).Where(dao.Word.In(words...)).Find()
	if err != nil {
		return err
	}
	if len(list) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Word.In(words...)).Delete()
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
func (r *SensitiveWordRepo) DeleteOneCacheByID(ctx context.Context, ID string) error {
	dao := fkratos_common_dao.Use(r.db).SensitiveWord
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
	err = r.DeleteUniqueIndexCache(ctx, []*fkratos_common_model.SensitiveWord{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (r *SensitiveWordRepo) DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error {
	dao := fkratos_common_dao.Use(r.db).SensitiveWord
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
func (r *SensitiveWordRepo) DeleteUniqueIndexCache(ctx context.Context, data []*fkratos_common_model.SensitiveWord) error {
	var err error
	cacheSensitiveWordByWord := CacheSensitiveWordByWord.NewSingleKey(r.rockscache)
	cacheSensitiveWordByID := CacheSensitiveWordByID.NewSingleKey(r.rockscache)

	for _, v := range data {
		err = cacheSensitiveWordByWord.SingleCacheDel(ctx, cacheSensitiveWordByWord.BuildKey(v.Word))
		if err != nil {
			return err
		}
		err = cacheSensitiveWordByID.SingleCacheDel(ctx, cacheSensitiveWordByID.BuildKey(v.ID))
		if err != nil {
			return err
		}

	}
	return nil
}

// FindOneCacheByWord 根据word查询一条数据并设置缓存
func (r *SensitiveWordRepo) FindOneCacheByWord(ctx context.Context, word string) (*fkratos_common_model.SensitiveWord, error) {
	resp := new(fkratos_common_model.SensitiveWord)
	cache := CacheSensitiveWordByWord.NewSingleKey(r.rockscache)
	cacheValue, err := cache.SingleCache(ctx, conv.String(word), func() (string, error) {
		dao := fkratos_common_dao.Use(r.db).SensitiveWord
		result, err := dao.WithContext(ctx).Where(dao.Word.Eq(word)).First()
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

// FindMultiCacheByWords 根据words查询多条数据并设置缓存
func (r *SensitiveWordRepo) FindMultiCacheByWords(ctx context.Context, words []string) ([]*fkratos_common_model.SensitiveWord, error) {
	resp := make([]*fkratos_common_model.SensitiveWord, 0)
	cacheKey := CacheSensitiveWordByWord.NewBatchKey(r.rockscache)
	batchKeys := make([]string, 0)
	for _, v := range words {
		batchKeys = append(batchKeys, conv.String(v))
	}
	cacheValue, err := cacheKey.BatchKeyCache(ctx, batchKeys, func() (map[string]string, error) {
		dao := fkratos_common_dao.Use(r.db).SensitiveWord
		result, err := dao.WithContext(ctx).Where(dao.Word.In(words...)).Find()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		value := make(map[string]string)
		for _, v := range result {
			marshal, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			value[conv.String(v.Word)] = string(marshal)
		}
		return value, nil
	})
	if err != nil {
		return nil, err
	}
	for _, v := range cacheValue {
		tmp := new(fkratos_common_model.SensitiveWord)
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}

// FindOneCacheByID 根据ID查询一条数据并设置缓存
func (r *SensitiveWordRepo) FindOneCacheByID(ctx context.Context, ID string) (*fkratos_common_model.SensitiveWord, error) {
	resp := new(fkratos_common_model.SensitiveWord)
	cache := CacheSensitiveWordByID.NewSingleKey(r.rockscache)
	cacheValue, err := cache.SingleCache(ctx, conv.String(ID), func() (string, error) {
		dao := fkratos_common_dao.Use(r.db).SensitiveWord
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

// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
func (r *SensitiveWordRepo) FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*fkratos_common_model.SensitiveWord, error) {
	resp := make([]*fkratos_common_model.SensitiveWord, 0)
	cacheKey := CacheSensitiveWordByID.NewBatchKey(r.rockscache)
	batchKeys := make([]string, 0)
	for _, v := range IDS {
		batchKeys = append(batchKeys, conv.String(v))
	}
	cacheValue, err := cacheKey.BatchKeyCache(ctx, batchKeys, func() (map[string]string, error) {
		dao := fkratos_common_dao.Use(r.db).SensitiveWord
		result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		value := make(map[string]string)
		for _, v := range result {
			marshal, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			value[conv.String(v.ID)] = string(marshal)
		}
		return value, nil
	})
	if err != nil {
		return nil, err
	}
	for _, v := range cacheValue {
		tmp := new(fkratos_common_model.SensitiveWord)
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}