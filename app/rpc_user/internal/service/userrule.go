package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
	"fkratos/app/rpc_user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewUserRuleService(logger log.Logger, userRuleUseCase *biz.UserRuleUseCase) *UserRuleService {
	l := log.NewHelper(log.With(logger, "module", "service/userRule"))
	return &UserRuleService{
		log:             l,
		userRuleUseCase: userRuleUseCase,
	}
}

type UserRuleService struct {
	pb.UnimplementedUserRuleServer
	log             *log.Helper
	userRuleUseCase *biz.UserRuleUseCase
}

func (s *UserRuleService) UserRuleList(ctx context.Context, req *pb.UserRuleListReq) (*pb.UserRuleListReply, error) {
	return s.userRuleUseCase.UserRuleList(ctx, req)
}
func (s *UserRuleService) UserRuleInfo(ctx context.Context, req *pb.UserRuleInfoReq) (*pb.UserRuleInfoReply, error) {
	return s.userRuleUseCase.UserRuleInfo(ctx, req)
}
func (s *UserRuleService) UserRuleStore(ctx context.Context, req *pb.UserRuleStoreReq) (*pb.UserRuleStoreReply, error) {
	return s.userRuleUseCase.UserRuleStore(ctx, req)
}
func (s *UserRuleService) UserRuleDel(ctx context.Context, req *pb.UserRuleDelReq) (*pb.UserRuleDelReply, error) {
	return s.userRuleUseCase.UserRuleDel(ctx, req)
}

func (s *UserRuleService) UserRules(ctx context.Context, req *pb.UserRulesReq) (*pb.UserRulesReply, error) {
	return s.userRuleUseCase.UserRules(ctx, req)
}
