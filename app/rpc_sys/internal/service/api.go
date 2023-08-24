package service

import (
	"context"
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewAPIService(logger log.Logger, apiUseCase *biz.APIUseCase) *APIService {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/service/api"))
	return &APIService{
		log:        l,
		apiUseCase: apiUseCase,
	}
}

type APIService struct {
	pb.UnimplementedAPIServer
	log        *log.Helper
	apiUseCase *biz.APIUseCase
}

func (a *APIService) SysAPIList(ctx context.Context, req *pb.SysAPIListReq) (*pb.SysAPIListReply, error) {
	return a.apiUseCase.SysAPIList(ctx, req)
}

func (a *APIService) SysAPIStore(ctx context.Context, req *pb.SysAPIStoreReq) (*pb.SysAPIStoreReply, error) {
	return a.apiUseCase.SysAPIStore(ctx, req)
}

func (a *APIService) SysAPIDel(ctx context.Context, req *pb.SysAPIDelReq) (*pb.SysAPIDelReply, error) {
	return a.apiUseCase.SysAPIDel(ctx, req)
}
