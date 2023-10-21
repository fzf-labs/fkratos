package biz

import (
	"context"
	pb "fkratos/api/rpc_device/v1"
	"fkratos/app/rpc_device/internal/data/gorm/fkratos_device_model"
	"fkratos/internal/errorx"

	"github.com/fzf-labs/fpkg/util/timeutil"
)

// DeviceStore 设备表-创建一条数据
func (d *DeviceUseCase) DeviceStore(ctx context.Context, req *pb.DeviceStoreReq) (*pb.DeviceStoreReply, error) {
	resp := &pb.DeviceStoreReply{}
	device := &fkratos_device_model.Device{
		ID:              req.GetID(),
		Sn:              req.GetSn(),
		DeviceName:      req.GetDeviceName(),
		DeviceType:      req.GetDeviceType(),
		DeviceModel:     req.GetDeviceModel(),
		Desc:            req.GetDesc(),
		Certificate:     req.GetCertificate(),
		SecureKey:       req.GetSecureKey(),
		FirmwareVersion: req.GetFirmwareVersion(),
		SoftwareVersion: req.GetSoftwareVersion(),
		RegistryTime:    timeutil.ToSQLNullTime(req.GetRegistryTime().AsTime()),
		Status:          1,
	}
	if device.ID != "" {
		err := d.deviceRepo.UpdateOneWithZero(ctx, device)
		if err != nil {
			return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
	} else {
		err := d.deviceRepo.CreateOne(ctx, device)
		if err != nil {
			return nil, errorx.DataSQLErr.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
		}
	}
	resp.ID = device.ID
	return resp, nil
}
