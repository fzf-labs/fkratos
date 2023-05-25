package cache

import (
	"time"

	"github.com/fzf-labs/fpkg/cache/cachekey"
)

var cache = cachekey.NewKeyPrefixes("user")

// 缓存key前缀
var (
	UUID = cache.AddKeyPrefix("uuid", time.Hour, "uuid")
	DL   = cache.AddKeyPrefix("dl", time.Second*5, "分布式锁")
)
