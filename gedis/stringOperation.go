package gedis

import (
	"context"
)

// 专门处理string类型的操作

type StringOperation struct {
	ctx  context.Context   // 为什么要这个
}

func NewStringOperation() *StringOperation {
	return &StringOperation{ctx: context.Background()}
}

func (s *StringOperation) Set(key interface{},value interface{})  {
	Redis().Do (s.ctx,key,value)
}


func (s *StringOperation) Get(key string) *StringResult  {
	return NewStringResult(Redis().Get(s.ctx,key).Result())
}

func (s *StringOperation) MGet(keys ...string) *SliceResult {
	return  NewSliceResult(Redis().MGet(s.ctx,keys...).Result())
}








