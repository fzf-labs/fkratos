package data

import (
	"fkratos/app/rpc_wechat/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.OfficialAccountRepo = (*OfficialAccountRepo)(nil)

func NewOfficialAccountRepo(
	logger log.Logger,
	data *Data,
) biz.OfficialAccountRepo {
	l := log.NewHelper(log.With(logger, "module", "data/officialAccount"))
	return &OfficialAccountRepo{
		log:  l,
		data: data,
	}
}

type OfficialAccountRepo struct {
	log  *log.Helper
	data *Data
}
