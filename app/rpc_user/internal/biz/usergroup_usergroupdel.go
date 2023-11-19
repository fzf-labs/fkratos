package biz

import (
	"context"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_user/v1"
)

// UserGroupDel 用户分组-删除
func (u *UserGroupUseCase) UserGroupDel(ctx context.Context, req *pb.UserGroupDelReq) (*pb.UserGroupDelReply, error) {
	var resp = &pb.UserGroupDelReply{}
	err := u.userGroupRepo.DeleteMultiByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, errorx.DataFormattingError.WithError(err).Err()
	}
	return resp, nil
}
