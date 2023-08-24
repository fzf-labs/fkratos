package cache

import (
	"fkratos/internal/constant"
	"time"

	"github.com/fzf-labs/fpkg/cache/cachekey"
)

var cache = cachekey.NewKeyManage(constant.RPCCommon)

// 缓存key前缀

var (
	DL                  = cache.AddKey("dl", time.Second*5, "分布式锁")
	SensitiveWordsCache = cache.AddKey("SensitiveWordsCache", time.Hour*24, "敏感词缓存")
)
