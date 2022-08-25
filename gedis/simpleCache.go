package gedis

import "time"

type DBGetterFunc func()  string
type SimpleCache struct {
	 Operation  *StringOperation
	 Expire 	time.Duration
	 DBGetter   DBGetterFunc
}

func NewSimpleCache(Operation *StringOperation, expire time.Duration) *SimpleCache {
	return &SimpleCache{Operation: Operation, Expire: expire}
}

// 设置缓存
func (s *SimpleCache) SetCache(key string,value interface{})  {
	s.Operation.Set(key,value,WithExpire(s.Expire)).Unwrap()
}

// 获取缓存
func (s *SimpleCache) GetCache(key string) (ret interface{})  {
	ret = s.Operation.Get(key).UnwrapOrElse(s.DBGetter)
	s.SetCache(key,ret)
	return ret
}
