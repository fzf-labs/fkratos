package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

// UserGroupList 用户分组-列表
func (u *UserGroupService) UserGroupList(ctx context.Context, req *pb.UserGroupListReq) (*pb.UserGroupListReply, error) {
	return u.userGroupUseCase.UserGroupList(ctx, req)
}
