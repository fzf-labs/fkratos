package data

import (
	"context"
	"fkratos/app/rpc_common/internal/biz"
	"fkratos/app/rpc_common/internal/data/cache"
	"fkratos/app/rpc_common/internal/data/gorm/fkratos_common_dao"
	"fkratos/app/rpc_common/internal/data/gorm/fkratos_common_model"
	"fkratos/app/rpc_common/internal/data/gorm/fkratos_common_repo"

	"github.com/fzf-labs/fpkg/util/jsonutil"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.SensitiveWordRepo = (*SensitiveWordRepo)(nil)

func NewSensitiveWordRepo(data *Data, logger log.Logger) biz.SensitiveWordRepo {
	l := log.NewHelper(log.With(logger, "module", "rpc_common/data/sensitive_word"))
	return &SensitiveWordRepo{
		data:              data,
		log:               l,
		SensitiveWordRepo: fkratos_common_repo.NewSensitiveWordRepo(data.db, data.rueidisdbcache),
	}
}

type SensitiveWordRepo struct {
	data *Data
	log  *log.Helper
	*fkratos_common_repo.SensitiveWordRepo
}

func (s *SensitiveWordRepo) SensitiveWordStore(ctx context.Context, data *fkratos_common_model.SensitiveWord) (*fkratos_common_model.SensitiveWord, error) {
	sensitiveWordDao := fkratos_common_dao.Use(s.data.db).SensitiveWord
	if data.ID == "" {
		err := sensitiveWordDao.WithContext(ctx).Create(data)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := sensitiveWordDao.WithContext(ctx).Omit(sensitiveWordDao.CreatedAt).Where(sensitiveWordDao.ID.Eq(data.ID)).Updates(data)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

func (s *SensitiveWordRepo) SensitiveWordsQuery(ctx context.Context) ([]string, error) {
	sensitiveWordDao := fkratos_common_dao.Use(s.data.db).SensitiveWord
	result := make([]string, 0)
	err := sensitiveWordDao.WithContext(ctx).Pluck(sensitiveWordDao.Word, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *SensitiveWordRepo) SensitiveWordsCache(ctx context.Context) ([]string, error) {
	key := cache.SensitiveWordsCache.Key()
	result := make([]string, 0)
	resp, err := s.data.rueidisdbcache.Fetch(ctx, key, func() (string, error) {
		sensitiveWordsQuery, err := s.SensitiveWordsQuery(ctx)
		if err != nil {
			return "", err
		}
		encode, err := jsonutil.Encode(sensitiveWordsQuery)
		if err != nil {
			return "", err
		}
		return string(encode), nil
	})
	if err != nil {
		return nil, err
	}
	err = jsonutil.DecodeString(resp, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *SensitiveWordRepo) SensitiveWordsCacheDel(ctx context.Context) error {
	key := cache.SensitiveWordsCache.Key()
	err := s.data.rueidisdbcache.Del(ctx, key)
	if err != nil {
		return err
	}
	return nil
}
