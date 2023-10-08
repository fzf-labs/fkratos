package biz

import (
	"context"

	pb "fkratos/api/rpc_device/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type DeviceRepo interface {
}

func NewDeviceUseCase(logger log.Logger, deviceRepo DeviceRepo) *DeviceUseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/device"))
	return &DeviceUseCase{
		log:        l,
		deviceRepo: deviceRepo,
	}
}

type DeviceUseCase struct {
	log        *log.Helper
	deviceRepo DeviceRepo
}

func (s *DeviceUseCase) CreateDevice(ctx context.Context, req *pb.CreateDeviceReq) (*pb.CreateDeviceReply, error) {
	resp := &pb.CreateDeviceReply{}
	return resp, nil
}
func (s *DeviceUseCase) UpdateDevice(ctx context.Context, req *pb.UpdateDeviceReq) (*pb.UpdateDeviceReply, error) {
	resp := &pb.UpdateDeviceReply{}
	return resp, nil
}
func (s *DeviceUseCase) DeleteDevice(ctx context.Context, req *pb.DeleteDeviceReq) (*pb.DeleteDeviceReply, error) {
	resp := &pb.DeleteDeviceReply{}
	return resp, nil
}
func (s *DeviceUseCase) GetDevice(ctx context.Context, req *pb.GetDeviceReq) (*pb.GetDeviceReply, error) {
	resp := &pb.GetDeviceReply{}
	return resp, nil
}
func (s *DeviceUseCase) ListDevice(ctx context.Context, req *pb.ListDeviceReq) (*pb.ListDeviceReply, error) {
	resp := &pb.ListDeviceReply{}
	return resp, nil
}
