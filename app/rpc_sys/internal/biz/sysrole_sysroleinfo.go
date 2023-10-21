package biz

import (
	"context"
	"strings"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/util/timeutil"
)

// SysRoleInfo 角色-信息
func (s *SysRoleUseCase) SysRoleInfo(ctx context.Context, req *pb.SysRoleInfoReq) (*pb.SysRoleInfoResp, error) {
	resp := &pb.SysRoleInfoResp{}
	sysRole, err := s.sysRoleRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	split := strings.Split(sysRole.PermissionIds, ",")
	resp.Info = &pb.SysRoleInfo{
		Id:            sysRole.ID,
		Pid:           sysRole.Pid,
		Name:          sysRole.Name,
		Remark:        sysRole.Remark,
		Status:        int32(sysRole.Status),
		Sort:          int32(sysRole.Sort),
		PermissionIds: split,
		CreatedAt:     timeutil.ToDateTimeStringByTime(sysRole.CreatedAt),
		UpdatedAt:     timeutil.ToDateTimeStringByTime(sysRole.UpdatedAt),
		Children:      make([]*pb.SysRoleInfo, 0),
	}
	return resp, nil
}
