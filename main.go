package main

import (
	"fmt"
	"goft-redis/gedis"
	"time"
)

func main()  {
	//iter := gedis.NewStringOperation().
	//	MGet("name","age","abc").Iter()   // 变成自己的迭代器
	//
	//for  iter.HasNext() {
	//	fmt.Println(iter.Next())
	//}

	fmt.Println(gedis.
		NewStringOperation().
		Set("name","xuchuan",
			gedis.WithExpire(time.Second*5),
			gedis.WithNx(),   // setnx
			))
}

