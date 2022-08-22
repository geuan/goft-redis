package main

import (
	"fmt"
	"goft-redis/gedis"
)

func main()  {
	iter := gedis.NewStringOperation().
		MGet("name","age","abc").Iter()   // 变成自己的迭代器

	for  iter.HasNext() {
		fmt.Println(iter.Next())
	}
}

