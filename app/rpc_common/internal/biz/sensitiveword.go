package biz

import (
	"context"
	v1 "fkratos/api/rpc_common/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type SensitiveWordRepo interface {
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

func (s *SensitiveWordUseCase) SensitiveCategoryList(ctx context.Context, req *v1.SensitiveCategoryListReq) (*v1.SensitiveCategoryListReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SensitiveWordUseCase) SensitiveCategoryInfo(ctx context.Context, req *v1.SensitiveCategoryInfoReq) (*v1.SensitiveCategoryInfoReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SensitiveWordUseCase) SensitiveCategoryStore(ctx context.Context, req *v1.SensitiveCategoryStoreReq) (*v1.SensitiveCategoryStoreReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SensitiveWordUseCase) SensitiveCategoryDel(ctx context.Context, req *v1.SensitiveCategoryDelReq) (*v1.SensitiveCategoryDelReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SensitiveWordUseCase) SensitiveWordList(ctx context.Context, req *v1.SensitiveWordListReq) (*v1.SensitiveWordListReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SensitiveWordUseCase) SensitiveWordInfo(ctx context.Context, req *v1.SensitiveWordInfoReq) (*v1.SensitiveWordInfoReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SensitiveWordUseCase) SensitiveWordStore(ctx context.Context, req *v1.SensitiveWordStoreReq) (*v1.SensitiveWordStoreReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SensitiveWordUseCase) SensitiveWordDel(ctx context.Context, req *v1.SensitiveWordDelReq) (*v1.SensitiveWordDelReply, error) {
	//TODO implement me
	panic("implement me")
}
