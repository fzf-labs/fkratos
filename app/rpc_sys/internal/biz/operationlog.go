package biz

import (
	"context"
	"fkratos/api/common"
	v1 "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

func NewOperationLogUseCase(logger log.Logger, sysOperationLogRepo SysOperationLogRepo, sysAdminRepo SysAdminRepo, sysApiRepo SysApiRepo) *OperationLogUseCase {
	l := log.NewHelper(log.With(logger, "module", "rpc_user/biz/log"))
	return &OperationLogUseCase{
		log:                 l,
		sysOperationLogRepo: sysOperationLogRepo,
		sysAdminRepo:        sysAdminRepo,
		sysApiRepo:          sysApiRepo,
	}
}

type OperationLogUseCase struct {
	log                 *log.Helper
	sysOperationLogRepo SysOperationLogRepo
	sysAdminRepo        SysAdminRepo
	sysApiRepo          SysApiRepo
}

func (l *OperationLogUseCase) SysOperationLogList(ctx context.Context, req *common.SearchListReq) (*v1.SysOperationLogListResp, error) {
	resp := new(v1.SysOperationLogListResp)
	sysLogs, paginator, err := l.sysOperationLogRepo.SysOperationLogListBySearch(ctx, req)
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
		adminIdToNameByIds, err := l.sysAdminRepo.GetAdminIdToNameByIds(ctx, adminIds)
		if err != nil {
			return nil, err
		}
		apiIdToNameByIds, err := l.sysApiRepo.GetApiIdToNameByIds(ctx, paths)
		if err != nil {
			return nil, err
		}
		for _, v := range sysLogs {
			resp.List = append(resp.List, &v1.SysOperationLogInfo{
				Id:        v.ID,
				AdminID:   v.AdminID,
				Username:  adminIdToNameByIds[v.AdminID],
				Ip:        v.IP,
				Uri:       v.URI,
				UriDesc:   apiIdToNameByIds[v.URI],
				Useragent: v.Useragent,
				Req:       v.Req.String(),
				Resp:      v.Resp.String(),
				CreatedAt: timeutil.ToDateTimeStringByTime(v.CreatedAt.Time),
			})
		}
	}
	err = copier.Copy(&resp.Paginator, paginator)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *OperationLogUseCase) SysOperationLogInfo(ctx context.Context, req *v1.SysOperationLogInfoReq) (*v1.SysOperationLogInfoResp, error) {
	resp := new(v1.SysOperationLogInfoResp)
	sysLog, err := l.sysOperationLogRepo.SysOperationLogInfoById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if sysLog == nil {
		return resp, nil
	}
	adminIdToNameByIds, err := l.sysAdminRepo.GetAdminIdToNameByIds(ctx, []string{req.GetId()})
	if err != nil {
		return nil, err
	}
	apiIdToNameByIds, err := l.sysApiRepo.GetApiIdToNameByIds(ctx, []string{sysLog.URI})
	if err != nil {
		return nil, err
	}
	resp.Info = &v1.SysOperationLogInfo{
		Id:        sysLog.ID,
		AdminID:   adminIdToNameByIds[sysLog.AdminID],
		Username:  sysLog.AdminID,
		Ip:        sysLog.IP,
		Uri:       sysLog.URI,
		UriDesc:   apiIdToNameByIds[sysLog.URI],
		Useragent: sysLog.Useragent,
		Req:       sysLog.Req.String(),
		Resp:      sysLog.Resp.String(),
		CreatedAt: timeutil.ToDateTimeStringByTime(sysLog.CreatedAt.Time),
	}
	return resp, nil
}

func (l *OperationLogUseCase) SysOperationLogStore(ctx context.Context, req *v1.SysOperationLogStoreReq) (*v1.SysOperationLogStoreResp, error) {
	resp := new(v1.SysOperationLogStoreResp)
	_, err := l.sysOperationLogRepo.SysOperationLogStore(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *OperationLogUseCase) SysOperationLogStoreProducer(ctx context.Context, req *v1.SysOperationLogStoreReq) error {
	err := l.sysOperationLogRepo.SysOperationLogStoreMQProducer(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
