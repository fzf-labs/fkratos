package biz

import (
	"context"
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"

	"github.com/go-kratos/kratos/v2/log"
)

type SysJobRepo interface {
	fkratos_sys_repo.ISysJobRepo
	GetJobIDToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
	SysJobStore(ctx context.Context, req *pb.SysJobStoreReq) (*fkratos_sys_model.SysJob, error)
}

func NewSysJobUseCase(
	logger log.Logger,
	sysJobRepo SysJobRepo,
) *SysJobUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/sysJob"))
	return &SysJobUseCase{
		log:        l,
		sysJobRepo: sysJobRepo,
	}
}

type SysJobUseCase struct {
	log        *log.Helper
	sysJobRepo SysJobRepo
}
