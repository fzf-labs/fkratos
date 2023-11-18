package service

import (
	"context"

	pb "fkratos/api/bff_api/v1"
)

// GetUserInfo 获取单个用户
func (u *UserService) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq) (*pb.GetUserInfoReply, error) {
	return u.userUseCase.GetUserInfo(ctx, req)
}
