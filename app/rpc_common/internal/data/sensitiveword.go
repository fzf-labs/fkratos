package data

import (
	"fkratos/app/rpc_common/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.SensitiveWordRepo = (*SensitiveWordRepo)(nil)

func NewSensitiveWordRepo(data *Data, logger log.Logger) biz.SensitiveWordRepo {
	l := log.NewHelper(log.With(logger, "module", "rpc_common/data/sensitive_word"))
	return &SensitiveWordRepo{
		data: data,
		log:  l,
	}
}

type SensitiveWordRepo struct {
	data *Data
	log  *log.Helper
}
