package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysManageStore 管理员-保存
func (s *SysAdminUseCase) SysManageStore(ctx context.Context, req *pb.SysManageStoreReq) (*pb.SysManageStoreReply, error) {
	resp := &pb.SysManageStoreReply{}
	_, err := s.sysAdminRepo.SysManageStore(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
