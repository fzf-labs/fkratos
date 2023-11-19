package biz

import (
	"context"
	"fkratos/api/paginator"
	"fkratos/internal/dto"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_user/v1"

	"github.com/fzf-labs/fpkg/orm"
)

// UserGroupList 用户分组-列表
func (u *UserGroupUseCase) UserGroupList(ctx context.Context, req *pb.UserGroupListReq) (*pb.UserGroupListReply, error) {
	resp := &pb.UserGroupListReply{
		Paginator: &paginator.PaginatorReply{},
		List:      make([]*pb.UserGroupInfo, 0),
	}
	paginatorReq := &orm.PaginatorReq{}
	err := dto.Copy(paginatorReq, req.GetPaginator())
	if err != nil {
		return nil, errorx.DataFormattingError.WithError(err).Err()
	}
	result, p, err := u.userGroupRepo.FindMultiByPaginator(ctx, paginatorReq)
	if err != nil {
		return nil, errorx.DataSQLErr.WithError(err).Err()
	}
	if len(result) > 0 {
		err = dto.Copy(&resp.List, &result)
		if err != nil {
			return nil, errorx.DataFormattingError.WithError(err).Err()
		}
		err = dto.Copy(&resp.Paginator, p)
		if err != nil {
			return nil, errorx.DataFormattingError.WithError(err).Err()
		}
	}
	return resp, nil
}
