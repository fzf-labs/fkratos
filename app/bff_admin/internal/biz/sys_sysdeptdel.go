package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// SysDeptDel 部门-删除
func (s *SysUseCase) SysDeptDel(ctx context.Context, req *pb.SysDeptDelReq) (*pb.SysDeptDelReply, error) {
	resp := new(pb.SysDeptDelReply)
	result, err := s.sysDeptClient.SysDeptDel(ctx, &sysV1.SysDeptDelReq{
		Ids: req.GetIds(),
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
