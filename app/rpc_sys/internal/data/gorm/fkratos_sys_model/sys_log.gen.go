// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package fkratos_sys_model

import (
	"time"

	"gorm.io/datatypes"
)

const TableNameSysLog = "sys_log"

// SysLog mapped from table <sys_log>
type SysLog struct {
	ID        string         `gorm:"column:id;primaryKey;comment:编号" json:"id"`                // 编号
	TenantID  string         `gorm:"column:tenant_id;comment:租户ID" json:"tenantId"`            // 租户ID
	AdminID   string         `gorm:"column:admin_id;not null;comment:管理员ID" json:"adminId"`    // 管理员ID
	IP        string         `gorm:"column:ip;not null;comment:ip" json:"ip"`                  // ip
	URI       string         `gorm:"column:uri;not null;comment:请求路径" json:"uri"`              // 请求路径
	Useragent string         `gorm:"column:useragent;comment:浏览器标识" json:"useragent"`          // 浏览器标识
	Header    datatypes.JSON `gorm:"column:header;comment:header" json:"header"`               // header
	Req       datatypes.JSON `gorm:"column:req;comment:请求数据" json:"req"`                       // 请求数据
	Resp      datatypes.JSON `gorm:"column:resp;comment:响应数据" json:"resp"`                     // 响应数据
	CreatedAt time.Time      `gorm:"column:created_at;not null;comment:创建时间" json:"createdAt"` // 创建时间
}

// TableName SysLog's table name
func (*SysLog) TableName() string {
	return TableNameSysLog
}
