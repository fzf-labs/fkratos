package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
	"fkratos/app/rpc_user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserGroupService struct {
	pb.UnimplementedUserGroupServer
	log              *log.Helper
	userGroupUseCase *biz.UserGroupUseCase
}

func NewUserGroupService(logger log.Logger, userGroupUseCase *biz.UserGroupUseCase) *UserGroupService {
	l := log.NewHelper(log.With(logger, "module", "service/userGroup"))
	return &UserGroupService{
		log:              l,
		userGroupUseCase: userGroupUseCase,
	}
}

func (s *UserGroupService) UserGroupList(ctx context.Context, req *pb.UserGroupListReq) (*pb.UserGroupListReply, error) {
	return s.userGroupUseCase.UserGroupList(ctx, req)
}
func (s *UserGroupService) UserGroupStore(ctx context.Context, req *pb.UserGroupStoreReq) (*pb.UserGroupStoreReply, error) {
	return s.userGroupUseCase.UserGroupStore(ctx, req)
}
func (s *UserGroupService) UserGroupDel(ctx context.Context, req *pb.UserGroupDelReq) (*pb.UserGroupDelReply, error) {
	return s.userGroupUseCase.UserGroupDel(ctx, req)
}
