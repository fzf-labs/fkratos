package biz

import (
	"context"
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"

	"github.com/go-kratos/kratos/v2/log"
)

type SysDeptRepo interface {
	fkratos_sys_repo.ISysDeptRepo
	GetDeptIDToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
	SysDeptList(ctx context.Context) ([]*pb.SysDeptInfo, error)
	SysDeptStore(ctx context.Context, req *pb.SysDeptStoreReq) (*fkratos_sys_model.SysDept, error)
}

func NewSysDeptUseCase(
	logger log.Logger,
	sysDeptRepo SysDeptRepo,
) *SysDeptUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/sysDept"))
	return &SysDeptUseCase{
		log:         l,
		sysDeptRepo: sysDeptRepo,
	}
}

type SysDeptUseCase struct {
	log         *log.Helper
	sysDeptRepo SysDeptRepo
}
