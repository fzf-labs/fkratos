package cache

import (
	"fkratos/internal/constant"
	"time"

	"github.com/fzf-labs/fpkg/keymanage"
)

var cache = keymanage.New(constant.RPCUser)

// 缓存key前缀
var (
	UUID = cache.AddKey("uuid", time.Hour, "uuid")
	DL   = cache.AddKey("dl", time.Second*5, "分布式锁")
)
