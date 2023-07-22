package data

import (
	"context"
	"encoding/json"
	"fkratos/api/paginator"
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_dao"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/mq"
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
	data *Data
	log  *log.Helper
}

func (s *SysLogRepo) SysLogStoreMQProducer(ctx context.Context, req *v1.SysLogStoreReq) error {
	marshal, err := json.Marshal(req)
	if err != nil {
		return err
	}
	err = s.data.aysnqClient.NewTask(mq.SysLogStore, marshal)
	if err != nil {
		return err
	}
	return nil
}

func (s *SysLogRepo) SysLogStore(ctx context.Context, req *v1.SysLogStoreReq) (*fkratos_sys_model.SysLog, error) {
	sysLogDao := fkratos_sys_dao.Use(s.data.gorm).SysLog
	reqJson, err := json.Marshal(req.Req)
	if err != nil {
		return nil, err
	}
	respJson, err := json.Marshal(req.Resp)
	if err != nil {
		return nil, err
	}
	model := &fkratos_sys_model.SysLog{
		AdminID:   req.AdminID,
		IP:        req.Ip,
		URI:       req.Uri,
		Useragent: req.Useragent,
		Header:    nil,
		Req:       reqJson,
		Resp:      respJson,
	}
	err = sysLogDao.WithContext(ctx).Create(model)
	if err != nil {
		return nil, errorx.DataSqlErr.WithCause(err)
	}
	return model, nil
}

func (s *SysLogRepo) SysLogListBySearch(ctx context.Context, req *paginator.PaginatorReq) ([]*fkratos_sys_model.SysLog, *page.Page, error) {
	sysLogDao := fkratos_sys_dao.Use(s.data.gorm).SysLog
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

func (s *SysLogRepo) SysLogInfoById(ctx context.Context, id string) (*fkratos_sys_model.SysLog, error) {
	sysLogDao := fkratos_sys_dao.Use(s.data.gorm).SysLog
	sysLog, err := sysLogDao.WithContext(ctx).Where(sysLogDao.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return sysLog, nil
}
