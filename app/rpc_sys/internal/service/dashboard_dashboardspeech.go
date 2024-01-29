package service

import (
	"context"
	pb "fkratos/api/rpc_sys/v1"
	"time"
)

// DashboardSpeech 仪表盘-一言
func (d *DashboardService) DashboardSpeech(ctx context.Context, req *pb.DashboardSpeechReq) (*pb.DashboardSpeechReply, error) {
	SleepContext(ctx, time.Second*10)
	return d.dashboardUseCase.DashboardSpeech(ctx, req)
}

func SleepContext(ctx context.Context, delay time.Duration) {
	timer := time.NewTimer(delay)
	select {
	case <-ctx.Done():
		if !timer.Stop() {
			<-timer.C
		}
	case <-timer.C:
	}
}
