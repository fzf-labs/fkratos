package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

func (u *UserRuleService) UserRuleStore(ctx context.Context, req *pb.UserRuleStoreReq) (*pb.UserRuleStoreReply, error) {
	return u.userRuleUseCase.UserRuleStore(ctx, req)
}
