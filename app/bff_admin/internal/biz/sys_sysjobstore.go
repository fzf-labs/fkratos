package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysJobStore 岗位-保存
func (s *SysUseCase) SysJobStore(ctx context.Context, req *pb.SysJobStoreReq) (*pb.SysJobStoreReply, error) {
	resp := new(pb.SysJobStoreReply)
	result, err := s.sysJobClient.SysJobStore(ctx, &sysV1.SysJobStoreReq{
		Id:     req.GetId(),
		Name:   req.GetName(),
		Code:   req.GetCode(),
		Remark: req.GetRemark(),
		Status: req.GetStatus(),
		Sort:   req.GetSort(),
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
