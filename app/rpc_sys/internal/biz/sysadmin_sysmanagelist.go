package biz

import (
	"context"
	"fkratos/internal/dto"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/orm"
	"github.com/fzf-labs/fpkg/util/jsonutil"
	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/jinzhu/copier"
)

// SysManageList 管理员-列表
func (s *SysAdminUseCase) SysManageList(ctx context.Context, req *pb.SysManageListReq) (*pb.SysManageListReply, error) {
	resp := &pb.SysManageListReply{}
	paginatorReq := &orm.PaginatorReq{}
	err := dto.Copy(paginatorReq, req.GetPaginator())
	if err != nil {
		return nil, err
	}
	sysAdmins, paginatorReply, err := s.sysAdminRepo.FindMultiByPaginator(ctx, paginatorReq)
	if err != nil {
		return nil, err
	}
	if len(sysAdmins) > 0 {
		list := make([]*pb.SysManageInfo, 0)
		roleIds := make([]string, 0)
		jobIds := make([]string, 0)
		deptIds := make([]string, 0)
		for _, v := range sysAdmins {
			tmpRoleIds := make([]string, 0)
			err2 := jsonutil.Unmarshal(v.RoleIds, &tmpRoleIds)
			if err2 != nil {
				return nil, err2
			}
			list = append(list, &pb.SysManageInfo{
				Id:        v.ID,
				Username:  v.Username,
				Nickname:  v.Nickname,
				Avatar:    v.Avatar,
				Gender:    int32(v.Gender),
				Email:     v.Email,
				Mobile:    v.Mobile,
				JobID:     v.JobID,
				DeptID:    v.DeptID,
				RoleIds:   tmpRoleIds,
				JobName:   "",
				DeptName:  "",
				RoleNames: make([]string, 0),
				Motto:     v.Motto,
				Status:    int32(v.Status),
				CreatedAt: timeutil.ToDateTimeStringByTime(v.CreatedAt),
				UpdatedAt: timeutil.ToDateTimeStringByTime(v.UpdatedAt),
			})
			roleIds = append(roleIds, tmpRoleIds...)
			jobIds = append(jobIds, v.JobID)
			deptIds = append(deptIds, v.DeptID)
		}
		roleIDToNameByIds, err2 := s.sysRoleRepo.GetRoleIDToNameByIds(ctx, roleIds)
		if err2 != nil {
			return nil, err2
		}

		jobIDToNameByIds, err2 := s.sysJobRepo.GetJobIDToNameByIds(ctx, jobIds)
		if err2 != nil {
			return nil, err2
		}
		deptIDToNameByIds, err2 := s.sysDeptRepo.GetDeptIDToNameByIds(ctx, deptIds)
		if err2 != nil {
			return nil, err2
		}
		for k := range list {
			list[k].JobName = jobIDToNameByIds[list[k].JobID]
			list[k].DeptName = deptIDToNameByIds[list[k].DeptID]
			if len(list[k].RoleIds) > 0 {
				for _, id := range list[k].RoleIds {
					list[k].RoleNames = append(list[k].RoleNames, roleIDToNameByIds[id])
				}
			}
		}
		resp.List = list
	}
	err = copier.Copy(&resp.Paginator, paginatorReply)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
