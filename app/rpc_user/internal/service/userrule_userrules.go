package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

func (u *UserRuleService) UserRules(ctx context.Context, req *pb.UserRulesReq) (*pb.UserRulesReply, error) {
	return u.userRuleUseCase.UserRules(ctx, req)
}
