package data

import (
	"context"
	"encoding/json"
	"errors"
	"fkratos/api/paginator"
	pb "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_dao"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"
	"fkratos/app/rpc_sys/internal/data/mq"
	"fkratos/internal/errorx"

	"github.com/fzf-labs/fpkg/page"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.SysLogRepo = (*SysLogRepo)(nil)

func NewSysLogRepo(
	data *Data,
	logger log.Logger,
	sysLogRepo *fkratos_sys_repo.SysLogRepo,
) biz.SysLogRepo {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/data/sys_log"))
	return &SysLogRepo{
		data:       data,
		log:        l,
		SysLogRepo: sysLogRepo,
	}
}

type SysLogRepo struct {
	data *Data
	log  *log.Helper
	*fkratos_sys_repo.SysLogRepo
}

func (s *SysLogRepo) SysLogStoreMQProducer(_ context.Context, req *pb.SysLogStoreReq) error {
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

func (s *SysLogRepo) SysLogStore(ctx context.Context, req *pb.SysLogStoreReq) (*fkratos_sys_model.SysLog, error) {
	sysLogDao := fkratos_sys_dao.Use(s.data.gorm).SysLog
	reqJSON, err := json.Marshal(req.Req)
	if err != nil {
		return nil, err
	}
	respJSON, err := json.Marshal(req.Resp)
	if err != nil {
		return nil, err
	}
	model := &fkratos_sys_model.SysLog{
		AdminID:   req.AdminID,
		IP:        req.Ip,
		URI:       req.Uri,
		Useragent: req.Useragent,
		Header:    nil,
		Req:       reqJSON,
		Resp:      respJSON,
	}
	err = sysLogDao.WithContext(ctx).Create(model)
	if err != nil {
		return nil, errorx.DataSQLErr.WithError(err).Err()
	}
	return model, nil
}

func (s *SysLogRepo) SysLogListBySearch(ctx context.Context, req *paginator.PaginatorReq) ([]*fkratos_sys_model.SysLog, *page.Page, error) {
	sysLogDao := fkratos_sys_dao.Use(s.data.gorm).SysLog
	count, err := sysLogDao.WithContext(ctx).Count()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, errorx.DataSQLErr.WithError(err).Err()
	}
	p := page.Paginator(int(req.Page), int(req.PageSize), int(count))
	sysLogs, err := sysLogDao.WithContext(ctx).Offset(p.Offset).Limit(p.Limit).Find()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, errorx.DataSQLErr.WithError(err).Err()
	}
	return sysLogs, p, nil
}

func (s *SysLogRepo) SysLogInfoByID(ctx context.Context, id string) (*fkratos_sys_model.SysLog, error) {
	sysLogDao := fkratos_sys_dao.Use(s.data.gorm).SysLog
	sysLog, err := sysLogDao.WithContext(ctx).Where(sysLogDao.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return sysLog, nil
}
