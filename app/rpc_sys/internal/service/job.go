package service

import (
	"context"
	"fkratos/api/paginator"
	"fkratos/app/rpc_sys/internal/biz"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type JobService struct {
	pb.UnimplementedJobServer

	log        *log.Helper
	jobUseCase *biz.JobUseCase
}

func NewJobService(logger log.Logger, jobUseCase *biz.JobUseCase) *JobService {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/service/job"))
	return &JobService{
		log:        l,
		jobUseCase: jobUseCase,
	}
}

func (s *JobService) SysJobList(ctx context.Context, req *paginator.PaginatorReq) (*pb.SysJobListReply, error) {
	return s.jobUseCase.SysJobList(ctx, req)
}
func (s *JobService) SysJobInfo(ctx context.Context, req *pb.SysJobInfoReq) (*pb.SysJobInfoReply, error) {
	return s.jobUseCase.SysJobInfo(ctx, req)
}
func (s *JobService) SysJobStore(ctx context.Context, req *pb.SysJobStoreReq) (*pb.SysJobStoreReply, error) {
	return s.jobUseCase.SysJobStore(ctx, req)
}
func (s *JobService) SysJobDel(ctx context.Context, req *pb.SysJobDelReq) (*pb.SysJobDelReply, error) {
	return s.jobUseCase.SysJobDel(ctx, req)
}
