package biz

import (
	"context"
	userV1 "fkratos/api/rpc_user/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
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

func (u *UserUseCase) UserList(ctx context.Context, req *userV1.UserListReq) (*userV1.UserListReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserUseCase) UserInfo(ctx context.Context, req *userV1.UserInfoReq) (*userV1.UserInfoReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserUseCase) UserStore(ctx context.Context, req *userV1.UserStoreReq) (*userV1.UserStoreReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserUseCase) UserDel(ctx context.Context, req *userV1.UserDelReq) (*userV1.UserDelReply, error) {
	//TODO implement me
	panic("implement me")
}
