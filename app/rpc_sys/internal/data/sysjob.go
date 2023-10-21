package data

import (
	"context"
	"errors"
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_dao"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"
	"fkratos/internal/errorx"

	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.SysJobRepo = (*SysJobRepo)(nil)

func NewSysJobRepo(
	data *Data,
	logger log.Logger,
	sysJobRepo *fkratos_sys_repo.SysJobRepo,
) biz.SysJobRepo {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/data/sys_job"))
	return &SysJobRepo{
		data:       data,
		log:        l,
		SysJobRepo: sysJobRepo,
	}
}

type SysJobRepo struct {
	data *Data
	log  *log.Helper
	*fkratos_sys_repo.SysJobRepo
}

func (s *SysJobRepo) SysJobInfoByID(ctx context.Context, id string) (*v1.SysJobInfo, error) {
	sysJobDao := fkratos_sys_dao.Use(s.data.gorm).SysJob
	sysJob, err := sysJobDao.WithContext(ctx).Where(sysJobDao.ID.Eq(id)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
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
	sysJobDao := fkratos_sys_dao.Use(s.data.gorm).SysJob
	_, err := sysJobDao.WithContext(ctx).Where(sysJobDao.ID.In(ids...)).Delete()
	if err != nil && err != gorm.ErrRecordNotFound {
		return errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return nil
}

func (s *SysJobRepo) SysJobStore(ctx context.Context, req *v1.SysJobStoreReq) (*fkratos_sys_model.SysJob, error) {
	sysJobDao := fkratos_sys_dao.Use(s.data.gorm).SysJob
	sysJob := &fkratos_sys_model.SysJob{
		Name:   req.Name,
		Code:   req.Code,
		Remark: req.Remark,
		Sort:   int64(req.Sort),
		Status: int16(req.Status),
	}
	if req.Id != "" {
		_, err := sysJobDao.WithContext(ctx).Select(sysJobDao.Name, sysJobDao.Code, sysJobDao.Remark, sysJobDao.Sort, sysJobDao.Status).Where(sysJobDao.ID.Eq(req.Id)).Updates(sysJob)
		if err != nil {
			return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
	} else {
		err := sysJobDao.WithContext(ctx).Create(sysJob)
		if err != nil {
			return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
	}
	return sysJob, nil
}

func (s *SysJobRepo) GetJobIDToNameByIds(ctx context.Context, ids []string) (map[string]string, error) {
	resp := make(map[string]string)
	dao := fkratos_sys_dao.Use(s.data.gorm).SysJob
	results, err := dao.WithContext(ctx).Where(dao.ID.In(ids...)).Find()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	for _, v := range results {
		resp[v.ID] = v.Name
	}
	return resp, nil
}
