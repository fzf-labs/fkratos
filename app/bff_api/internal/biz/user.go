package biz

import "github.com/go-kratos/kratos/v2/log"

type UserRepo interface {
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
