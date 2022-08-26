package gedis

import (
	"encoding/json"
	"time"
)

const   (
	Serilizer_JSON = "json"
)

type DBGetterFunc func()  interface{}
type SimpleCache struct {
	 Operation  *StringOperation // 操作类
	 Expire 	time.Duration // 过期时间
	 DBGetter   DBGetterFunc  // 一旦缓存没有，DB获取的方法
	 Serilizer string  //  序列化方式
	 Policy  CachePolicy
}

func NewSimpleCache(operation *StringOperation, expire time.Duration, serilizer string, policy CachePolicy) *SimpleCache {
	policy.SetOperation(operation)
	return &SimpleCache{Operation: operation, Expire: expire, Serilizer: serilizer,Policy:policy}
}



// 设置缓存
func (s *SimpleCache) SetCache(key string,value interface{})  {
	s.Operation.Set(key,value,WithExpire(s.Expire)).Unwrap()
}

// 获取缓存
func (s *SimpleCache) GetCache(key string) (ret interface{})  {
	if s.Policy != nil {  // 检查策略
		s.Policy.Before(key)
	}
	if s.Serilizer == Serilizer_JSON {
		f := func() string {
			obj := s.DBGetter()
			b,err := json.Marshal(obj)
			if err != nil {
				return ""
			}
			return  string(b)
		}
		ret = s.Operation.Get(key).UnwrapOrElse(f)
		s.SetCache(key,ret)
	}

	if ret.(string) == "" && s.Policy != nil {
		s.Policy.IfNil(key,"")
	} else {
		s.SetCache(key,ret)
	}

	return ret
}
