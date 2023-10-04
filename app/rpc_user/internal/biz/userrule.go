package biz

import (
	"context"

	pb "fkratos/api/rpc_user/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRuleRepo interface {
}

type UserRuleUseCase struct {
	log *log.Helper
}

func NewUserRuleUseCase(logger log.Logger) *UserRuleUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/userRule"))
	return &UserRuleUseCase{
		log: l,
	}
}

func (s *UserRuleUseCase) UserRuleList(ctx context.Context, req *pb.UserRuleListReq) (*pb.UserRuleListReply, error) {
	return &pb.UserRuleListReply{}, nil
}
func (s *UserRuleUseCase) UserRuleInfo(ctx context.Context, req *pb.UserRuleInfoReq) (*pb.UserRuleInfoReply, error) {
	return &pb.UserRuleInfoReply{}, nil
}
func (s *UserRuleUseCase) UserRuleStore(ctx context.Context, req *pb.UserRuleStoreReq) (*pb.UserRuleStoreReply, error) {
	return &pb.UserRuleStoreReply{}, nil
}
func (s *UserRuleUseCase) UserRuleDel(ctx context.Context, req *pb.UserRuleDelReq) (*pb.UserRuleDelReply, error) {
	return &pb.UserRuleDelReply{}, nil
}
