package biz

import "github.com/go-kratos/kratos/v2/log"

type OfficialAccountRepo interface {
}

type OfficialAccountUseCase struct {
	log                 *log.Helper
	officialAccountRepo OfficialAccountRepo
}

func NewOfficialAccountUseCase(logger log.Logger, officialAccountRepo OfficialAccountRepo) *OfficialAccountUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/officialAccount"))
	return &OfficialAccountUseCase{
		log:                 l,
		officialAccountRepo: officialAccountRepo,
	}
}
