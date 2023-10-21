package biz

import (
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_repo"

	"github.com/go-kratos/kratos/v2/log"
)

type UserGroupRepo interface {
	fkratos_user_repo.IUserGroupRepo
}

func NewUserGroupUseCase(
	logger log.Logger,
	userGroupRepo UserGroupRepo,
) *UserGroupUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/userGroup"))
	return &UserGroupUseCase{
		log:           l,
		userGroupRepo: userGroupRepo,
	}
}

type UserGroupUseCase struct {
	log           *log.Helper
	userGroupRepo UserGroupRepo
}
