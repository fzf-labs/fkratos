// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package fkratos_sys_model

import (
	"database/sql"
	"time"
)

const TableNameSysRole = "sys_role"

// SysRole mapped from table <sys_role>
type SysRole struct {
	ID            string       `gorm:"column:id;primaryKey;comment:编号" json:"id"`                 // 编号
	TenantID      string       `gorm:"column:tenant_id;comment:租户ID" json:"tenantId"`             // 租户ID
	Pid           string       `gorm:"column:pid;not null;comment:父级id" json:"pid"`               // 父级id
	Name          string       `gorm:"column:name;not null;comment:名称" json:"name"`               // 名称
	PermissionIds string       `gorm:"column:permission_ids;comment:菜单权限集合" json:"permissionIds"` // 菜单权限集合
	Remark        string       `gorm:"column:remark;comment:备注" json:"remark"`                    // 备注
	Status        int16        `gorm:"column:status;not null;comment:0=禁用 1=开启" json:"status"`    // 0=禁用 1=开启
	Sort          int64        `gorm:"column:sort;not null;comment:排序值" json:"sort"`              // 排序值
	CreatedAt     time.Time    `gorm:"column:created_at;not null;comment:创建时间" json:"createdAt"`  // 创建时间
	UpdatedAt     time.Time    `gorm:"column:updated_at;not null;comment:更新时间" json:"updatedAt"`  // 更新时间
	DeletedAt     sql.NullTime `gorm:"column:deleted_at;comment:删除时间" json:"deletedAt"`           // 删除时间
}

// TableName SysRole's table name
func (*SysRole) TableName() string {
	return TableNameSysRole
}