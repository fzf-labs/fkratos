package service

import (
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewSysAPIService(
	logger log.Logger,
	sysAPIUseCase *biz.SysAPIUseCase,
) *SysAPIService {
	l := log.NewHelper(log.With(logger, "module", "service/sysAPI"))
	return &SysAPIService{
		log:           l,
		sysAPIUseCase: sysAPIUseCase,
	}
}

type SysAPIService struct {
	pb.UnimplementedSysAPIServer
	log           *log.Helper
	sysAPIUseCase *biz.SysAPIUseCase
}
