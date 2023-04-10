package data

import (
	"context"
	"fkratos/api/rpc/user/v1"
	"fkratos/app/rpc/user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*UserRepo)(nil)

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	l := log.NewHelper(log.With(logger, "module", "user/repo/user-service"))
	return &UserRepo{
		data: data,
		log:  l,
	}
}

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func (u *UserRepo) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (*v1.UpdateUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (*v1.DeleteUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) ListUser(ctx context.Context, req *v1.ListUserReq) (*v1.ListUserReply, error) {
	//TODO implement me
	panic("implement me")
}
