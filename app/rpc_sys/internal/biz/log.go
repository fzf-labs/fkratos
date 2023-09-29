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

type SysLogRepo interface {
	fkratos_sys_repo.ISysLogRepo
	SysLogStore(ctx context.Context, req *v1.SysLogStoreReq) (*fkratos_sys_model.SysLog, error)
	SysLogStoreMQProducer(ctx context.Context, req *v1.SysLogStoreReq) error
}

func NewLogUseCase(logger log.Logger, sysLogRepo SysLogRepo, sysAdminRepo SysAdminRepo, sysAPIRepo SysAPIRepo) *LogUseCase {
	l := log.NewHelper(log.With(logger, "module", "rpc_user/biz/log"))
	return &LogUseCase{
		log:          l,
		sysLogRepo:   sysLogRepo,
		sysAdminRepo: sysAdminRepo,
		sysAPIRepo:   sysAPIRepo,
	}
}

type LogUseCase struct {
	log          *log.Helper
	sysLogRepo   SysLogRepo
	sysAdminRepo SysAdminRepo
	sysAPIRepo   SysAPIRepo
}

func (l *LogUseCase) SysLogList(ctx context.Context, req *paginator.PaginatorReq) (*v1.SysLogListResp, error) {
	resp := new(v1.SysLogListResp)
	paginatorReq := &orm.PaginatorReq{}
	err := dto.Copy(req, paginatorReq)
	if err != nil {
		return nil, err
	}
	sysLogs, p, err := l.sysLogRepo.FindMultiByPaginator(ctx, paginatorReq)
	if err != nil {
		return nil, err
	}
	if len(sysLogs) > 0 {
		adminIds := make([]string, 0)
		paths := make([]string, 0)
		for _, v := range sysLogs {
			adminIds = append(adminIds, v.AdminID)
			paths = append(paths, v.URI)
		}
		adminIDToNameByIds, err2 := l.sysAdminRepo.GetAdminIDToNameByIds(ctx, adminIds)
		if err2 != nil {
			return nil, err2
		}
		apiIDToNameByIds, err2 := l.sysAPIRepo.GetAPIIdToNameByIds(ctx, paths)
		if err2 != nil {
			return nil, err2
		}
		for _, v := range sysLogs {
			resp.List = append(resp.List, &v1.SysLogInfo{
				Id:        v.ID,
				AdminID:   v.AdminID,
				Username:  adminIDToNameByIds[v.AdminID],
				Ip:        v.IP,
				Uri:       v.URI,
				UriDesc:   apiIDToNameByIds[v.URI],
				Useragent: v.Useragent,
				Req:       v.Req.String(),
				Resp:      v.Resp.String(),
				CreatedAt: timeutil.ToDateTimeStringByTime(v.CreatedAt),
			})
		}
	}
	err = copier.Copy(&resp.Paginator, p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LogUseCase) SysLogInfo(ctx context.Context, req *v1.SysLogInfoReq) (*v1.SysLogInfoResp, error) {
	resp := new(v1.SysLogInfoResp)
	sysLog, err := l.sysLogRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if sysLog == nil {
		return resp, nil
	}
	adminIDToNameByIds, err := l.sysAdminRepo.GetAdminIDToNameByIds(ctx, []string{req.GetId()})
	if err != nil {
		return nil, err
	}
	apiIDToNameByIds, err := l.sysAPIRepo.GetAPIIdToNameByIds(ctx, []string{sysLog.URI})
	if err != nil {
		return nil, err
	}
	resp.Info = &v1.SysLogInfo{
		Id:        sysLog.ID,
		AdminID:   adminIDToNameByIds[sysLog.AdminID],
		Username:  sysLog.AdminID,
		Ip:        sysLog.IP,
		Uri:       sysLog.URI,
		UriDesc:   apiIDToNameByIds[sysLog.URI],
		Useragent: sysLog.Useragent,
		Req:       sysLog.Req.String(),
		Resp:      sysLog.Resp.String(),
		CreatedAt: timeutil.ToDateTimeStringByTime(sysLog.CreatedAt),
	}
	return resp, nil
}

func (l *LogUseCase) SysLogStore(ctx context.Context, req *v1.SysLogStoreReq) (*v1.SysLogStoreResp, error) {
	resp := new(v1.SysLogStoreResp)
	_, err := l.sysLogRepo.SysLogStore(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LogUseCase) SysLogStoreProducer(ctx context.Context, req *v1.SysLogStoreReq) error {
	err := l.sysLogRepo.SysLogStoreMQProducer(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
