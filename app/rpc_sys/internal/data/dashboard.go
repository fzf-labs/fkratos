package data

import (
	"fkratos/app/rpc_sys/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.DashboardRepo = (*DashboardRepo)(nil)

func NewDashboardRepo(logger log.Logger, data *Data) biz.DashboardRepo {
	l := log.NewHelper(log.With(logger, "module", "data/dashboard"))
	return &DashboardRepo{
		log:  l,
		data: data,
	}
}

type DashboardRepo struct {
	log  *log.Helper
	data *Data
}
