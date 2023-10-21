package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysManageDel 管理员-删除
func (s *SysAdminUseCase) SysManageDel(ctx context.Context, req *pb.SysManageDelReq) (*pb.SysManageDelReply, error) {
	resp := &pb.SysManageDelReply{}
	err := s.sysAdminRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
