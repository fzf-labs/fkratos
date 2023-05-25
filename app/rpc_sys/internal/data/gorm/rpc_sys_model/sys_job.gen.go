// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package rpc_sys_model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysJob = "sys_job"

// SysJob mapped from table <sys_job>
type SysJob struct {
	ID        string         `gorm:"column:id;primaryKey" json:"id"`              // 编号
	Name      string         `gorm:"column:name;not null" json:"name"`            // 岗位名称
	Code      string         `gorm:"column:code" json:"code"`                     // 岗位编码
	Remark    string         `gorm:"column:remark" json:"remark"`                 // 备注
	Sort      int64          `gorm:"column:sort;not null" json:"sort"`            // 排序值
	Status    int16          `gorm:"column:status;not null" json:"status"`        // 0=禁用 1=开启
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"createdAt"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`          // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`          // 删除时间
}

// TableName SysJob's table name
func (*SysJob) TableName() string {
	return TableNameSysJob
}
