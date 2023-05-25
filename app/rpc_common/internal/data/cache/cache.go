package cache

import (
	"fkratos/internal/constant"
	"time"

	"github.com/fzf-labs/fpkg/cache/cachekey"
)

var cache = cachekey.NewKeyManage(constant.RpcCommon)

// 缓存key前缀

var (
	DL = cache.AddKey("dl", time.Second*5, "分布式锁")
)
