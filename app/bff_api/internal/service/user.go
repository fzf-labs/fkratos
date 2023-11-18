package service

import (
	pb "fkratos/api/bff_api/v1"
	"fkratos/app/bff_api/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewUserService(
	logger log.Logger,
	userUseCase *biz.UserUseCase,
) *UserService {
	l := log.NewHelper(log.With(logger, "module", "service/user"))
	return &UserService{
		log:         l,
		userUseCase: userUseCase,
	}
}

type UserService struct {
	pb.UnimplementedUserServer
	log         *log.Helper
	userUseCase *biz.UserUseCase
}
