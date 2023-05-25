package cache

import (
	"fkratos/internal/service"
	"time"

	"github.com/fzf-labs/fpkg/cache/cachekey"
)

var cache = cachekey.NewKeyPrefixes(service.RpcCommon)

// 缓存key前缀

var (
	DL = cache.AddKeyPrefix("dl", time.Second*5, "分布式锁")
)
