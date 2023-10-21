package biz

import (
	"context"
	"fkratos/app/rpc_common/internal/data/gorm/fkratos_common_model"
	"fkratos/app/rpc_common/internal/data/gorm/fkratos_common_repo"

	"github.com/go-kratos/kratos/v2/log"
)

type SensitiveWordRepo interface {
	fkratos_common_repo.ISensitiveWordRepo
	SensitiveWordStore(ctx context.Context, data *fkratos_common_model.SensitiveWord) (*fkratos_common_model.SensitiveWord, error)
	SensitiveWordsCache(ctx context.Context) ([]string, error)
	SensitiveWordsCacheDel(ctx context.Context) error
}

func NewSensitiveWordUseCase(
	logger log.Logger,
	sensitiveWordRepo SensitiveWordRepo,
) *SensitiveWordUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/sensitiveWord"))
	return &SensitiveWordUseCase{
		log:               l,
		sensitiveWordRepo: sensitiveWordRepo,
	}
}

type SensitiveWordUseCase struct {
	log               *log.Helper
	sensitiveWordRepo SensitiveWordRepo
}
