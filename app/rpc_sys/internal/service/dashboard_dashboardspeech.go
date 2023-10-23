package service

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"
)

// DashboardSpeech 仪表盘-一言
func (d *DashboardService) DashboardSpeech(ctx context.Context, req *pb.DashboardSpeechReq) (*pb.DashboardSpeechReply, error) {
	return d.dashboardUseCase.DashboardSpeech(ctx, req)
}
