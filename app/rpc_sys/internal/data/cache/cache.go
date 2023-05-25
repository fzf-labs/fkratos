package cache

import (
	"fkratos/internal/constant"
	"time"

	"github.com/fzf-labs/fpkg/cache/cachekey"
)

var cache = cachekey.NewKeyManage(constant.RpcSys)

// 缓存key前缀
var (
	UUID          = cache.AddKey("uuid", time.Hour, "uuid")
	DL            = cache.AddKey("dl", time.Second*5, "分布式锁")
	Sms           = cache.AddKey("sms", time.Minute*5, "短信验证")
	SmsDayNum     = cache.AddKey("sms_day_num", time.Minute*5, "短信发送次数")
	SensitiveWord = cache.AddKey("sensitive_word", time.Hour*24, "敏感词")
	TinyUrl       = cache.AddKey("tiny_url", time.Hour*24, "短连接")
)

// admin
var (
	SysAdminInfo       = cache.AddKey("sys_admin_info", time.Minute*5, "管理员信息")
	SysAdminPermission = cache.AddKey("sys_admin_permission", time.Hour*1, "管理员权限")
)
