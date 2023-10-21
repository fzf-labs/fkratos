package service

import (
	"context"

	pb "fkratos/api/bff_admin/v1"
)

func (s *SysService) DashboardSpeech(ctx context.Context, req *pb.DashboardSpeechReq) (*pb.DashboardSpeechReply, error) {
	return s.sysUseCase.DashboardSpeech(ctx, req)
}
