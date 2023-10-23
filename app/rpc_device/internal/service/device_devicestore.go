package service

import (
	"context"

	pb "fkratos/api/rpc_device/v1"
)

// DeviceStore 设备表-创建一条数据
func (d *DeviceService) DeviceStore(ctx context.Context, req *pb.DeviceStoreReq) (*pb.DeviceStoreReply, error) {
	return d.deviceUseCase.DeviceStore(ctx, req)
}
