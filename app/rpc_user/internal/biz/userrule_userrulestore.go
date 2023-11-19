package biz

import (
	"context"
	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_model"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_user/v1"
)

// UserRuleStore 用户规则-保存
func (s *UserRuleUseCase) UserRuleStore(ctx context.Context, req *pb.UserRuleStoreReq) (*pb.UserRuleStoreReply, error) {
	resp := &pb.UserRuleStoreReply{
		Id: "",
	}
	userRule := &fkratos_user_model.UserRule{
		ID:        req.GetId(),
		Pid:       req.GetPid(),
		Type:      req.GetType(),
		Title:     req.GetTitle(),
		Name:      req.GetName(),
		Path:      req.GetPath(),
		Icon:      req.GetIcon(),
		MenuType:  req.GetMenuType(),
		URL:       req.GetUrl(),
		Component: req.GetComponent(),
		Extend:    req.GetExtend(),
		Remark:    req.GetRemark(),
		Status:    int16(req.GetStatus()),
	}
	err := s.userRuleRepo.UpsertOne(ctx, userRule)
	if err != nil {
		return nil, errorx.DataSQLErr.WithError(err).Err()
	}
	resp.Id = userRule.ID
	return resp, nil
}
