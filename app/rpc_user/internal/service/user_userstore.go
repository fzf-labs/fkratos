package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

func (u *UserService) UserStore(ctx context.Context, req *pb.UserStoreReq) (*pb.UserStoreReply, error) {
	return u.userUseCase.UserStore(ctx, req)
}
