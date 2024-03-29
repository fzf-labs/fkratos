// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package fkratos_wallet_model

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePay = "pay"

// Pay mapped from table <pay>
type Pay struct {
	ID        string         `gorm:"column:id;primaryKey;default:gen_random_uuid();comment:Id" json:"id"` // Id
	UID       string         `gorm:"column:uid;not null;comment:uid" json:"uid"`                          // uid
	Status    int16          `gorm:"column:status;not null;comment:状态" json:"status"`                     // 状态
	CreatedAt time.Time      `gorm:"column:created_at;not null;comment:创建时间" json:"createdAt"`            // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;comment:更新时间" json:"updatedAt"`            // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deletedAt"`                     // 删除时间
}

// TableName Pay's table name
func (*Pay) TableName() string {
	return TableNamePay
}
