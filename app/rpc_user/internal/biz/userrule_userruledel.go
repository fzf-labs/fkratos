package biz

import (
	"context"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_user/v1"
)

// UserRuleDel 用户规则-删除
func (s *UserRuleUseCase) UserRuleDel(ctx context.Context, req *pb.UserRuleDelReq) (*pb.UserRuleDelReply, error) {
	var resp = &pb.UserRuleDelReply{}
	err := s.userRuleRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, errorx.DataSQLErr.WithError(err).Err()
	}
	return resp, nil
}
