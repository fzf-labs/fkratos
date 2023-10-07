package biz

import (
	"context"

	pb "fkratos/api/rpc_wechat/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type OfficialAccountRepo interface {
}

func NewOfficialAccountUseCase(logger log.Logger, officialAccountRepo OfficialAccountRepo) *OfficialAccountUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/officialAccount"))
	return &OfficialAccountUseCase{
		log:                 l,
		officialAccountRepo: officialAccountRepo,
	}
}

type OfficialAccountUseCase struct {
	log                 *log.Helper
	officialAccountRepo OfficialAccountRepo
}

func (s *OfficialAccountUseCase) OfficialAccountMenu(ctx context.Context, req *pb.OfficialAccountMenuReq) (*pb.OfficialAccountMenuReply, error) {
	return &pb.OfficialAccountMenuReply{}, nil
}
