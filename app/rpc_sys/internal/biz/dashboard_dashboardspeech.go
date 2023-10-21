package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/third_api/speech"
)

// DashboardSpeech 仪表盘-一言
func (d *DashboardUseCase) DashboardSpeech(_ context.Context, _ *pb.DashboardSpeechReq) (*pb.DashboardSpeechReply, error) {
	resp := &pb.DashboardSpeechReply{}
	word, _ := speech.GetWord()
	resp.Word = word
	return resp, nil
}
