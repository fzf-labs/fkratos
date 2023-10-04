package biz

import (
	"context"

	pb "fkratos/api/rpc_user/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type UserGroupRepo interface {
}

type UserGroupUseCase struct {
	log *log.Helper
}

func NewUserGroupUseCase(logger log.Logger) *UserGroupUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/userGroup"))
	return &UserGroupUseCase{
		log: l,
	}
}

func (s *UserGroupUseCase) UserGroupList(ctx context.Context, req *pb.UserGroupListReq) (*pb.UserGroupListReply, error) {
	return &pb.UserGroupListReply{}, nil
}
func (s *UserGroupUseCase) UserGroupStore(ctx context.Context, req *pb.UserGroupStoreReq) (*pb.UserGroupStoreReply, error) {
	return &pb.UserGroupStoreReply{}, nil
}
func (s *UserGroupUseCase) UserGroupDel(ctx context.Context, req *pb.UserGroupDelReq) (*pb.UserGroupDelReply, error) {
	return &pb.UserGroupDelReply{}, nil
}
