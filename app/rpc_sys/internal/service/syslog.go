package service

import (
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewSysLogService(
	logger log.Logger,
	sysLogUseCase *biz.SysLogUseCase,
) *SysLogService {
	l := log.NewHelper(log.With(logger, "module", "service/sysLog"))
	return &SysLogService{
		log:           l,
		sysLogUseCase: sysLogUseCase,
	}
}

type SysLogService struct {
	pb.UnimplementedSysLogServer
	log           *log.Helper
	sysLogUseCase *biz.SysLogUseCase
}

func (s *SysLogService) SysLogStoreConsumer(payload []byte) error {
	return s.sysLogUseCase.SysLogStoreConsumer(payload)
}
