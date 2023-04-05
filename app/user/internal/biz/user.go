package biz

import (
	"context"
	v1 "fkratos/api/user/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserReply, error)
	UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserReply, error)
	DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserReply, error)
	GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error)
	ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error)
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	l := log.NewHelper(log.With(logger, "module", "User/usecase/User-service"))
	return &UserUseCase{
		repo: repo,
		log:  l,
	}
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func (u *UserUseCase) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	return u.repo.CreateUser(ctx, req)
}

func (u *UserUseCase) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserUseCase) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserUseCase) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserUseCase) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	//TODO implement me
	panic("implement me")
}
