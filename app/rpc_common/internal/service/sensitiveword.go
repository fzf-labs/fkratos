package service

import (
	"context"
	"fkratos/api/paginator"
	pb "fkratos/api/rpc_common/v1"
	"fkratos/app/rpc_common/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type SensitiveWordService struct {
	pb.UnimplementedSensitiveWordServer
	log                  *log.Helper
	sensitiveWordUseCase *biz.SensitiveWordUseCase
}

func NewSensitiveWordService(logger log.Logger, sensitiveWordUseCase *biz.SensitiveWordUseCase) *SensitiveWordService {
	l := log.NewHelper(log.With(logger, "module", "rpc_common/service/sensitive_word"))
	return &SensitiveWordService{
		log:                  l,
		sensitiveWordUseCase: sensitiveWordUseCase,
	}
}
func (s *SensitiveWordService) SensitiveWordList(ctx context.Context, req *paginator.PaginatorReq) (*pb.SensitiveWordListReply, error) {
	return s.sensitiveWordUseCase.SensitiveWordList(ctx, req)
}
func (s *SensitiveWordService) SensitiveWordStore(ctx context.Context, req *pb.SensitiveWordStoreReq) (*pb.SensitiveWordStoreReply, error) {
	return s.sensitiveWordUseCase.SensitiveWordStore(ctx, req)
}
func (s *SensitiveWordService) SensitiveWordDel(ctx context.Context, req *pb.SensitiveWordDelReq) (*pb.SensitiveWordDelReply, error) {
	return s.sensitiveWordUseCase.SensitiveWordDel(ctx, req)
}

// SensitiveWordCheck 敏感词-检测
func (s *SensitiveWordService) SensitiveWordCheck(ctx context.Context, req *pb.SensitiveWordCheckReq) (*pb.SensitiveWordCheckResp, error) {
	return s.sensitiveWordUseCase.SensitiveWordCheck(ctx, req)
}
