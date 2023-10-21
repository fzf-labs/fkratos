package biz

import (
	"context"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_device/v1"
)

// DeviceDel 设备表-删除多条数据
func (d *DeviceUseCase) DeviceDel(ctx context.Context, req *pb.DeviceDelReq) (*pb.DeviceDelReply, error) {
	resp := &pb.DeviceDelReply{}
	err := d.deviceRepo.DeleteMultiCacheByIDS(ctx, req.GetIDS())
	if err != nil {
		return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
