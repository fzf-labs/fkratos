package biz

import (
	"context"
	rpcUserV1 "fkratos/api/rpc_user/v1"

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

func (u *UserUseCase) CreateUser(ctx context.Context, req *rpcUserV1.CreateUserReq) (*rpcUserV1.CreateUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserUseCase) UpdateUser(ctx context.Context, req *rpcUserV1.UpdateUserReq) (*rpcUserV1.UpdateUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserUseCase) DeleteUser(ctx context.Context, req *rpcUserV1.DeleteUserReq) (*rpcUserV1.DeleteUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserUseCase) GetUser(ctx context.Context, req *rpcUserV1.GetUserReq) (*rpcUserV1.GetUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserUseCase) ListUser(ctx context.Context, req *rpcUserV1.ListUserReq) (*rpcUserV1.ListUserReply, error) {
	//TODO implement me
	panic("implement me")
}
