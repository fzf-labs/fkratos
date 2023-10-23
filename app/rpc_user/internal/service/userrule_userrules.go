package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

// UserRules 用户规则-用户所有规则
func (u *UserRuleService) UserRules(ctx context.Context, req *pb.UserRulesReq) (*pb.UserRulesReply, error) {
	return u.userRuleUseCase.UserRules(ctx, req)
}
