package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

// UserGroupStore 用户分组-保存
func (u *UserGroupService) UserGroupStore(ctx context.Context, req *pb.UserGroupStoreReq) (*pb.UserGroupStoreReply, error) {
	return u.userGroupUseCase.UserGroupStore(ctx, req)
}
