package service

import (
	"context"
	"fkratos/app/rpc_common/internal/biz"

	pb "fkratos/api/rpc_common/v1"

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

func (s *SensitiveWordService) SensitiveCategoryList(ctx context.Context, req *pb.SensitiveCategoryListReq) (*pb.SensitiveCategoryListReply, error) {
	return s.sensitiveWordUseCase.SensitiveCategoryList(ctx, req)
}
func (s *SensitiveWordService) SensitiveCategoryInfo(ctx context.Context, req *pb.SensitiveCategoryInfoReq) (*pb.SensitiveCategoryInfoReply, error) {
	return s.sensitiveWordUseCase.SensitiveCategoryInfo(ctx, req)
}
func (s *SensitiveWordService) SensitiveCategoryStore(ctx context.Context, req *pb.SensitiveCategoryStoreReq) (*pb.SensitiveCategoryStoreReply, error) {
	return s.sensitiveWordUseCase.SensitiveCategoryStore(ctx, req)
}
func (s *SensitiveWordService) SensitiveCategoryDel(ctx context.Context, req *pb.SensitiveCategoryDelReq) (*pb.SensitiveCategoryDelReply, error) {
	return s.sensitiveWordUseCase.SensitiveCategoryDel(ctx, req)
}
func (s *SensitiveWordService) SensitiveWordList(ctx context.Context, req *pb.SensitiveWordListReq) (*pb.SensitiveWordListReply, error) {
	return s.sensitiveWordUseCase.SensitiveWordList(ctx, req)
}
func (s *SensitiveWordService) SensitiveWordInfo(ctx context.Context, req *pb.SensitiveWordInfoReq) (*pb.SensitiveWordInfoReply, error) {
	return s.sensitiveWordUseCase.SensitiveWordInfo(ctx, req)
}
func (s *SensitiveWordService) SensitiveWordStore(ctx context.Context, req *pb.SensitiveWordStoreReq) (*pb.SensitiveWordStoreReply, error) {
	return s.sensitiveWordUseCase.SensitiveWordStore(ctx, req)
}
func (s *SensitiveWordService) SensitiveWordDel(ctx context.Context, req *pb.SensitiveWordDelReq) (*pb.SensitiveWordDelReply, error) {
	return s.sensitiveWordUseCase.SensitiveWordDel(ctx, req)
}
