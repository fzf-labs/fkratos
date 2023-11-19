package biz

import (
	"context"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"

	pb "fkratos/api/bff_admin/v1"

	"github.com/jinzhu/copier"
)

// DashboardSpeech 仪表盘-一言
func (s *SysUseCase) DashboardSpeech(ctx context.Context, _ *pb.DashboardSpeechReq) (*pb.DashboardSpeechReply, error) {
	resp := &pb.DashboardSpeechReply{}
	result, err := s.sysDashboardClient.DashboardSpeech(ctx, &sysV1.DashboardSpeechReq{})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithError(err).Err()
	}
	return resp, nil
}
