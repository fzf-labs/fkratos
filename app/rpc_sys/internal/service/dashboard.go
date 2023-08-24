package service

import (
	"context"
	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/third_api/speech"
	"github.com/go-kratos/kratos/v2/log"
)

type DashboardService struct {
	pb.UnimplementedDashboardServer

	log *log.Helper
}

func NewDashboardService(logger log.Logger) *DashboardService {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/service/dashboard"))
	return &DashboardService{
		log: l,
	}
}
func (s *DashboardService) DashboardSpeech(_ context.Context, _ *pb.DashboardSpeechReq) (*pb.DashboardSpeechReply, error) {
	resp := new(pb.DashboardSpeechReply)
	word, _ := speech.GetWord()
	resp.Word = word
	return resp, nil
}
