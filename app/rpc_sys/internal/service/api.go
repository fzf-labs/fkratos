package service

import (
	"context"
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewApiService(logger log.Logger, apiUseCase *biz.ApiUseCase) *ApiService {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/service/api"))
	return &ApiService{
		log:        l,
		apiUseCase: apiUseCase,
	}
}

type ApiService struct {
	pb.UnimplementedApiServer
	log        *log.Helper
	apiUseCase *biz.ApiUseCase
}

func (a *ApiService) SysApiList(ctx context.Context, req *pb.SysApiListReq) (*pb.SysApiListReply, error) {
	return a.apiUseCase.SysApiList(ctx, req)
}

func (a *ApiService) SysApiStore(ctx context.Context, req *pb.SysApiStoreReq) (*pb.SysApiStoreReply, error) {
	return a.apiUseCase.SysApiStore(ctx, req)
}

func (a *ApiService) SysApiDel(ctx context.Context, req *pb.SysApiDelReq) (*pb.SysApiDelReply, error) {
	return a.apiUseCase.SysApiDel(ctx, req)
}
