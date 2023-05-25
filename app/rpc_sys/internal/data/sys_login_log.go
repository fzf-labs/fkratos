package data

import (
	"context"
	"encoding/json"
	"fkratos/api/common"
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"
	"fkratos/app/rpc_sys/internal/data/gorm/rpc_sys_dao"
	"fkratos/app/rpc_sys/internal/data/gorm/rpc_sys_model"
	"fkratos/app/rpc_sys/internal/data/mq"
	"fkratos/internal/errorx"

	"github.com/fzf-labs/fpkg/page"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.SysLoginLogRepo = (*SysLoginLogRepo)(nil)

func NewSysLoginLogRepo(data *Data, logger log.Logger) biz.SysLoginLogRepo {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/data/sys_log"))
	return &SysLoginLogRepo{
		data: data,
		log:  l,
	}
}

type SysLoginLogRepo struct {
	data *Data
	log  *log.Helper
}

func (s *SysLoginLogRepo) SysLoginLogStoreMQProducer(ctx context.Context, req *v1.SysLoginLogStoreReq) error {
	marshal, err := json.Marshal(req)
	if err != nil {
		return err
	}
	err = s.data.aysnqClient.NewTask(mq.SysLoginLogStore, marshal)
	if err != nil {
		return err
	}
	return nil
}

func (s *SysLoginLogRepo) SysLoginLogStore(ctx context.Context, req *v1.SysLoginLogStoreReq) (*rpc_sys_model.SysLoginLog, error) {
	sysLogDao := rpc_sys_dao.Use(s.data.gorm).SysLoginLog
	model := &rpc_sys_model.SysLoginLog{
		AdminID:   req.AdminID,
		IP:        req.Ip,
		Useragent: req.Useragent,
	}
	err := sysLogDao.WithContext(ctx).Create(model)
	if err != nil {
		return nil, errorx.DataSqlErr.WithCause(err)
	}
	return model, nil
}

func (s *SysLoginLogRepo) SysLoginLogListBySearch(ctx context.Context, req *common.SearchListReq) ([]*rpc_sys_model.SysLoginLog, *page.Page, error) {
	sysLogDao := rpc_sys_dao.Use(s.data.gorm).SysLoginLog
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

func (s *SysLoginLogRepo) SysLoginLogInfoById(ctx context.Context, id string) (*rpc_sys_model.SysLoginLog, error) {
	sysLogDao := rpc_sys_dao.Use(s.data.gorm).SysLoginLog
	sysLog, err := sysLogDao.WithContext(ctx).Where(sysLogDao.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return sysLog, nil
}
