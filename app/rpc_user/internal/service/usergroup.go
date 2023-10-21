package service

import (
	pb "fkratos/api/rpc_user/v1"
	"fkratos/app/rpc_user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewUserGroupService(
	logger log.Logger,
	userGroupUseCase *biz.UserGroupUseCase,
) *UserGroupService {
	l := log.NewHelper(log.With(logger, "module", "service/userGroup"))
	return &UserGroupService{
		log:              l,
		userGroupUseCase: userGroupUseCase,
	}
}

type UserGroupService struct {
	pb.UnimplementedUserGroupServer
	log              *log.Helper
	userGroupUseCase *biz.UserGroupUseCase
}
