package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

// UserList 用户-列表
func (u *UserService) UserList(ctx context.Context, req *pb.UserListReq) (*pb.UserListReply, error) {
	return u.userUseCase.UserList(ctx, req)
}
