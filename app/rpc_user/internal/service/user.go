package service

import (
	"context"
	rpcUserV1 "fkratos/api/rpc_user/v1"
	"fkratos/app/rpc_user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewUserService(logger log.Logger, userUseCase *biz.UserUseCase) *UserService {
	l := log.NewHelper(log.With(logger, "module", "user/service/user-service"))
	return &UserService{
		log:         l,
		userUseCase: userUseCase,
	}
}

type UserService struct {
	rpcUserV1.UnimplementedUserServer
	log *log.Helper

	userUseCase *biz.UserUseCase
}

func (u *UserService) UserList(ctx context.Context, req *rpcUserV1.UserListReq) (*rpcUserV1.UserListReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) UserInfo(ctx context.Context, req *rpcUserV1.UserInfoReq) (*rpcUserV1.UserInfoReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) UserStore(ctx context.Context, req *rpcUserV1.UserStoreReq) (*rpcUserV1.UserStoreReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) UserDel(ctx context.Context, req *rpcUserV1.UserDelReq) (*rpcUserV1.UserDelReply, error) {
	//TODO implement me
	panic("implement me")
}
