package data

import (
	"fkratos/app/rpc_user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*UserRepo)(nil)

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	l := log.NewHelper(log.With(logger, "module", "user/data/user"))
	return &UserRepo{
		data: data,
		log:  l,
	}
}

type UserRepo struct {
	data *Data
	log  *log.Helper
}
