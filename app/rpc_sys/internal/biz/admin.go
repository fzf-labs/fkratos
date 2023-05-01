package biz

import (
	"context"
	"fkratos/api/common"
	v1 "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/third_api/avatar"
	"github.com/fzf-labs/fpkg/util/jsonutil"
	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

func NewAdminUseCase(logger log.Logger, sysAdminRepo SysAdminRepo, sysRoleRepo SysRoleRepo, sysJobRepo SysJobRepo, sysDeptRepo SysDeptRepo) *AdminUseCase {
	l := log.NewHelper(log.With(logger, "module", "rpc_user/biz/admin"))
	return &AdminUseCase{
		log:          l,
		sysAdminRepo: sysAdminRepo,
		sysRoleRepo:  sysRoleRepo,
		sysJobRepo:   sysJobRepo,
		sysDeptRepo:  sysDeptRepo,
	}
}

type AdminUseCase struct {
	log          *log.Helper
	sysAdminRepo SysAdminRepo
	sysRoleRepo  SysRoleRepo
	sysJobRepo   SysJobRepo
	sysDeptRepo  SysDeptRepo
}

func (a *AdminUseCase) SysAdminInfo(ctx context.Context, req *v1.SysAdminInfoReq) (*v1.SysAdminInfoReply, error) {
	resp := new(v1.SysAdminInfoReply)

	sysAdmin, err := a.sysAdminRepo.SysAdminInfoCacheByAdminId(ctx, req.AdminId)
	if err != nil {
		return nil, err
	}
	resp.Info = sysAdmin
	return resp, nil
}

func (a *AdminUseCase) SysAdminInfoUpdate(ctx context.Context, req *v1.SysAdminInfoUpdateReq) (*v1.SysAdminInfoUpdateReply, error) {
	return a.sysAdminRepo.SysAdminInfoUpdate(ctx, req)
}

func (a *AdminUseCase) SysAdminGenerateAvatar(ctx context.Context, req *v1.SysAdminGenerateAvatarReq) (*v1.SysAdminGenerateAvatarReply, error) {
	return &v1.SysAdminGenerateAvatarReply{AvatarUrl: avatar.Url()}, nil
}

func (a *AdminUseCase) SysManageList(ctx context.Context, req *common.SearchListReq) (*v1.SysManageListReply, error) {
	resp := new(v1.SysManageListReply)
	sysAdmins, p, err := a.sysAdminRepo.SysManageListBySearch(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(sysAdmins) > 0 {
		list := make([]*v1.SysManageInfo, 0)
		roleIds := make([]string, 0)
		jobIds := make([]string, 0)
		deptIds := make([]string, 0)
		for _, v := range sysAdmins {
			tmpRoleIds := make([]string, 0)
			err := jsonutil.Decode(v.RoleIds, &tmpRoleIds)
			if err != nil {
				return nil, err
			}
			list = append(list, &v1.SysManageInfo{
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
		roleIdToNameByIds, err := a.sysRoleRepo.GetRoleIdToNameByIds(ctx, roleIds)
		if err != nil {
			return nil, err
		}

		jobIdToNameByIds, err := a.sysJobRepo.GetJobIdToNameByIds(ctx, jobIds)
		if err != nil {
			return nil, err
		}
		deptIdToNameByIds, err := a.sysDeptRepo.GetDeptIdToNameByIds(ctx, deptIds)
		if err != nil {
			return nil, err
		}
		for k := range list {
			list[k].JobName = jobIdToNameByIds[list[k].JobID]
			list[k].DeptName = deptIdToNameByIds[list[k].DeptID]
			if len(list[k].RoleIds) > 0 {
				for _, id := range list[k].RoleIds {
					list[k].RoleNames = append(list[k].RoleNames, roleIdToNameByIds[id])
				}
			}
		}
		resp.List = list

	}
	err = copier.Copy(&resp.Paginator, p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AdminUseCase) SysManageInfo(ctx context.Context, req *v1.SysManageInfoReq) (*v1.SysManageInfoReply, error) {
	resp := new(v1.SysManageInfoReply)
	sysAdmin, err := a.sysAdminRepo.SysAdminInfoByAdminId(ctx, req.GetAdminId())
	if err != nil {
		return nil, err
	}
	if sysAdmin != nil {
		tmpRoleIds := make([]string, 0)
		err := jsonutil.Decode(sysAdmin.RoleIds, &tmpRoleIds)
		if err != nil {
			return nil, err
		}
		roleIdToNameByIds, err := a.sysRoleRepo.GetRoleIdToNameByIds(ctx, tmpRoleIds)
		if err != nil {
			return nil, err
		}
		jobIdToNameByIds, err := a.sysJobRepo.GetJobIdToNameByIds(ctx, []string{sysAdmin.JobID})
		if err != nil {
			return nil, err
		}
		deptIdToNameByIds, err := a.sysDeptRepo.GetDeptIdToNameByIds(ctx, []string{sysAdmin.DeptID})
		if err != nil {
			return nil, err
		}
		roleNames := make([]string, 0)
		for _, v := range tmpRoleIds {
			roleNames = append(roleNames, roleIdToNameByIds[v])
		}
		resp.Info = &v1.SysManageInfo{
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
			JobName:   jobIdToNameByIds[sysAdmin.JobID],
			DeptName:  deptIdToNameByIds[sysAdmin.JobID],
			RoleNames: roleNames,
			Motto:     sysAdmin.Motto,
			Status:    int32(sysAdmin.Status),
			CreatedAt: timeutil.ToDateTimeStringByTime(sysAdmin.CreatedAt),
			UpdatedAt: timeutil.ToDateTimeStringByTime(sysAdmin.UpdatedAt),
		}
	}
	return resp, nil
}

func (a *AdminUseCase) SysManageStore(ctx context.Context, req *v1.SysManageStoreReq) (*v1.SysManageStoreReply, error) {
	resp := new(v1.SysManageStoreReply)
	_, err := a.sysAdminRepo.SysManageStore(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AdminUseCase) SysManageDel(ctx context.Context, req *v1.SysManageDelReq) (*v1.SysManageDelReply, error) {
	resp := new(v1.SysManageDelReply)
	err := a.sysAdminRepo.SysAdminDel(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
