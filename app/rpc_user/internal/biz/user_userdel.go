package biz

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

// UserDel 用户-删除
func (u *UserUseCase) UserDel(ctx context.Context, req *pb.UserDelReq) (*pb.UserDelReply, error) {
	resp := new(pb.UserDelReply)
	err := u.userRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
