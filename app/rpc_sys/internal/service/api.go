package service

import (
	"context"
	"fkratos/app/rpc_sys/internal/biz"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type ApiService struct {
	pb.UnimplementedApiServer

	log        *log.Helper
	apiUseCase *biz.ApiUseCase
}

func NewApiService(logger log.Logger, apiUseCase *biz.ApiUseCase) *ApiService {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/service/api"))
	return &ApiService{
		log:        l,
		apiUseCase: apiUseCase,
	}
}

func (s *ApiService) SysApiList(ctx context.Context, req *pb.SysApiListReq) (*pb.SysApiListReply, error) {
	return s.apiUseCase.SysApiList(ctx, req)
}
