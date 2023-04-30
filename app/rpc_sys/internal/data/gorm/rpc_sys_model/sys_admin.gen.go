// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package rpc_sys_model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const TableNameSysAdmin = "sys_admin"

// SysAdmin mapped from table <sys_admin>
type SysAdmin struct {
	ID        string         `gorm:"column:id;primaryKey" json:"id"`                 // 编号
	Username  string         `gorm:"column:username;not null" json:"username"`       // 用户名
	Password  string         `gorm:"column:password;not null" json:"password"`       // 密码
	Nickname  string         `gorm:"column:nickname;not null" json:"nickname"`       // 昵称
	Avatar    string         `gorm:"column:avatar" json:"avatar"`                    // 头像
	Gender    int16          `gorm:"column:gender;not null" json:"gender"`           // 0=保密 1=女 2=男
	Email     string         `gorm:"column:email" json:"email"`                      // 邮件
	Mobile    string         `gorm:"column:mobile" json:"mobile"`                    // 手机号
	JobID     string         `gorm:"column:job_id" json:"jobId"`                     // 岗位
	DeptID    string         `gorm:"column:dept_id" json:"deptId"`                   // 部门
	RoleIds   datatypes.JSON `gorm:"column:role_ids" json:"roleIds"`                 // 角色集
	Salt      string         `gorm:"column:salt;not null" json:"salt"`               // 盐值
	Status    int16          `gorm:"column:status;not null;default:1" json:"status"` // 0=禁用 1=开启
	Motto     string         `gorm:"column:motto" json:"motto"`                      // 个性签名
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`             // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`             // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`             // 删除时间
}

// TableName SysAdmin's table name
func (*SysAdmin) TableName() string {
	return TableNameSysAdmin
}
