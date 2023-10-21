package biz

import (
	"context"
	"fkratos/api/paginator"
	"fkratos/internal/dto"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_user/v1"

	"github.com/fzf-labs/fpkg/orm"
)

// UserRuleList 用户规则-列表
func (s *UserRuleUseCase) UserRuleList(ctx context.Context, req *pb.UserRuleListReq) (*pb.UserRuleListReply, error) {
	resp := &pb.UserRuleListReply{
		Paginator: &paginator.PaginatorReply{},
		List:      make([]*pb.UserRuleInfo, 0),
	}
	paginatorReq := &orm.PaginatorReq{}
	err := dto.Copy(paginatorReq, req.GetPaginator())
	if err != nil {
		return nil, errorx.DataFormattingError.WithMetadata(errorx.SetErrMetadata(err))
	}
	result, p, err := s.userRuleRepo.FindMultiByPaginator(ctx, paginatorReq)
	if err != nil {
		return nil, errorx.DataSQLErr.WithMetadata(errorx.SetErrMetadata(err))
	}
	if len(result) > 0 {
		err = dto.Copy(&resp.List, &result)
		if err != nil {
			return nil, errorx.DataFormattingError.WithMetadata(errorx.SetErrMetadata(err))
		}
		err = dto.Copy(&resp.Paginator, p)
		if err != nil {
			return nil, errorx.DataFormattingError.WithMetadata(errorx.SetErrMetadata(err))
		}
	}
	return resp, nil
}
