package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysJobList 岗位-列表
func (s *SysUseCase) SysJobList(ctx context.Context, req *pb.SysJobListReq) (*pb.SysJobListReply, error) {
	resp := new(pb.SysJobListReply)
	result, err := s.sysJobClient.SysJobList(ctx, &sysV1.SysJobListReq{
		Paginator: req.GetPaginator(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithError(err).Err()
	}
	return resp, nil
}
