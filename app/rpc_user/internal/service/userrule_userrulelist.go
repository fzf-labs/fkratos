package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

// UserRuleList 用户规则-列表
func (u *UserRuleService) UserRuleList(ctx context.Context, req *pb.UserRuleListReq) (*pb.UserRuleListReply, error) {
	return u.userRuleUseCase.UserRuleList(ctx, req)
}
