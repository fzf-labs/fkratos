package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// SysJobDel 岗位-删除
func (s *SysJobUseCase) SysJobDel(ctx context.Context, req *pb.SysJobDelReq) (*pb.SysJobDelReply, error) {
	resp := &pb.SysJobDelReply{}
	err := s.sysJobRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
