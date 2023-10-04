package data

import (
	"fkratos/app/rpc_user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*UserRepo)(nil)

func NewUserRepo(logger log.Logger, data *Data) *UserRepo {
	l := log.NewHelper(log.With(logger, "module", "data/user"))
	return &UserRepo{
		log:  l,
		data: data,
	}
}

type UserRepo struct {
	log  *log.Helper
	data *Data
}
