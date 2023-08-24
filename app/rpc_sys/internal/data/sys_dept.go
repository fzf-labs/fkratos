package data

import (
	"context"
	v1 "fkratos/api/rpc_sys/v1"
	"fkratos/app/rpc_sys/internal/biz"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_dao"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_repo"
	"fkratos/internal/errorx"

	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.SysDeptRepo = (*SysDeptRepo)(nil)

func NewSysDeptRepo(data *Data, logger log.Logger) biz.SysDeptRepo {
	l := log.NewHelper(log.With(logger, "module", "rpc_sys/data/sys_dept"))
	return &SysDeptRepo{
		data:        data,
		log:         l,
		SysDeptRepo: fkratos_sys_repo.NewSysDeptRepo(data.gorm, data.rocksCache),
	}
}

type SysDeptRepo struct {
	data *Data
	log  *log.Helper
	*fkratos_sys_repo.SysDeptRepo
}

func (s *SysDeptRepo) SysDeptStore(ctx context.Context, req *v1.SysDeptStoreReq) (*fkratos_sys_model.SysDept, error) {
	sysDeptDao := fkratos_sys_dao.Use(s.data.gorm).SysDept
	sysDept := &fkratos_sys_model.SysDept{
		Pid:         req.Pid,
		Name:        req.Name,
		FullName:    req.FullName,
		Responsible: req.Responsible,
		Phone:       req.Phone,
		Email:       req.Email,
		Type:        int16(req.Type),
		Status:      int16(req.Status),
		Sort:        int64(req.Sort),
	}
	if req.Id != "" {
		_, err := sysDeptDao.WithContext(ctx).Select(sysDeptDao.Pid, sysDeptDao.Name, sysDeptDao.FullName, sysDeptDao.Responsible, sysDeptDao.Phone, sysDeptDao.Email, sysDeptDao.Type, sysDeptDao.Status, sysDeptDao.Sort).Where(sysDeptDao.ID.Eq(req.Id)).Updates(sysDept)
		if err != nil {
			return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
	} else {
		err := sysDeptDao.WithContext(ctx).Create(sysDept)
		if err != nil {
			return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
	}
	return sysDept, nil
}

func (s *SysDeptRepo) SysDeptList(ctx context.Context) ([]*v1.SysDeptInfo, error) {
	resp := make([]*v1.SysDeptInfo, 0)
	sysDeptDao := fkratos_sys_dao.Use(s.data.gorm).SysDept
	results, err := sysDeptDao.WithContext(ctx).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	for _, role := range results {
		resp = append(resp, &v1.SysDeptInfo{
			Id:          role.ID,
			Pid:         role.Pid,
			Name:        role.Name,
			FullName:    role.FullName,
			Responsible: role.Responsible,
			Phone:       role.Phone,
			Email:       role.Email,
			Type:        int32(role.Type),
			Status:      int32(role.Status),
			Sort:        int32(role.Sort),
			CreatedAt:   role.CreatedAt.Format(timeutil.TimeLayout),
			UpdatedAt:   role.UpdatedAt.Format(timeutil.TimeLayout),
			Children:    nil,
		})
	}
	return sysDeptGenerateTree(resp), nil
}

func sysDeptGenerateTree(list []*v1.SysDeptInfo) []*v1.SysDeptInfo {
	var trees []*v1.SysDeptInfo
	// Define the top-level root and child nodes
	var roots, childs []*v1.SysDeptInfo
	for _, v := range list {
		if v.Pid == "" {
			// Determine the top-level root node
			roots = append(roots, v)
		}
		childs = append(childs, v)
	}

	for _, v := range roots {
		childTree := &v1.SysDeptInfo{
			Id:          v.Id,
			Pid:         v.Pid,
			Name:        v.Name,
			FullName:    v.FullName,
			Responsible: v.Responsible,
			Phone:       v.Phone,
			Email:       v.Email,
			Type:        v.Type,
			Status:      v.Status,
			Sort:        v.Sort,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			Children:    make([]*v1.SysDeptInfo, 0),
		}
		// recursive
		sysDeptRecursiveTree(childTree, childs)

		trees = append(trees, childTree)
	}
	return trees
}
func sysDeptRecursiveTree(tree *v1.SysDeptInfo, allNodes []*v1.SysDeptInfo) {
	for _, v := range allNodes {
		if v.Pid == "" {
			// If the current node is the top-level root node, skip
			continue
		}
		if tree.Id == v.Pid {
			childTree := &v1.SysDeptInfo{
				Id:          v.Id,
				Pid:         v.Pid,
				Name:        v.Name,
				FullName:    v.FullName,
				Responsible: v.Responsible,
				Phone:       v.Phone,
				Email:       v.Email,
				Type:        v.Type,
				Status:      v.Status,
				Sort:        v.Sort,
				CreatedAt:   v.CreatedAt,
				UpdatedAt:   v.UpdatedAt,
				Children:    make([]*v1.SysDeptInfo, 0),
			}
			sysDeptRecursiveTree(childTree, allNodes)
			tree.Children = append(tree.Children, childTree)
		}
	}
}

func (s *SysDeptRepo) GetDeptIDToNameByIds(ctx context.Context, ids []string) (map[string]string, error) {
	resp := make(map[string]string)
	dao := fkratos_sys_dao.Use(s.data.gorm).SysDept
	results, err := dao.WithContext(ctx).Where(dao.ID.In(ids...)).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	for _, v := range results {
		resp[v.ID] = v.Name
	}
	return resp, nil
}
