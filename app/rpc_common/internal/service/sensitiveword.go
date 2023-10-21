package service

import (
	pb "fkratos/api/rpc_common/v1"
	"fkratos/app/rpc_common/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewSensitiveWordService(
	logger log.Logger,
	sensitiveWordUseCase *biz.SensitiveWordUseCase,
) *SensitiveWordService {
	l := log.NewHelper(log.With(logger, "module", "service/sensitiveWord"))
	return &SensitiveWordService{
		log:                  l,
		sensitiveWordUseCase: sensitiveWordUseCase,
	}
}

type SensitiveWordService struct {
	pb.UnimplementedSensitiveWordServer
	log                  *log.Helper
	sensitiveWordUseCase *biz.SensitiveWordUseCase
}
