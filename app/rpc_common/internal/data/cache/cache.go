package cache

import (
	"fkratos/internal/service"
	"time"

	"github.com/fzf-labs/fpkg/cache/cachekey"
)

var cache = cachekey.NewKeyManage(service.RpcCommon)

// 缓存key前缀

var (
	DL = cache.AddKey("dl", time.Second*5, "分布式锁")
)
