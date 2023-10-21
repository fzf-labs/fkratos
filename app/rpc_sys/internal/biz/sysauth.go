package biz

import "github.com/go-kratos/kratos/v2/log"

type SysAuthRepo interface {
}

func NewSysAuthUseCase(
	logger log.Logger,
	sysAdminRepo SysAdminRepo,
) *SysAuthUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/sysAuth"))
	return &SysAuthUseCase{
		log:          l,
		sysAdminRepo: sysAdminRepo,
	}
}

type SysAuthUseCase struct {
	log          *log.Helper
	sysAdminRepo SysAdminRepo
}
