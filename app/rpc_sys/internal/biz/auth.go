package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type AuthRepo interface {
}

func NewAuthUseCase(repo AuthRepo, logger log.Logger) *AuthUseCase {
	l := log.NewHelper(log.With(logger, "module", "auth/usecase/auth-service"))
	return &AuthUseCase{
		repo: repo,
		log:  l,
	}
}

type AuthUseCase struct {
	repo AuthRepo
	log  *log.Helper
}
