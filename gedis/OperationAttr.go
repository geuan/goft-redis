package gedis

import (
	"fmt"
	"time"
)

const  (
	ATTR_EXPIRE = "expr"   //过期时间
	ATTR_NX = "nx"         // setnx
	ATTR_XX = "xx"         // setxx
)

type empty struct {}  // 空结构不占用空间
// 属性
type OperationAttr struct {
	Name string
	Value interface{}
}

type OperationAttrs []*OperationAttr

func (o OperationAttrs) Find(name string) *InterfaceResult  {
	for _,attr := range o {
		if attr.Name == name {
			return NewInterfaceResult(attr.Value,nil)
		}
	}
	return NewInterfaceResult(nil,fmt.Errorf("OperationAttrs found error:%s",name))
}

func WithExpire(t time.Duration) *OperationAttr  {
	return &OperationAttr{
		Name:  ATTR_EXPIRE,
		Value: t,
	}
}

func WithNx() *OperationAttr  {
	return &OperationAttr{
		Name:  ATTR_NX,
		Value: empty{},
	}

}