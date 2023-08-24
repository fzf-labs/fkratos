package service

import (
	rpcUserV1 "fkratos/api/rpc_user/v1"
	"fkratos/app/rpc_user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	rpcUserV1.UnimplementedUserServer
	log *log.Helper

	userUseCase *biz.UserUseCase
}

func NewUserService(logger log.Logger, userUseCase *biz.UserUseCase) *UserService {
	l := log.NewHelper(log.With(logger, "module", "user/service/user-service"))
	return &UserService{
		log:         l,
		userUseCase: userUseCase,
	}
}
