package data

import (
	"fkratos/app/rpc_user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserGroupRepo = (*UserGroupRepo)(nil)

func NewUserGroupRepo(logger log.Logger, data *Data) *UserGroupRepo {
	l := log.NewHelper(log.With(logger, "module", "data/userGroup"))
	return &UserGroupRepo{
		log:  l,
		data: data,
	}
}

type UserGroupRepo struct {
	log  *log.Helper
	data *Data
}
