package data

import (
	"fkratos/app/rpc_user/internal/biz"
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_repo"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*UserRepo)(nil)

func NewUserRepo(logger log.Logger, data *Data, userRepo *fkratos_user_repo.UserRepo) biz.UserRepo {
	l := log.NewHelper(log.With(logger, "module", "data/user"))
	return &UserRepo{
		log:      l,
		data:     data,
		UserRepo: userRepo,
	}
}

type UserRepo struct {
	log  *log.Helper
	data *Data
	*fkratos_user_repo.UserRepo
}
