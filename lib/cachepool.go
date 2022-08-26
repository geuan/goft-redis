package lib

import (
	"goft-redis/gedis"
	"sync"
	"time"
)

var NewsCachePool *sync.Pool

func init()  {
	NewsCachePool = &sync.Pool{
		New: func() interface{} {
			return gedis.NewSimpleCache(gedis.NewStringOperation(),time.Second*15,
				gedis.Serilizer_JSON,gedis.NewCrossPolicy("^news\\d{1,2}$")) // 指定序列化 是 JSON
		},
	}
}

func NewsCache()  *gedis.SimpleCache {
	return NewsCachePool.Get().(*gedis.SimpleCache)

}

func ReleaseNewsCache(cache *gedis.SimpleCache)  {
	NewsCachePool.Put(cache)
	
}