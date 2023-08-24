package service

import (
	userV1 "fkratos/api/rpc_user/v1"

	pb "fkratos/api/bff_api/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	pb.UnimplementedUserServer
	log        *log.Helper
	userClient userV1.UserClient
}

func NewUserService(logger log.Logger, userClient userV1.UserClient) *UserService {
	l := log.NewHelper(log.With(logger, "module", "bff_admin/service/sys"))
	return &UserService{
		log:        l,
		userClient: userClient,
	}
}
