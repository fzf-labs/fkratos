package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

func (u *UserGroupService) UserGroupDel(ctx context.Context, req *pb.UserGroupDelReq) (*pb.UserGroupDelReply, error) {
	return u.userGroupUseCase.UserGroupDel(ctx, req)
}
