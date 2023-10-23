package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

// UserDel 用户-删除
func (u *UserService) UserDel(ctx context.Context, req *pb.UserDelReq) (*pb.UserDelReply, error) {
	return u.userUseCase.UserDel(ctx, req)
}
