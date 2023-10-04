package biz

import (
	"context"

	pb "fkratos/api/rpc_user/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
}

type UserUseCase struct {
	log *log.Helper
}

func NewUserUseCase(logger log.Logger) *UserUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/user"))
	return &UserUseCase{
		log: l,
	}
}

func (s *UserUseCase) UserList(ctx context.Context, req *pb.UserListReq) (*pb.UserListReply, error) {
	return &pb.UserListReply{}, nil
}
func (s *UserUseCase) UserInfo(ctx context.Context, req *pb.UserInfoReq) (*pb.UserInfoReply, error) {
	return &pb.UserInfoReply{}, nil
}
func (s *UserUseCase) UserStore(ctx context.Context, req *pb.UserStoreReq) (*pb.UserStoreReply, error) {
	return &pb.UserStoreReply{}, nil
}
func (s *UserUseCase) UserDel(ctx context.Context, req *pb.UserDelReq) (*pb.UserDelReply, error) {
	return &pb.UserDelReply{}, nil
}
