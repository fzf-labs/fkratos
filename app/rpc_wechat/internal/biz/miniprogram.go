package biz

import "github.com/go-kratos/kratos/v2/log"

type MiniProgramRepo interface {
}

func NewMiniProgramUseCase(
	logger log.Logger,
	miniProgramRepo MiniProgramRepo,
) *MiniProgramUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/miniProgram"))
	return &MiniProgramUseCase{
		log:             l,
		miniProgramRepo: miniProgramRepo,
	}
}

type MiniProgramUseCase struct {
	log             *log.Helper
	miniProgramRepo MiniProgramRepo
}
