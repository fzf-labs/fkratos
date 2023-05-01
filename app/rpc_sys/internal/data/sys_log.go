package data

import (
	"context"
	"fkratos/api/common"
	"fkratos/app/rpc_sys/internal/biz"
	"fkratos/app/rpc_sys/internal/data/gorm/rpc_sys_dao"
	"fkratos/app/rpc_sys/internal/data/gorm/rpc_sys_model"
	"fkratos/internal/bootstrap/conf"
	"fkratos/internal/errorx"

	"github.com/fzf-labs/fpkg/page"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.SysLogRepo = (*SysLogRepo)(nil)

func NewSysLogRepo(data *Data, logger log.Logger) biz.SysLogRepo {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/data/sys_log"))
	return &SysLogRepo{
		data: data,
		log:  l,
	}
}

type SysLogRepo struct {
	config *conf.Bootstrap
	data   *Data
	log    *log.Helper
}

func (s *SysLogRepo) SysLogListBySearch(ctx context.Context, req *common.SearchListReq) ([]*rpc_sys_model.SysLog, *page.Page, error) {
	sysLogDao := rpc_sys_dao.Use(s.data.gorm).SysLog
	count, err := sysLogDao.WithContext(ctx).Count()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, nil, errorx.DataSqlErr.WithCause(err)
	}
	paginator := page.Paginator(int(req.Page), int(req.PageSize), int(count))
	sysLogs, err := sysLogDao.WithContext(ctx).Offset(paginator.Offset).Limit(paginator.Limit).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, nil, errorx.DataSqlErr.WithCause(err)
	}
	return sysLogs, paginator, nil
}

func (s *SysLogRepo) SysLogInfoById(ctx context.Context, id string) (*rpc_sys_model.SysLog, error) {
	sysLogDao := rpc_sys_dao.Use(s.data.gorm).SysLog
	sysLog, err := sysLogDao.WithContext(ctx).Where(sysLogDao.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return sysLog, nil
}
