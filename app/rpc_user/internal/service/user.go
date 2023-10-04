package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
	"fkratos/app/rpc_user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	pb.UnimplementedUserServer
	log         *log.Helper
	userUseCase *biz.UserUseCase
}

func NewUserService(logger log.Logger, userUseCase *biz.UserUseCase) *UserService {
	l := log.NewHelper(log.With(logger, "module", "service/user"))
	return &UserService{
		log:         l,
		userUseCase: userUseCase,
	}
}

func (s *UserService) UserList(ctx context.Context, req *pb.UserListReq) (*pb.UserListReply, error) {
	return s.userUseCase.UserList(ctx, req)
}
func (s *UserService) UserInfo(ctx context.Context, req *pb.UserInfoReq) (*pb.UserInfoReply, error) {
	return s.userUseCase.UserInfo(ctx, req)
}
func (s *UserService) UserStore(ctx context.Context, req *pb.UserStoreReq) (*pb.UserStoreReply, error) {
	return s.userUseCase.UserStore(ctx, req)
}
func (s *UserService) UserDel(ctx context.Context, req *pb.UserDelReq) (*pb.UserDelReply, error) {
	return s.userUseCase.UserDel(ctx, req)
}
