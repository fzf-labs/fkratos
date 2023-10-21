package biz

import (
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_repo"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRuleRepo interface {
	fkratos_user_repo.IUserRuleRepo
}

func NewUserRuleUseCase(
	logger log.Logger,
	userRuleRepo UserRuleRepo,
	userRepo UserRepo,
	userGroupRepo UserGroupRepo,
) *UserRuleUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/userRule"))
	return &UserRuleUseCase{
		log:           l,
		userRuleRepo:  userRuleRepo,
		userRepo:      userRepo,
		userGroupRepo: userGroupRepo,
	}
}

type UserRuleUseCase struct {
	log           *log.Helper
	userRuleRepo  UserRuleRepo
	userRepo      UserRepo
	userGroupRepo UserGroupRepo
}
