package service

import (
	"context"

	pb "fkratos/api/rpc_device/v1"
)

func (d *DeviceService) DeviceDel(ctx context.Context, req *pb.DeviceDelReq) (*pb.DeviceDelReply, error) {
	return d.deviceUseCase.DeviceDel(ctx, req)
}
