// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package rpc_sys_model

import (
	"time"

	"gorm.io/datatypes"
)

const TableNameSysLog = "sys_log"

// SysLog mapped from table <sys_log>
type SysLog struct {
	ID        string         `gorm:"column:id;primaryKey" json:"id"`              // 编号
	AdminID   string         `gorm:"column:admin_id;not null" json:"adminId"`     // 管理员ID
	IP        string         `gorm:"column:ip;not null" json:"ip"`                // ip
	URI       string         `gorm:"column:uri;not null" json:"uri"`              // 请求路径
	Useragent string         `gorm:"column:useragent" json:"useragent"`           // 浏览器标识
	Header    datatypes.JSON `gorm:"column:header" json:"header"`                 // header
	Req       datatypes.JSON `gorm:"column:req" json:"req"`                       // 请求数据
	Resp      datatypes.JSON `gorm:"column:resp" json:"resp"`                     // 响应数据
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"createdAt"` // 创建时间
}

// TableName SysLog's table name
func (*SysLog) TableName() string {
	return TableNameSysLog
}
