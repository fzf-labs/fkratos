package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

func (u *UserRuleService) UserRuleInfo(ctx context.Context, req *pb.UserRuleInfoReq) (*pb.UserRuleInfoReply, error) {
	return u.userRuleUseCase.UserRuleInfo(ctx, req)
}
