package data

import (
	"fkratos/app/rpc_user/internal/biz"
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_repo"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserGroupRepo = (*UserGroupRepo)(nil)

func NewUserGroupRepo(
	logger log.Logger,
	data *Data,
	userGroupRepo *fkratos_user_repo.UserGroupRepo,
) biz.UserGroupRepo {
	l := log.NewHelper(log.With(logger, "module", "data/userGroup"))
	return &UserGroupRepo{
		log:           l,
		data:          data,
		UserGroupRepo: userGroupRepo,
	}
}

type UserGroupRepo struct {
	log  *log.Helper
	data *Data
	*fkratos_user_repo.UserGroupRepo
}
