package biz

import (
	"context"
	pb "fkratos/api/bff_api/v1"
)

// GetUserInfo 获取单个用户
func (u *UserUseCase) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq) (*pb.GetUserInfoReply, error) {
	resp := &pb.GetUserInfoReply{}
	return resp, nil
}
