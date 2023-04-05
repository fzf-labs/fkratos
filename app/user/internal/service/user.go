package service

import (
	"context"
	"fkratos/app/user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	pb "fkratos/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
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

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
