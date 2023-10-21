package service

import (
	"context"

	pb "fkratos/api/rpc_common/v1"
)

func (s *SensitiveWordService) SensitiveWordCheck(ctx context.Context, req *pb.SensitiveWordCheckReq) (*pb.SensitiveWordCheckResp, error) {
	return s.sensitiveWordUseCase.SensitiveWordCheck(ctx, req)
}
