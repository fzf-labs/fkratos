package data

import (
	"context"
	"fkratos/api/common"
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"
	"fkratos/app/rpc_sys/internal/data/gorm/rpc_sys_dao"
	"fkratos/app/rpc_sys/internal/data/gorm/rpc_sys_model"
	"fkratos/internal/bootstrap/conf"
	"fkratos/internal/errorx"
	"strings"

	"github.com/fzf-labs/fpkg/conv"
	"github.com/fzf-labs/fpkg/page"
	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.SysJobRepo = (*SysJobRepo)(nil)

func NewSysJobRepo(data *Data, logger log.Logger) biz.SysJobRepo {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/data/sys_job"))
	return &SysJobRepo{
		data: data,
		log:  l,
	}
}

type SysJobRepo struct {
	config *conf.Bootstrap
	data   *Data
	log    *log.Helper
}

func (s *SysJobRepo) SysJobInfoById(ctx context.Context, id string) (*v1.SysJobInfo, error) {
	sysJobDao := rpc_sys_dao.Use(s.data.gorm).SysJob
	sysJob, err := sysJobDao.WithContext(ctx).Where(sysJobDao.ID.Eq(id)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithCause(err)
	}
	return &v1.SysJobInfo{
		Id:        sysJob.ID,
		Name:      sysJob.Name,
		Code:      sysJob.Code,
		Remark:    sysJob.Remark,
		Status:    int32(sysJob.Status),
		Sort:      int32(sysJob.Sort),
		CreatedAt: timeutil.ToDateTimeStringByTime(sysJob.CreatedAt),
		UpdatedAt: timeutil.ToDateTimeStringByTime(sysJob.UpdatedAt),
	}, nil
}

func (s *SysJobRepo) SysJobDelByIds(ctx context.Context, ids []string) error {
	sysJobDao := rpc_sys_dao.Use(s.data.gorm).SysJob
	_, err := sysJobDao.WithContext(ctx).Where(sysJobDao.ID.In(ids...)).Delete()
	if err != nil && err != gorm.ErrRecordNotFound {
		return errorx.DataSqlErr.WithCause(err)
	}
	return nil
}

func (s *SysJobRepo) SysJobStore(ctx context.Context, req *v1.SysJobStoreReq) (*rpc_sys_model.SysJob, error) {
	sysJobDao := rpc_sys_dao.Use(s.data.gorm).SysJob
	sysJob := &rpc_sys_model.SysJob{
		Name:   req.Name,
		Code:   req.Code,
		Remark: req.Remark,
		Sort:   int64(req.Sort),
		Status: int16(req.Status),
	}
	if req.Id != "" {
		_, err := sysJobDao.WithContext(ctx).Select(sysJobDao.Name, sysJobDao.Code, sysJobDao.Remark, sysJobDao.Sort, sysJobDao.Status).Where(sysJobDao.ID.Eq(req.Id)).Updates(sysJob)
		if err != nil {
			return nil, errorx.DataSqlErr.WithCause(err)
		}
	} else {
		err := sysJobDao.WithContext(ctx).Create(sysJob)
		if err != nil {
			return nil, errorx.DataSqlErr.WithCause(err)
		}
	}
	return sysJob, nil
}

func (s *SysJobRepo) SysJobListBySearch(ctx context.Context, req *common.SearchListReq) ([]*rpc_sys_model.SysJob, *page.Page, error) {

	sysJobDao := rpc_sys_dao.Use(s.data.gorm).SysJob
	query := sysJobDao.WithContext(ctx)
	if req.QuickSearch != "" {
		query = query.Where(sysJobDao.Name.Like(req.QuickSearch))
	} else {
		for _, search := range req.Search {
			if search.Field == "id" {
				query = query.Where(sysJobDao.ID.Eq(search.Val))
			}
			if search.Field == "name" {
				query = query.Where(sysJobDao.Name.Eq(search.Val))
			}
			if search.Field == "status" {
				query = query.Where(sysJobDao.Status.Eq(conv.Int16(search.Val)))
			}
			if search.Field == "createdAt" {
				ss := strings.Split(search.Val, ",")
				query = query.Where(sysJobDao.CreatedAt.Gte(timeutil.Carbon().Parse(ss[0]).Carbon2Time()), sysJobDao.CreatedAt.Lte(timeutil.Carbon().Parse(ss[1]).Carbon2Time()))
			}
			if search.Field == "updatedAt" {
				ss := strings.Split(search.Val, ",")
				query = query.Where(sysJobDao.UpdatedAt.Gte(timeutil.Carbon().Parse(ss[0]).Carbon2Time()), sysJobDao.UpdatedAt.Lte(timeutil.Carbon().Parse(ss[1]).Carbon2Time()))
			}
		}
	}
	if req.Order != "" {
		orderParam := strings.Split(req.Order, ",")
		if len(orderParam) != 2 {
			return nil, nil, errorx.ParamErr
		}
		switch orderParam[0] {
		case "id":
			if orderParam[1] == "desc" {
				query = query.Order(sysJobDao.ID.Desc())
			}
		case "createdAt":
			if orderParam[1] == "desc" {
				query = query.Order(sysJobDao.CreatedAt.Desc())
			}
		case "updatedAt":
			if orderParam[1] == "desc" {
				query = query.Order(sysJobDao.UpdatedAt.Desc())
			}
		}
	}
	queryCount := query
	total, err := queryCount.Count()
	if err != nil {
		return nil, nil, errorx.DataSqlErr.WithCause(err)
	}
	paginator := page.Paginator(int(req.Page), int(req.PageSize), int(total))
	sysJobs, err := query.Offset(paginator.Offset).Limit(paginator.Limit).Find()
	if err != nil {
		return nil, nil, errorx.DataSqlErr.WithCause(err)
	}
	return sysJobs, paginator, nil
}

func (s *SysJobRepo) GetJobIdToNameByIds(ctx context.Context, ids []string) (map[string]string, error) {
	resp := make(map[string]string)
	dao := rpc_sys_dao.Use(s.data.gorm).SysJob
	results, err := dao.WithContext(ctx).Where(dao.ID.In(ids...)).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithCause(err)
	}
	for _, v := range results {
		resp[v.ID] = v.Name
	}
	return resp, nil
}
