package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

func (u *UserService) UserList(ctx context.Context, req *pb.UserListReq) (*pb.UserListReply, error) {
	return u.userUseCase.UserList(ctx, req)
}
