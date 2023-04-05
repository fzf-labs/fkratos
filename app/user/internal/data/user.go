package data

import (
	"context"
	v1 "fkratos/api/user/v1"
	"fkratos/app/user/internal/biz"
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

func (u *UserRepo) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	//TODO implement me
	panic("implement me")
}
