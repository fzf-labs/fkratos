package biz

import (
	"context"
	rpcDeviceV1 "fkratos/api/rpc_device/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type DeviceRepo interface {
	CreateDevice()
}

func NewDeviceUseCase(repo DeviceRepo, logger log.Logger) *DeviceUseCase {
	l := log.NewHelper(log.With(logger, "module", "rpc_Device/biz"))
	return &DeviceUseCase{
		repo: repo,
		log:  l,
	}
}

type DeviceUseCase struct {
	repo DeviceRepo
	log  *log.Helper
}

func (u *DeviceUseCase) CreateDevice(ctx context.Context, req *rpcDeviceV1.CreateDeviceReq) (*rpcDeviceV1.CreateDeviceReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *DeviceUseCase) UpdateDevice(ctx context.Context, req *rpcDeviceV1.UpdateDeviceReq) (*rpcDeviceV1.UpdateDeviceReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *DeviceUseCase) DeleteDevice(ctx context.Context, req *rpcDeviceV1.DeleteDeviceReq) (*rpcDeviceV1.DeleteDeviceReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *DeviceUseCase) GetDevice(ctx context.Context, req *rpcDeviceV1.GetDeviceReq) (*rpcDeviceV1.GetDeviceReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u *DeviceUseCase) ListDevice(ctx context.Context, req *rpcDeviceV1.ListDeviceReq) (*rpcDeviceV1.ListDeviceReply, error) {
	//TODO implement me
	panic("implement me")
}
