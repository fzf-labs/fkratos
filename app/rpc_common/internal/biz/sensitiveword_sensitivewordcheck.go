package biz

import (
	"context"
	"sync"

	pb "fkratos/api/rpc_common/v1"

	"github.com/importcjj/sensitive"
)

type SensitiveCheck struct {
	len    int
	filter *sensitive.Filter
}

var lock sync.Mutex
var sc = new(SensitiveCheck)

// SensitiveWordCheck 敏感词-检测
func (s *SensitiveWordUseCase) SensitiveWordCheck(ctx context.Context, req *pb.SensitiveWordCheckReq) (*pb.SensitiveWordCheckResp, error) {
	resp := &pb.SensitiveWordCheckResp{}
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
