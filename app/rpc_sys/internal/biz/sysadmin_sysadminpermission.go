package biz

import (
	"context"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/internal/constant"
	"fkratos/internal/errorx"
	"strings"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/util/jsonutil"
)

// SysAdminPermission 管理员-查询权限
func (s *SysAdminUseCase) SysAdminPermission(ctx context.Context, req *pb.SysAdminPermissionReq) (*pb.SysAdminPermissionReply, error) {
	resp := &pb.SysAdminPermissionReply{}
	sysAdmin, err := s.sysAdminRepo.FindOneCacheByID(ctx, req.GetAdminId())
	if err != nil {
		return nil, err
	}
	if sysAdmin.RoleIds == nil {
		return nil, errorx.AccountNotBoundRole
	}
	roleIds := make([]string, 0)
	err = jsonutil.Unmarshal(sysAdmin.RoleIds, &roleIds)
	if err != nil {
		return nil, err
	}
	sysRoles, err := s.sysRoleRepo.FindMultiCacheByIDS(ctx, roleIds)
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
		permissions, err = s.sysPermissionRepo.SysPermissionByStatus(ctx, constant.StatusEnable)
		if err != nil {
			return nil, err
		}
	} else {
		permissions, err = s.sysPermissionRepo.SysPermissionByIdsAndStatus(ctx, permissionIds, constant.StatusEnable)
		if err != nil {
			return nil, err
		}
	}
	if len(permissions) == 0 {
		return nil, errorx.AccountNotBoundRole
	}
	menus := make([]*pb.SysPermissionInfo, 0)
	for _, v := range permissions {
		menus = append(menus, &pb.SysPermissionInfo{
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

func sysAdminPermissionGenerateTree(list []*pb.SysPermissionInfo) []*pb.SysPermissionInfo {
	var trees []*pb.SysPermissionInfo
	// Define the top-level root and child nodes
	var roots, childs []*pb.SysPermissionInfo
	for _, v := range list {
		if v.Pid == "" {
			// Determine the top-level root node
			roots = append(roots, v)
		}
		childs = append(childs, v)
	}

	for _, v := range roots {
		childTree := &pb.SysPermissionInfo{
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
			Children:  make([]*pb.SysPermissionInfo, 0),
		}
		// recursive
		sysAdminPermissionRecursiveTree(childTree, childs)

		trees = append(trees, childTree)
	}
	return trees
}
func sysAdminPermissionRecursiveTree(tree *pb.SysPermissionInfo, allNodes []*pb.SysPermissionInfo) {
	for _, v := range allNodes {
		if v.Pid == "" {
			// If the current node is the top-level root node, skip
			continue
		}
		if tree.Id == v.Pid {
			childTree := &pb.SysPermissionInfo{
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
				Children:  make([]*pb.SysPermissionInfo, 0),
			}
			sysAdminPermissionRecursiveTree(childTree, allNodes)
			tree.Children = append(tree.Children, childTree)
		}
	}
}
