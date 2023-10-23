package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

// UserRuleDel 用户规则-删除
func (u *UserRuleService) UserRuleDel(ctx context.Context, req *pb.UserRuleDelReq) (*pb.UserRuleDelReply, error) {
	return u.userRuleUseCase.UserRuleDel(ctx, req)
}
