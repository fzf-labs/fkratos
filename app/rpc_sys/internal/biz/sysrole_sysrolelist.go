package biz

import (
	"context"
	v1 "fkratos/api/bff_admin/v1"

	pb "fkratos/api/rpc_sys/v1"

	"github.com/fzf-labs/fpkg/util/timeutil"
)

// SysRoleList 角色-列表
func (s *SysRoleUseCase) SysRoleList(ctx context.Context, _ *pb.SysRoleListReq) (*pb.SysRoleListResp, error) {
	resp := &pb.SysRoleListResp{}
	list, err := s.sysRoleRepo.SysRoleList(ctx)
	if err != nil {
		return nil, err
	}
	roles := make([]*v1.SysRoleInfo, 0)
	for _, role := range list {
		roles = append(roles, &v1.SysRoleInfo{
			Id:        role.ID,
			Pid:       role.Pid,
			Name:      role.Name,
			Remark:    role.Remark,
			Status:    int32(role.Status),
			Sort:      int32(role.Sort),
			CreatedAt: role.CreatedAt.Format(timeutil.TimeLayout),
			UpdatedAt: role.UpdatedAt.Format(timeutil.TimeLayout),
			Children:  nil,
		})
	}
	sysRoleGenerateTree(roles)
	return resp, nil
}
func sysRoleGenerateTree(list []*v1.SysRoleInfo) []*v1.SysRoleInfo {
	var trees []*v1.SysRoleInfo
	// Define the top-level root and child nodes
	var roots, childs []*v1.SysRoleInfo
	for _, v := range list {
		if v.Pid != "" {
			// Determine the top-level root node
			roots = append(roots, v)
		}
		childs = append(childs, v)
	}

	for _, v := range roots {
		childTree := &v1.SysRoleInfo{
			Id:        v.Id,
			Pid:       v.Pid,
			Name:      v.Name,
			Remark:    v.Remark,
			Status:    v.Status,
			Sort:      v.Sort,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Children:  make([]*v1.SysRoleInfo, 0),
		}
		// recursive
		sysRoleRecursiveTree(childTree, childs)

		trees = append(trees, childTree)
	}
	return trees
}
func sysRoleRecursiveTree(tree *v1.SysRoleInfo, allNodes []*v1.SysRoleInfo) {
	for _, v := range allNodes {
		if v.Pid != "" {
			// If the current node is the top-level root node, skip
			continue
		}
		if tree.Id == v.Pid {
			childTree := &v1.SysRoleInfo{
				Id:        v.Id,
				Pid:       v.Pid,
				Name:      v.Name,
				Remark:    v.Remark,
				Status:    v.Status,
				Sort:      v.Sort,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
				Children:  make([]*v1.SysRoleInfo, 0),
			}
			sysRoleRecursiveTree(childTree, allNodes)
			tree.Children = append(tree.Children, childTree)
		}
	}
}
