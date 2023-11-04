// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package fkratos_wechat_model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameOfficialaccount = "officialaccount"

// Officialaccount mapped from table <officialaccount>
type Officialaccount struct {
	ID                 string         `gorm:"column:id;primaryKey;default:gen_random_uuid();comment:ID" json:"id"`                     // ID
	AppID              string         `gorm:"column:app_id;not null;comment:开发者ID(AppID)" json:"appId"`                                // 开发者ID(AppID)
	AppSecret          string         `gorm:"column:app_secret;not null;comment:开发者密码(AppSecret)" json:"appSecret"`                    // 开发者密码(AppSecret)
	AccountID          string         `gorm:"column:account_id;not null;comment:账号id" json:"accountId"`                                // 账号id
	BusinessID         string         `gorm:"column:business_id;not null;comment:业务主账号id" json:"businessId"`                           // 业务主账号id
	Name               string         `gorm:"column:name;not null;comment:公众号名称" json:"name"`                                          // 公众号名称
	ExpiresAccessToken string         `gorm:"column:expires_access_token;not null;comment:获取access_token时间" json:"expiresAccessToken"` // 获取access_token时间
	ExpiresJsapiTicket string         `gorm:"column:expires_jsapi_ticket;not null;comment:获取jsapi_ticket时间" json:"expiresJsapiTicket"` // 获取jsapi_ticket时间
	AccessToken        string         `gorm:"column:access_token;not null;comment:access_token" json:"accessToken"`                    // access_token
	JsapiTicket        string         `gorm:"column:jsapi_ticket;not null;comment:jsapi_ticket" json:"jsapiTicket"`                    // jsapi_ticket
	Qrcode             string         `gorm:"column:qrcode;not null;comment:二维码" json:"qrcode"`                                        // 二维码
	Token              string         `gorm:"column:token;not null;comment:Token 长度为3-32字符" json:"token"`                              // Token 长度为3-32字符
	EncodingAseKey     string         `gorm:"column:encoding_ase_key;not null;comment:消息加密密钥由43位字符组成" json:"encodingAseKey"`           // 消息加密密钥由43位字符组成
	Remark             string         `gorm:"column:remark;not null;comment:备注" json:"remark"`                                         // 备注
	CreatedAt          time.Time      `gorm:"column:created_at;not null;comment:创建时间" json:"createdAt"`                                // 创建时间
	UpdatedAt          time.Time      `gorm:"column:updated_at;not null;comment:更新时间" json:"updatedAt"`                                // 更新时间
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deletedAt"`                                         // 删除时间
}

// TableName Officialaccount's table name
func (*Officialaccount) TableName() string {
	return TableNameOfficialaccount
}