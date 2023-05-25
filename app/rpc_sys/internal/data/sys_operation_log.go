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

var _ biz.SysOperationLogRepo = (*SysOperationLogRepo)(nil)

func NewSysOperationLogRepo(data *Data, logger log.Logger) biz.SysOperationLogRepo {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/data/sys_log"))
	return &SysOperationLogRepo{
		data: data,
		log:  l,
	}
}

type SysOperationLogRepo struct {
	data *Data
	log  *log.Helper
}

func (s *SysOperationLogRepo) SysOperationLogStoreMQProducer(ctx context.Context, req *v1.SysOperationLogStoreReq) error {
	marshal, err := json.Marshal(req)
	if err != nil {
		return err
	}
	err = s.data.aysnqClient.NewTask(mq.SysOperationLogStore, marshal)
	if err != nil {
		return err
	}
	return nil
}

func (s *SysOperationLogRepo) SysOperationLogStore(ctx context.Context, req *v1.SysOperationLogStoreReq) (*rpc_sys_model.SysOperationLog, error) {
	sysLogDao := rpc_sys_dao.Use(s.data.gorm).SysOperationLog
	reqJson, err := json.Marshal(req.Req)
	if err != nil {
		return nil, err
	}
	respJson, err := json.Marshal(req.Resp)
	if err != nil {
		return nil, err
	}
	model := &rpc_sys_model.SysOperationLog{
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

func (s *SysOperationLogRepo) SysOperationLogListBySearch(ctx context.Context, req *common.SearchListReq) ([]*rpc_sys_model.SysOperationLog, *page.Page, error) {
	sysLogDao := rpc_sys_dao.Use(s.data.gorm).SysOperationLog
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

func (s *SysOperationLogRepo) SysOperationLogInfoById(ctx context.Context, id string) (*rpc_sys_model.SysOperationLog, error) {
	sysLogDao := rpc_sys_dao.Use(s.data.gorm).SysOperationLog
	sysLog, err := sysLogDao.WithContext(ctx).Where(sysLogDao.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return sysLog, nil
}
