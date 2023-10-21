package service

import (
	pb "fkratos/api/rpc_user/v1"
	"fkratos/app/rpc_user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewUserRuleService(
	logger log.Logger,
	userRuleUseCase *biz.UserRuleUseCase,
) *UserRuleService {
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
