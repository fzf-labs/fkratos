package biz

import (
	"context"
	"encoding/json"
	"fkratos/api/paginator"
	v1 "fkratos/api/rpc_common/v1"
	"fkratos/app/rpc_common/internal/data/gorm/fkratos_common_model"
	"fkratos/app/rpc_common/internal/data/gorm/fkratos_common_repo"
	"fkratos/internal/errorx"
	"sync"

	"github.com/fzf-labs/fpkg/page"
	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/importcjj/sensitive"
	"github.com/jinzhu/copier"

	"github.com/go-kratos/kratos/v2/log"
)

type SensitiveWordRepo interface {
	fkratos_common_repo.ISensitiveWordRepo
	SensitiveWordListBySearch(ctx context.Context, req *paginator.PaginatorReq) ([]*fkratos_common_model.SensitiveWord, *page.Page, error)
	SensitiveWordStore(ctx context.Context, data *fkratos_common_model.SensitiveWord) (*fkratos_common_model.SensitiveWord, error)
	SensitiveWordsCache(ctx context.Context) ([]string, error)
	SensitiveWordsCacheDel(ctx context.Context) error
}

func NewSensitiveWordUseCase(logger log.Logger, sensitiveWordRepo SensitiveWordRepo) *SensitiveWordUseCase {
	l := log.NewHelper(log.With(logger, "module", "rpc_common/biz/sensitive_word"))
	return &SensitiveWordUseCase{
		log:               l,
		sensitiveWordRepo: sensitiveWordRepo,
	}
}

type SensitiveWordUseCase struct {
	log               *log.Helper
	sensitiveWordRepo SensitiveWordRepo
}

func (s *SensitiveWordUseCase) SensitiveWordList(ctx context.Context, req *paginator.PaginatorReq) (*v1.SensitiveWordListReply, error) {
	resp := new(v1.SensitiveWordListReply)
	result, p, err := s.sensitiveWordRepo.SensitiveWordListBySearch(ctx, req)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		labs := make([]string, 0)
		err := json.Unmarshal(v.Labs, &labs)
		if err != nil {
			return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
		resp.List = append(resp.List, &v1.SensitiveWordInfo{
			Id:        v.ID,
			Word:      v.Word,
			Labs:      labs,
			Desc:      v.Desc,
			CreatedAt: timeutil.ToDateTimeStringByTime(v.CreatedAt),
			UpdatedAt: timeutil.ToDateTimeStringByTime(v.UpdatedAt),
		})
	}
	err = copier.Copy(&resp.Paginator, p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SensitiveWordUseCase) SensitiveWordStore(ctx context.Context, req *v1.SensitiveWordStoreReq) (*v1.SensitiveWordStoreReply, error) {
	resp := new(v1.SensitiveWordStoreReply)
	labs, err := json.Marshal(req.GetLabs())
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	_, err = s.sensitiveWordRepo.SensitiveWordStore(ctx, &fkratos_common_model.SensitiveWord{
		ID:   req.GetId(),
		Word: req.GetWord(),
		Labs: labs,
		Desc: req.GetDesc(),
	})
	if err != nil {
		return nil, err
	}
	err = s.sensitiveWordRepo.SensitiveWordsCacheDel(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SensitiveWordUseCase) SensitiveWordDel(ctx context.Context, req *v1.SensitiveWordDelReq) (*v1.SensitiveWordDelReply, error) {
	resp := new(v1.SensitiveWordDelReply)
	err := s.sensitiveWordRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type SensitiveCheck struct {
	len    int
	filter *sensitive.Filter
}

var lock sync.Mutex
var sc = new(SensitiveCheck)

func (s *SensitiveWordUseCase) SensitiveWordCheck(ctx context.Context, req *v1.SensitiveWordCheckReq) (*v1.SensitiveWordCheckResp, error) {
	resp := new(v1.SensitiveWordCheckResp)
	sensitiveWords, err := s.sensitiveWordRepo.SensitiveWordsCache(ctx)
	if err != nil {
		return nil, err
	}
	if sc.len != len(sensitiveWords) {
		lock.Lock()
		defer lock.Unlock()
		sc = &SensitiveCheck{
			len:    len(sensitiveWords),
			filter: sensitive.New(),
		}
		sc.filter.AddWord(sensitiveWords...)
		sc.filter.UpdateNoisePattern(`x`)
	}
	validate, _ := sc.filter.Validate(req.Word)
	if validate {
		resp.Result = false
		return resp, nil
	}
	replace := sc.filter.Replace(req.Word, '*')
	filterStr := sc.filter.Filter(req.Word)

	resp.Result = false
	resp.Replace = replace
	resp.Filter = filterStr
	return resp, nil
}
