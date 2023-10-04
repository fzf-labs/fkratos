// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package fkratos_device_model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameDevice = "device"

// Device mapped from table <device>
type Device struct {
	ID          string         `gorm:"column:id;primaryKey;default:gen_random_uuid();comment:记录ID" json:"id"` // 记录ID
	SNID        string         `gorm:"column:SNId;not null;comment:设备SN唯一标识码" json:"snid"`                    // 设备SN唯一标识码
	DeviceBrand string         `gorm:"column:deviceBrand;comment:设备品牌" json:"deviceBrand"`                    // 设备品牌
	ModelDevice string         `gorm:"column:modelDevice;comment:设备型号" json:"modelDevice"`                    // 设备型号
	Status      int16          `gorm:"column:status;comment:状态" json:"status"`                                // 状态
	CreatedAt   time.Time      `gorm:"column:created_at;not null;comment:创建时间" json:"createdAt"`              // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;comment:更新时间" json:"updatedAt"`              // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deletedAt"`                       // 删除时间
}

// TableName Device's table name
func (*Device) TableName() string {
	return TableNameDevice
}
