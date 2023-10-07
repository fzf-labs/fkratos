package service

import (
	"context"

	pb "fkratos/api/rpc_wechat/v1"
	"fkratos/app/rpc_wechat/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewOfficialAccountService(logger log.Logger, officialAccountUseCase *biz.OfficialAccountUseCase) *OfficialAccountService {
	l := log.NewHelper(log.With(logger, "module", "service/officialAccount"))
	return &OfficialAccountService{
		log:                    l,
		officialAccountUseCase: officialAccountUseCase,
	}
}

type OfficialAccountService struct {
	pb.UnimplementedOfficialAccountServer
	log                    *log.Helper
	officialAccountUseCase *biz.OfficialAccountUseCase
}

func (s *OfficialAccountService) OfficialAccountMenu(ctx context.Context, req *pb.OfficialAccountMenuReq) (*pb.OfficialAccountMenuReply, error) {
	return s.officialAccountUseCase.OfficialAccountMenu(ctx, req)
}
