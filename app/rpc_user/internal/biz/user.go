package biz

import (
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_repo"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	fkratos_user_repo.IUserRepo
}

func NewUserUseCase(
	logger log.Logger,
	userRepo UserRepo,
) *UserUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/user"))
	return &UserUseCase{
		log:      l,
		userRepo: userRepo,
	}
}

type UserUseCase struct {
	log      *log.Helper
	userRepo UserRepo
}
