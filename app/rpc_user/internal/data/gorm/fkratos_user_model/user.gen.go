// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package fkratos_user_model

import (
	"database/sql"

	"gorm.io/datatypes"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID        string         `gorm:"column:id;primaryKey;comment:ID" json:"id"`                 // ID
	UID       int64          `gorm:"column:uid;not null;comment:uid" json:"uid"`                // uid
	Username  string         `gorm:"column:username;comment:用户名" json:"username"`               // 用户名
	Phone     string         `gorm:"column:phone;comment:手机" json:"phone"`                      // 手机
	Email     string         `gorm:"column:email;comment:邮箱" json:"email"`                      // 邮箱
	Password  string         `gorm:"column:password;not null;comment:密码" json:"password"`       // 密码
	Salt      string         `gorm:"column:salt;not null;comment:盐值" json:"salt"`               // 盐值
	Nickname  string         `gorm:"column:nickname;default;comment:昵称" json:"nickname"`        // 昵称
	Sex       int16          `gorm:"column:sex;comment:性别（0未知 1男 2女）" json:"sex"`               // 性别（0未知 1男 2女）
	Avatar    string         `gorm:"column:avatar;default;comment:头像" json:"avatar"`            // 头像
	Profile   string         `gorm:"column:profile;comment:简介" json:"profile"`                  // 简介
	Other     datatypes.JSON `gorm:"column:other;comment:其他" json:"other"`                      // 其他
	Status    int16          `gorm:"column:status;not null;default:1;comment:状态" json:"status"` // 状态
	CreatedAt sql.NullTime   `gorm:"column:created_at;comment:创建时间" json:"createdAt"`           // 创建时间
	UpdatedAt sql.NullTime   `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`           // 更新时间
	DeletedAt sql.NullTime   `gorm:"column:deleted_at;comment:删除时间" json:"deletedAt"`           // 删除时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
