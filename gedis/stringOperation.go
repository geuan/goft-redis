package gedis

import (
	"context"
	"time"
)

// 专门处理string类型的操作

type StringOperation struct {
	ctx context.Context // 接口类型，需要实现context的四个方法
}

func NewStringOperation() *StringOperation {
	return &StringOperation{ctx: context.Background()}
}

func (s *StringOperation) Set(key string, value interface{}, attrs ...*OperationAttr) *InterfaceResult {
	exp := OperationAttrs(attrs).Find(ATTR_EXPIRE).UnwrapOr(time.Second*0).(time.Duration)
	nx := OperationAttrs(attrs).Find(ATTR_NX)
	xx := OperationAttrs(attrs).Find(ATTR_XX)
	if nx != nil {
		return  NewInterfaceResult(Redis().SetNX(s.ctx,key,value,exp).Result())
	}
	if xx != nil {
		return NewInterfaceResult(Redis().SetXX(s.ctx,key,value,exp).Result())
	}
	return NewInterfaceResult(Redis().Set(s.ctx, key, value,
		exp).Result())
}

func (s *StringOperation) Get(key string) *StringResult {
	return NewStringResult(Redis().Get(s.ctx, key).Result())
}

func (s *StringOperation) MGet(keys ...string) *SliceResult {
	return NewSliceResult(Redis().MGet(s.ctx, keys...).Result())
}
