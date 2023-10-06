package data

import (
	"fkratos/app/rpc_wechat/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.MiniProgramRepo = (*MiniProgramRepo)(nil)

func NewMiniProgramRepo(logger log.Logger, data *Data) biz.MiniProgramRepo {
	l := log.NewHelper(log.With(logger, "module", "data/miniProgram"))
	return &MiniProgramRepo{
		log:  l,
		data: data,
	}
}

type MiniProgramRepo struct {
	log  *log.Helper
	data *Data
}
