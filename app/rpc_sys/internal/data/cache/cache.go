package cache

import (
	"time"

	fCache "github.com/fzf-labs/fpkg/cache"
)

var cache = fCache.NewKeyPrefixes("user")

// 缓存key前缀
var (
	UUID          = cache.AddKeyPrefix("uuid", time.Hour, "uuid")
	DL            = cache.AddKeyPrefix("dl", time.Second*5, "分布式锁")
	Sms           = cache.AddKeyPrefix("sms", time.Minute*5, "短信验证")
	SmsDayNum     = cache.AddKeyPrefix("sms_day_num", time.Minute*5, "短信发送次数")
	SensitiveWord = cache.AddKeyPrefix("sensitive_word", time.Hour*24, "敏感词")
)
