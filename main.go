package main

import (
	"github.com/gin-gonic/gin"
	"goft-redis/lib"
)

func main()  {
	//iter := gedis.NewStringOperation().
	//	MGet("name","age","abc").Iter()   // 变成自己的迭代器
	//
	//for  iter.HasNext() {
	//	fmt.Println(iter.Next())
	//}

	//fmt.Println(gedis.
	//	NewStringOperation().
	//	Set("name","xuchuan",
	//		gedis.WithExpire(time.Second*5),
	//		gedis.WithNx(),   // setnx
	//		))

	/*
	// 新闻缓存，假设我们认为它的过期时间 = 15s
	newsCache := gedis.NewSimpleCache(gedis.NewStringOperation(),time.Second*15)
	// 新闻缓存key : news123 news101

	newsID := 2
	newsCache.DBGetter = func() string {
		log.Println("get from db")
		newsModel := lib.NewNewsModel()
		lib.Gorm.Table("mynews").Where("id=?",newsID).Find(newsModel)
		b, _ := json.Marshal(newsModel)
		return string(b)
	}
	//fmt.Println(newsCache.GetCache("news123"))
	//fmt.Println(newsCache.GetCache("news123").(*lib.NewsModel).NewsTitle)
	fmt.Println(newsCache.GetCache("news123"))
	*/

	r := gin.New()
	r.Use(func(context *gin.Context) {
		defer func() {
			if e := recover();e!=nil {
				context.JSON(400,gin.H{"message":e})
			}
		}()
		context.Next()
	})
	r.Handle("GET","/news/:id", func(context *gin.Context) {
		// 1、从对象池 获取新闻缓存 对象
		newsCache := lib.NewsCache()
		defer lib.ReleaseNewsCache(newsCache)
		// 2、获取参数，设置DBGetter
		newsID := context.Param("id")
		newsCache.DBGetter = lib.NewsDBGetter(newsID) // 一旦缓存没有，则需要从数据库中去取
		// 3、取缓存输出（一旦没有，上面的DBGetter会被调用）
		context.Header("Context-type","application/json")
		context.String(200,newsCache.GetCache("news"+newsID).(string))

	})
	r.Run(":8080")

}

