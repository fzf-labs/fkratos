package biz

import (
	"context"
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"

	"github.com/go-kratos/kratos/v2/log"
)

type SysLogRepo interface {
	fkratos_sys_repo.ISysLogRepo
	SysLogStore(ctx context.Context, req *pb.SysLogStoreReq) (*fkratos_sys_model.SysLog, error)
	SysLogStoreMQProducer(ctx context.Context, req *pb.SysLogStoreReq) error
}

func NewSysLogUseCase(
	logger log.Logger,
	sysLogRepo SysLogRepo,
	sysAdminRepo SysAdminRepo,
	sysAPIRepo SysAPIRepo,
) *SysLogUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/sysLog"))
	return &SysLogUseCase{
		log:          l,
		sysLogRepo:   sysLogRepo,
		sysAdminRepo: sysAdminRepo,
		sysAPIRepo:   sysAPIRepo,
	}
}

type SysLogUseCase struct {
	log          *log.Helper
	sysLogRepo   SysLogRepo
	sysAdminRepo SysAdminRepo
	sysAPIRepo   SysAPIRepo
}
