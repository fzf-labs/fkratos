package service

import (
	"context"
	"fkratos/api/rpc/user/v1"
	"fkratos/app/rpc_user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	v1.UnimplementedUserServer
	log *log.Helper

	userUseCase *biz.UserUseCase
}

func NewUserService(logger log.Logger, userUseCase *biz.UserUseCase) *UserService {
	l := log.NewHelper(log.With(logger, "module", "user/service/user-service"))
	return &UserService{
		log:         l,
		userUseCase: userUseCase,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	return &v1.CreateUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (*v1.UpdateUserReply, error) {
	return &v1.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (*v1.DeleteUserReply, error) {
	return &v1.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	return &v1.GetUserReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *v1.ListUserReq) (*v1.ListUserReply, error) {
	return &v1.ListUserReply{}, nil
}
