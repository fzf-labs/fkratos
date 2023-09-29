package data

import (
	"fkratos/app/rpc_user/internal/biz"
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_repo"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*UserRepo)(nil)

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	l := log.NewHelper(log.With(logger, "module", "user/data/user"))
	return &UserRepo{
		data:     data,
		log:      l,
		UserRepo: fkratos_user_repo.NewUserRepo(data.gorm, data.rueidisdbcache),
	}
}

type UserRepo struct {
	data *Data
	log  *log.Helper
	*fkratos_user_repo.UserRepo
}
