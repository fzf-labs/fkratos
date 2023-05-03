package cache

import (
	"fkratos/internal/service"
	"time"

	"github.com/fzf-labs/fpkg/cache/cachekey"
)

var cache = cachekey.NewKeyPrefixes(service.RpcSys)

// 缓存key前缀
var (
	UUID          = cache.AddKeyPrefix("uuid", time.Hour, "uuid")
	DL            = cache.AddKeyPrefix("dl", time.Second*5, "分布式锁")
	Sms           = cache.AddKeyPrefix("sms", time.Minute*5, "短信验证")
	SmsDayNum     = cache.AddKeyPrefix("sms_day_num", time.Minute*5, "短信发送次数")
	SensitiveWord = cache.AddKeyPrefix("sensitive_word", time.Hour*24, "敏感词")
	TinyUrl       = cache.AddKeyPrefix("tiny_url", time.Hour*24, "短连接")
)

// admin
var (
	SysAdminInfo       = cache.AddKeyPrefix("sys_admin_info", time.Minute*5, "管理员信息")
	SysAdminPermission = cache.AddKeyPrefix("sys_admin_permission", time.Hour*1, "管理员权限")
)
