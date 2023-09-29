package biz

import (
	"context"
	"fkratos/api/paginator"
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"
	"fkratos/internal/dto"

	"github.com/fzf-labs/fpkg/orm"
	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type SysJobRepo interface {
	fkratos_sys_repo.ISysJobRepo
	GetJobIDToNameByIds(ctx context.Context, ids []string) (map[string]string, error)
	SysJobStore(ctx context.Context, req *v1.SysJobStoreReq) (*fkratos_sys_model.SysJob, error)
}

func NewJobUseCase(logger log.Logger, sysJobRepo SysJobRepo) *JobUseCase {
	l := log.NewHelper(log.With(logger, "module", "rpc_user/biz/job"))
	return &JobUseCase{
		log:        l,
		sysJobRepo: sysJobRepo,
	}
}

type JobUseCase struct {
	log        *log.Helper
	sysJobRepo SysJobRepo
}

func (d *JobUseCase) SysJobList(ctx context.Context, req *paginator.PaginatorReq) (*v1.SysJobListReply, error) {
	resp := new(v1.SysJobListReply)
	paginatorReq := &orm.PaginatorReq{}
	err := dto.Copy(req, paginatorReq)
	if err != nil {
		return nil, err
	}
	sysJobs, p, err := d.sysJobRepo.FindMultiByPaginator(ctx, paginatorReq)
	if err != nil {
		return nil, err
	}
	if len(sysJobs) > 0 {
		for _, v := range sysJobs {
			resp.List = append(resp.List, &v1.SysJobInfo{
				Id:        v.ID,
				Name:      v.Name,
				Code:      v.Code,
				Remark:    v.Remark,
				Status:    int32(v.Status),
				Sort:      int32(v.Sort),
				CreatedAt: timeutil.ToDateTimeStringByTime(v.CreatedAt),
				UpdatedAt: timeutil.ToDateTimeStringByTime(v.UpdatedAt),
			})
		}
	}
	err = copier.Copy(&resp.Paginator, p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *JobUseCase) SysJobInfo(ctx context.Context, req *v1.SysJobInfoReq) (*v1.SysJobInfoReply, error) {
	resp := new(v1.SysJobInfoReply)
	sysJob, err := d.sysJobRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	resp.Info = &v1.SysJobInfo{
		Id:        sysJob.ID,
		Name:      sysJob.Name,
		Code:      sysJob.Code,
		Remark:    sysJob.Remark,
		Status:    int32(sysJob.Status),
		Sort:      int32(sysJob.Sort),
		CreatedAt: timeutil.ToDateTimeStringByTime(sysJob.CreatedAt),
		UpdatedAt: timeutil.ToDateTimeStringByTime(sysJob.UpdatedAt),
	}
	return resp, nil
}

func (d *JobUseCase) SysJobStore(ctx context.Context, req *v1.SysJobStoreReq) (*v1.SysJobStoreReply, error) {
	resp := new(v1.SysJobStoreReply)
	_, err := d.sysJobRepo.SysJobStore(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *JobUseCase) SysJobDel(ctx context.Context, req *v1.SysJobDelReq) (*v1.SysJobDelReply, error) {
	resp := new(v1.SysJobDelReply)
	err := d.sysJobRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
