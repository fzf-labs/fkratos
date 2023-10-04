package data

import (
	"fkratos/app/rpc_user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRuleRepo = (*UserRuleRepo)(nil)

func NewUserRuleRepo(logger log.Logger, data *Data) *UserRuleRepo {
	l := log.NewHelper(log.With(logger, "module", "data/userRule"))
	return &UserRuleRepo{
		log:  l,
		data: data,
	}
}

type UserRuleRepo struct {
	log  *log.Helper
	data *Data
}
