package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysDeptList 部门-列表
func (s *SysUseCase) SysDeptList(ctx context.Context, _ *pb.SysDeptListReq) (*pb.SysDeptListReply, error) {
	resp := new(pb.SysDeptListReply)
	result, err := s.sysDeptClient.SysDeptList(ctx, &sysV1.SysDeptListReq{})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithError(err).Err()
	}
	return resp, nil
}
