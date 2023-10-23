package service

import (
	"context"

	pb "fkratos/api/rpc_device/v1"
)

// DeviceList 设备表-列表数据查询
func (d *DeviceService) DeviceList(ctx context.Context, req *pb.DeviceListReq) (*pb.DeviceListReply, error) {
	return d.deviceUseCase.DeviceList(ctx, req)
}
