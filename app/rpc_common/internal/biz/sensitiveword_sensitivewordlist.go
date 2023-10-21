package biz

import (
	"context"
	"fkratos/internal/dto"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_common/v1"

	"github.com/fzf-labs/fpkg/orm"
	"github.com/fzf-labs/fpkg/util/jsonutil"
	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/jinzhu/copier"
)

// SensitiveWordList 敏感词-列表
func (s *SensitiveWordUseCase) SensitiveWordList(ctx context.Context, req *pb.SensitiveWordListReq) (*pb.SensitiveWordListReply, error) {
	resp := &pb.SensitiveWordListReply{}
	paginatorReq := &orm.PaginatorReq{}
	err := dto.Copy(paginatorReq, req)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	result, paginatorReply, err := s.sensitiveWordRepo.FindMultiByPaginator(ctx, paginatorReq)
	if err != nil {
		return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	for _, v := range result {
		labs := make([]string, 0)
		if err = jsonutil.Unmarshal(v.Labs, &labs); err != nil {
			return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
		resp.List = append(resp.List, &pb.SensitiveWordInfo{
			Id:        v.ID,
			Word:      v.Word,
			Labs:      labs,
			Desc:      v.Desc,
			CreatedAt: timeutil.ToDateTimeStringByTime(v.CreatedAt),
			UpdatedAt: timeutil.ToDateTimeStringByTime(v.UpdatedAt),
		})
	}
	err = copier.Copy(&resp.Paginator, paginatorReply)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
