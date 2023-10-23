package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

// UserInfo 用户-信息
func (u *UserService) UserInfo(ctx context.Context, req *pb.UserInfoReq) (*pb.UserInfoReply, error) {
	return u.userUseCase.UserInfo(ctx, req)
}
