package gedis

import "regexp"

type CachePolicy interface {
	Before(key string)
}

type CrossPolicy struct {
	KeyRegx string       // 检查key的正则
}

func NewCrossPolicy(keyRegx string) *CrossPolicy {
	return &CrossPolicy{KeyRegx: keyRegx}
}

func (c CrossPolicy) Before(key string)  {
	if !regexp.MustCompile(c.KeyRegx).MatchString(key){
		panic("error cache key")
	}

}