package data

import (
	"fkratos/app/rpc_user/internal/biz"
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_repo"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRuleRepo = (*UserRuleRepo)(nil)

func NewUserRuleRepo(
	logger log.Logger,
	data *Data,
	userRuleRepo *fkratos_user_repo.UserRuleRepo,
) biz.UserRuleRepo {
	l := log.NewHelper(log.With(logger, "module", "data/userRule"))
	return &UserRuleRepo{
		log:          l,
		data:         data,
		UserRuleRepo: userRuleRepo,
	}
}

type UserRuleRepo struct {
	log  *log.Helper
	data *Data
	*fkratos_user_repo.UserRuleRepo
}
