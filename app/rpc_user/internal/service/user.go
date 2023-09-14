package service

import (
	"context"
	v1 "fkratos/api/bff_api/v1"
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

func (u *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (*v1.UpdateUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (*v1.DeleteUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) ListUser(ctx context.Context, req *v1.ListUserReq) (*v1.ListUserReply, error) {
	//TODO implement me
	panic("implement me")
}
