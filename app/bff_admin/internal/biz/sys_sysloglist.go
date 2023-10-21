package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysLogList 日志-列表
func (s *SysUseCase) SysLogList(ctx context.Context, req *pb.SysLogListReq) (*pb.SysLogListResp, error) {
	resp := new(pb.SysLogListResp)
	result, err := s.sysLogClient.SysLogList(ctx, &sysV1.SysLogListReq{
		Paginator: req.GetPaginator(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
