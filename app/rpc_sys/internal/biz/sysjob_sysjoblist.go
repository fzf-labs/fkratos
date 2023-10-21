package biz

import (
	"context"
	"fkratos/internal/dto"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/orm"
	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/jinzhu/copier"
)

// SysJobList 岗位-列表
func (s *SysJobUseCase) SysJobList(ctx context.Context, req *pb.SysJobListReq) (*pb.SysJobListReply, error) {
	resp := &pb.SysJobListReply{}
	paginatorReq := &orm.PaginatorReq{}
	err := dto.Copy(req, paginatorReq)
	if err != nil {
		return nil, err
	}
	sysJobs, p, err := s.sysJobRepo.FindMultiByPaginator(ctx, paginatorReq)
	if err != nil {
		return nil, err
	}
	if len(sysJobs) > 0 {
		for _, v := range sysJobs {
			resp.List = append(resp.List, &pb.SysJobInfo{
				Id:        v.ID,
				Name:      v.Name,
				Code:      v.Code,
				Remark:    v.Remark,
				Status:    int32(v.Status),
				Sort:      int32(v.Sort),
				CreatedAt: timeutil.ToDateTimeStringByTime(v.CreatedAt),
				UpdatedAt: timeutil.ToDateTimeStringByTime(v.UpdatedAt),
			})
		}
	}
	err = copier.Copy(&resp.Paginator, p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
