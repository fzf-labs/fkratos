package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

// UserRuleInfo 用户规则-信息
func (u *UserRuleService) UserRuleInfo(ctx context.Context, req *pb.UserRuleInfoReq) (*pb.UserRuleInfoReply, error) {
	return u.userRuleUseCase.UserRuleInfo(ctx, req)
}
