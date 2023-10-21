package biz

import (
	"context"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/util/timeutil"
)

// SysPermissionList 权限-列表
func (s *SysPermissionUseCase) SysPermissionList(ctx context.Context, _ *pb.SysPermissionListReq) (*pb.SysPermissionListResp, error) {
	resp := &pb.SysPermissionListResp{}
	list, err := s.sysPermissionRepo.SysPermissionList(ctx)
	if err != nil {
		return nil, err
	}
	roles := make([]*pb.SysPermissionInfo, 0)
	for _, v := range list {
		roles = append(roles, &pb.SysPermissionInfo{
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
			Remark:    v.Remark,
			Sort:      int32(v.Sort),
			Status:    int32(v.Status),
			CreatedAt: timeutil.ToDateTimeStringByTime(v.CreatedAt),
			UpdatedAt: timeutil.ToDateTimeStringByTime(v.UpdatedAt),
			Children:  nil,
		})
	}
	resp.List = sysPermMenuGenerateTree(roles)
	return resp, nil
}
func sysPermMenuGenerateTree(list []*pb.SysPermissionInfo) []*pb.SysPermissionInfo {
	var trees []*pb.SysPermissionInfo
	// Define the top-level root and child nodes
	var roots, childs []*pb.SysPermissionInfo
	for _, v := range list {
		if v.Pid != "" {
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
			Keepalive: v.Keepalive,
			Extend:    v.Extend,
			Remark:    v.Remark,
			Sort:      v.Sort,
			Status:    v.Status,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Children:  make([]*pb.SysPermissionInfo, 0),
		}
		// recursive
		sysPermMenuRecursiveTree(childTree, childs)

		trees = append(trees, childTree)
	}
	return trees
}
func sysPermMenuRecursiveTree(tree *pb.SysPermissionInfo, allNodes []*pb.SysPermissionInfo) {
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
				Keepalive: v.Keepalive,
				Extend:    v.Extend,
				Remark:    v.Remark,
				Sort:      v.Sort,
				Status:    v.Status,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
				Children:  make([]*pb.SysPermissionInfo, 0),
			}
			sysPermMenuRecursiveTree(childTree, allNodes)
			tree.Children = append(tree.Children, childTree)
		}
	}
}
