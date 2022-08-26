package gedis

import (
	"regexp"
	"time"
)

type CachePolicy interface {
	Before(key string)
	IfNil (key string,v interface{})
	SetOperation (opt *StringOperation)
}

type CrossPolicy struct {
	KeyRegx string       // 检查key的正则
	Expire time.Duration
	opt *StringOperation
}

func NewCrossPolicy(keyRegx string, expire time.Duration) *CrossPolicy {
	return &CrossPolicy{KeyRegx: keyRegx, Expire: expire}
}



func (c *CrossPolicy) Before(key string)  {
	if !regexp.MustCompile(c.KeyRegx).MatchString(key){
		panic("error cache key")
	}
}

func (c *CrossPolicy) IfNil(key string,v interface{})  {
	c.opt.Set(key,v,WithExpire(c.Expire)).Unwrap()
}

func (c *CrossPolicy) SetOperation(opt *StringOperation)  {
	c.opt = opt

}