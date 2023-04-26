package data

import (
	"fkratos/app/rpc_sys/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.AuthRepo = (*AuthRepo)(nil)

func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	l := log.NewHelper(log.With(logger, "module", "auth/repo/auth-service"))
	return &AuthRepo{
		data: data,
		log:  l,
	}
}

type AuthRepo struct {
	data *Data
	log  *log.Helper
}
