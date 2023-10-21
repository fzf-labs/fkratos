package data

import (
	"fkratos/app/rpc_sys/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.SysAuthRepo = (*SysAuthRepo)(nil)

func NewSysAuthRepo(
	logger log.Logger,
	data *Data,
) biz.SysAuthRepo {
	l := log.NewHelper(log.With(logger, "module", "data/sysAuth"))
	return &SysAuthRepo{
		log:  l,
		data: data,
	}
}

type SysAuthRepo struct {
	log  *log.Helper
	data *Data
}
