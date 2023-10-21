package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

func (u *UserService) UserInfo(ctx context.Context, req *pb.UserInfoReq) (*pb.UserInfoReply, error) {
	return u.userUseCase.UserInfo(ctx, req)
}
