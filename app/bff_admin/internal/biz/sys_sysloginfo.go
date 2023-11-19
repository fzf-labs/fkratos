package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysLogInfo 日志-信息
func (s *SysUseCase) SysLogInfo(ctx context.Context, req *pb.SysLogInfoReq) (*pb.SysLogInfoResp, error) {
	resp := new(pb.SysLogInfoResp)
	result, err := s.sysLogClient.SysLogInfo(ctx, &sysV1.SysLogInfoReq{
		Id: req.GetId(),
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
