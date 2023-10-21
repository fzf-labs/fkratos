package service

import (
	"context"

	pb "fkratos/api/rpc_common/v1"
)

func (s *SensitiveWordService) SensitiveWordDel(ctx context.Context, req *pb.SensitiveWordDelReq) (*pb.SensitiveWordDelReply, error) {
	return s.sensitiveWordUseCase.SensitiveWordDel(ctx, req)
}
