package biz

import (
	"context"
	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/util/timeutil"
)

// SysLogInfo 日志-信息
func (s *SysLogUseCase) SysLogInfo(ctx context.Context, req *pb.SysLogInfoReq) (*pb.SysLogInfoResp, error) {
	resp := &pb.SysLogInfoResp{}
	sysLog, err := s.sysLogRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if sysLog == nil {
		return resp, nil
	}
	adminIDToNameByIds, err := s.sysAdminRepo.GetAdminIDToNameByIds(ctx, []string{req.GetId()})
	if err != nil {
		return nil, err
	}
	apiIDToNameByIds, err := s.sysAPIRepo.GetAPIIdToNameByIds(ctx, []string{sysLog.URI})
	if err != nil {
		return nil, err
	}
	resp.Info = &pb.SysLogInfo{
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
