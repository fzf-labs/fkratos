package service

import (
	"context"

	pb "fkratos/api/rpc_common/v1"
)

func (s *SensitiveWordService) SensitiveWordStore(ctx context.Context, req *pb.SensitiveWordStoreReq) (*pb.SensitiveWordStoreReply, error) {
	return s.sensitiveWordUseCase.SensitiveWordStore(ctx, req)
}
