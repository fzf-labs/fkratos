package data

import (
	"context"
	"fkratos/api/paginator"
	"fkratos/app/rpc_common/internal/biz"
	"fkratos/app/rpc_common/internal/data/cache"
	"fkratos/app/rpc_common/internal/data/gorm/fkratos_common_dao"
	"fkratos/app/rpc_common/internal/data/gorm/fkratos_common_model"
	"fkratos/app/rpc_common/internal/data/gorm/fkratos_common_repo"
	"fkratos/internal/errorx"
	"strings"

	"github.com/fzf-labs/fpkg/page"
	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.SensitiveWordRepo = (*SensitiveWordRepo)(nil)

func NewSensitiveWordRepo(data *Data, logger log.Logger) biz.SensitiveWordRepo {
	l := log.NewHelper(log.With(logger, "module", "rpc_common/data/sensitive_word"))
	return &SensitiveWordRepo{
		data:              data,
		log:               l,
		SensitiveWordRepo: fkratos_common_repo.NewSensitiveWordRepo(data.db, data.rockscache),
	}
}

type SensitiveWordRepo struct {
	data *Data
	log  *log.Helper
	*fkratos_common_repo.SensitiveWordRepo
}

func (s *SensitiveWordRepo) SensitiveWordListBySearch(ctx context.Context, req *paginator.PaginatorReq) ([]*fkratos_common_model.SensitiveWord, *page.Page, error) {
	sensitiveWordDao := fkratos_common_dao.Use(s.data.db).SensitiveWord
	query := sensitiveWordDao.WithContext(ctx)
	if req.QuickSearch != "" {
		query = query.Where(sensitiveWordDao.Word.Like(req.QuickSearch))
	} else {
		for _, search := range req.Search {
			if search.Field == "id" {
				query = query.Where(sensitiveWordDao.ID.Eq(search.Val))
			}
			if search.Field == "Word" {
				query = query.Where(sensitiveWordDao.Word.Eq(search.Val))
			}
			if search.Field == "createdAt" {
				ss := strings.Split(search.Val, ",")
				query = query.Where(sensitiveWordDao.CreatedAt.Gte(timeutil.Carbon().Parse(ss[0]).Carbon2Time()), sensitiveWordDao.CreatedAt.Lte(timeutil.Carbon().Parse(ss[1]).Carbon2Time()))
			}
		}
	}

	queryCount := query
	total, err := queryCount.Count()
	if err != nil {
		return nil, nil, errorx.DataSqlErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	p := page.Paginator(int(req.Page), int(req.PageSize), int(total))
	result, err := query.Offset(p.Offset).Limit(p.Limit).Find()
	if err != nil {
		return nil, nil, errorx.DataSqlErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return result, p, nil
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
	exists, err := s.data.redis.Exists(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	//存在
	if exists == 1 {
		result, err := s.data.redis.HKeys(ctx, key).Result()
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	result, err := s.SensitiveWordsQuery(ctx)
	if err != nil {
		return nil, err
	}
	values := make([]interface{}, 0)
	for _, v := range result {
		values = append(values, v)
		values = append(values, "1")
	}
	err = s.data.redis.HSet(ctx, key, values...).Err()
	if err != nil {
		return nil, err
	}
	return result, nil
}
