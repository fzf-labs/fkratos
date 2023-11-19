package biz

import (
	userV1 "fkratos/api/rpc_user/v1"

	"github.com/go-kratos/kratos/v2/log"
)

func NewUserUseCase(
	logger log.Logger,
	userClient userV1.UserClient,
) *UserUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/user"))
	return &UserUseCase{
		log:        l,
		userClient: userClient,
	}
}

type UserUseCase struct {
	log        *log.Helper
	userClient userV1.UserClient
}
