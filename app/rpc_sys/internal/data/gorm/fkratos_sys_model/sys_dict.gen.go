// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package fkratos_sys_model

import (
	"database/sql"
	"time"
)

const TableNameSysDict = "sys_dict"

// SysDict mapped from table <sys_dict>
type SysDict struct {
	ID        string       `gorm:"column:id;primaryKey;comment:编号" json:"id"`                                                     // 编号
	Pid       string       `gorm:"column:pid;not null;comment:0=配置集 !0=父级id" json:"pid"`                                          // 0=配置集 !0=父级id
	Name      string       `gorm:"column:name;not null;comment:名称" json:"name"`                                                   // 名称
	Type      int16        `gorm:"column:type;not null;comment:1文本 2数字 3数组 4单选 5多选 6下拉 7日期 8时间 9单图 10多图 11单文件 12多文件" json:"type"` // 1文本 2数字 3数组 4单选 5多选 6下拉 7日期 8时间 9单图 10多图 11单文件 12多文件
	UniqueKey string       `gorm:"column:unique_key;not null;comment:唯一值" json:"uniqueKey"`                                       // 唯一值
	Value     string       `gorm:"column:value;not null;comment:配置值" json:"value"`                                                // 配置值
	Status    int16        `gorm:"column:status;not null;comment:0=禁用 1=开启" json:"status"`                                        // 0=禁用 1=开启
	Sort      float64      `gorm:"column:sort;not null;comment:排序值" json:"sort"`                                                  // 排序值
	Remark    string       `gorm:"column:remark;not null;comment:备注" json:"remark"`                                               // 备注
	CreatedAt time.Time    `gorm:"column:created_at;not null;comment:创建时间" json:"createdAt"`                                      // 创建时间
	UpdatedAt time.Time    `gorm:"column:updated_at;not null;comment:更新时间" json:"updatedAt"`                                      // 更新时间
	DeletedAt sql.NullTime `gorm:"column:deleted_at;comment:删除时间" json:"deletedAt"`                                               // 删除时间
}

// TableName SysDict's table name
func (*SysDict) TableName() string {
	return TableNameSysDict
}
