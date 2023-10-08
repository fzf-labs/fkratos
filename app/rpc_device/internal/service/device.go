package service

import (
	"context"

	pb "fkratos/api/rpc_device/v1"
	"fkratos/app/rpc_device/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewDeviceService(logger log.Logger, deviceUseCase *biz.DeviceUseCase) *DeviceService {
	l := log.NewHelper(log.With(logger, "module", "service/device"))
	return &DeviceService{
		log:           l,
		deviceUseCase: deviceUseCase,
	}
}

type DeviceService struct {
	pb.UnimplementedDeviceServer
	log           *log.Helper
	deviceUseCase *biz.DeviceUseCase
}

func (s *DeviceService) CreateDevice(ctx context.Context, req *pb.CreateDeviceReq) (*pb.CreateDeviceReply, error) {
	return s.deviceUseCase.CreateDevice(ctx, req)
}
func (s *DeviceService) UpdateDevice(ctx context.Context, req *pb.UpdateDeviceReq) (*pb.UpdateDeviceReply, error) {
	return s.deviceUseCase.UpdateDevice(ctx, req)
}
func (s *DeviceService) DeleteDevice(ctx context.Context, req *pb.DeleteDeviceReq) (*pb.DeleteDeviceReply, error) {
	return s.deviceUseCase.DeleteDevice(ctx, req)
}
func (s *DeviceService) GetDevice(ctx context.Context, req *pb.GetDeviceReq) (*pb.GetDeviceReply, error) {
	return s.deviceUseCase.GetDevice(ctx, req)
}
func (s *DeviceService) ListDevice(ctx context.Context, req *pb.ListDeviceReq) (*pb.ListDeviceReply, error) {
	return s.deviceUseCase.ListDevice(ctx, req)
}
