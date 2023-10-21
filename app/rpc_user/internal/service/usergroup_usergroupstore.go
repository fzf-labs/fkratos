package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

func (u *UserGroupService) UserGroupStore(ctx context.Context, req *pb.UserGroupStoreReq) (*pb.UserGroupStoreReply, error) {
	return u.userGroupUseCase.UserGroupStore(ctx, req)
}
