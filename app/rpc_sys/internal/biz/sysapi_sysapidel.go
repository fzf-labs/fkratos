package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysAPIDel API-删除
func (s *SysAPIUseCase) SysAPIDel(ctx context.Context, req *pb.SysAPIDelReq) (*pb.SysAPIDelReply, error) {
	resp := &pb.SysAPIDelReply{}
	err := s.sysAPIRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
