package biz

import (
	"context"
	v1 "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

func NewLogUseCase(logger log.Logger, sysLogRepo SysLogRepo, sysAdminRepo SysAdminRepo, sysApiRepo SysApiRepo) *LogUseCase {
	l := log.NewHelper(log.With(logger, "module", "rpc_user/biz/log"))
	return &LogUseCase{
		log:          l,
		sysLogRepo:   sysLogRepo,
		sysAdminRepo: sysAdminRepo,
		sysApiRepo:   sysApiRepo,
	}
}

type LogUseCase struct {
	log          *log.Helper
	sysLogRepo   SysLogRepo
	sysAdminRepo SysAdminRepo
	sysApiRepo   SysApiRepo
}

func (l *LogUseCase) SysLogList(ctx context.Context, req *common.SearchListReq) (*v1.SysLogListResp, error) {
	resp := new(v1.SysLogListResp)
	sysLogs, paginator, err := l.sysLogRepo.SysLogListBySearch(ctx, req)
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
			resp.List = append(resp.List, &v1.SysLogInfo{
				Id:        v.ID,
				AdminID:   v.AdminID,
				Username:  adminIdToNameByIds[v.AdminID],
				Ip:        v.IP,
				Uri:       v.URI,
				UriDesc:   apiIdToNameByIds[v.URI],
				Useragent: v.Useragent,
				Req:       v.Req.String(),
				Resp:      v.Resp.String(),
				CreatedAt: timeutil.ToDateTimeStringByTime(v.CreatedAt),
			})
		}
	}
	err = copier.Copy(&resp.Paginator, paginator)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LogUseCase) SysLogInfo(ctx context.Context, req *v1.SysLogInfoReq) (*v1.SysLogInfoResp, error) {
	resp := new(v1.SysLogInfoResp)
	sysLog, err := l.sysLogRepo.SysLogInfoById(ctx, req.GetId())
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
	resp.Info = &v1.SysLogInfo{
		Id:        sysLog.ID,
		AdminID:   adminIdToNameByIds[sysLog.AdminID],
		Username:  sysLog.AdminID,
		Ip:        sysLog.IP,
		Uri:       sysLog.URI,
		UriDesc:   apiIdToNameByIds[sysLog.URI],
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
