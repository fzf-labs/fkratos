package biz

import (
	"context"

	pb "fkratos/api/rpc_common/v1"
)

// SensitiveWordDel 敏感词-删除
func (s *SensitiveWordUseCase) SensitiveWordDel(ctx context.Context, req *pb.SensitiveWordDelReq) (*pb.SensitiveWordDelReply, error) {
	resp := &pb.SensitiveWordDelReply{}
	err := s.sensitiveWordRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
