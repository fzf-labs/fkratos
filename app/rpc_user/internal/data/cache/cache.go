package cache

import (
	"time"

	"github.com/fzf-labs/fpkg/cache/cachekey"
)

var cache = cachekey.NewKeyManage("user")

// 缓存key前缀
var (
	UUID = cache.AddKey("uuid", time.Hour, "uuid")
	DL   = cache.AddKey("dl", time.Second*5, "分布式锁")
)
