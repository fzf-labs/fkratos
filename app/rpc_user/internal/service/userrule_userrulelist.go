package service

import (
	"context"

	pb "fkratos/api/rpc_user/v1"
)

func (u *UserRuleService) UserRuleList(ctx context.Context, req *pb.UserRuleListReq) (*pb.UserRuleListReply, error) {
	return u.userRuleUseCase.UserRuleList(ctx, req)
}
