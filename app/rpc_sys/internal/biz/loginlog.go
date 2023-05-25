package biz

import (
	"context"
	"fkratos/api/common"
	v1 "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

func NewLoginLogUseCase(logger log.Logger, sysLoginLogRepo SysLoginLogRepo, sysAdminRepo SysAdminRepo) *LoginLogUseCase {
	l := log.NewHelper(log.With(logger, "module", "rpc_user/biz/login_log"))
	return &LoginLogUseCase{
		log:             l,
		sysLoginLogRepo: sysLoginLogRepo,
		sysAdminRepo:    sysAdminRepo,
	}
}

type LoginLogUseCase struct {
	log             *log.Helper
	sysLoginLogRepo SysLoginLogRepo
	sysAdminRepo    SysAdminRepo
}

func (l *LoginLogUseCase) SysLoginLogList(ctx context.Context, req *common.SearchListReq) (*v1.SysLoginLogListResp, error) {
	resp := new(v1.SysLoginLogListResp)
	sysLogs, paginator, err := l.sysLoginLogRepo.SysLoginLogListBySearch(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(sysLogs) > 0 {
		adminIds := make([]string, 0)
		for _, v := range sysLogs {
			adminIds = append(adminIds, v.AdminID)
		}
		adminIdToNameByIds, err := l.sysAdminRepo.GetAdminIdToNameByIds(ctx, adminIds)
		if err != nil {
			return nil, err
		}
		for _, v := range sysLogs {
			resp.List = append(resp.List, &v1.SysLoginLogInfo{
				Id:        v.ID,
				TenantId:  v.TenantID,
				AdminID:   v.AdminID,
				Username:  adminIdToNameByIds[v.AdminID],
				Ip:        v.IP,
				Useragent: v.Useragent,
				LoginTime: timeutil.ToDateTimeStringByTime(v.LoginTime.Time),
			})
		}
	}
	err = copier.Copy(&resp.Paginator, paginator)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LoginLogUseCase) SysLoginLogInfo(ctx context.Context, req *v1.SysLoginLogInfoReq) (*v1.SysLoginLogInfoResp, error) {
	resp := new(v1.SysLoginLogInfoResp)
	sysLoginLog, err := l.sysLoginLogRepo.SysLoginLogInfoById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if sysLoginLog == nil {
		return resp, nil
	}
	adminIdToNameByIds, err := l.sysAdminRepo.GetAdminIdToNameByIds(ctx, []string{req.GetId()})
	if err != nil {
		return nil, err
	}
	resp.Info = &v1.SysLoginLogInfo{
		Id:        sysLoginLog.ID,
		AdminID:   sysLoginLog.AdminID,
		Username:  adminIdToNameByIds[sysLoginLog.AdminID],
		Ip:        sysLoginLog.IP,
		Useragent: sysLoginLog.Useragent,
		LoginTime: timeutil.ToDateTimeStringByTime(sysLoginLog.LoginTime.Time),
	}
	return resp, nil
}
