package biz

import (
	"context"
	"fkratos/api/paginator"
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/internal/constant"
	"fkratos/internal/errorx"
	"strings"

	"github.com/fzf-labs/fpkg/crypt"
	"github.com/fzf-labs/fpkg/third_api/avatar"
	"github.com/fzf-labs/fpkg/util/jsonutil"
	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

func NewAdminUseCase(logger log.Logger, sysAdminRepo SysAdminRepo, sysRoleRepo SysRoleRepo, sysJobRepo SysJobRepo, sysDeptRepo SysDeptRepo, sysPermissionRepo SysPermissionRepo) *AdminUseCase {
	l := log.NewHelper(log.With(logger, "module", "rpc_user/biz/admin"))
	return &AdminUseCase{
		log:               l,
		sysAdminRepo:      sysAdminRepo,
		sysRoleRepo:       sysRoleRepo,
		sysJobRepo:        sysJobRepo,
		sysDeptRepo:       sysDeptRepo,
		sysPermissionRepo: sysPermissionRepo,
	}
}

type AdminUseCase struct {
	log               *log.Helper
	sysAdminRepo      SysAdminRepo
	sysRoleRepo       SysRoleRepo
	sysJobRepo        SysJobRepo
	sysDeptRepo       SysDeptRepo
	sysPermissionRepo SysPermissionRepo
}

// SysAdminInfo 查询管理员信息
func (a *AdminUseCase) SysAdminInfo(ctx context.Context, req *v1.SysAdminInfoReq) (*v1.SysAdminInfoReply, error) {
	resp := new(v1.SysAdminInfoReply)
	sysAdmin, err := a.sysAdminRepo.FindOneCacheByID(ctx, req.GetAdminId())
	if err != nil {
		return nil, err
	}
	if sysAdmin == nil {
		return resp, nil
	}
	roleIds := make([]string, 0)
	if sysAdmin.RoleIds.String() != "" {
		err = jsonutil.Decode(sysAdmin.RoleIds, &roleIds)
		if err != nil {
			return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
	}
	resp.Info = &v1.SysAdminInfo{
		Id:       sysAdmin.ID,
		Username: sysAdmin.Username,
		Nickname: sysAdmin.Nickname,
		Avatar:   sysAdmin.Avatar,
		Gender:   int32(sysAdmin.Gender),
		Email:    sysAdmin.Email,
		Mobile:   sysAdmin.Mobile,
		JobID:    sysAdmin.JobID,
		DeptID:   sysAdmin.DeptID,
		RoleIds:  roleIds,
		Motto:    sysAdmin.Motto,
	}
	return resp, nil
}

// SysAdminInfoUpdate 管理员更新
func (a *AdminUseCase) SysAdminInfoUpdate(ctx context.Context, req *v1.SysAdminInfoUpdateReq) (*v1.SysAdminInfoUpdateReply, error) {
	resp := new(v1.SysAdminInfoUpdateReply)
	sysAdminInfo, err := a.sysAdminRepo.FindOneCacheByID(ctx, req.GetAdminId())
	if err != nil {
		return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	if sysAdminInfo == nil {
		return nil, errorx.AccountNotExist
	}
	var pwd string
	if req.Password != "" {
		pwd, err = crypt.Encrypt(req.Password + sysAdminInfo.Salt)
		if err != nil {
			return resp, errorx.DataProcessingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
	}
	sysAdminInfo.Password = pwd
	sysAdminInfo.Nickname = req.Nickname
	sysAdminInfo.Email = req.Email
	sysAdminInfo.Mobile = req.Mobile
	sysAdminInfo.Motto = req.Motto
	err = a.sysAdminRepo.UpdateOne(ctx, sysAdminInfo)
	if err != nil {
		return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}

func (a *AdminUseCase) SysAdminGenerateAvatar(_ context.Context, _ *v1.SysAdminGenerateAvatarReq) (*v1.SysAdminGenerateAvatarReply, error) {
	return &v1.SysAdminGenerateAvatarReply{AvatarUrl: avatar.URL()}, nil
}

func (a *AdminUseCase) SysManageList(ctx context.Context, req *paginator.PaginatorReq) (*v1.SysManageListReply, error) {
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
			err2 := jsonutil.Decode(v.RoleIds, &tmpRoleIds)
			if err2 != nil {
				return nil, err2
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
		roleIDToNameByIds, err2 := a.sysRoleRepo.GetRoleIDToNameByIds(ctx, roleIds)
		if err2 != nil {
			return nil, err2
		}

		jobIDToNameByIds, err2 := a.sysJobRepo.GetJobIDToNameByIds(ctx, jobIds)
		if err2 != nil {
			return nil, err2
		}
		deptIDToNameByIds, err2 := a.sysDeptRepo.GetDeptIDToNameByIds(ctx, deptIds)
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
	err = copier.Copy(&resp.Paginator, p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AdminUseCase) SysManageInfo(ctx context.Context, req *v1.SysManageInfoReq) (*v1.SysManageInfoReply, error) {
	resp := new(v1.SysManageInfoReply)
	sysAdmin, err := a.sysAdminRepo.FindOneCacheByID(ctx, req.GetAdminId())
	if err != nil {
		return nil, err
	}
	if sysAdmin != nil {
		tmpRoleIds := make([]string, 0)
		err2 := jsonutil.Decode(sysAdmin.RoleIds, &tmpRoleIds)
		if err2 != nil {
			return nil, err2
		}
		roleIDToNameByIds, err2 := a.sysRoleRepo.GetRoleIDToNameByIds(ctx, tmpRoleIds)
		if err2 != nil {
			return nil, err2
		}
		jobIDToNameByIds, err2 := a.sysJobRepo.GetJobIDToNameByIds(ctx, []string{sysAdmin.JobID})
		if err2 != nil {
			return nil, err2
		}
		deptIDToNameByIds, err2 := a.sysDeptRepo.GetDeptIDToNameByIds(ctx, []string{sysAdmin.DeptID})
		if err2 != nil {
			return nil, err2
		}
		roleNames := make([]string, 0)
		for _, v := range tmpRoleIds {
			roleNames = append(roleNames, roleIDToNameByIds[v])
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
	err := a.sysAdminRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AdminUseCase) SysAdminPermission(ctx context.Context, req *v1.SysAdminPermissionReq) (*v1.SysAdminPermissionReply, error) {
	resp := new(v1.SysAdminPermissionReply)
	sysAdmin, err := a.sysAdminRepo.FindOneCacheByID(ctx, req.GetAdminId())
	if err != nil {
		return nil, err
	}
	if sysAdmin.RoleIds == nil {
		return nil, errorx.AccountNotBoundRole
	}
	roleIds := make([]string, 0)
	err = jsonutil.Decode(sysAdmin.RoleIds, &roleIds)
	if err != nil {
		return nil, err
	}
	sysRoles, err := a.sysRoleRepo.FindMultiCacheByIDS(ctx, roleIds)
	if err != nil {
		return nil, err
	}
	if len(sysRoles) == 0 {
		return nil, errorx.AccountNotBoundRole
	}
	var super bool
	permissionIds := make([]string, 0)
	for _, role := range sysRoles {
		if role.PermissionIds == "" {
			continue
		}
		if role.PermissionIds == "*" {
			super = true
			break
		}
		int64s := strings.Split(role.PermissionIds, ",")
		permissionIds = append(permissionIds, int64s...)
	}
	var permissions []*fkratos_sys_model.SysPermission
	if super {
		permissions, err = a.sysPermissionRepo.SysPermissionByStatus(ctx, constant.StatusEnable)
		if err != nil {
			return nil, err
		}
	} else {
		permissions, err = a.sysPermissionRepo.SysPermissionByIdsAndStatus(ctx, permissionIds, constant.StatusEnable)
		if err != nil {
			return nil, err
		}
	}
	if len(permissions) == 0 {
		return nil, errorx.AccountNotBoundRole
	}
	menus := make([]*v1.SysPermissionInfo, 0)
	for _, v := range permissions {
		menus = append(menus, &v1.SysPermissionInfo{
			Id:        v.ID,
			Pid:       v.Pid,
			Type:      v.Type,
			Title:     v.Title,
			Name:      v.Name,
			Path:      v.Path,
			Icon:      v.Icon,
			MenuType:  v.MenuType,
			Url:       v.URL,
			Component: v.Component,
			Extend:    v.Extend,
		})
	}
	menus = sysAdminPermissionGenerateTree(menus)
	resp.List = menus
	return resp, nil
}
func sysAdminPermissionGenerateTree(list []*v1.SysPermissionInfo) []*v1.SysPermissionInfo {
	var trees []*v1.SysPermissionInfo
	// Define the top-level root and child nodes
	var roots, childs []*v1.SysPermissionInfo
	for _, v := range list {
		if v.Pid == "" {
			// Determine the top-level root node
			roots = append(roots, v)
		}
		childs = append(childs, v)
	}

	for _, v := range roots {
		childTree := &v1.SysPermissionInfo{
			Id:        v.Id,
			Pid:       v.Pid,
			Type:      v.Type,
			Title:     v.Title,
			Name:      v.Name,
			Path:      v.Path,
			Icon:      v.Icon,
			MenuType:  v.MenuType,
			Url:       v.Url,
			Component: v.Component,
			Extend:    v.Extend,
			Children:  make([]*v1.SysPermissionInfo, 0),
		}
		// recursive
		sysAdminPermissionRecursiveTree(childTree, childs)

		trees = append(trees, childTree)
	}
	return trees
}
func sysAdminPermissionRecursiveTree(tree *v1.SysPermissionInfo, allNodes []*v1.SysPermissionInfo) {
	for _, v := range allNodes {
		if v.Pid == "" {
			// If the current node is the top-level root node, skip
			continue
		}
		if tree.Id == v.Pid {
			childTree := &v1.SysPermissionInfo{
				Id:        v.Id,
				Pid:       v.Pid,
				Type:      v.Type,
				Title:     v.Title,
				Name:      v.Name,
				Path:      v.Path,
				Icon:      v.Icon,
				MenuType:  v.MenuType,
				Url:       v.Url,
				Component: v.Component,
				Extend:    v.Extend,
				Children:  make([]*v1.SysPermissionInfo, 0),
			}
			sysAdminPermissionRecursiveTree(childTree, allNodes)
			tree.Children = append(tree.Children, childTree)
		}
	}
}
