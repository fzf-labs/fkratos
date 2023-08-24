package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	CreateUSer()
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	l := log.NewHelper(log.With(logger, "module", "rpc_user/biz"))
	return &UserUseCase{
		repo: repo,
		log:  l,
	}
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}
