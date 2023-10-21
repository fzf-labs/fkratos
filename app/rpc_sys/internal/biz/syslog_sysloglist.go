package biz

import (
	"context"
	"fkratos/internal/dto"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/orm"
	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/jinzhu/copier"
)

// SysLogList 日志-列表
func (s *SysLogUseCase) SysLogList(ctx context.Context, req *pb.SysLogListReq) (*pb.SysLogListResp, error) {
	resp := &pb.SysLogListResp{}
	paginatorReq := &orm.PaginatorReq{}
	err := dto.Copy(req, paginatorReq)
	if err != nil {
		return nil, err
	}
	sysLogs, p, err := s.sysLogRepo.FindMultiByPaginator(ctx, paginatorReq)
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
		adminIDToNameByIds, err2 := s.sysAdminRepo.GetAdminIDToNameByIds(ctx, adminIds)
		if err2 != nil {
			return nil, err2
		}
		apiIDToNameByIds, err2 := s.sysAPIRepo.GetAPIIdToNameByIds(ctx, paths)
		if err2 != nil {
			return nil, err2
		}
		for _, v := range sysLogs {
			resp.List = append(resp.List, &pb.SysLogInfo{
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
