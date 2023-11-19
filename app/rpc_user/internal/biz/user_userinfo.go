package biz

import (
	"context"
	"fkratos/internal/dto"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_user/v1"
)

// UserInfo 用户-信息
func (u *UserUseCase) UserInfo(ctx context.Context, req *pb.UserInfoReq) (*pb.UserInfoReply, error) {
	resp := &pb.UserInfoReply{
		Info: &pb.UserInfo{},
	}
	result, err := u.userRepo.FindOneCacheByUID(ctx, req.GetUid())
	if err != nil {
		return nil, err
	}
	err = dto.Copy(&resp.Info, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithError(err).Err()
	}
	return resp, nil
}
