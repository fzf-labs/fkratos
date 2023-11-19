package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/util/jsonutil"
	"github.com/fzf-labs/fpkg/util/timeutil"
)

// SysManageInfo 管理员-信息
func (s *SysAdminUseCase) SysManageInfo(ctx context.Context, req *pb.SysManageInfoReq) (*pb.SysManageInfoReply, error) {
	resp := &pb.SysManageInfoReply{}
	sysAdmin, err := s.sysAdminRepo.FindOneCacheByID(ctx, req.GetAdminId())
	if err != nil {
		return nil, err
	}
	if sysAdmin != nil {
		roleIDToNameByIds := make(map[string]string)
		jobIDToNameByIds := make(map[string]string)
		deptIDToNameByIds := make(map[string]string)
		tmpRoleIds := make([]string, 0)
		err2 := jsonutil.Unmarshal(sysAdmin.RoleIds, &tmpRoleIds)
		if err2 != nil {
			return nil, err2
		}
		if len(tmpRoleIds) > 0 {
			roleIDToNameByIds, err2 = s.sysRoleRepo.GetRoleIDToNameByIds(ctx, tmpRoleIds)
			if err2 != nil {
				return nil, err2
			}
		}
		if sysAdmin.JobID != "" {
			jobIDToNameByIds, err2 = s.sysJobRepo.GetJobIDToNameByIds(ctx, []string{sysAdmin.JobID})
			if err2 != nil {
				return nil, err2
			}
		}
		if sysAdmin.DeptID != "" {
			deptIDToNameByIds, err2 = s.sysDeptRepo.GetDeptIDToNameByIds(ctx, []string{sysAdmin.DeptID})
			if err2 != nil {
				return nil, err2
			}
		}
		roleNames := make([]string, 0)
		for _, v := range tmpRoleIds {
			roleNames = append(roleNames, roleIDToNameByIds[v])
		}
		resp.Info = &pb.SysManageInfo{
			Id:        sysAdmin.ID,
			Username:  sysAdmin.Username,
			Nickname:  sysAdmin.Nickname,
			Avatar:    sysAdmin.Avatar,
			Gender:    int32(sysAdmin.Gender),
			Email:     sysAdmin.Email,
			Mobile:    sysAdmin.Mobile,
			JobID:     sysAdmin.JobID,
			DeptID:    sysAdmin.DeptID,
			RoleIds:   tmpRoleIds,
			JobName:   jobIDToNameByIds[sysAdmin.JobID],
			DeptName:  deptIDToNameByIds[sysAdmin.JobID],
			RoleNames: roleNames,
			Motto:     sysAdmin.Motto,
			Status:    int32(sysAdmin.Status),
			CreatedAt: timeutil.ToDateTimeStringByTime(sysAdmin.CreatedAt),
			UpdatedAt: timeutil.ToDateTimeStringByTime(sysAdmin.UpdatedAt),
		}
	}
	return resp, nil
}
