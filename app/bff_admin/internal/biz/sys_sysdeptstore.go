package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysDeptStore 部门-保存
func (s *SysUseCase) SysDeptStore(ctx context.Context, req *pb.SysDeptStoreReq) (*pb.SysDeptStoreReply, error) {
	resp := new(pb.SysDeptStoreReply)
	result, err := s.sysDeptClient.SysDeptStore(ctx, &sysV1.SysDeptStoreReq{
		Id:          req.GetId(),
		Pid:         req.GetPid(),
		Name:        req.GetName(),
		FullName:    req.GetFullName(),
		Responsible: req.GetResponsible(),
		Phone:       req.GetPhone(),
		Email:       req.GetEmail(),
		Type:        req.GetType(),
		Status:      req.GetStatus(),
		Sort:        req.GetSort(),
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
