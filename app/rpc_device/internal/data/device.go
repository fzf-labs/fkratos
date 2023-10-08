package data

import (
	"fkratos/app/rpc_device/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.DeviceRepo = (*DeviceRepo)(nil)

func NewDeviceRepo(logger log.Logger, data *Data) biz.DeviceRepo {
	l := log.NewHelper(log.With(logger, "module", "data/device"))
	return &DeviceRepo{
		log:  l,
		data: data,
	}
}

type DeviceRepo struct {
	log  *log.Helper
	data *Data
}
