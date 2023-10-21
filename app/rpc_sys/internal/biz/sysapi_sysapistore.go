package biz

import (
	"context"
	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"

	pb "fkratos/api/rpc_sys/v1"
)

// SysAPIStore API-保存
func (s *SysAPIUseCase) SysAPIStore(ctx context.Context, req *pb.SysAPIStoreReq) (*pb.SysAPIStoreReply, error) {
	resp := &pb.SysAPIStoreReply{}
	_, err := s.sysAPIRepo.SysAPIStore(ctx, &fkratos_sys_model.SysAPI{
		PermissionID: req.GetPermissionID(),
		Method:       req.GetMethod(),
		Path:         req.GetPath(),
		Desc:         req.GetDesc(),
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
