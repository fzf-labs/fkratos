package service

import (
	"context"

	pb "fkratos/api/rpc_common/v1"
)

func (s *SensitiveWordService) SensitiveWordList(ctx context.Context, req *pb.SensitiveWordListReq) (*pb.SensitiveWordListReply, error) {
	return s.sensitiveWordUseCase.SensitiveWordList(ctx, req)
}
