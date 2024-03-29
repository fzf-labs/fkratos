package biz

import (
	"context"
	"fkratos/api/paginator"
	"fkratos/internal/dto"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_user/v1"

	"github.com/fzf-labs/fpkg/orm"
)

// UserList 用户-列表
func (u *UserUseCase) UserList(ctx context.Context, req *pb.UserListReq) (*pb.UserListReply, error) {
	resp := &pb.UserListReply{
		Paginator: &paginator.PaginatorReply{},
		List:      make([]*pb.UserInfo, 0),
	}
	paginatorReq := &orm.PaginatorReq{}
	err := dto.Copy(paginatorReq, req.GetPaginator())
	if err != nil {
		return nil, errorx.DataFormattingError.WithError(err).Err()
	}
	result, p, err := u.userRepo.FindMultiByPaginator(ctx, paginatorReq)
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
