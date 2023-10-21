package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/util/timeutil"
)

// SysAPIList API-列表
func (s *SysAPIUseCase) SysAPIList(ctx context.Context, req *pb.SysAPIListReq) (*pb.SysAPIListReply, error) {
	resp := &pb.SysAPIListReply{}
	list, err := s.sysAPIRepo.FindMultiByPermissionID(ctx, req.GetPermissionId())
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		resp.List = append(resp.List, &pb.SysAPIInfo{
			Id:           v.ID,
			PermissionID: v.PermissionID,
			Method:       v.Method,
			Path:         v.Path,
			Desc:         v.Desc,
			CreatedAt:    timeutil.ToDateTimeStringByTime(v.CreatedAt),
			UpdatedAt:    timeutil.ToDateTimeStringByTime(v.UpdatedAt),
		})
	}
	return resp, nil
}
