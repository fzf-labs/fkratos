package service

import (
	pb "fkratos/api/rpc_wechat/v1"
	"fkratos/app/rpc_wechat/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type MiniProgramService struct {
	pb.UnimplementedMiniProgramServer
	log                *log.Helper
	miniProgramUseCase *biz.MiniProgramUseCase
}

func NewMiniProgramService(logger log.Logger, miniProgramUseCase *biz.MiniProgramUseCase) *MiniProgramService {
	l := log.NewHelper(log.With(logger, "module", "service/miniProgram"))
	return &MiniProgramService{
		log:                l,
		miniProgramUseCase: miniProgramUseCase,
	}
}
