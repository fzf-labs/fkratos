package biz

import (
	"context"
	"fkratos/internal/dto"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_user/v1"

	"github.com/fzf-labs/fpkg/util/jsonutil"
)

// UserRules 用户规则-用户所有规则
func (s *UserRuleUseCase) UserRules(ctx context.Context, req *pb.UserRulesReq) (*pb.UserRulesReply, error) {
	resp := &pb.UserRulesReply{
		List: make([]*pb.UserRuleInfo, 0),
	}
	user, err := s.userRepo.FindOneCacheByUID(ctx, req.GetUid())
	if err != nil {
		return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	if user == nil {
		return nil, errorx.AccountError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	if user.UserGroupID == "" {
		return resp, nil
	}
	userGroup, err := s.userGroupRepo.FindOneCacheByID(ctx, user.UserGroupID)
	if err != nil {
		return nil, err
	}
	if userGroup == nil || userGroup.Roles == nil {
		return resp, nil
	}
	userRuleIds := make([]string, 0)
	err = jsonutil.Unmarshal(userGroup.Roles, &userRuleIds)
	if err != nil {
		return nil, err
	}
	userRules, err := s.userRuleRepo.FindMultiCacheByIDS(ctx, userRuleIds)
	if err != nil {
		return nil, err
	}
	err = dto.Copy(&resp.List, userRules)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
